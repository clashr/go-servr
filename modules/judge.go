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

	"github.com/clashr/judgrpcd/api"
)

func judgr() {
	//make connection to rpc server
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	//make arguments object
	bin, err := ioutil.ReadFile("a.out")

	tests := make([]api.Test, 1)
	tests[0] = api.Test{In: "asdf", Out: ""}

	args := &api.Args{Language: "c", Binary: bin, TestData: tests}
	//this will store returned result
	var result api.Result
	//call remote procedure with args
	if err = client.Call("Judge.Runner", args, &result); err != nil {
		log.Fatalf("Error in running: %s", err)
	}
	//Print out the result
	log.Printf("Mem Used: %dKB\nTime Used: %s\nScore: %d\n", result.Data[0].TotalMem,
		result.Data[0].TotalTime, result.Score)
}
