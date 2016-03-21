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

package models

import "time"

type Challenge struct {
	Id      int       `json:"-"`
	Name    string    `json:"name"`
	Details string    `json:"details"`
	Tests   []Test    `json:"-"`
	Upload  time.Time `json:"upload" db:"upload_date"`
}

type Challenges []Challenge

type Test struct {
	Id          int
	Input       string
	Output      string
	ChallengeId int
}
