package cmd

import (
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/app/config"
	expirecode "github.com/nadhiljanitra/nadhil-golang-training-beginner/expirepaymentcode"
)

func InitExpireCode() {
	// init postgres database
	db, err := config.InitPostgres()
	if err != nil {
		panic(err)
	}

	repo := expirecode.NewRepository(db)
	svc := expirecode.NewService(repo)

	job := expirecode.NewJob(svc)
	job.Execute()
}
