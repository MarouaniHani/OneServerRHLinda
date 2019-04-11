package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	http1 "net/http"
	endpoint "onServicemgo/rh/pkg/endpoint"
	"onServicemgo/rh/pkg/io"

	"gopkg.in/mgo.v2/bson"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeGetHandler creates the handler logic
func makeGetHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/employees/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)),
	)
}

// decodeGetRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetRequest{}
	return req, nil
}

// encodeGetResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddHandler creates the handler logicJ
func makeAddHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST", "OPTIONS").Path("/employees/").Handler(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"POST"}),
		)(http.NewServer(endpoints.AddEndpoint, decodeAddRequest, encodeAddResponse, options...)))
}

// decodeAddRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Employee); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteHandler creates the handler logic
func makeDeleteHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/employees/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeDeleteResponse, options...)))
}

// decodeDeleteRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateHandler creates the handler logic
func makeUpdateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/employees/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"PUT"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.UpdateEndpoint, decodeUpdateRequest, encodeUpdateResponse, options...)))
}

// decodeUpdateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fmt.Println("docs ids:", r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateRequest{
		io.Employee{
			Id: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.Employee); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDHandler creates the handler logic
func makeGetByIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET", "OPTIONS").Path("/employees/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDEndpoint, decodeGetByIDRequest, encodeGetByIDResponse, options...)))
}

// decodeGetByIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByCreteriaHandler creates the handler logic
func makeGetByCreteriaHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/employees/{name}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByCreteriaEndpoint, decodeGetByCreteriaRequest, encodeGetByCreteriaResponse, options...)))
}

// decodeGetByCreteriaRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByCreteriaRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		return nil, errors.New("not a valid creteria")
	}
	req := endpoint.GetByCreteriaRequest{
		Creteria: name,
	}
	return req, nil
}

// encodeGetByCreteriaResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByCreteriaResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByMultiCriteriaHandler creates the handler logic
func makeGetByMultiCriteriaHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET", "OPTIONS").Path("/employees/criteria/").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByMultiCriteriaEndpoint, decodeGetByMultiCriteriaRequest, encodeGetByMultiCriteriaResponse, options...)))
}

// decodeGetByMultiCriteriaRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByMultiCriteriaRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetByMultiCriteriaRequest{
		UrlMap: r.URL.String(),
	}
	return req, nil
}

// encodeGetByMultiCriteriaResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByMultiCriteriaResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

//ErrorEncoder ...
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

//ErrorDecoder ...
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeGetDepartmentHandler creates the handler logic
func makeGetDepartmentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/departments/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetDepartmentEndpoint, decodeGetDepartmentRequest, encodeGetDepartmentResponse, options...)),
	)
}

// decodeGetDepartmentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetDepartmentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetDepartmentRequest{}
	return req, nil
}

// encodeGetDepartmentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetDepartmentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddDepartmentHandler creates the handler logic
func makeAddDepartmentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST", "OPTIONS").Path("/departments/").Handler(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"POST"}),
		)(http.NewServer(endpoints.AddDepartmentEndpoint, decodeAddDepartmentRequest, encodeAddDepartmentResponse, options...)))
}

// decodeAddDepartmentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddDepartmentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddDepartmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Department); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddDepartmentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddDepartmentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteDepartmentHandler creates the handler logic
func makeDeleteDepartmentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/departments/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteDepartmentEndpoint, decodeDeleteDepartmentRequest, encodeDeleteDepartmentResponse, options...)))
}

// decodeDeleteDepartmentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteDepartmentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteDepartmentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteDepartmentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDDepartmentHandler creates the handler logic
func makeGetByIDDepartmentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET", "OPTIONS").Path("/departments/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDDepartmentEndpoint, decodeGetByIDDepartmentRequest, encodeGetByIDDepartmentResponse, options...)))
}

// decodeGetByIDDepartmentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDDepartmentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDDepartmentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDDepartmentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateDepartmentHandler creates the handler logic
func makeUpdateDepartmentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/departments/{id}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateDepartmentEndpoint, decodeUpdateDepartmentRequest, encodeUpdateDepartmentResponse, options...)))
}

// decodeUpdateDepartmentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateDepartmentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateDepartmentRequest{
		io.Department{
			ID: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.Depatment); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateDepartmentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateDepartmentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetEventHandler creates the handler logic
func makeGetEventHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/events/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetEventEndpoint, decodeGetEventRequest, encodeGetEventResponse, options...)))
}

// decodeGetEventRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetEventRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetEventRequest{}
	return req, nil
}

// encodeGetEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetEventResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddEventHandler creates the handler logic
func makeAddEventHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/events/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.AddEventEndpoint, decodeAddEventRequest, encodeAddEventResponse, options...)))
}

// decodeAddEventRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddEventRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Event); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddEventResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteEventHandler creates the handler logic
func makeDeleteEventHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE").Path("/events/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteEventEndpoint, decodeDeleteEventRequest, encodeDeleteEventResponse, options...)))
}

// decodeDeleteEventRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteEventRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteEventRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteEventResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDEventHandler creates the handler logic
func makeGetByIDEventHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/events/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDEventEndpoint, decodeGetByIDEventRequest, encodeGetByIDEventResponse, options...)))
}

// decodeGetByIDEventRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDEventRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDEventRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDEventResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetEventByMultiCriteriaHandler creates the handler logic
func makeGetEventByMultiCriteriaHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/events/criteria/").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetEventByMultiCriteriaEndpoint, decodeGetEventByMultiCriteriaRequest, encodeGetEventByMultiCriteriaResponse, options...)))
}

// decodeGetEventByMultiCriteriaRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetEventByMultiCriteriaRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetEventByMultiCriteriaRequest{
		UrlMap: r.URL.String(),
	}
	return req, nil
}

// encodeGetEventByMultiCriteriaResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetEventByMultiCriteriaResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateEventHandler creates the handler logic
func makeUpdateEventHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/events/{id}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateEventEndpoint, decodeUpdateEventRequest, encodeUpdateEventResponse, options...)))
}

// decodeUpdateEventRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateEventRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateEventRequest{
		io.Event{
			Id: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.Event); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateEventResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetAdminRequestHandler creates the handler logic
func makeGetAdminRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/admin_request/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetAdminRequestEndpoint, decodeGetAdminRequestRequest, encodeGetAdminRequestResponse, options...)))
}

// decodeGetAdminRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAdminRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetAdminRequestRequest{}
	return req, nil
}

// encodeGetAdminRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAdminRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddAdminRequestHandler creates the handler logic
func makeAddAdminRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/admin_request/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.AddAdminRequestEndpoint, decodeAddAdminRequestRequest, encodeAddAdminRequestResponse, options...)))
}

// decodeAddAdminRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddAdminRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddAdminRequestRequest
	if err := json.NewDecoder(r.Body).Decode(&req.AdminRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddAdminRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddAdminRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteAdminRequestHandler creates the handler logic
func makeDeleteAdminRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/admin_request/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteAdminRequestEndpoint, decodeDeleteAdminRequestRequest, encodeDeleteAdminRequestResponse, options...)))
}

// decodeDeleteAdminRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteAdminRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteAdminRequestRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteAdminRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteAdminRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDAdminRequestHandler creates the handler logic
func makeGetByIDAdminRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET", "OPTIONS").Path("/admin_request/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDAdminRequestEndpoint, decodeGetByIDAdminRequestRequest, encodeGetByIDAdminRequestResponse, options...)))
}

// decodeGetByIDAdminRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDAdminRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDAdminRequestRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDAdminRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDAdminRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetAdminRequestByMultiCriteriaHandler creates the handler logic
func makeGetAdminRequestByMultiCriteriaHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET", "OPTIONS").Path("/admin_request/criteria/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetAdminRequestByMultiCriteriaEndpoint, decodeGetAdminRequestByMultiCriteriaRequest, encodeGetAdminRequestByMultiCriteriaResponse, options...)))
}

// decodeGetAdminRequestByMultiCriteriaRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAdminRequestByMultiCriteriaRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetAdminRequestByMultiCriteriaRequest{
		UrlMap: r.URL.String(),
	}
	return req, nil
}

// encodeGetAdminRequestByMultiCriteriaResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAdminRequestByMultiCriteriaResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateAdminRequestHandler creates the handler logic
func makeUpdateAdminRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/admin_request/{id}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateAdminRequestEndpoint, decodeUpdateAdminRequestRequest, encodeUpdateAdminRequestResponse, options...)))
}

// decodeUpdateAdminRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateAdminRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateAdminRequestRequest{
		io.AdminRequest{
			Id: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.AdminRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateAdminRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateAdminRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetLeaveRequestHandler creates the handler logic
func makeGetLeaveRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/leave_request/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetLeaveRequestEndpoint, decodeGetLeaveRequestRequest, encodeGetLeaveRequestResponse, options...)))
}

// decodeGetLeaveRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetLeaveRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetLeaveRequestRequest{}
	return req, nil
}

// encodeGetLeaveRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetLeaveRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddLeaveRequestHandler creates the handler logic
func makeAddLeaveRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST", "OPTIONS").Path("/leave_request/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.AddLeaveRequestEndpoint, decodeAddLeaveRequestRequest, encodeAddLeaveRequestResponse, options...)))
}

// decodeAddLeaveRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddLeaveRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddLeaveRequestRequest
	if err := json.NewDecoder(r.Body).Decode(&req.LeaveRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddLeaveRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddLeaveRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteLeaveRequestHandler creates the handler logic
func makeDeleteLeaveRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/leave_request/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteLeaveRequestEndpoint, decodeDeleteLeaveRequestRequest, encodeDeleteLeaveRequestResponse, options...)))
}

// decodeDeleteLeaveRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteLeaveRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteLeaveRequestRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteLeaveRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteLeaveRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDLeaveRequestHandler creates the handler logic
func makeGetByIDLeaveRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/leave_request/{id}").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDLeaveRequestEndpoint, decodeGetByIDLeaveRequestRequest, encodeGetByIDLeaveRequestResponse, options...)))
}

// decodeGetByIDLeaveRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDLeaveRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDLeaveRequestRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDLeaveRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDLeaveRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetLeaveRequestByMultiCriteriaHandler creates the handler logic
func makeGetLeaveRequestByMultiCriteriaHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET", "OPTIONS").Path("/leave_request/criteria/").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetLeaveRequestByMultiCriteriaEndpoint, decodeGetLeaveRequestByMultiCriteriaRequest, encodeGetLeaveRequestByMultiCriteriaResponse, options...)))
}

// decodeGetLeaveRequestByMultiCriteriaRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetLeaveRequestByMultiCriteriaRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetLeaveRequestByMultiCriteriaRequest{
		UrlMap: r.URL.String(),
	}
	return req, nil
}

// encodeGetLeaveRequestByMultiCriteriaResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetLeaveRequestByMultiCriteriaResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateLeaveRequestHandler creates the handler logic
func makeUpdateLeaveRequestHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/leave_request/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"PUT"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.UpdateLeaveRequestEndpoint, decodeUpdateLeaveRequestRequest, encodeUpdateLeaveRequestResponse, options...)))
}

// decodeUpdateLeaveRequestRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateLeaveRequestRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateLeaveRequestRequest{
		io.LeaveRequest{
			Id: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.LeaveRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateLeaveRequestResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateLeaveRequestResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetConventionHandler creates the handler logic
func makeGetConventionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/conventions/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetConventionEndpoint, decodeGetConventionRequest, encodeGetConventionResponse, options...)))
}

// decodeGetConventionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetConventionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetConventionRequest{}
	return req, nil
}

// encodeGetConventionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetConventionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddConventionHandler creates the handler logic
func makeAddConventionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST", "OPTIONS").Path("/conventions/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.AddConventionEndpoint, decodeAddConventionRequest, encodeAddConventionResponse, options...)))
}

// decodeAddConventionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddConventionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddConventionRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Convention); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddConventionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddConventionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteConventionHandler creates the handler logic
func makeDeleteConventionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/conventions/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteConventionEndpoint, decodeDeleteConventionRequest, encodeDeleteConventionResponse, options...)))
}

// decodeDeleteConventionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteConventionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteConventionRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteConventionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteConventionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDConventionHandler creates the handler logic
func makeGetByIDConventionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET", "OPTIONS").Path("/conventions/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDConventionEndpoint, decodeGetByIDConventionRequest, encodeGetByIDConventionResponse, options...)))
}

// decodeGetByIDConventionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDConventionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDConventionRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDConventionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDConventionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetConventionByMultiCriteriaHandler creates the handler logic
func makeGetConventionByMultiCriteriaHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET", "OPTIONS").Path("/conventions/criteria/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetConventionByMultiCriteriaEndpoint, decodeGetConventionByMultiCriteriaRequest, encodeGetConventionByMultiCriteriaResponse, options...)))
}

// decodeGetConventionByMultiCriteriaRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetConventionByMultiCriteriaRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetConventionByMultiCriteriaRequest{
		UrlMap: r.URL.String(),
	}
	return req, nil
}

// encodeGetConventionByMultiCriteriaResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetConventionByMultiCriteriaResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateConventionHandler creates the handler logic
func makeUpdateConventionHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/conventions/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"PUT"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.UpdateConventionEndpoint, decodeUpdateConventionRequest, encodeUpdateConventionResponse, options...)))
}

// decodeUpdateConventionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateConventionRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateConventionRequest{
		io.Convention{
			Id: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.Convention); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateConventionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateConventionResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetContractTypeHandler creates the handler logic
func makeGetContractTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/contract_type/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetContractTypeEndpoint, decodeGetContractTypeRequest, encodeGetContractTypeResponse, options...)))
}

// decodeGetContractTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetContractTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetContractTypeRequest{}
	return req, nil
}

// encodeGetContractTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetContractTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddContractTypeHandler creates the handler logic
func makeAddContractTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/contract_type/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.AddContractTypeEndpoint, decodeAddContractTypeRequest, encodeAddContractTypeResponse, options...)))
}

// decodeAddContractTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddContractTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddContractTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&req.ContractType); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddContractTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddContractTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteContractTypeHandler creates the handler logic
func makeDeleteContractTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/contract_type/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteContractTypeEndpoint, decodeDeleteContractTypeRequest, encodeDeleteContractTypeResponse, options...)))
}

// decodeDeleteContractTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteContractTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteContractTypeRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteContractTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteContractTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDContractTypeHandler creates the handler logic
func makeGetByIDContractTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/contract_type/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDContractTypeEndpoint, decodeGetByIDContractTypeRequest, encodeGetByIDContractTypeResponse, options...)))
}

// decodeGetByIDContractTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDContractTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDContractTypeRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDContractTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDContractTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateContractTypeHandler creates the handler logic
func makeUpdateContractTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/contract_type/{id}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateContractTypeEndpoint, decodeUpdateContractTypeRequest, encodeUpdateContractTypeResponse, options...)))
}

// decodeUpdateContractTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateContractTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateContractTypeRequest{
		io.ContractType{
			Id: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.ContractType); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateContractTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateContractTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetEmployeeRoleHandler creates the handler logic
func makeGetEmployeeRoleHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/employee_role/").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetEmployeeRoleEndpoint, decodeGetEmployeeRoleRequest, encodeGetEmployeeRoleResponse, options...)))
}

// decodeGetEmployeeRoleRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetEmployeeRoleRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetEmployeeRoleRequest{}
	return req, nil
}

// encodeGetEmployeeRoleResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetEmployeeRoleResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddEmployeeRoleHandler creates the handler logic
func makeAddEmployeeRoleHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/employee_role/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.AddEmployeeRoleEndpoint, decodeAddEmployeeRoleRequest, encodeAddEmployeeRoleResponse, options...)))
}

// decodeAddEmployeeRoleRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddEmployeeRoleRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddEmployeeRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req.EmployeeRole); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddEmployeeRoleResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddEmployeeRoleResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteEmployeeRoleHandler creates the handler logic
func makeDeleteEmployeeRoleHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/employee_role/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteEmployeeRoleEndpoint, decodeDeleteEmployeeRoleRequest, encodeDeleteEmployeeRoleResponse, options...)))
}

// decodeDeleteEmployeeRoleRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteEmployeeRoleRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteEmployeeRoleRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteEmployeeRoleResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteEmployeeRoleResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDEmployeeRoleHandler creates the handler logic
func makeGetByIDEmployeeRoleHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/employee_role/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDEmployeeRoleEndpoint, decodeGetByIDEmployeeRoleRequest, encodeGetByIDEmployeeRoleResponse, options...)))
}

// decodeGetByIDEmployeeRoleRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDEmployeeRoleRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDEmployeeRoleRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDEmployeeRoleResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDEmployeeRoleResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateEmployeeRoleHandler creates the handler logic
func makeUpdateEmployeeRoleHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/employee_role/{id}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateEmployeeRoleEndpoint, decodeUpdateEmployeeRoleRequest, encodeUpdateEmployeeRoleResponse, options...)))
}

// decodeUpdateEmployeeRoleRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateEmployeeRoleRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateEmployeeRoleRequest{
		io.EmployeeRole{
			ID: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.EmployeeRole); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateEmployeeRoleResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateEmployeeRoleResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetRequestTypeHandler creates the handler logic
func makeGetRequestTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/request_type/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetRequestTypeEndpoint, decodeGetRequestTypeRequest, encodeGetRequestTypeResponse, options...)))
}

// decodeGetRequestTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetRequestTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetRequestTypeRequest{}
	return req, nil
}

// encodeGetRequestTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetRequestTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddRequestTypeHandler creates the handler logic
func makeAddRequestTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/request_type/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.AddRequestTypeEndpoint, decodeAddRequestTypeRequest, encodeAddRequestTypeResponse, options...)))
}

// decodeAddRequestTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddRequestTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddRequestTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&req.RequestType); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddRequestTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddRequestTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteRequestTypeHandler creates the handler logic
func makeDeleteRequestTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/request_type/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteRequestTypeEndpoint, decodeDeleteRequestTypeRequest, encodeDeleteRequestTypeResponse, options...)))
}

// decodeDeleteRequestTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteRequestTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteRequestTypeRequest{
		Id: id,
	}
	return req, nil

}

// encodeDeleteRequestTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteRequestTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDRequestTypeHandler creates the handler logic
func makeGetByIDRequestTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/request_type/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDRequestTypeEndpoint, decodeGetByIDRequestTypeRequest, encodeGetByIDRequestTypeResponse, options...)))
}

// decodeGetByIDRequestTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDRequestTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDRequestTypeRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDRequestTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDRequestTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateRequestTypeHandler creates the handler logic
func makeUpdateRequestTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/request_type/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"PUT"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.UpdateRequestTypeEndpoint, decodeUpdateRequestTypeRequest, encodeUpdateRequestTypeResponse, options...)))
}

// decodeUpdateRequestTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateRequestTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateRequestTypeRequest{
		io.RequestType{
			ID: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.RequestType); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateRequestTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateRequestTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetDocumentTypeHandler creates the handler logic
func makeGetDocumentTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/document_types/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetDocumentTypeEndpoint, decodeGetDocumentTypeRequest, encodeGetDocumentTypeResponse, options...)))
}

// decodeGetDocumentTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetDocumentTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetDocumentTypeRequest{}
	return req, nil
}

// encodeGetDocumentTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetDocumentTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddDocumentTypeHandler creates the handler logic
func makeAddDocumentTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/document_types/").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.AddDocumentTypeEndpoint, decodeAddDocumentTypeRequest, encodeAddDocumentTypeResponse, options...)))
}

// decodeAddDocumentTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddDocumentTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	var req endpoint.AddDocumentTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&req.DocumentType); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeAddDocumentTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddDocumentTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteDocumentTypeHandler creates the handler logic
func makeDeleteDocumentTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE", "OPTIONS").Path("/document_types/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteDocumentTypeEndpoint, decodeDeleteDocumentTypeRequest, encodeDeleteDocumentTypeResponse, options...)))
}

// decodeDeleteDocumentTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteDocumentTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteDocumentTypeRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteDocumentTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteDocumentTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetByIDDocumentTypeHandler creates the handler logic
func makeGetByIDDocumentTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/document_types/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetByIDDocumentTypeEndpoint, decodeGetByIDDocumentTypeRequest, encodeGetByIDDocumentTypeResponse, options...)))
}

// decodeGetByIDDocumentTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByIDDocumentTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetByIDDocumentTypeRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetByIDDocumentTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByIDDocumentTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateDocumentTypeHandler creates the handler logic
func makeUpdateDocumentTypeHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/document_types/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"PUT"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.UpdateDocumentTypeEndpoint, decodeUpdateDocumentTypeRequest, encodeUpdateDocumentTypeResponse, options...)))
}

// decodeUpdateDocumentTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateDocumentTypeRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.UpdateDocumentTypeRequest{
		io.DocumentType{
			ID: bson.ObjectIdHex(id),
		},
	}
	if err := json.NewDecoder(r.Body).Decode(&req.DocumentType); err != nil {
		return nil, err
	}
	return req, nil
}

// encodeUpdateDocumentTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateDocumentTypeResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
