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

package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/clashr/go-servr/routes"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	r := routes.Router()
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	address := "localhost:" + port
	log.Fatal(http.ListenAndServe(address, r))
}
