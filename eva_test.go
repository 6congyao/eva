/*******************************************************************************
 * Copyright (c) 2018. LuCongyao <6congyao@gmail.com> .
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this work except in compliance with the License.
 * You may obtain a copy of the License in the LICENSE file, or at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 ******************************************************************************/

package eva

import (
	"testing"
	"fmt"
	"eva/agent"
	"eva/manager/sql"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"eva/utils"
	"github.com/jmoiron/sqlx"
	"os"
	"math/rand"
	"eva/policy"
)

const checkPass = "\u2713"
const checkFail = "\u2717"

type testCase struct {
	keys            []string
	rcs             []*agent.RequestContext
	AuthorizeResult error
}

//Only test the type of error, not value, because FindCandidates has been tested.
// 8 cases
var testCases = []testCase{
	{
		keys: []string{"qrn:partition::iam:usr-Vtl3VCfF:policy/IAMPolicyAccess", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
		rcs: []*agent.RequestContext{
			{"", "k8s:list", "k8s:pods"},
			{"", "k8s:get", "k8s:pods/log"},
		},
		AuthorizeResult: utils.NewErrDefaultDenied(nil),
	},
	{
		keys: []string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Lucy"},
		rcs: []*agent.RequestContext{
			{"", "qstor:modify", "qstor:*"},
		},
		AuthorizeResult: utils.NewErrExplicitlyDenied(nil, nil),
	},
	{
		keys: []string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Lucy"},
		rcs: []*agent.RequestContext{
			{"", "qstor:modify", "qstor:sdad"},
		},
		AuthorizeResult: utils.NewErrExplicitlyDenied(nil, nil),
	},
	{
		keys: []string{"qrn:partition::iam:usr-Vtl3VCfF:group/OpenPitrix/dev", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
		rcs: []*agent.RequestContext{
			{"", "k8s:list", "k8s:pods"},
			{"", "iam:CreatePolicy", "k8s:pods/log"},
		},
		AuthorizeResult: nil,
	},
	{
		keys: []string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Tom", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
		rcs: []*agent.RequestContext{
			{"", "k8s:list", "k8s:pods"},
			{"", "k8s:watch", "k8s:pods/log"},
		},
		AuthorizeResult: nil,
	},

	{
		keys: []string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/To", "qrn:partition::iam:usr-Vtl3VCfF:user/To"},
		rcs: []*agent.RequestContext{
			{"", "k8s:list", "k8s:pods"},
			{"", "k8s:watch", "k8s:pods/log"},
		},
		AuthorizeResult: utils.NewErrDefaultDenied(nil),
	},
	{
		keys: []string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Tom", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
		rcs: []*agent.RequestContext{
			{"", "k8s:list ", "k8s:pods"}, //(Action)may be forget trim?
			{"", "k8s:watch", "k8s:pods/log"},
		},
		//AuthorizeResult: nil, //trimmed
		AuthorizeResult: utils.NewErrDefaultDenied(nil), //do not trim
	},
	{
		keys: []string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Tom", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
		rcs: []*agent.RequestContext{
			{"", "k8s:list", "k8s:pods "}, //(Resource)may be forget trim?
			{"", "k8s:watch", "k8s:pods/log"},
		},
		//AuthorizeResult: nil,//trimmed
		AuthorizeResult: utils.NewErrDefaultDenied(nil), // do not trim
	},
}
//connect, authorize...etc
var warden *Eva00

func init() {
	warden = &Eva00{
		Manager: sqlInit(),
	}
}
func TestEva00_Authorize(t *testing.T) {
	for c, k := range testCases {
		err := warden.Authorize(k.rcs, k.keys)
		if err == k.AuthorizeResult {
			//log.Printf("case %d PASSED", c)
			continue
		}
		if typeof(err) == typeof(k.AuthorizeResult) {
			//log.Printf("case %d PASSED", c)
			continue
		}
		correct := switchType(k.AuthorizeResult)
		output := switchType(err)
		t.Errorf("FAIL! case %d ,authorize result error! output is %s,but correct is %s  %s", c+1, output, correct, checkFail)
	}

}
func BenchmarkEva00_Authorize(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := testCases[rand.Intn(len(testCases))]
		warden.Authorize(c.rcs, c.keys)
	}
}

func BenchmarkEva00_Evaluate(b *testing.B) {
	b.ReportAllocs()
	//begin:
	//c := testCases[rand.Intn(len(testCases))]
	//polices, err := warden.Manager.FindCandidates(c.keys)
	//if err != nil {
	//	goto begin
	//}
	//prepare testcase
	var policesSlice []policy.Policies
	var rcsSlice [][]*agent.RequestContext
	for _, x := range testCases {
		polices, err := warden.Manager.FindCandidates(x.keys)
		if err != nil {
			continue
		}
		policesSlice = append(policesSlice, polices)
		rcsSlice = append(rcsSlice, x.rcs)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for s:=range policesSlice{
			warden.Evaluate(rcsSlice[s], policesSlice[s])
			i++
		}

	}

}

//judge err type: nil(Allow) default Explicit
func switchType(err error) string {
	switch e := err.(type) {

	case utils.ErrorWithContext:
		return e.Causer()
	case nil:
		return "Allow"
	default:
		return "Error"
	}

}

//get error type
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

//connect database
func sqlInit() *sql.PgSqlManager {

	dbDrv := os.Getenv(EnvDBDriver)
	dbSrc := os.Getenv(EnvDBSource)

	db, err := sqlx.Connect(dbDrv, dbSrc)

	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	//createTables(db)
	//insertPolicies(db, mock.Jps)
	//insertBindings(db, mock.Jps)

	return sql.NewPgSqlManager(db)
}
