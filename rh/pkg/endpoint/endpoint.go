package endpoint

import (
	"context"
	io "onServicemgo/rh/pkg/io"
	service "onServicemgo/rh/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	E     []io.Employee `json:"e"`
	Error error         `json:"error"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		e, error := s.Get(ctx)
		return GetResponse{
			E:     e,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Error
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Employee io.Employee `json:"employee"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	E     io.Employee `json:"e"`
	Error error       `json:"error"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		e, error := s.Add(ctx, req.Employee)
		return AddResponse{
			E:     e,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Error
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Error error `json:"error"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		error := s.Delete(ctx, req.Id)
		return DeleteResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Error
}

// GetByIDRequest collects the request parameters for the GetByID method.
type GetByIDRequest struct {
	Id string `json:"id"`
}

// GetByIDResponse collects the response parameters for the GetByID method.
type GetByIDResponse struct {
	E     io.Employee `json:"e"`
	Error error       `json:"error"`
}

// MakeGetByIDEndpoint returns an endpoint that invokes GetByID on the service.
func MakeGetByIDEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		e, error := s.GetByID(ctx, req.Id)
		return GetByIDResponse{
			E:     e,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDResponse) Failed() error {
	return r.Error
}

// GetByCreteriaRequest collects the request parameters for the GetByCreteria method.
type GetByCreteriaRequest struct {
	Creteria string `json:"creteria"`
}

// GetByCreteriaResponse collects the response parameters for the GetByCreteria method.
type GetByCreteriaResponse struct {
	E     []io.Employee `json:"e"`
	Error error         `json:"error"`
}

// MakeGetByCreteriaEndpoint returns an endpoint that invokes GetByCreteria on the service.
func MakeGetByCreteriaEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByCreteriaRequest)
		e, error := s.GetByCreteria(ctx, req.Creteria)
		return GetByCreteriaResponse{
			E:     e,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByCreteriaResponse) Failed() error {
	return r.Error
}

// GetByMultiCriteriaRequest collects the request parameters for the GetByMultiCriteria method.
type GetByMultiCriteriaRequest struct {
	UrlMap string `json:"url_map"`
}

// GetByMultiCriteriaResponse collects the response parameters for the GetByMultiCriteria method.
type GetByMultiCriteriaResponse struct {
	E     []io.Employee `json:"e"`
	Error error         `json:"error"`
}

// MakeGetByMultiCriteriaEndpoint returns an endpoint that invokes GetByMultiCriteria on the service.
func MakeGetByMultiCriteriaEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByMultiCriteriaRequest)
		e, error := s.GetByMultiCriteria(ctx, req.UrlMap)
		return GetByMultiCriteriaResponse{
			E:     e,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByMultiCriteriaResponse) Failed() error {
	return r.Error
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (en Endpoints) Get(ctx context.Context) (e []io.Employee, error error) {
	request := GetRequest{}
	response, err := en.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).E, response.(GetResponse).Error
}

// Add implements Service. Primarily useful in a client.
func (en Endpoints) Add(ctx context.Context, employee io.Employee) (e io.Employee, error error) {
	request := AddRequest{Employee: employee}
	response, err := en.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).E, response.(AddResponse).Error
}

// Delete implements Service. Primarily useful in a client.
func (en Endpoints) Delete(ctx context.Context, id string) (error error) {
	request := DeleteRequest{Id: id}
	response, err := en.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Error
}

// GetByID implements Service. Primarily useful in a client.
func (en Endpoints) GetByID(ctx context.Context, id string) (e io.Employee, error error) {
	request := GetByIDRequest{Id: id}
	response, err := en.GetByIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDResponse).E, response.(GetByIDResponse).Error
}

// GetByCreteria implements Service. Primarily useful in a client.
func (en Endpoints) GetByCreteria(ctx context.Context, creteria string) (e []io.Employee, error error) {
	request := GetByCreteriaRequest{Creteria: creteria}
	response, err := en.GetByCreteriaEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByCreteriaResponse).E, response.(GetByCreteriaResponse).Error
}

// GetByMultiCriteria implements Service. Primarily useful in a client.
func (en Endpoints) GetByMultiCriteria(ctx context.Context, urlMap string) (e []io.Employee, error error) {
	request := GetByMultiCriteriaRequest{UrlMap: urlMap}
	response, err := en.GetByMultiCriteriaEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByMultiCriteriaResponse).E, response.(GetByMultiCriteriaResponse).Error
}

