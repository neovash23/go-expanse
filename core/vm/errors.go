// Copyright 2014 The go-ethereum Authors && Copyright 2015 go-expanse Authors
// This file is part of the go-expanse library.
//
// The go-expanse library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-expanse library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-expanse library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"errors"
	"fmt"

	"github.com/expanse-project/go-expanse/params"
)

var OutOfGasError = errors.New("Out of gas")
var DepthError = fmt.Errorf("Max call depth exceeded (%d)", params.CallCreateDepth)

type StackError struct {
	req, has int
}

func StackErr(req, has int) StackError {
	return StackError{req, has}
}

func (self StackError) Error() string {
	return fmt.Sprintf("stack error! require %v, have %v", self.req, self.has)
}

func IsStackErr(err error) bool {
	_, ok := err.(StackError)
	return ok
}
