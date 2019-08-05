package main

import (
	"database/sql"
	"flag"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/777777miSSU7777777/gaming-website/api"
	"github.com/777777miSSU7777777/gaming-website/repository"
	service "github.com/777777miSSU7777777/gaming-website/service"
)

func main() {
	var listenAddr string
	var connectionString string

	flag.StringVar(&listenAddr, "listen_addr", ":8080", "Defines listen address")
	flag.StringVar(&connectionString, "connection_string", "", "Defines connection string for MySQL")
	flag.Parse()

	logger := log.New()
	jsonFormatter := &log.JSONFormatter{}
	jsonFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logger.SetFormatter(jsonFormatter)
	logger.SetReportCaller(true)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Fatalln(err)
	}

	userRepo := repository.NewUserRepository(db)
	userSvc := service.New(userRepo)

	handler := api.NewHttpServer(userSvc, logger)
	logger.Infof("Server started on %s", listenAddr)
	err = http.ListenAndServe(listenAddr, handler)
	if err != nil {
		logger.Fatalln(err)
	}
}
