package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"buffalo_go_gplgen_user_registration/graph/model"
)

type Resolver struct {
	UserStore map[string]model.User
}
