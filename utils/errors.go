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

package utils

import (
	"errors"
	"net/http"
)

const DefaultDenied = "Default"
const ExplicitlyDenied = "Explicit"

type ErrorWithContext interface {

	// StatusCode returns the status code of this error.
	Code() int

	// Causer returns the info that caused the error, if applicable.
	Causer() string

	// Details returns deny type.
	Details() string

	// Source returns the request source.
	Source() interface{}

	// Decider returns the specific rule which take effect.
	Decider() interface{}

	// Error returns the messages.
	Error() string
}

type ErrDefaultDenied struct {
	code    int
	causer  string
	details string
	source  interface{}
	decider interface{}
	error
}

func NewErrDefaultDenied(source interface{}) error {
	return &ErrDefaultDenied{
		error:   errors.New("request was denied by default (no matching statements)"),
		code:    http.StatusForbidden,
		causer:  DefaultDenied,
		details: http.StatusText(http.StatusForbidden),
		source: source,
	}
}

// StatusCode returns the status code of this error.
func (edd ErrDefaultDenied) Code() int {
	return edd.code
}

// Causer returns the info that caused the error, if applicable.
func (edd ErrDefaultDenied) Causer() string {
	return edd.causer
}

// Details returns deny type.
func (edd ErrDefaultDenied) Details() string {
	return edd.details
}

// Source returns the request source.
func (edd ErrDefaultDenied) Source() interface{} {
	return edd.source
}

// Decider returns the specific rule which take effect.
func (edd ErrDefaultDenied) Decider() interface{} {
	return edd.decider
}

type ErrExplicitlyDenied struct {
	code    int
	causer  string
	details string
	source  interface{}
	decider interface{}
	error
}

func NewErrExplicitlyDenied(source interface{}, decider interface{}) error {
	return &ErrExplicitlyDenied{
		error:   errors.New("request was explicitly denied"),
		code:    http.StatusForbidden,
		causer:  ExplicitlyDenied,
		details: http.StatusText(http.StatusForbidden),
		source: source,
		decider: decider,
	}
}

// StatusCode returns the status code of this error.
func (eed ErrExplicitlyDenied) Code() int {
	return eed.code
}

// Causer returns the info that caused the error, if applicable.
func (eed ErrExplicitlyDenied) Causer() string {
	return eed.causer
}

// Details returns deny type.
func (eed ErrExplicitlyDenied) Details() string {
	return eed.details
}

// Source returns the request source.
func (eed ErrExplicitlyDenied) Source() interface{} {
	return eed.source
}

// Decider returns the specific rule which take effect.
func (eed ErrExplicitlyDenied) Decider() interface{} {
	return eed.decider
}
