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
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/clashr/go-servr/models"
)

func ChallengeIndex(w http.ResponseWriter, r *http.Request) {
	challenges := new(models.Challenges)
	if err := engine.Find(challenges); err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(challenges); err != nil {
		log.Fatalln(err)
	}
}

func ChallengeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["challengeId"]
	challenge := new(models.Challenge)
	exists, err := engine.Id(challengeId).Get(challenge)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln(err)
		}
		return
	}
	if !exists {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(challenge); err != nil {
		log.Fatalln(err)
	}
}

func ChallengeCreate(w http.ResponseWriter, r *http.Request) {
	var challenge models.Challenge
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(body, &challenge); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln(err)
		}
		return
	}

	affected, err := engine.Insert(&challenge)
	if err != nil {
		log.Fatalln(err)
	}

	target := fmt.Sprintf("/challenges/%d", affected)
	http.Redirect(w, r, target, http.StatusSeeOther)
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
