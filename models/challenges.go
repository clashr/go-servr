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
	Id      int       `json:"-" xorm:"pk autoincr not null"`
	Name    string    `json:"name" xorm:"unique"`
	Details string    `json:"details" xorm:"text"`
	Tests   []Test    `json:"-" xorm:"-"`
	Upload  time.Time `json:"upload" xorm:"created"`
}

type Challenges []Challenge

type Test struct {
	Id          int    `xorm: pk autoincr`
	Input       string `xorm: text`
	Output      string `xorm: text`
	ChallengeId int
}
