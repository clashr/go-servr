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
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"time"

	"github.com/clashr/buildrpcd/api"
)

func buildr(lang, dialect, src string) ([]byte, error) {
	//make connection to rpc server
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		return nil, fmt.Errorf("Error in dialing. %s", err)
	}
	//make arguments object
	args := &api.Args{
		Language: lang,
		Dialect:  dialect,
		Contents: src,
	}
	//this will store returned result
	var result api.Result
	//call remote procedure with args
	err = client.Call("Build.Compile", args, &result)
	if err != nil {
		return nil, fmt.Errorf("error in Build: %v", err)
	}
	//we got our result in result
	log.Printf("RPC Recieved: Language:%s\n, args.Language")
	log.Printf("RPC Recieved: Binary Length:%d\n", len(result.Binary))

	return result.Binary, nil
}
