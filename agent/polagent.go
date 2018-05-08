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
	"encoding/json"
	"eva/policy"

	"eva/utils"
)

// PolicyInput handles the json text for policy part.
type PolicyInput struct {
	Version     string           `json:"version,omitempty"`
	Description string           `json:"description,omitempty"`
	Statements  []StatementInput `json:"statement,omitempty" binding:"required"`
}

// StatementInput handles the json text for statement part.
type StatementInput struct {
	Principals interface{} `json:"principal,omitempty"`
	Effect     string      `json:"effect,omitempty" binding:"required"`
	Actions    interface{} `json:"action,omitempty" binding:"required"`
	Resources  interface{} `json:"resource,omitempty"`

	//todo: Conditions  Conditions `json:"conditions,omitempty"`
}

type PolAgent struct {
	Policies []string
}

func NewPolAgent(ps []string) *PolAgent {
	if ps != nil {
		return &PolAgent{Policies: ps}
	}
	return &PolAgent{Policies: []string{}}
}

func (pa PolAgent) NormalizePolicies() (policy.Policies, error) {
	var policies policy.Policies = nil

	for _, p := range pa.Policies {
		pi := &PolicyInput{}
		if err := json.Unmarshal([]byte(p), pi); err != nil {
			return nil, err
		}

		var statements []policy.DefaultStatement = nil

		for _, s := range pi.Statements {
			statement := &policy.DefaultStatement{
				Principals: utils.ItoS(s.Principals),
				Effect:     s.Effect,
				Actions:    utils.ItoS(s.Actions),
				Resources:  utils.ItoS(s.Resources),
			}

			statements = append(statements, *statement)
		}

		policy := &policy.DefaultPolicy{
			Version:     pi.Version,
			Description: pi.Description,
			Statements:  statements,
		}

		policies = append(policies, policy)
	}

	return policies, nil
}

func (pa *PolAgent) Payload() interface{} {
	return &pa.Policies
}
