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

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/clashr/go-servr/models"
)

func ChallengeIndex(w http.ResponseWriter, r *http.Request) {
	challenges := models.Challenges{
		models.Challenge{
			Name: "Unnamed",
		},
		models.Challenge{
			Name: "The Second Unnamed",
		},
	}

	if err := json.NewEncoder(w).Encode(challenges); err != nil {
		log.Fatalln(err)
	}
}

func ChallengeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["challengeId"]
	fmt.Fprintf(w, "Challenge Show:", challengeId)
}

func ChallengeCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Yet Implemented")
}
func ChallengeUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["challengeId"]
	fmt.Fprintf(w, "Challenge Update:", challengeId)
}

func ChallengeDestroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["challengeId"]
	fmt.Fprintf(w, "Challenge Destroy:", challengeId)
}
