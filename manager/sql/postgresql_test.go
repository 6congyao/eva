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
	"eva/policy"
	"fmt"
)

const checkPass = "\u2713"
const checkFail = "\u2717"

func TestPgSqlManager_FindCandidates(t *testing.T) {
	db:=sqlInit()
	if db==nil{
		t.Errorf("FAIL! unable connect database ! %s",checkFail)
	}
	var testCases=[]struct{
		input []string
		output policy.Policies
	}{
		{
			[]string{"qrn:partition::iam:usr-Vtl3VCfF:group/OpenPitrix/dev"},
			policy.Policies{
				&policy.DefaultPolicy{
					ID:"",
					Version:"2018-5-4",
					Name:"",
					Urn:"",
					Description:"",
					Statements:[]policy.DefaultStatement{
						policy.DefaultStatement{
							Principals:[]string(nil),
							Effect:"allow",
							Actions:[]string{"iam:CreatePolicy",
											"iam:DeletePolicy",
							},
							Resources:[]string{"*"},
						},
						policy.DefaultStatement{
							Principals:[]string(nil),
							Effect:"allow",
							Actions:[]string{"iam:CreateRole"},
							Resources:[]string{"*"},
						},
					},
				},

			},

		},
		{
			[]string{"qrn:partition::iam:usr-Vtl3VCfF:user/Lucy"},
			policy.Policies{
				&policy.DefaultPolicy{
					ID:"",
					Version:"2018-5-4",
					Name:"",
					Urn:"",
					Description:"",
					Statements:[]policy.DefaultStatement{
						policy.DefaultStatement{
							Principals:[]string(nil),
							Effect:"allow",
							Actions:[]string{"iam:CreatePolicy",
								"iam:DeletePolicy",
							},
							Resources:[]string{"*"},
						},
						policy.DefaultStatement{
							Principals:[]string(nil),
							Effect:"allow",
							Actions:[]string{"iam:CreateRole"},
							Resources:[]string{"*"},
						},
					},
				},

			},
		},
		{
			[]string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Lucy"},
			policy.Policies{
				&policy.DefaultPolicy{ID:"", Version:"2012-04-01", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"deny", Actions:[]string{"qstor:modify", "qstor:delete"}, Resources:[]string{"qstor:*"}}}},
			},
		},
		{
			[]string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Tom"},
			policy.Policies{
				&policy.DefaultPolicy{ID:"", Version:"2018-5-5", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:list", "k8s:get"}, Resources:[]string{"k8s:pods"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:watch"}, Resources:[]string{"k8s:pods", "k8s:pods/log"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2012-04-01", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"deny", Actions:[]string{"qstor:modify", "qstor:delete"}, Resources:[]string{"qstor:*"}}}},
			},
		},
		{
			[]string{"qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
			policy.Policies{
				&policy.DefaultPolicy{ID:"", Version:"2018-5-5", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:list", "k8s:get"}, Resources:[]string{"k8s:pods"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:watch"}, Resources:[]string{"k8s:pods", "k8s:pods/log"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2018-5-4", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreatePolicy", "iam:DeletePolicy"}, Resources:[]string{"*"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreateRole"}, Resources:[]string{"*"}}}},
			},
		},
		{
			[]string{"1","qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
			policy.Policies{
				&policy.DefaultPolicy{ID:"", Version:"2018-5-5", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:list", "k8s:get"}, Resources:[]string{"k8s:pods"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:watch"}, Resources:[]string{"k8s:pods", "k8s:pods/log"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2018-5-4", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreatePolicy", "iam:DeletePolicy"}, Resources:[]string{"*"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreateRole"}, Resources:[]string{"*"}}}},
			},
		},
		{
			[]string{"qrn:partition::iam:usr-Vtl3VCfF:user/Tom","qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Tom"},
			policy.Policies{
				&policy.DefaultPolicy{ID:"", Version:"2018-5-5", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:list", "k8s:get"}, Resources:[]string{"k8s:pods"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:watch"}, Resources:[]string{"k8s:pods", "k8s:pods/log"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2018-5-5", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:list", "k8s:get"}, Resources:[]string{"k8s:pods"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:watch"}, Resources:[]string{"k8s:pods", "k8s:pods/log"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2012-04-01", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"deny", Actions:[]string{"qstor:modify", "qstor:delete"}, Resources:[]string{"qstor:*"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2018-5-4", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreatePolicy", "iam:DeletePolicy"}, Resources:[]string{"*"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreateRole"}, Resources:[]string{"*"}}}},
			},
		},
		{
			[]string{ "qrn:partition::iam:usr-Vtl3VCfF:user/Tom",
							"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Tom",
							"qrn:partition::iam:usr-Vtl3VCfF:user/Lucy"},
			policy.Policies{
				&policy.DefaultPolicy{ID:"", Version:"2018-5-4", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreatePolicy", "iam:DeletePolicy"}, Resources:[]string{"*"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreateRole"}, Resources:[]string{"*"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2018-5-4", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreatePolicy", "iam:DeletePolicy"}, Resources:[]string{"*"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"iam:CreateRole"}, Resources:[]string{"*"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2012-04-01", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"deny", Actions:[]string{"qstor:modify", "qstor:delete"}, Resources:[]string{"qstor:*"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2018-5-5", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:list", "k8s:get"}, Resources:[]string{"k8s:pods"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:watch"}, Resources:[]string{"k8s:pods", "k8s:pods/log"}}}},
				&policy.DefaultPolicy{ID:"", Version:"2018-5-5", Name:"", Urn:"", Description:"", Statements:[]policy.DefaultStatement{policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:list", "k8s:get"}, Resources:[]string{"k8s:pods"}}, policy.DefaultStatement{Principals:[]string(nil), Effect:"allow", Actions:[]string{"k8s:watch"}, Resources:[]string{"k8s:pods", "k8s:pods/log"}}}},

			},
		},
		{
			[]string{"qrn:partition::iam:usr-Vtl3VCfF:user/*"},
			nil,
		},
		{
			[]string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/*"},
			nil,
		},
		{
			[]string{"qrn:partition::iam:usr-Vtl3VCfF:user/*","qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/*"},
			nil,
		},
	}
	for k,testCase:=range testCases{
		p,err:=db.FindCandidates(testCase.input)

		if err!=nil{
			t.Error(err)
		}
		if len(p)!=len(testCase.output){
			t.Errorf("FAIL! case %d length not equal ! %s",k,checkFail)
			continue
		}
		for k1,_:=range p {
			//statement:=c.GetStatements()
			output:=fmt.Sprintf("%#v",p[k1])
			correct:=fmt.Sprintf("%#v",testCase.output[k1])
			if output !=correct{
				t.Errorf("FAIL! case %d Policy not equal !\n Policy %d: output is:%s but correct is %s %s\n",k,k1,output,correct,checkFail)
				break
			}

		}
		//t.Logf("PASS ! case %d %s",k,checkPass)
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
