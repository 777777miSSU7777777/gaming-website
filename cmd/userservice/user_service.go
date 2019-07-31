package main

import (
	"context"
	"database/sql"
	"flag"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/777777miSSU7777777/gaming-website/internal/api/userapi"
	"github.com/777777miSSU7777777/gaming-website/pkg/repository/userrepository"
	"github.com/777777miSSU7777777/gaming-website/pkg/services/userservice"
)

func main() {
	var (
		listenAddr       string
		connectionString string
	)

	flag.StringVar(&listenAddr, "listen_addr", ":8080", "Defines listen address")
	flag.StringVar(&connectionString, "connection_string", "", "Defines connection string for MySQL")
	flag.Parse()

	logger := log.New()
	jsonFormatter := &log.JSONFormatter{}
	jsonFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logger.SetFormatter(jsonFormatter)
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Fatalln(err)
	}

	userRepo := userrepository.New(db)
	userSvc := userservice.WrapLoggingMiddleware(userservice.New(userRepo), logger)
	ctx := context.Background()

	handler := userapi.NewHttpServer(ctx, userSvc, logger)
	logger.Infof("Server started on %s", listenAddr)
	err = http.ListenAndServe(listenAddr, handler)
	if err != nil {
		logger.Fatalln(err)
	}
}