// GetDepartmentRequest collects the request parameters for the GetDepartment method.
type GetDepartmentRequest struct{}

// GetDepartmentResponse collects the response parameters for the GetDepartment method.
type GetDepartmentResponse struct {
	D     []io.Department `json:"d"`
	Error error           `json:"error"`
}

// MakeGetDepartmentEndpoint returns an endpoint that invokes GetDepartment on the service.
func MakeGetDepartmentEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d, error := s.GetDepartment(ctx)
		return GetDepartmentResponse{
			D:     d,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetDepartmentResponse) Failed() error {
	return r.Error
}

// AddDepartmentRequest collects the request parameters for the AddDepartment method.
type AddDepartmentRequest struct {
	Department io.Department `json:"department"`
}

// AddDepartmentResponse collects the response parameters for the AddDepartment method.
type AddDepartmentResponse struct {
	D     io.Department `json:"d"`
	Error error         `json:"error"`
}

// MakeAddDepartmentEndpoint returns an endpoint that invokes AddDepartment on the service.
func MakeAddDepartmentEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddDepartmentRequest)
		d, error := s.AddDepartment(ctx, req.Department)
		return AddDepartmentResponse{
			D:     d,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r AddDepartmentResponse) Failed() error {
	return r.Error
}

// DeleteDepartmentRequest collects the request parameters for the DeleteDepartment method.
type DeleteDepartmentRequest struct {
	Id string `json:"id"`
}

// DeleteDepartmentResponse collects the response parameters for the DeleteDepartment method.
type DeleteDepartmentResponse struct {
	Error error `json:"error"`
}

// MakeDeleteDepartmentEndpoint returns an endpoint that invokes DeleteDepartment on the service.
func MakeDeleteDepartmentEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteDepartmentRequest)
		error := s.DeleteDepartment(ctx, req.Id)
		return DeleteDepartmentResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteDepartmentResponse) Failed() error {
	return r.Error
}

// GetByIDDepartmentRequest collects the request parameters for the GetByIDDepartment method.
type GetByIDDepartmentRequest struct {
	Id string `json:"id"`
}

// GetByIDDepartmentResponse collects the response parameters for the GetByIDDepartment method.
type GetByIDDepartmentResponse struct {
	D     io.Department `json:"d"`
	Error error         `json:"error"`
}

// MakeGetByIDDepartmentEndpoint returns an endpoint that invokes GetByIDDepartment on the service.
func MakeGetByIDDepartmentEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDDepartmentRequest)
		d, error := s.GetByIDDepartment(ctx, req.Id)
		return GetByIDDepartmentResponse{
			D:     d,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDDepartmentResponse) Failed() error {
	return r.Error
}

// GetDepartment implements Service. Primarily useful in a client.
func (en Endpoints) GetDepartment(ctx context.Context) (d []io.Department, error error) {
	request := GetDepartmentRequest{}
	response, err := en.GetDepartmentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetDepartmentResponse).D, response.(GetDepartmentResponse).Error
}

// AddDepartment implements Service. Primarily useful in a client.
func (en Endpoints) AddDepartment(ctx context.Context, department io.Department) (d io.Department, error error) {
	request := AddDepartmentRequest{Department: department}
	response, err := en.AddDepartmentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddDepartmentResponse).D, response.(AddDepartmentResponse).Error
}

// DeleteDepartment implements Service. Primarily useful in a client.
func (en Endpoints) DeleteDepartment(ctx context.Context, id string) (error error) {
	request := DeleteDepartmentRequest{Id: id}
	response, err := en.DeleteDepartmentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteDepartmentResponse).Error
}

// GetByIDDepartment implements Service. Primarily useful in a client.
func (en Endpoints) GetByIDDepartment(ctx context.Context, id string) (d io.Department, error error) {
	request := GetByIDDepartmentRequest{Id: id}
	response, err := en.GetByIDDepartmentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDDepartmentResponse).D, response.(GetByIDDepartmentResponse).Error
}

// GetEventRequest collects the request parameters for the GetEvent method.
type GetEventRequest struct{}

