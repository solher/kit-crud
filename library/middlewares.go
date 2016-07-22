package library

import (
	"time"

	"golang.org/x/net/context"

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

func (mw serviceLoggingMiddleware) CreateDocument(ctx context.Context, userID string, document *Document) (new *Document, err error) {
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
	return mw.next.CreateDocument(ctx, userID, document)
}

func (mw serviceLoggingMiddleware) FindDocuments(ctx context.Context, userID string) (documents []Document, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "FindDocuments",
			"userID", userID,
			"foundDocuments", documents,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.FindDocuments(ctx, userID)
}

func (mw serviceLoggingMiddleware) FindDocumentsByID(ctx context.Context, userID string, ids []string) (documents []Document, err error) {
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
	return mw.next.FindDocumentsByID(ctx, userID, ids)
}

func (mw serviceLoggingMiddleware) ReplaceDocumentByID(ctx context.Context, userID string, id string, document *Document) (new *Document, err error) {
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
	return mw.next.ReplaceDocumentByID(ctx, userID, id, document)
}

func (mw serviceLoggingMiddleware) DeleteDocumentsByID(ctx context.Context, userID string, ids []string) (documents []Document, err error) {
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
	return mw.next.DeleteDocumentsByID(ctx, userID, ids)
}
