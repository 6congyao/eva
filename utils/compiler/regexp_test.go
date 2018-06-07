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

package compiler

import (
	"regexp"
	"testing"
)

const checkPass = "\u2713"
const checkFail = "\u2717"
func TestRegexCompiler(t *testing.T) {
//	t.Log("Test Regex Compiler")
	for k, c := range []struct {
		template       string
		delimiterStart byte
		delimiterEnd   byte
		failCompile    bool
		matchAgainst   string
		failMatch      bool
	}{
		{"urn:foo:{.*}", '{', '}', false, "urn:foo:bar:baz", false},
		{"urn:foo.bar.com:{.*}", '{', '}', false, "urn:foo.bar.com:bar:baz", false},
		{"urn:foo.bar.com:{.*}", '{', '}', false, "urn:foo.com:bar:baz", true},
		{"urn:foo.bar.com:{.*}", '{', '}', false, "foobar", true},
		{"urn:foo.bar.com:{.{1,2}}", '{', '}', false, "urn:foo.bar.com:aa", false},
		{"urn:foo.bar.com:{.*{}", '{', '}', true, "", true},
		{"urn:foo:<.*>", '<', '>', false, "urn:foo:bar:baz", false},

		// Ignoring this case for now...
		//{"urn:foo.bar.com:{.*\\{}", '{', '}', false, "", true},
	} {
		k++
		result, err := CompileRegex(c.template, c.delimiterStart, c.delimiterEnd)
		//assert.Equal(t, c.failCompile, err != nil, "Case %d", k)
		if c.failCompile ==( err != nil) {
			if !c.failCompile{
				//if err is not nil,ok will be false
				ok, _ := regexp.MatchString(result.String(), c.matchAgainst)
				if ok{
					//t.Logf("case %d success compiled to %s %s",k,result.String(),checkPass)
				}

				if ok!=!c.failMatch {
					t.Errorf("\tError match! case %d template:%s,matchAgainst:%s,failMatch:%t %s",k,c.template,c.matchAgainst,c.failMatch,checkFail)
				}
			}else {
//				t.Logf("case %d fail compiled %s %s",k,c.template,checkPass)

			}
		}else {
			t.Errorf("\tError compile! case %d should comile %t,but %t %s",k,c.failCompile,err!=nil,checkFail)
		}


		//assert.Nil(t, err, "Case %d", k)
		//assert.Equal(t, !c.failMatch, ok, "Case %d", k)
	}
}