// GetEventResponse collects the response parameters for the GetEvent method.
type GetEventResponse struct {
	D     []io.Event `json:"d"`
	Error error      `json:"error"`
}

// MakeGetEventEndpoint returns an endpoint that invokes GetEvent on the service.
func MakeGetEventEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d, error := s.GetEvent(ctx)
		return GetEventResponse{
			D:     d,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetEventResponse) Failed() error {
	return r.Error
}

// AddEventRequest collects the request parameters for the AddEvent method.
type AddEventRequest struct {
	Event io.Event `json:"event"`
}

// AddEventResponse collects the response parameters for the AddEvent method.
type AddEventResponse struct {
	D     io.Event `json:"d"`
	Error error    `json:"error"`
}

// MakeAddEventEndpoint returns an endpoint that invokes AddEvent on the service.
func MakeAddEventEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddEventRequest)
		d, error := s.AddEvent(ctx, req.Event)
		return AddEventResponse{
			D:     d,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r AddEventResponse) Failed() error {
	return r.Error
}

// DeleteEventRequest collects the request parameters for the DeleteEvent method.
type DeleteEventRequest struct {
	Id string `json:"id"`
}

// DeleteEventResponse collects the response parameters for the DeleteEvent method.
type DeleteEventResponse struct {
	Error error `json:"error"`
}

// MakeDeleteEventEndpoint returns an endpoint that invokes DeleteEvent on the service.
func MakeDeleteEventEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteEventRequest)
		error := s.DeleteEvent(ctx, req.Id)
		return DeleteEventResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteEventResponse) Failed() error {
	return r.Error
}

// GetByIDEventRequest collects the request parameters for the GetByIDEvent method.
type GetByIDEventRequest struct {
	Id string `json:"id"`
}

// GetByIDEventResponse collects the response parameters for the GetByIDEvent method.
type GetByIDEventResponse struct {
	D     io.Event `json:"d"`
	Error error    `json:"error"`
}

// MakeGetByIDEventEndpoint returns an endpoint that invokes GetByIDEvent on the service.
func MakeGetByIDEventEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDEventRequest)
		d, error := s.GetByIDEvent(ctx, req.Id)
		return GetByIDEventResponse{
			D:     d,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDEventResponse) Failed() error {
	return r.Error
}

// GetEvent implements Service. Primarily useful in a client.
func (en Endpoints) GetEvent(ctx context.Context) (d []io.Event, error error) {
	request := GetEventRequest{}
	response, err := en.GetEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetEventResponse).D, response.(GetEventResponse).Error
}

// AddEvent implements Service. Primarily useful in a client.
func (en Endpoints) AddEvent(ctx context.Context, event io.Event) (d io.Event, error error) {
	request := AddEventRequest{Event: event}
	response, err := en.AddEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddEventResponse).D, response.(AddEventResponse).Error
}

// DeleteEvent implements Service. Primarily useful in a client.
func (en Endpoints) DeleteEvent(ctx context.Context, id string) (error error) {
	request := DeleteEventRequest{Id: id}
	response, err := en.DeleteEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteEventResponse).Error
}

// GetByIDEvent implements Service. Primarily useful in a client.
func (en Endpoints) GetByIDEvent(ctx context.Context, id string) (d io.Event, error error) {
	request := GetByIDEventRequest{Id: id}
	response, err := en.GetByIDEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDEventResponse).D, response.(GetByIDEventResponse).Error
}

// GetEventByMultiCriteriaRequest collects the request parameters for the GetEventByMultiCriteria method.
type GetEventByMultiCriteriaRequest struct {
	UrlMap string `json:"url_map"`
}

// GetEventByMultiCriteriaResponse collects the response parameters for the GetEventByMultiCriteria method.
type GetEventByMultiCriteriaResponse struct {
	E     []io.Event `json:"e"`
	Error error      `json:"error"`
}

// MakeGetEventByMultiCriteriaEndpoint returns an endpoint that invokes GetEventByMultiCriteria on the service.
func MakeGetEventByMultiCriteriaEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetEventByMultiCriteriaRequest)
		e, error := s.GetEventByMultiCriteria(ctx, req.UrlMap)
		return GetEventByMultiCriteriaResponse{
			E:     e,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetEventByMultiCriteriaResponse) Failed() error {
	return r.Error
}

