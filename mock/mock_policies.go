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
var Policies = []policy.Policy{
	&policy.DefaultPolicy{
		ID:          "1",
		Description: `This policy allows entity to put and get on qstor service at path /foo/*.`,
		Version:     "2018-4-18",
		Name:        "foo",
		Urn:         "qrn:qws:iam::12321:policy/foo",
		Statements: []policy.DefaultStatement{
			{
				Principals: []string{"*"},
				Effect:     policy.AllowAccess,
				Actions:    []string{"qstor:GetObject"},
				Resources:  []string{"qrn:qcs:qstor:::foo/*"},
			},
			{
				Principals: []string{"*"},
				Effect:     policy.AllowAccess,
				Actions:    []string{"qstor:PutObject"},
				Resources:  []string{"qrn:qcs:qstor:::foo/*"},
			},
		},
	},
	&policy.DefaultPolicy{
		ID:          "1",
		Description: "This policy denies entity to perform api get* on qstor service at path /bar/*.",
		Version:     "2018-4-19",
		Name:        "bar",
		Urn:         "qrn:qws:iam::12321:policy/bar",
		Statements: []policy.DefaultStatement{
			{
				Principals: []string{"*"},
				Effect:     policy.AllowAccess,
				Actions:    []string{"qstor:Get*"},
				Resources:  []string{"qrn:qcs:qstor:::bar/*"},
			},
			{
				Principals: []string{"*"},
				Effect:     policy.DenyAccess,
				Actions:    []string{"qstor:Put*"},
				Resources:  []string{"qrn:qcs:qstor:::bar/*"},
			},
		},
	},
	&policy.DefaultPolicy{
		ID:          "2",
		Description: "This policy allows entity to do everything on IAM & STS service.",
		Version:     "2018-4-20",
		Name:        "max",
		Urn:         "qrn:qws:iam::12321:policy/max",
		Statements: []policy.DefaultStatement{
			{
				Principals: []string{"*"},
				Effect:     policy.AllowAccess,
				Actions:    []string{"iam:*"},
				Resources:  []string{"*"},
			},
			{
				Principals: []string{"*"},
				Effect:     policy.AllowAccess,
				Actions:    []string{"sts:*"},
				Resources:  []string{"*"},
			},
		},
	},
	&policy.DefaultPolicy{
		ID:          "2",
		Description: "This policy denies entity to perform some actions on IAM & STS service",
		Version:     "2018-4-21",
		Name:        "min",
		Urn:         "qrn:qws:iam::12321:policy/min",
		Statements: []policy.DefaultStatement{
			{
				Principals: []string{"*"},
				Effect:     policy.DenyAccess,
				Actions:    []string{"sts:AssumeRole", "sts:UpdateToken"},
				Resources:  []string{"*"},
			},
			{
				Principals: []string{"<.*>"},
				Effect:     policy.DenyAccess,
				Actions:    []string{"iam:Create*", "sts:Update*"},
				Resources:  []string{"*"},
			},
		},
	},
	&policy.DefaultPolicy{
		ID:          "3",
		Description: "This policy allows entity to describe ec2 service and run instances",
		Version:     "2018-4-22",
		Name:        "ec2op",
		Urn:         "qrn:qws:iam::12321:policy/ec2op",
		Statements: []policy.DefaultStatement{
			{
				//Principals: []string{"*"},
				Effect:     policy.AllowAccess,
				Actions:    []string{"ec2:describe*", "ec2:RunInstances"},
				Resources:  []string{"*"},
			},
		},
	},
}
