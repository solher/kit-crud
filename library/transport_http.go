package library

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ansel1/merry"
	"github.com/codegangsta/negroni"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
	"github.com/solher/kit-crud/errs"

	"golang.org/x/net/context"
)

func MakeHTTPHandler(ctx context.Context, s Service) http.Handler {
	opts := []httptransport.ServerOption{}

	createDocumentHandler := httptransport.NewServer(
		ctx,
		makeCreateDocumentEndpoint(s),
		decodeHTTPCreateDocumentRequest,
		encodeHTTPCreateDocumentResponse,
		opts...,
	)
	findDocumentsHandler := httptransport.NewServer(
		ctx,
		makeFindDocumentsEndpoint(s),
		decodeHTTPFindDocumentsRequest,
		encodeHTTPFindDocumentsResponse,
		opts...,
	)
	findDocumentsByIDHandler := httptransport.NewServer(
		ctx,
		makeFindDocumentsByIDEndpoint(s),
		decodeHTTPFindDocumentsByIDRequest,
		encodeHTTPFindDocumentsByIDResponse,
		opts...,
	)
	replaceDocumentByIDHandler := httptransport.NewServer(
		ctx,
		makeReplaceDocumentByIDEndpoint(s),
		decodeHTTPReplaceDocumentByIDRequest,
		encodeHTTPReplaceDocumentByIDResponse,
		opts...,
	)
	deleteDocumentsByIDHandler := httptransport.NewServer(
		ctx,
		makeDeleteDocumentsByIDEndpoint(s),
		decodeHTTPDeleteDocumentsByIDRequest,
		encodeHTTPDeleteDocumentsByIDResponse,
		opts...,
	)

	r := bone.New()
	r.SubRoute("/documents", func() bone.Router {
		r := bone.New()
		m := negroni.New()
		m.Use(negroni.HandlerFunc(gate))
		m.UseHandler(r)

		r.Post("", createDocumentHandler)
		r.Get("", findDocumentsHandler)
		r.Get("/:ids", findDocumentsByIDHandler)
		r.Put("/:id", replaceDocumentByIDHandler)
		r.Delete("/:ids", deleteDocumentsByIDHandler)

		return m
	}())

	return r
}

func gate(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(rw, r)
}

func decodeHTTPCreateDocumentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var document *Document
	if err := json.NewDecoder(r.Body).Decode(document); err != nil {
		return nil, err
	}
	req := createDocumentRequest{
		UserID:   "admin",
		Document: document,
	}
	return req, nil
}

func encodeHTTPCreateDocumentResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(createDocumentResponse)
	if res.Err != nil {
		return encodeHTTPError(ctx, w, res.Err)
	}
	return encodeHTTPResponse(ctx, w, http.StatusCreated, res.Document)
}

func decodeHTTPFindDocumentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := findDocumentsRequest{
		UserID: "admin",
	}
	return req, nil
}

func encodeHTTPFindDocumentsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(findDocumentsResponse)
	if res.Err != nil {
		return encodeHTTPError(ctx, w, res.Err)
	}
	return encodeHTTPResponse(ctx, w, http.StatusOK, res.Documents)
}

func decodeHTTPFindDocumentsByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	val := bone.GetValue(r, ":ids")
	ids := strings.Split(val, ",")
	req := findDocumentsByIDRequest{
		UserID: "admin",
		IDs:    ids,
	}
	return req, nil
}

func encodeHTTPFindDocumentsByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(findDocumentsByIDResponse)
	if res.Err != nil {
		return encodeHTTPError(ctx, w, res.Err)
	}
	if len(res.Documents) == 1 {
		return encodeHTTPResponse(ctx, w, http.StatusOK, res.Documents[0])
	}
	return encodeHTTPResponse(ctx, w, http.StatusOK, res.Documents)
}

func decodeHTTPReplaceDocumentByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var document *Document
	if err := json.NewDecoder(r.Body).Decode(document); err != nil {
		return nil, err
	}
	req := replaceDocumentByIDRequest{
		UserID:   "admin",
		ID:       bone.GetValue(r, ":id"),
		Document: document,
	}
	return req, nil
}

func encodeHTTPReplaceDocumentByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(replaceDocumentByIDResponse)
	if res.Err != nil {
		return encodeHTTPError(ctx, w, res.Err)
	}
	return encodeHTTPResponse(ctx, w, http.StatusOK, res.Document)
}

func decodeHTTPDeleteDocumentsByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	val := bone.GetValue(r, ":ids")
	ids := strings.Split(val, ",")
	req := deleteDocumentsByIDRequest{
		UserID: "admin",
		IDs:    ids,
	}
	return req, nil
}

func encodeHTTPDeleteDocumentsByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(deleteDocumentsByIDResponse)
	if res.Err != nil {
		return encodeHTTPError(ctx, w, res.Err)
	}
	if len(res.Documents) == 1 {
		return encodeHTTPResponse(ctx, w, http.StatusOK, res.Documents[0])
	}
	return encodeHTTPResponse(ctx, w, http.StatusOK, res.Documents)
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, code int, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(response)
}

func encodeHTTPError(_ context.Context, w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var apiErr errs.APIError
	switch {
	case merry.Is(err, errs.NotFound):
		w.WriteHeader(http.StatusForbidden)
		apiErr = errs.APIForbidden
	default:
		w.WriteHeader(http.StatusInternalServerError)
		apiErr = errs.APIInternal
	}
	return json.NewEncoder(w).Encode(apiErr)
}