// GetEventByMultiCriteria implements Service. Primarily useful in a client.
func (en Endpoints) GetEventByMultiCriteria(ctx context.Context, urlMap string) (e []io.Event, error error) {
	request := GetEventByMultiCriteriaRequest{UrlMap: urlMap}
	response, err := en.GetEventByMultiCriteriaEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetEventByMultiCriteriaResponse).E, response.(GetEventByMultiCriteriaResponse).Error
}

// GetAdminRequestRequest collects the request parameters for the GetAdminRequest method.
type GetAdminRequestRequest struct{}

// GetAdminRequestResponse collects the response parameters for the GetAdminRequest method.
type GetAdminRequestResponse struct {
	A     []io.AdminRequest `json:"a"`
	Error error             `json:"error"`
}

// MakeGetAdminRequestEndpoint returns an endpoint that invokes GetAdminRequest on the service.
func MakeGetAdminRequestEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		a, error := s.GetAdminRequest(ctx)
		return GetAdminRequestResponse{
			A:     a,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAdminRequestResponse) Failed() error {
	return r.Error
}

// AddAdminRequestRequest collects the request parameters for the AddAdminRequest method.
type AddAdminRequestRequest struct {
	AdminRequest io.AdminRequest `json:"admin_request"`
}

// AddAdminRequestResponse collects the response parameters for the AddAdminRequest method.
type AddAdminRequestResponse struct {
	A     io.AdminRequest `json:"a"`
	Error error           `json:"error"`
}

// MakeAddAdminRequestEndpoint returns an endpoint that invokes AddAdminRequest on the service.
func MakeAddAdminRequestEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddAdminRequestRequest)
		a, error := s.AddAdminRequest(ctx, req.AdminRequest)
		return AddAdminRequestResponse{
			A:     a,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r AddAdminRequestResponse) Failed() error {
	return r.Error
}

// DeleteAdminRequestRequest collects the request parameters for the DeleteAdminRequest method.
type DeleteAdminRequestRequest struct {
	Id string `json:"id"`
}

// DeleteAdminRequestResponse collects the response parameters for the DeleteAdminRequest method.
type DeleteAdminRequestResponse struct {
	Error error `json:"error"`
}

// MakeDeleteAdminRequestEndpoint returns an endpoint that invokes DeleteAdminRequest on the service.
func MakeDeleteAdminRequestEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAdminRequestRequest)
		error := s.DeleteAdminRequest(ctx, req.Id)
		return DeleteAdminRequestResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteAdminRequestResponse) Failed() error {
	return r.Error
}

// GetByIDAdminRequestRequest collects the request parameters for the GetByIDAdminRequest method.
type GetByIDAdminRequestRequest struct {
	Id string `json:"id"`
}

// GetByIDAdminRequestResponse collects the response parameters for the GetByIDAdminRequest method.
type GetByIDAdminRequestResponse struct {
	A     io.AdminRequest `json:"a"`
	Error error           `json:"error"`
}

// MakeGetByIDAdminRequestEndpoint returns an endpoint that invokes GetByIDAdminRequest on the service.
func MakeGetByIDAdminRequestEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDAdminRequestRequest)
		a, error := s.GetByIDAdminRequest(ctx, req.Id)
		return GetByIDAdminRequestResponse{
			A:     a,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDAdminRequestResponse) Failed() error {
	return r.Error
}

// GetAdminRequestByMultiCriteriaRequest collects the request parameters for the GetAdminRequestByMultiCriteria method.
type GetAdminRequestByMultiCriteriaRequest struct {
	UrlMap string `json:"url_map"`
}

// GetAdminRequestByMultiCriteriaResponse collects the response parameters for the GetAdminRequestByMultiCriteria method.
type GetAdminRequestByMultiCriteriaResponse struct {
	A     []io.AdminRequest `json:"a"`
	Error error             `json:"error"`
}

// MakeGetAdminRequestByMultiCriteriaEndpoint returns an endpoint that invokes GetAdminRequestByMultiCriteria on the service.
func MakeGetAdminRequestByMultiCriteriaEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAdminRequestByMultiCriteriaRequest)
		a, error := s.GetAdminRequestByMultiCriteria(ctx, req.UrlMap)
		return GetAdminRequestByMultiCriteriaResponse{
			A:     a,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAdminRequestByMultiCriteriaResponse) Failed() error {
	return r.Error
}

