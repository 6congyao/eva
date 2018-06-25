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

package agent

import (
	"testing"
	"encoding/json"
)

const checkPass = "\u2713"
const checkFail = "\u2717"

var dataSet = []string{
	`{
		"id": "1",
		"subject": ["qrn:partition::iam:usr-Vtl3VCfF:policy/IAMPolicyAccess", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"],
		"payload": [
	{
		"action": "k8s:list",
		"resource": "k8s:pods"
	},
	{
		"action": "k8s:get",
		"resource": "k8s:pods/log"
	}
	]
	}`,

	`{
		"id": "2",
		"subject": ["qrn:partition::iam:usr-Vtl3VCfF:group/OpenPitrix/dev", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"],
		"payload": [
		{
			"action": "k8s:list",
			"resource": "k8s:pods"
		},
		{
			"action": "iam:CreatePolicy",
			"resource": "k8s:pods/log"
		}
					]
	}`,

	`{
		"id": "3",
		"subject": ["qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Tom", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"],
		"payload": [
	{
		"action": "k8s:list",
		"action": "k8s:list",
		"resource": "k8s:pods"
	},
	{
		"action": "k8s:watch",
		"resource": "k8s:pods/log"
	}
				]
	}`,
}

type testCase struct {
	input *ReqAgent
	keys  []string
	rcs   []*RequestContext
}

var testCases = []testCase{
	{
		input: addReqagent(dataSet[0]),
		keys:  []string{"qrn:partition::iam:usr-Vtl3VCfF:policy/IAMPolicyAccess", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
		rcs: []*RequestContext{
			{"", "k8s:list", "k8s:pods"},
			{"", "k8s:get", "k8s:pods/log"},
		},
	},
	{
		input: addReqagent(dataSet[1]),
		keys:  []string{"qrn:partition::iam:usr-Vtl3VCfF:group/OpenPitrix/dev", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
		rcs: []*RequestContext{
			{"", "k8s:list", "k8s:pods"},
			{"", "iam:CreatePolicy", "k8s:pods/log"},
		},
	},
	{
		addReqagent(dataSet[2]),
		[]string{"qrn:partition::iam:usr-Vtl3VCfF:user/OpenPitrix/Tom", "qrn:partition::iam:usr-Vtl3VCfF:user/Tom"},
		[]*RequestContext{
			{"", "k8s:list", "k8s:pods"},
			{"", "k8s:watch", "k8s:pods/log"},
		},
	},
}

func TestNormalizeRequests(t *testing.T) {

	for c, k := range testCases {
		keyCheck := true
		rcsCheck := true
		keys, rcs, _ := k.input.NormalizeRequests()
		if !keyEqual(keys, k.keys) {
			keyCheck = false
			t.Errorf("Error! case %d keys not equal! output:%s correct:%s %s", c+1, keys, k.keys, checkFail)
		}

		if !rcsEqual(rcs, k.rcs) {
			rcsCheck = false
			output := []RequestContext{}
			for _, k := range rcs {
				output = append(output, *k)
			}
			correct := []RequestContext{}
			for _, k := range k.rcs {
				correct = append(correct, *k)
			}
			t.Errorf("Error! %d RequestContext not equal! output:%#v correct:%#v %s", c+1, output, correct, checkFail)
		}
		if keyCheck && rcsCheck {
			//t.Logf("Pass! case %d,%s", c+1, checkPass)
		}

	}
}

func BenchmarkReqAgent_NormalizeRequests(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {

		for _, x := range testCases {
			//b.StartTimer()
			x.input.NormalizeRequests()
			//b.StopTimer()
			i++
		}
	}
}

func keyEqual(s, t []string) bool {
	if s == nil && t == nil {
		return true
	}
	if len(s) != len(t) {
		return false
	}
	for k := range s {
		if s[k] != t[k] {
			return false
		}
	}
	return true
}
func rcsEqual(s, t []*RequestContext) bool {

	if s == nil && t == nil {
		return true
	}
	if len(s) != len(t) {
		return false
	}
	for k := range s {
		if *t[k] != *s[k] {
			return false
		}
	}
	return true
}
func addReqagent(data string) (*ReqAgent) {
	reg := NewReqAgent()
	json.Unmarshal([]byte(data), reg.Payload())
	return reg

}
