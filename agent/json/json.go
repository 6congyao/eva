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

package json

import (
	"eva/policy"
	"eva/agent"
	"eva/utils"
)

type JsonAgent struct {
	RequestInput agent.AuthRequestInput
}

func NewJsonAgent() *JsonAgent{
	return &JsonAgent{}
}

func (ja *JsonAgent)NormalizeRequests() ([]string, []*agent.RequestContext, error) {
	keys, _ := utils.ItoS(ja.RequestInput.Subject)
	var rcs []*agent.RequestContext = nil
	for _, v := range ja.RequestInput.Payload {
		request := &agent.RequestContext{
			Principal: v.Principal,
			Action:    v.Action,
			Resource:  v.Resource,
		}
		rcs = append(rcs, request)
	}
	return keys, rcs, nil
}

func (ja *JsonAgent)NormalizePolicies() (policy.Policies, error) {
	return nil, nil
}

func (ja *JsonAgent)Payload() (interface{}) {
	return &ja.RequestInput
}
