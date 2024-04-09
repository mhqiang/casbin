package casbin

import (
	"log"

	casbinused "github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbinused.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "dajun", "user/dajun/1", "read")
	check(e, "lizi", "user/lizi/2", "read")
	check(e, "dajun", "user/lizi/1", "read")
}