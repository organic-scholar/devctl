package pages

import (
	"io/fs"

	"github.com/organic-scholar/devctl/common"
)

func Router(f fs.FS) *common.PageRouter {
	r := common.NewPageRouter(f)
	return r
}
