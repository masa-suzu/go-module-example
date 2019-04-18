package main

import (
	"github.com/masa-suzu/myservice/core/domain"
	"github.com/masa-suzu/myservice/core/infra"
)

func main() {
	db := infra.Db{}
	db.Save(domain.User{1, "a"})
	db.Save(domain.User{2, "b"})
	db.Save(domain.User{3, "c"})
}
