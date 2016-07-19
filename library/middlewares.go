package library

import (
	"time"

	"github.com/go-kit/kit/log"
)

type Middleware func(Service) Service

func ServiceLoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return serviceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type serviceLoggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw serviceLoggingMiddleware) CreateDocument(userID string, document *Document) (new *Document, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "CreateDocument",
			"userID", userID,
			"document", document,
			"newDocument", new,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.CreateDocument(userID, document)
}

func (mw serviceLoggingMiddleware) FindDocuments(userID string) (documents []Document, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "FindDocuments",
			"userID", userID,
			"foundDocuments", documents,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.FindDocuments(userID)
}

func (mw serviceLoggingMiddleware) FindDocumentsByID(userID string, ids []string) (documents []Document, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "FindDocumentsByID",
			"userID", userID,
			"ids", ids,
			"foundDocuments", documents,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.FindDocumentsByID(userID, ids)
}

func (mw serviceLoggingMiddleware) ReplaceDocumentByID(userID string, id string, document *Document) (new *Document, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "ReplaceDocumentByID",
			"userID", userID,
			"id", id,
			"document", document,
			"newDocument", new,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.ReplaceDocumentByID(userID, id, document)
}

func (mw serviceLoggingMiddleware) DeleteDocumentsByID(userID string, ids []string) (documents []Document, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "DeleteDocumentsByID",
			"userID", userID,
			"ids", ids,
			"deletedDocuments", documents,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.DeleteDocumentsByID(userID, ids)
}
