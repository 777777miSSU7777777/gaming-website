package userservice

import (
	"github.com/777777miSSU7777777/gaming-website/pkg/entity"

	log "github.com/sirupsen/logrus"
)

type loggingMiddleware struct {
	logger *log.Logger
	next   UserService
}

func WrapLoggingMiddleware(svc UserService, logger *log.Logger) loggingMiddleware {
	return loggingMiddleware{logger, svc}
}

func (mw loggingMiddleware) NewUser(username string, balance int64) (user entity.User, err error) {
	defer func() {
		mw.logger.WithFields(log.Fields{
			"name":    username,
			"balance": balance,
			"user":    user,
			"error":   err,
		}).Infoln()
	}()

	user, err = mw.next.NewUser(username, balance)
	return
}

func (mw loggingMiddleware) GetUser(id int64) (user entity.User, err error) {
	defer func() {
		mw.logger.WithFields(log.Fields{
			"id":    id,
			"user":  user,
			"error": err,
		}).Infoln()
	}()

	user, err = mw.next.GetUser(id)
	return
}

func (mw loggingMiddleware) DeleteUser(id int64) (err error) {
	defer func() {
		mw.logger.WithFields(log.Fields{
			"id":    id,
			"error": err,
		}).Infoln()
	}()

	err = mw.next.DeleteUser(id)
	return
}

func (mw loggingMiddleware) UserTake(id int64, points int64) (user entity.User, err error) {
	defer func() {
		mw.logger.WithFields(log.Fields{
			"id":     id,
			"points": points,
			"user":   user,
			"error":  err,
		}).Infoln()
	}()

	user, err = mw.next.UserTake(id, points)
	return
}

func (mw loggingMiddleware) UserFund(id int64, points int64) (user entity.User, err error) {
	defer func() {
		mw.logger.WithFields(log.Fields{
			"id":     id,
			"points": id,
			"user":   user,
			"error":  err,
		}).Infoln()
	}()

	user, err = mw.next.UserFund(id, points)
	return
}
