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

// RequestContext is the expected request object.
type RequestContext struct {
	// Principal is the subject that is requesting access.
	Principal string `json:"principal"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Condition is the request's environmental context.
	Condition string `json:"condition"`
}

// Eva is responsible for deciding if principal p can perform action a on resource r with condition c.
type Eva interface {
	// IsAllowed returns nil if principal p can perform action a on resource r with condition c or an error otherwise.
	//  if err := guard.Authorize(&Request{Resource: "article/1234", Action: "update", Principal: "peter"}); err != nil {
	//    return errors.New("Not allowed")
	//  }
	Authorize(rc *[]RequestContext) error
}
