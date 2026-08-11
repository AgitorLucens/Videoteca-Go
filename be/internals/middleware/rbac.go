package middleware

import (
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
    "fmt"
)

type Rbac struct {
	DB *gorm.DB
}

func NewRBAC(conn string ) (*Rbac,error) {
	
	db, err := gorm.Open(postgres.Open(conn),&gorm.Config{})
	if err!= nil{
		fmt.Print("Error while connecting to postgres. Err: ",err)
	}
	return &Rbac{
		DB: db,
	}, nil
}


