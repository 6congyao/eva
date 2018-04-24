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
	"eva/policy"
)

// Binding from JSON
type AuthRequestInput struct {
	// Subject is the Query keys that is requesting access.
	// Support both string and []string
	Subject interface{}          `json:"subject" binding:"required"`
	Payload []AuthRequestPayload `json:"payload" binding:"required"`
}

// Binding from JSON
type AuthRequestPayload struct {
	// Action is the action that is requested on the resource.
	Action string `json:"action" binding:"required"`

	// Principal is the principal who sent the request.
	Principal string `json:"principal"`

	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// todo condition
}

// RequestContext is the expected request object.
type RequestContext struct {
	// Principal is the subject that is requesting access.
	Principal string `json:"principal,omitempty"`

	// Action is the action that is requested on the resource.
	Action string `json:"action,omitempty"`

	// Resource is the resource that access is requested to.
	Resource string `json:"resource,omitempty"`

	// todo:Condition is the request's environmental context.
	//Condition string `json:"condition,omitempty"`
}

type Agent interface {
	NormalizeRequests() ([]string, []*RequestContext, error)
	NormalizePolicies() (policy.Policies, error)
	Payload() interface{}
}
