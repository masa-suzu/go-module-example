package main

import (
	"github.com/masa-suzu/go-module-example/core/domain"
	"github.com/masa-suzu/go-module-example/core/infra"
)

func main() {
	db := infra.Db{}
	db.Save(domain.User{1, "a"})
	db.Save(domain.User{2, "b"})
	db.Save(domain.User{3, "c"})
}
