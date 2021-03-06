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

package sql

import (
	"testing"
	"os"
	"eva"
	"github.com/jmoiron/sqlx"
	"log"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const checkPass = "\u2713"
const checkFail = "\u2717"

func TestPgSqlManager_FindCandidates(t *testing.T) {
	db := sqlInit()
	if db == nil {
		t.Errorf("FAIL! unable connect database ! %s", checkFail)
	}

	for k, testCase := range testCases {
		p, err := db.FindCandidates(testCase.input)

		if err != nil {
			t.Error(err)
		}
		if len(p) != len(testCase.output) {
			t.Errorf("FAIL! case %d length not equal ! %s", k, checkFail)
			continue
		}
		fail := false
		var output []string
		var correct []string
		for k1 := range p {
			//statement:=c.GetStatements()
			output = append(output, fmt.Sprintf("%#v", p[k1]))
			correct = append(correct, fmt.Sprintf("%#v", testCase.output[k1]))
		}
		if !isSamePolicy(output, correct) {
			t.Errorf("FAIL! case %d Policy not equal !\noutput is:%s but correct is %s %s\n", k, output, correct, checkFail)
			fail = true
			break

		}
		if !fail {
			//t.Logf("PASS ! case %d %s",k,checkPass)
		}

	}

}
//only about 2000/s on localhost?
func BenchmarkPgSqlManager_FindCandidates(b *testing.B) {
	db := sqlInit()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, x := range testCases {
			//testcase := testCases[rand.Intn(len(testCases))]
			db.FindCandidates(x.input)
			i++
		}
	}
}

func sqlInit() *PgSqlManager {

	dbDrv := os.Getenv(eva.EnvDBDriver)
	dbSrc := os.Getenv(eva.EnvDBSource)

	db, err := sqlx.Connect(dbDrv, dbSrc)

	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	//createTables(db)
	//insertPolicies(db, mock.Jps)
	//insertBindings(db, mock.Jps)

	return NewPgSqlManager(db)
}

//Judge two string slience is same but may be not in order,
//algorithm may be can improve
func isSamePolicy(s, t []string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 && len(t) == 0 {
		return true
	}
	for _, sk := range s {
		found := false
		for _, tk := range t {
			if sk == tk {
				found = true
				break
			}
		}
		if found {
			continue
		}
		if sk != t[len(t)-1] {
			return false
		}
	}
	return true
}
