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
	"net/rpc"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/clashr/buildrpcd/api"
)

func buildr() {
	//make connection to rpc server
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	//make arguments object
	args := &api.Args{
		Language: "c",
		Dialect:  "ansi",
		Contents: "#include<stdio.h>\nint main() {\nprintf(\"Hello World\");\nreturn 0; }",
	}
	//this will store returned result
	var result api.Result
	//call remote procedure with args
	log.Printf("%s ", args.Language)
	err = client.Call("Build.Compile", args, &result)
	if err != nil {
		log.Fatalf("error in Build: %v", err)
	}
	//we got our result in result
	log.Printf("Language:%s\n %s\nResult: %d", args.Language, args.Contents, len(result.Binary))

	ioutil.WriteFile("a.out", result.Binary, 0755)
	if err != nil {
		log.Fatalln("could not write result")
	}
	log.Println("Wrote Binary")
}

func DoesLittle() {
	time.Sleep(5 * time.Second)
	log.Println("Finished Does Little")
}
