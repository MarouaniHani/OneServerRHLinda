package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"
	endpoint "onServicemgo/rh/pkg/endpoint"
	"onServicemgo/rh/pkg/io"
	"strconv"

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

// makeAddHandler creates the handler logic
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
	zip, _ := strconv.Atoi(r.FormValue("ZipCode"))
	tel, _ := strconv.Atoi(r.FormValue("EmployeeNumTel"))
	EmergencyTel, _ := strconv.Atoi(r.FormValue("EmergencyContactTel"))
	salary, _ := strconv.ParseFloat(r.FormValue("EmployeeSalary"), 32)
	iban, _ := strconv.Atoi(r.FormValue("EmployeeIban"))
	bic, _ := strconv.Atoi(r.FormValue("EmployeeBic"))
	req := endpoint.AddRequest{
		io.Employee{
			EmployeeName:         r.FormValue("EmployeeName"),
			EmployeeEmail:        r.FormValue("EmployeeEmail"),
			Address:              r.FormValue("Address"),
			ZipCode:              zip,
			EmployeeBirthDate:    r.FormValue("EmployeeBirthDate"),
			EmployeeNumTel:       tel,
			EmergencyContactName: r.FormValue("EmergencyContactName"),
			EmergencyContactTel:  EmergencyTel,
			EmployeeStartDate:    r.FormValue("EmployeeStartDate"),
			EmployeeSalary:       salary,
			EmployeeIban:         iban,
			EmployeeBic:          bic,
		},
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
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
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
	req := endpoint.AddDepartmentRequest{
		io.Department{
			DepartmentName: r.FormValue("DepartmentName"),
		},
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
	req := endpoint.AddEventRequest{
		io.Event{
			EventName:        r.FormValue("EventName"),
			EventDescription: r.FormValue("EventDescription"),
			EventStartDate:   r.FormValue("EventStartDate"),
		},
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
	numberofpaper, _ := strconv.Atoi(r.FormValue("NumberOfPaper"))
	requeststatus, _ := strconv.ParseBool(r.FormValue("RequestStatus"))
	req := endpoint.AddAdminRequestRequest{
		io.AdminRequest{
			NumberOfPaper:   numberofpaper,
			ApplicationDate: r.FormValue("ApplicationDate"),
			RequestReason:   r.FormValue("RequestReason"),
			RequestStatus:   requeststatus,
		},
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
