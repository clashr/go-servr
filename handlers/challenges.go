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
	"github.com/clashr/go-servr/modules"
)

func ChallengeIndex(w http.ResponseWriter, r *http.Request) {
	challenges := new(models.Challenges)
	if err := db.Select(challenges, "SELECT * FROM challenges"); err != nil {
		log.Fatalln(err)
	}

	w.Header().Set(ct, headerJSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(challenges); err != nil {
		log.Fatalln(err)
	}
}

func ChallengeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["challengeId"]
	challenge := new(models.Challenge)
	err := db.Get(challenge, "SELECT * FROM challenges WHERE id=$1", challengeId)
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
		w.Header().Set(ct, headerJSON)
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln(err)
		}
		return
	}

	stmt, err := db.PrepareNamed("INSERT INTO challenges(name, details) VALUES (:name, :details)")
	if err != nil {
		log.Fatalln(err)
	}
	res, err := stmt.Exec(challenge)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: challenges.name" {
			w.Header().Set(ct, headerJSON)
			w.WriteHeader(http.StatusConflict)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				log.Fatalln(err)
			}
			return
		} else {
			log.Fatalln(err)
		}
	}

	go modules.DoesLittle()

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	target := fmt.Sprintf("/challenges/%d", id)
	http.Redirect(w, r, target, http.StatusSeeOther)
}

func ChallengeUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["challengeId"]
	fmt.Fprintf(w, "Challenge Update: %s", challengeId)
	fmt.Fprintln(w, "Feature not yet Implemented.")
}

func ChallengeDestroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	challengeId := vars["challengeId"]
	fmt.Fprintf(w, "Challenge Destroy: %s", challengeId)
	fmt.Fprintln(w, "Feature not yet Implemented.")
}
