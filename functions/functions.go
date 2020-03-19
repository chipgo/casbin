package functions

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/chipgo/casbin/util"
)

func LoadByCSV() {
	e, err := casbin.NewEnforcer("./config/casbin/model.conf", "./config/casbin/policy.csv")
	util.FailedOnError(err, "Failed to load policy")

	roles := e.GetAllRoles()
	log.Println("roles: ", roles)
}