// GetAdminRequest implements Service. Primarily useful in a client.
func (en Endpoints) GetAdminRequest(ctx context.Context) (a []io.AdminRequest, error error) {
	request := GetAdminRequestRequest{}
	response, err := en.GetAdminRequestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAdminRequestResponse).A, response.(GetAdminRequestResponse).Error
}

// AddAdminRequest implements Service. Primarily useful in a client.
func (en Endpoints) AddAdminRequest(ctx context.Context, adminRequest io.AdminRequest) (a io.AdminRequest, error error) {
	request := AddAdminRequestRequest{AdminRequest: adminRequest}
	response, err := en.AddAdminRequestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddAdminRequestResponse).A, response.(AddAdminRequestResponse).Error
}

// DeleteAdminRequest implements Service. Primarily useful in a client.
func (en Endpoints) DeleteAdminRequest(ctx context.Context, id string) (error error) {
	request := DeleteAdminRequestRequest{Id: id}
	response, err := en.DeleteAdminRequestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteAdminRequestResponse).Error
}

// GetByIDAdminRequest implements Service. Primarily useful in a client.
func (en Endpoints) GetByIDAdminRequest(ctx context.Context, id string) (a io.AdminRequest, error error) {
	request := GetByIDAdminRequestRequest{Id: id}
	response, err := en.GetByIDAdminRequestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDAdminRequestResponse).A, response.(GetByIDAdminRequestResponse).Error
}

// GetAdminRequestByMultiCriteria implements Service. Primarily useful in a client.
func (en Endpoints) GetAdminRequestByMultiCriteria(ctx context.Context, urlMap string) (a []io.AdminRequest, error error) {
	request := GetAdminRequestByMultiCriteriaRequest{UrlMap: urlMap}
	response, err := en.GetAdminRequestByMultiCriteriaEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAdminRequestByMultiCriteriaResponse).A, response.(GetAdminRequestByMultiCriteriaResponse).Error
}

// GetLeaveRequestRequest collects the request parameters for the GetLeaveRequest method.
type GetLeaveRequestRequest struct{}

// GetLeaveRequestResponse collects the response parameters for the GetLeaveRequest method.
type GetLeaveRequestResponse struct {
	L     []io.LeaveRequest `json:"l"`
	Error error             `json:"error"`
}

// MakeGetLeaveRequestEndpoint returns an endpoint that invokes GetLeaveRequest on the service.
func MakeGetLeaveRequestEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		l, error := s.GetLeaveRequest(ctx)
		return GetLeaveRequestResponse{
			Error: error,
			L:     l,
		}, nil
	}
}

// Failed implements Failer.
func (r GetLeaveRequestResponse) Failed() error {
	return r.Error
}

// AddLeaveRequestRequest collects the request parameters for the AddLeaveRequest method.
type AddLeaveRequestRequest struct {
	LeaveRequest io.LeaveRequest `json:"leave_request"`
}

// AddLeaveRequestResponse collects the response parameters for the AddLeaveRequest method.
type AddLeaveRequestResponse struct {
	L     io.LeaveRequest `json:"l"`
	Error error           `json:"error"`
}

// MakeAddLeaveRequestEndpoint returns an endpoint that invokes AddLeaveRequest on the service.
func MakeAddLeaveRequestEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddLeaveRequestRequest)
		l, error := s.AddLeaveRequest(ctx, req.LeaveRequest)
		return AddLeaveRequestResponse{
			Error: error,
			L:     l,
		}, nil
	}
}

// Failed implements Failer.
func (r AddLeaveRequestResponse) Failed() error {
	return r.Error
}

// DeleteLeaveRequestRequest collects the request parameters for the DeleteLeaveRequest method.
type DeleteLeaveRequestRequest struct {
	Id string `json:"id"`
}

// DeleteLeaveRequestResponse collects the response parameters for the DeleteLeaveRequest method.
type DeleteLeaveRequestResponse struct {
	Error error `json:"error"`
}

