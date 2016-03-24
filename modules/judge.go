//
// Copyright (c) 2016 Dennis Chen
//
// This file is part of Clashr.
//
// Clashr is free software: you can redistribute it and/or modify it under the
// terms of the GNU Affero General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// Clashr is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License for
// more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Clashr.  If not, see <http://www.gnu.org/licenses/>.
//

package modules

import (
	"io/ioutil"
	"log"
	"net/rpc"

	"github.com/clashr/go-servr/models"
	"github.com/clashr/judgrpcd/api"
)

func judgr(lang string, bin []byte, tests models.Tests) (int, error) {
	//make connection to rpc server
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		return -1, fmt.Errorf("Error in dialing. %s", err)
	}
	testdata := make([]api.Test, len(tests))
	for i, test := range tests {
		testdata[i] = api.Test{
			In: test.Input,
			Out: test.Output,
		}
	//make arguments object
	args := &api.Args{
		Language: lang,
		Binary: bin,
		TestData: tests,
	}

	//this will store returned result
	var result api.Result
	//call remote procedure with args
	if err = client.Call("Judge.Runner", args, &result); err != nil {
		return -1, fmt.Errorf("Error in running: %s", err)
	}
	//return the result
	return result.Score, nil
}
