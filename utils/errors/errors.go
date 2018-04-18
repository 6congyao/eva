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

package errors

import (
	"net/http"
	"errors"
)

const DefaultDenied = "Default"
const ExplicitlyDenied = "Explicit"

type errorWithContext interface {

	// StatusCode returns the status code of this error.
	StatusCode() int

	// Causer returns the info that caused the error, if applicable.
	Causer() string

	Details() string
}

type ErrDefaultDenied struct {
	code   int
	causer string
	details string
	error
}

func NewErrDefaultDenied() error {
	return &ErrDefaultDenied{
		error:  errors.New("request was denied by default"),
		code:   http.StatusForbidden,
		causer: http.StatusText(http.StatusForbidden),
		details: DefaultDenied,
	}
}

// StatusCode returns the status code of this error.
func (edd *ErrDefaultDenied) StatusCode() int {
	return edd.code
}

// Causer returns the causer of this error.
func (edd *ErrDefaultDenied) Causer() string {
	return edd.causer
}

// Details returns details on the error, if applicable.
func (edd *ErrDefaultDenied) Details() string {
	return edd.details
}

type ErrExplicitlyDenied struct {
	code   int
	causer string
	details string
	error
}

func NewErrExplicitlyDenied() error {
	return &ErrExplicitlyDenied{
		error:  errors.New("request was explicitly denied"),
		code:   http.StatusForbidden,
		causer: http.StatusText(http.StatusForbidden),
		details: ExplicitlyDenied,
	}
}

// StatusCode returns the status code of this error.
func (eed *ErrExplicitlyDenied) StatusCode() int {
	return eed.code
}

// StatusCode returns the status code of this error.
func (eed *ErrExplicitlyDenied) Causer() string {
	return eed.causer
}

// Details returns details on the error, if applicable.
func (eed *ErrExplicitlyDenied) Details() string {
	return eed.details
}