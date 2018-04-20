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
	"bytes"
	"fmt"
	"regexp"
)

// delimiterIndices returns the first level delimiter indices from a string.
// It returns an error in case of unbalanced delimiters.
func delimiterIndices(s string, delimiterStart, delimiterEnd byte) ([]int, error) {
	var level, idx int
	idxs := make([]int, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case delimiterStart:
			if level++; level == 1 {
				idx = i
			}
		case delimiterEnd:
			if level--; level == 0 {
				idxs = append(idxs, idx, i+1)
			} else if level < 0 {
				return nil, fmt.Errorf(`unbalanced braces in "%q"`, s)
			}
		}
	}

	if level != 0 {
		return nil, fmt.Errorf(`unbalanced braces in "%q"`, s)
	}

	return idxs, nil
}

// CompileRegex parses a template and returns a Regexp.
//
// You can define your own delimiters. It is e.g. common to use curly braces {} but I recommend using characters
// which have no special meaning in Regex, e.g.: <, >
//
//  reg, err := compiler.CompileRegex("foo:bar.baz:<[0-9]{2,10}>", '<', '>')
//  // if err != nil ...
//  reg.MatchString("foo:bar.baz:123")
func CompileRegex(tpl string, delimiterStart, delimiterEnd byte) (*regexp.Regexp, error) {
	// Check if it is well-formed.
	idxs, errBraces := delimiterIndices(tpl, delimiterStart, delimiterEnd)
	if errBraces != nil {
		return nil, errBraces
	}
	varsR := make([]*regexp.Regexp, len(idxs)/2)
	pattern := bytes.NewBufferString("")
	pattern.WriteByte('^')

	var end int
	var err error
	for i := 0; i < len(idxs); i += 2 {
		// Set all values we are interested in.
		raw := tpl[end:idxs[i]]
		end = idxs[i+1]
		patt := tpl[idxs[i]+1 : end-1]
		// Build the regexp pattern.
		varIdx := i / 2
		fmt.Fprintf(pattern, "%s(%s)", regexp.QuoteMeta(raw), patt)
		varsR[varIdx], err = regexp.Compile(fmt.Sprintf("^%s$", patt))
		if err != nil {
			return nil, err
		}
	}

	// Add the remaining.
	raw := tpl[end:]
	pattern.WriteString(regexp.QuoteMeta(raw))
	pattern.WriteByte('$')

	// Compile full regexp.
	reg, errCompile := regexp.Compile(pattern.String())
	if errCompile != nil {
		return nil, errCompile
	}

	return reg, nil
}
