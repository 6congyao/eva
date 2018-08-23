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

package policy

import (
	"strings"
)

const (
	// AllowAccess should be used as effect for policies that allow access.
	AllowAccess = "allow"
	// DenyAccess should be used as effect for policies that deny access.
	DenyAccess = "deny"
)

// Policies is an array of policies.
type Policies []Policy

// Policy represent a policy model.
type Policy interface {
	// GetID returns the policy id.
	GetID() string

	// GetVersion returns the policy version.
	GetVersion() string

	// GetName returns the policy name.
	GetName() string

	// GetUrn returns the policy urn.
	GetUrn() string

	// GetDescription returns the policy description.
	GetDescription() string

	// GetStatement returns the policy statements.
	GetStatements() Statements

	// GetStartDelimiter returns the delimiter which identifies the beginning of a regular expression.
	GetStartDelimiter() byte

	// GetEndDelimiter returns the delimiter which identifies the end of a regular expression.
	GetEndDelimiter() byte

	// Support '*' wildcard
	ReplaceWildcard(string) string
}

// Statements is an array of Statement.
type Statements []Statement

type Statement interface {
	// GetPrincipal returns the policies principals.
	GetPrincipals() []string

	// GetEffect returns the policies effect which might be 'allow' or 'deny'.
	GetEffect() string

	// GetActions returns the policies actions.
	GetActions() []string

	// GetResources returns the policies resources.
	GetResources() []string

	// AllowAccess returns true if the effect is allow, otherwise false.
	AllowAccess() bool

	// todo: GetConditions returns the policies conditions.
	//GetConditions() Conditions
}

// DefaultPolicy is the default implementation of the policy interface.
type DefaultPolicy struct {
	ID          string             `json:"id,omitempty"`
	Version     string             `json:"version,omitempty"`
	Name        string             `json:"name,omitempty"`
	Urn         string             `json:"urn,omitempty"`
	Description string             `json:"description,omitempty"`
	Statements  []DefaultStatement `json:"statement,omitempty" binding:"required"`
}

// GetID returns the policy id.
func (p *DefaultPolicy) GetID() string {
	return p.ID
}

// GetVersion returns the policy version.
func (p *DefaultPolicy) GetVersion() string {
	return p.Version
}

// GetName returns the policy name.
func (p *DefaultPolicy) GetName() string {
	return p.Name
}

// GetUrn returns the policy urn.
func (p *DefaultPolicy) GetUrn() string {
	return p.Urn
}

// GetDescription returns the policy description.
func (p *DefaultPolicy) GetDescription() string {
	return p.Description
}

// GetStatements returns the policy Statements.
func (p *DefaultPolicy) GetStatements() Statements {
	i := Statements{}
	for n := range p.Statements {
		i = append(i, &p.Statements[n])
	}

	return i
}

func (p *DefaultPolicy) GetStartDelimiter() byte {
	return '<'
}

func (p *DefaultPolicy) GetEndDelimiter() byte {
	return '>'
}

func (p *DefaultPolicy) ReplaceWildcard(s string) string {
	if strings.Count(s, "<.*>") == 0 {
		return strings.Replace(s, "*", "<.*>", -1)
	}

	return s
}

type DefaultStatement struct {
	Principals []string `json:"principal,omitempty"`
	Effect     string   `json:"effect,omitempty" binding:"required"`
	Actions    []string `json:"action,omitempty" binding:"required"`
	Resources  []string `json:"resource,omitempty"`

	//todo: Conditions  Conditions `json:"conditions,omitempty"`
}

// GetPrincipals returns the policies principals.
func (s *DefaultStatement) GetPrincipals() []string {
	return s.Principals
}

// GetEffect returns the policies effect which might be 'allow' or 'deny'.
func (s *DefaultStatement) GetEffect() string {
	return s.Effect
}

// GetActions returns the policies actions.
func (s *DefaultStatement) GetActions() []string {
	return s.Actions
}

// GetResources returns the policies resources.
func (s *DefaultStatement) GetResources() []string {
	return s.Resources
}

// AllowAccess returns true if the policy effect is allow, otherwise false.
func (s *DefaultStatement) AllowAccess() bool {
	return strings.ToLower(s.Effect) == AllowAccess
}
