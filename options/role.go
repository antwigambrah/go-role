package options

import (
	"github.com/antwigambrah/go-role/utils"
)

// RoleOption represents options when fetching roles.
type RoleOption struct {
	WithPermissions bool
	Pagination      *utils.Pagination
}
