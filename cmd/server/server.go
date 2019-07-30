package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

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
	conString := fmt.Sprintf("%s:%s@/%s", dbuser, dbpass, dbname)
	fmt.Println(conString)
	db, err := sql.Open("mysql", conString)
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	userRepo := userrepository.New(db)

	userSvc := userservice.New(userRepo)

	newUserhandler := api.MakeNewUserHandler(userSvc)
	getUserHandler := api.MakeGetUserHandler(userSvc)
	deleteUserHandler := api.MakeDeleteUserHandler(userSvc)
	userTakeHandler := api.MakeUserTakeHandler(userSvc)
	userFundHandler := api.MakeUserFundHandler(userSvc)

	router := mux.NewRouter()

	router.Handle("/user", newUserhandler).Methods("POST")
	router.Handle("/user/{id}", getUserHandler).Methods("GET")
	router.Handle("/user/{id}", deleteUserHandler).Methods("DELETE")
	router.Handle("/user/{id}/take", userTakeHandler).Methods("POST")
	router.Handle("/user/{id}/fund", userFundHandler).Methods("POST")

	http.Handle("/", router)

	log.Println("Server started")

	addr := fmt.Sprintf("%s:%s", host, port)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
