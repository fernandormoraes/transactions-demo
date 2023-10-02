package main

import (
	"net/http"

	api "github.com/fernandormoraes/transaction-demo/internal/api"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/database"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := api.LoadConfig(".")

	if err != nil {
		logrus.Fatal(err)
	}

	db, err := database.OpenDatabase()

	if err != nil {
		logrus.Fatal(err)
	}

	server := api.NewServer(db, &cfg, http.DefaultClient)

	if err = server.Run(); err != nil {
		logrus.Fatal(err)
	}
}
