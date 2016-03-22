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
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/clashr/go-servr/models"
)

func AttemptIndex(w http.ResponseWriter, r *http.Request) {
	attempts := new(models.Attempts)
	if err := db.Select(attempts, "SELECT * FROM attempts"); err != nil {
		log.Fatalln(err)
	}

	w.Header().Set(ct, headerJSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(attempts); err != nil {
		log.Fatalln(err)
	}
}

func AttemptShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attemptId := vars["attemptId"]
	challengeId := vars["challengeId"]
	attempt := new(models.Attempt)
	err := db.Get(attempt, "SELECT * FROM attempts WHERE id=$1 AND challenge_id=$2", attemptId, challengeId)
	if err != nil {
		if err == sql.ErrNoRows {
			w.Header().Set(ct, headerPLAIN)
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			log.Println(err)
			w.Header().Set(ct, headerJSON)
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				log.Fatalln(err)
			}
			return
		}
	}

	w.Header().Set(ct, headerJSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(attempt); err != nil {
		log.Fatalln(err)
	}
}

func AttemptCreate(w http.ResponseWriter, r *http.Request) {
	var attempt models.Attempt
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(body, &attempt); err != nil {
		w.Header().Set(ct, headerJSON)
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln(err)
		}
		return
	}

	stmt, err := db.PrepareNamed("INSERT INTO attempts(username, source, language, dialect, challenge_id) VALUES (:username, :source, :language, :dialect, :challenge_id)")
	if err != nil {
		log.Fatalln(err)
	}
	res, err := stmt.Exec(attempt)
	if err != nil {
		log.Fatalln(err)
	}

	target := fmt.Sprintf("/challenges/%d/attempts/%d", attempt.ChallengeId, res.LastInsertId)
	http.Redirect(w, r, target, http.StatusSeeOther)
}

func AttemptUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attemptId := vars["attemptId"]
	challengeId := vars["challengeId"]
	fmt.Fprintf(w, "Attempt Update: %d\n", attemptId)
	fmt.Fprintf(w, "For Challenge: %d\n", challengeId)
	fmt.Fprintln(w, "Feature not yet Implemented.")
}

func AttemptDestroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attemptId := vars["attemptId"]
	challengeId := vars["challengeId"]
	fmt.Fprintf(w, "Attempt Destroy: %d\n", attemptId)
	fmt.Fprintf(w, "For Challenge: %d\n", challengeId)
	fmt.Fprintln(w, "Feature not yet Implemented.")
}
