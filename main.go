package main

import (
	"fmt"

	infra "github.com/bbcyyb/goarch/infrastructure"
	repo "github.com/bbcyyb/goarch/repository"
	svc "github.com/bbcyyb/goarch/service"
)

func main() {
	fmt.Println("Hello World")
}

func init() {
	conf = infra.NewConfiguration("conf", "sys")
	conn = infra.NewOrm("dbDefault")
	infra.C.Set("infra_config", conf)
	infra.C.Set("infra_orm", conn)

	infra.C.Register("svc_account", svc.NewUserService)

	infra.C.Register("repo_user", repo.NewUserRepository)
}
