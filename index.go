package session_buntdb

import (
	cb "github.com/chefsgo/cache-buntdb"
	"github.com/chefsgo/chef"
)

func init() {
	chef.Register("file", cb.Driver("store/cache.db"))
}
