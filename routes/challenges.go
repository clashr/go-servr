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
package routes

import "github.com/clashr/go-servr/handlers"

var challengeRoutes = Routes{
	Route{
		"challengeIndex",
		"GET",
		"/challenges",
		handlers.ChallengeIndex,
	},
	Route{
		"challengeShow",
		"GET",
		"/challenges/{challengeId}",
		handlers.ChallengeShow,
	},
	Route{
		"challengeCreate",
		"POST",
		"/challenges",
		handlers.ChallengeCreate,
	},
	Route{
		"challengeUpdate",
		"PATCH",
		"/challenges/{challengeId}",
		handlers.ChallengeUpdate,
	},
	Route{
		"challengeDestroy",
		"DELETE",
		"/challenges/{challengeId}",
		handlers.ChallengeDestroy,
	},
}
