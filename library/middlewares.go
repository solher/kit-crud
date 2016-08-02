package library

import (
	stdopentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"
)

func ServiceTracingMiddleware(next Service) Service {
	return serviceTracingMiddleware{
		next: next,
	}
}

type serviceTracingMiddleware struct {
	next Service
}

func (mw serviceTracingMiddleware) CreateDocument(ctx context.Context, userID string, document *Document) (new *Document, err error) {
	if span := stdopentracing.SpanFromContext(ctx); span != nil {
		span.SetTag("in.userId", userID)
		span.SetTag("in.document", document)
		defer func() {
			if err != nil {
				span.SetTag("out.err", err.Error())
				TraceIfUnexpectedError(span, err)
				return
			}
			span.SetTag("out.document", new)
		}()
	}
	return mw.next.CreateDocument(ctx, userID, document)
}

func (mw serviceTracingMiddleware) FindDocuments(ctx context.Context, userID string) (documents []Document, err error) {
	if span := stdopentracing.SpanFromContext(ctx); span != nil {
		span.SetTag("in.userId", userID)
		defer func() {
			if err != nil {
				span.SetTag("out.err", err.Error())
				TraceIfUnexpectedError(span, err)
				return
			}
			span.SetTag("out.documents", documents)
		}()
	}
	return mw.next.FindDocuments(ctx, userID)
}

func (mw serviceTracingMiddleware) FindDocumentsByID(ctx context.Context, userID string, ids []string) (documents []Document, err error) {
	if span := stdopentracing.SpanFromContext(ctx); span != nil {
		span.SetTag("in.userId", userID)
		span.SetTag("in.ids", ids)
		defer func() {
			if err != nil {
				span.SetTag("out.err", err.Error())
				TraceIfUnexpectedError(span, err)
				return
			}
			span.SetTag("out.documents", documents)
		}()
	}
	return mw.next.FindDocumentsByID(ctx, userID, ids)
}

func (mw serviceTracingMiddleware) ReplaceDocumentByID(ctx context.Context, userID string, id string, document *Document) (new *Document, err error) {
	if span := stdopentracing.SpanFromContext(ctx); span != nil {
		span.SetTag("in.userId", userID)
		span.SetTag("in.id", id)
		span.SetTag("in.document", document)
		defer func() {
			if err != nil {
				span.SetTag("out.err", err.Error())
				TraceIfUnexpectedError(span, err)
				return
			}
			span.SetTag("out.document", new)
		}()
	}
	return mw.next.ReplaceDocumentByID(ctx, userID, id, document)
}

func (mw serviceTracingMiddleware) DeleteDocumentsByID(ctx context.Context, userID string, ids []string) (documents []Document, err error) {
	if span := stdopentracing.SpanFromContext(ctx); span != nil {
		span.SetTag("in.userId", userID)
		span.SetTag("in.ids", ids)
		defer func() {
			if err != nil {
				span.SetTag("out.err", err.Error())
				TraceIfUnexpectedError(span, err)
				return
			}
			span.SetTag("out.documents", documents)
		}()
	}
	return mw.next.DeleteDocumentsByID(ctx, userID, ids)
}