// MakeDeleteLeaveRequestEndpoint returns an endpoint that invokes DeleteLeaveRequest on the service.
func MakeDeleteLeaveRequestEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteLeaveRequestRequest)
		error := s.DeleteLeaveRequest(ctx, req.Id)
		return DeleteLeaveRequestResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteLeaveRequestResponse) Failed() error {
	return r.Error
}

// GetByIDLeaveRequestRequest collects the request parameters for the GetByIDLeaveRequest method.
type GetByIDLeaveRequestRequest struct {
	Id string `json:"id"`
}

// GetByIDLeaveRequestResponse collects the response parameters for the GetByIDLeaveRequest method.
type GetByIDLeaveRequestResponse struct {
	L     io.LeaveRequest `json:"l"`
	Error error           `json:"error"`
}

// MakeGetByIDLeaveRequestEndpoint returns an endpoint that invokes GetByIDLeaveRequest on the service.
func MakeGetByIDLeaveRequestEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDLeaveRequestRequest)
		l, error := s.GetByIDLeaveRequest(ctx, req.Id)
		return GetByIDLeaveRequestResponse{
			Error: error,
			L:     l,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDLeaveRequestResponse) Failed() error {
	return r.Error
}

// GetLeaveRequestByMultiCriteriaRequest collects the request parameters for the GetLeaveRequestByMultiCriteria method.
type GetLeaveRequestByMultiCriteriaRequest struct {
	UrlMap string `json:"url_map"`
}

// GetLeaveRequestByMultiCriteriaResponse collects the response parameters for the GetLeaveRequestByMultiCriteria method.
type GetLeaveRequestByMultiCriteriaResponse struct {
	L     []io.LeaveRequest `json:"l"`
	Error error             `json:"error"`
}

// MakeGetLeaveRequestByMultiCriteriaEndpoint returns an endpoint that invokes GetLeaveRequestByMultiCriteria on the service.
func MakeGetLeaveRequestByMultiCriteriaEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetLeaveRequestByMultiCriteriaRequest)
		l, error := s.GetLeaveRequestByMultiCriteria(ctx, req.UrlMap)
		return GetLeaveRequestByMultiCriteriaResponse{
			Error: error,
			L:     l,
		}, nil
	}
}

// Failed implements Failer.
func (r GetLeaveRequestByMultiCriteriaResponse) Failed() error {
	return r.Error
}

// GetLeaveRequest implements Service. Primarily useful in a client.
func (en Endpoints) GetLeaveRequest(ctx context.Context) (l []io.LeaveRequest, error error) {
	request := GetLeaveRequestRequest{}
	response, err := en.GetLeaveRequestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetLeaveRequestResponse).L, response.(GetLeaveRequestResponse).Error
}

// AddLeaveRequest implements Service. Primarily useful in a client.
func (en Endpoints) AddLeaveRequest(ctx context.Context, leaveRequest io.LeaveRequest) (l io.LeaveRequest, error error) {
	request := AddLeaveRequestRequest{LeaveRequest: leaveRequest}
	response, err := en.AddLeaveRequestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddLeaveRequestResponse).L, response.(AddLeaveRequestResponse).Error
}

// DeleteLeaveRequest implements Service. Primarily useful in a client.
func (en Endpoints) DeleteLeaveRequest(ctx context.Context, id string) (error error) {
	request := DeleteLeaveRequestRequest{Id: id}
	response, err := en.DeleteLeaveRequestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteLeaveRequestResponse).Error
}

// GetByIDLeaveRequest implements Service. Primarily useful in a client.
func (en Endpoints) GetByIDLeaveRequest(ctx context.Context, id string) (l io.LeaveRequest, error error) {
	request := GetByIDLeaveRequestRequest{Id: id}
	response, err := en.GetByIDLeaveRequestEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDLeaveRequestResponse).L, response.(GetByIDLeaveRequestResponse).Error
}

// GetLeaveRequestByMultiCriteria implements Service. Primarily useful in a client.
func (en Endpoints) GetLeaveRequestByMultiCriteria(ctx context.Context, urlMap string) (l []io.LeaveRequest, error error) {
	request := GetLeaveRequestByMultiCriteriaRequest{UrlMap: urlMap}
	response, err := en.GetLeaveRequestByMultiCriteriaEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetLeaveRequestByMultiCriteriaResponse).L, response.(GetLeaveRequestByMultiCriteriaResponse).Error
}
