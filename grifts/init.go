package grifts

import (
	"buffalo_go_gplgen_user_registration/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
