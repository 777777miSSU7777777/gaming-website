package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/777777miSSU7777777/gaming-website/internal/api/userapi"
	"github.com/777777miSSU7777777/gaming-website/pkg/repository/userrepository"
	"github.com/777777miSSU7777777/gaming-website/pkg/services/userservice"
)

var (
	host   string
	port   string
	dbhost string
	dbuser string
	dbpass string
	dbname string
)

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "Defines host ip")
	flag.StringVar(&host, "h", "0.0.0.0", "Defines host ip")
	flag.StringVar(&port, "port", "8080", "Defines host port")
	flag.StringVar(&port, "p", "8080", "Defines host port")
	flag.StringVar(&dbhost, "db_host", "mysql", "Defines db host")
	flag.StringVar(&dbuser, "db_user", "root", "Defines db user")
	flag.StringVar(&dbpass, "db_pass", "", "Defines db user's password")
	flag.StringVar(&dbname, "db_name", "GAMING_WEBSITE", "Defines db name")
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
	db, err := sql.Open(dbhost, conString)
	if err != nil {
		logger.Fatalln(err)
	}
	defer db.Close()

	userRepo := userrepository.New(db)
	userSvc := userservice.WrapLoggingMiddleware(userservice.New(userRepo), logger)
	ctx := context.Background()
	addr := fmt.Sprintf("%s:%s", host, port)

	handler := userapi.NewHttpServer(ctx, userSvc, logger)
	logger.Infof("Server started on %s", addr)
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		logger.Fatalln(err)
	}
}
