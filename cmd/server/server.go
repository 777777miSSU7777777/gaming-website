package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/777777miSSU7777777/gaming-website/internal/api"
	"github.com/777777miSSU7777777/gaming-website/pkg/repository/userrepository"
	"github.com/777777miSSU7777777/gaming-website/pkg/services/userservice"
)

var (
	host   string
	port   string
	dbuser string
	dbpass string
	dbname string
)

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "Defines host ip")
	flag.StringVar(&host, "h", "0.0.0.0", "Defines host ip")
	flag.StringVar(&port, "port", "8080", "Defines host port")
	flag.StringVar(&port, "p", "8080", "Defines host port")
	flag.StringVar(&dbuser, "user", "root", "Defines db user")
	flag.StringVar(&dbpass, "pass", "", "Defines db user's password")
	flag.StringVar(&dbname, "name", "GAMING_WEBSITE", "Defines db name")
	flag.Parse()
}

func main() {
	logger := log.New()
	jsonFormatter := &log.JSONFormatter{}
	jsonFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logger.SetFormatter(jsonFormatter)
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)

	conString := fmt.Sprintf("%s:%s@/%s", dbuser, dbpass, dbname)
	db, err := sql.Open("mysql", conString)
	if err != nil {
		logger.Fatalln(err)
	}
	defer db.Close()

	userRepo := userrepository.New(db)
	userSvc := userservice.WrapLoggingMiddleware(userservice.New(userRepo), logger)

	newUserhandler := api.MakeNewUserHandler(userSvc, logger)
	getUserHandler := api.MakeGetUserHandler(userSvc, logger)
	deleteUserHandler := api.MakeDeleteUserHandler(userSvc, logger)
	userTakeHandler := api.MakeUserTakeHandler(userSvc, logger)
	userFundHandler := api.MakeUserFundHandler(userSvc, logger)

	router := mux.NewRouter()
	router.Handle("/user", newUserhandler).Methods("POST")
	router.Handle("/user/{id}", getUserHandler).Methods("GET")
	router.Handle("/user/{id}", deleteUserHandler).Methods("DELETE")
	router.Handle("/user/{id}/take", userTakeHandler).Methods("POST")
	router.Handle("/user/{id}/fund", userFundHandler).Methods("POST")

	http.Handle("/", router)
	logger.Infoln("Server started")
	addr := fmt.Sprintf("%s:%s", host, port)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Fatalln(err)
	}
}
