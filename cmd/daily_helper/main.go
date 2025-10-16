package main

import (
	"fmt"

	"github.com/deadman360/daily_helper/internal/infrastructure"
	"github.com/deadman360/daily_helper/internal/infrastructure/config/db"
	"github.com/deadman360/daily_helper/internal/interface/cli"
)

func main() {
	db, err := db.ConnectDatabase("./file")
	if err != nil {
		fmt.Println(err)
	}

	repository := infrastructure.NewRepository(db)
	cli.Execute(repository)
}
