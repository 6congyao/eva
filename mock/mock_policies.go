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

package mock

import "eva/policy"

// A bunch of exemplary policies
var Polices = []policy.Policy{
	&policy.DefaultPolicy{
		ID: "1",
		Description: `This policy allows max, peter, zac and ken to create, delete and get the listed resources,
			but only if the client ip matches and the request states that they are the owner of those resources as well.`,
		Version: "2018-4-23",
		Name:   "foo",
		Urn:    "qrn:qws:iam::12321:user/Lucas",
	},
	&policy.DefaultPolicy{
		ID:          "2",
		Description: "This policy allows max to put object to QingStor bucket 'max'",
		Version: "2018-4-23",
		Name:   "bar",
		Urn:    "qrn:qws:iam::12321:user/Lucas",
	},
	&policy.DefaultPolicy{
		ID:          "3",
		Description: "This policy denies max to put object to object name 'min' on bucket 'max'",
		Version: "2018-4-23",
		Name:   "hello",
		Urn:    "qrn:qws:iam::12321:user/Lucas",
	},
	&policy.DefaultPolicy{
		ID:          "4",
		Description: "This policy allows lucas to perform actions STS:* on any of the resources",
		Version: "2018-4-23",
		Name:   "world",
		Urn:    "qrn:qws:iam::12321:user/Lucas",
	},
	&policy.DefaultPolicy{
		ID:          "5",
		Description: "This policy revoke all the access requests from lucas",
		Version: "2018-4-23",
		Name:   "share",
		Urn:    "qrn:qws:iam::12321:user/Lucas",
	},
}
