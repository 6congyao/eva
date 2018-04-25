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

import (
	"eva/manager"
	"eva/matcher"
	"eva/policy"
	"eva/utils"
	"eva/agent"
)

// Eva is responsible for deciding if principal p can perform action a on resource r with condition c.
type Eva interface {
	// Authorize returns nil if principal p can perform action a on resource r with condition c or an error otherwise.
	//  if err := guard.Authorize(&Request{Resource: "article/1234", Action: "update", Principal: "peter"}); err != nil {
	//    return errors.New("Not allowed")
	//  }
	Authorize(rcs []*agent.RequestContext, keys []string) error
}

// Eva instance Eva00 inspired by "Ayanami Rei" :P.
type Eva00 struct {
	Manager manager.Manager
	matcher matcher.Matcher
}

func (e *Eva00) Matcher() matcher.Matcher {
	if e.matcher == nil {
		e.matcher = matcher.DefaultMatcher
	}
	return e.matcher
}

// Authorize returns nil if principal p can perform action a on resource r with condition c or an error otherwise.
func (e *Eva00) Authorize(rcs []*agent.RequestContext, keys []string) error {
	policies, err := e.Manager.FindCandidates(keys)
	if err != nil {
		return err
	}

	// Although the manager is responsible of matching the policies, it might decide to just scan for
	// principal, it might return all policies, or it might have a different pattern matching.
	// Thus, we need to make sure that we actually matched the right policies.
	return e.Evaluate(rcs, policies)
}

// Evaluate returns nil if principal p has permission p on resource r with condition c for a given policy list or an error otherwise.
// The Authorize interface should be preferred since it uses the manager directly. This is a lower level interface for when you don't want to use the eva manager.
func (e *Eva00) Evaluate(rcs []*agent.RequestContext, policies policy.Policies) error {
	var deciders = policy.Policies{}

	// Iterate through all RequestContexts
	for _, r := range rcs {
		var allowed = false
		// Iterate through all policies
		for _, p := range policies {
			for _, s := range p.GetStatements() {
				// Does the action match with one of the statements?
				// This is the first check because usually actions are a superset of get|update|delete|set
				// and thus match faster.
				if am, err := e.Matcher().Matches(p, s.GetActions(), r.Action); err != nil {
					return err
				} else if !am {
					// no, continue to next statement
					continue
				}

				// Does the principal match with one of the statements?
				// There are usually less principals than resources which is why this is checked
				// before checking for resources.
				// Principal is optionally match in entity-based policy.
				if principals := s.GetPrincipals(); len(principals) > 0 {
					if pm, err := e.Matcher().Matches(p, principals, r.Principal); err != nil {
						return err
					} else if !pm {
						// no, continue to next statement
						continue
					}
				}

				// Does the resource match with one of the statements?
				// Resource is optionally match in resource-based policy.
				if resources := s.GetResources(); len(resources) > 0 {
					if rm, err := e.Matcher().Matches(p, s.GetResources(), r.Resource); err != nil {
						return err
					} else if !rm {
						// no, continue to next policy
						continue
					}
				}

				// todo conditions
				// Are the policies conditions met?
				// This is checked first because it usually has a small complexity.

				// Is the policies effect deny? If yes, this overrides all allow policies -> access denied.
				if !s.AllowAccess() {
					deciders = append(deciders, p)
					//l.auditLogger().LogRejectedAccessRequest(r, policies, deciders)
					return utils.NewErrExplicitlyDenied()
				}

				allowed = true
				deciders = append(deciders, p)
			}

		}

		// Request was denied by default
		if !allowed {
			//l.auditLogger().LogRejectedAccessRequest(r, policies, deciders)
			return utils.NewErrDefaultDenied()
		}
	}

	return nil
}
