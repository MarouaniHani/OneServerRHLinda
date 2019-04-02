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

// GetConventionRequest collects the request parameters for the GetConvention method.
type GetConventionRequest struct{}

// GetConventionResponse collects the response parameters for the GetConvention method.
type GetConventionResponse struct {
	C     []io.Convention `json:"c"`
	Error error           `json:"error"`
}

// MakeGetConventionEndpoint returns an endpoint that invokes GetConvention on the service.
func MakeGetConventionEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		c, error := s.GetConvention(ctx)
		return GetConventionResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetConventionResponse) Failed() error {
	return r.Error
}

// AddConventionRequest collects the request parameters for the AddConvention method.
type AddConventionRequest struct {
	Convention io.Convention `json:"convention"`
}

// AddConventionResponse collects the response parameters for the AddConvention method.
type AddConventionResponse struct {
	C     io.Convention `json:"c"`
	Error error         `json:"error"`
}

// MakeAddConventionEndpoint returns an endpoint that invokes AddConvention on the service.
func MakeAddConventionEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddConventionRequest)
		c, error := s.AddConvention(ctx, req.Convention)
		return AddConventionResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r AddConventionResponse) Failed() error {
	return r.Error
}

// DeleteConventionRequest collects the request parameters for the DeleteConvention method.
type DeleteConventionRequest struct {
	Id string `json:"id"`
}

// DeleteConventionResponse collects the response parameters for the DeleteConvention method.
type DeleteConventionResponse struct {
	Error error `json:"error"`
}

// MakeDeleteConventionEndpoint returns an endpoint that invokes DeleteConvention on the service.
func MakeDeleteConventionEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteConventionRequest)
		error := s.DeleteConvention(ctx, req.Id)
		return DeleteConventionResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteConventionResponse) Failed() error {
	return r.Error
}

// GetByIDConventionRequest collects the request parameters for the GetByIDConvention method.
type GetByIDConventionRequest struct {
	Id string `json:"id"`
}

// GetByIDConventionResponse collects the response parameters for the GetByIDConvention method.
type GetByIDConventionResponse struct {
	C     io.Convention `json:"c"`
	Error error         `json:"error"`
}

// MakeGetByIDConventionEndpoint returns an endpoint that invokes GetByIDConvention on the service.
func MakeGetByIDConventionEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDConventionRequest)
		c, error := s.GetByIDConvention(ctx, req.Id)
		return GetByIDConventionResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDConventionResponse) Failed() error {
	return r.Error
}

// GetConventionByMultiCriteriaRequest collects the request parameters for the GetConventionByMultiCriteria method.
type GetConventionByMultiCriteriaRequest struct {
	UrlMap string `json:"url_map"`
}

// GetConventionByMultiCriteriaResponse collects the response parameters for the GetConventionByMultiCriteria method.
type GetConventionByMultiCriteriaResponse struct {
	C     []io.Convention `json:"c"`
	Error error           `json:"error"`
}

// MakeGetConventionByMultiCriteriaEndpoint returns an endpoint that invokes GetConventionByMultiCriteria on the service.
func MakeGetConventionByMultiCriteriaEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetConventionByMultiCriteriaRequest)
		c, error := s.GetConventionByMultiCriteria(ctx, req.UrlMap)
		return GetConventionByMultiCriteriaResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetConventionByMultiCriteriaResponse) Failed() error {
	return r.Error
}

// GetConvention implements Service. Primarily useful in a client.
func (en Endpoints) GetConvention(ctx context.Context) (c []io.Convention, error error) {
	request := GetConventionRequest{}
	response, err := en.GetConventionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetConventionResponse).C, response.(GetConventionResponse).Error
}

// AddConvention implements Service. Primarily useful in a client.
func (en Endpoints) AddConvention(ctx context.Context, convention io.Convention) (c io.Convention, error error) {
	request := AddConventionRequest{Convention: convention}
	response, err := en.AddConventionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddConventionResponse).C, response.(AddConventionResponse).Error
}

// DeleteConvention implements Service. Primarily useful in a client.
func (en Endpoints) DeleteConvention(ctx context.Context, id string) (error error) {
	request := DeleteConventionRequest{Id: id}
	response, err := en.DeleteConventionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteConventionResponse).Error
}

// GetByIDConvention implements Service. Primarily useful in a client.
func (en Endpoints) GetByIDConvention(ctx context.Context, id string) (c io.Convention, error error) {
	request := GetByIDConventionRequest{Id: id}
	response, err := en.GetByIDConventionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDConventionResponse).C, response.(GetByIDConventionResponse).Error
}

// GetConventionByMultiCriteria implements Service. Primarily useful in a client.
func (en Endpoints) GetConventionByMultiCriteria(ctx context.Context, urlMap string) (c []io.Convention, error error) {
	request := GetConventionByMultiCriteriaRequest{UrlMap: urlMap}
	response, err := en.GetConventionByMultiCriteriaEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetConventionByMultiCriteriaResponse).C, response.(GetConventionByMultiCriteriaResponse).Error
}

// GetContractTypeRequest collects the request parameters for the GetContractType method.
type GetContractTypeRequest struct{}

// GetContractTypeResponse collects the response parameters for the GetContractType method.
type GetContractTypeResponse struct {
	C     []io.ContractType `json:"c"`
	Error error             `json:"error"`
}

// MakeGetContractTypeEndpoint returns an endpoint that invokes GetContractType on the service.
func MakeGetContractTypeEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		c, error := s.GetContractType(ctx)
		return GetContractTypeResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetContractTypeResponse) Failed() error {
	return r.Error
}

// AddContractTypeRequest collects the request parameters for the AddContractType method.
type AddContractTypeRequest struct {
	ContractType io.ContractType `json:"contract_type"`
}

// AddContractTypeResponse collects the response parameters for the AddContractType method.
type AddContractTypeResponse struct {
	C     io.ContractType `json:"c"`
	Error error           `json:"error"`
}

// MakeAddContractTypeEndpoint returns an endpoint that invokes AddContractType on the service.
func MakeAddContractTypeEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddContractTypeRequest)
		c, error := s.AddContractType(ctx, req.ContractType)
		return AddContractTypeResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r AddContractTypeResponse) Failed() error {
	return r.Error
}

// DeleteContractTypeRequest collects the request parameters for the DeleteContractType method.
type DeleteContractTypeRequest struct {
	Id string `json:"id"`
}

// DeleteContractTypeResponse collects the response parameters for the DeleteContractType method.
type DeleteContractTypeResponse struct {
	Error error `json:"error"`
}

// MakeDeleteContractTypeEndpoint returns an endpoint that invokes DeleteContractType on the service.
func MakeDeleteContractTypeEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteContractTypeRequest)
		error := s.DeleteContractType(ctx, req.Id)
		return DeleteContractTypeResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteContractTypeResponse) Failed() error {
	return r.Error
}

// GetByIDContractTypeRequest collects the request parameters for the GetByIDContractType method.
type GetByIDContractTypeRequest struct {
	Id string `json:"id"`
}

// GetByIDContractTypeResponse collects the response parameters for the GetByIDContractType method.
type GetByIDContractTypeResponse struct {
	C     io.ContractType `json:"c"`
	Error error           `json:"error"`
}

// MakeGetByIDContractTypeEndpoint returns an endpoint that invokes GetByIDContractType on the service.
func MakeGetByIDContractTypeEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDContractTypeRequest)
		c, error := s.GetByIDContractType(ctx, req.Id)
		return GetByIDContractTypeResponse{
			C:     c,
			Error: error,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDContractTypeResponse) Failed() error {
	return r.Error
}

// GetContractType implements Service. Primarily useful in a client.
func (en Endpoints) GetContractType(ctx context.Context) (c []io.ContractType, error error) {
	request := GetContractTypeRequest{}
	response, err := en.GetContractTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetContractTypeResponse).C, response.(GetContractTypeResponse).Error
}

// AddContractType implements Service. Primarily useful in a client.
func (en Endpoints) AddContractType(ctx context.Context, contractType io.ContractType) (c io.ContractType, error error) {
	request := AddContractTypeRequest{ContractType: contractType}
	response, err := en.AddContractTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddContractTypeResponse).C, response.(AddContractTypeResponse).Error
}

// DeleteContractType implements Service. Primarily useful in a client.
func (en Endpoints) DeleteContractType(ctx context.Context, id string) (error error) {
	request := DeleteContractTypeRequest{Id: id}
	response, err := en.DeleteContractTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteContractTypeResponse).Error
}

// GetByIDContractType implements Service. Primarily useful in a client.
func (en Endpoints) GetByIDContractType(ctx context.Context, id string) (c io.ContractType, error error) {
	request := GetByIDContractTypeRequest{Id: id}
	response, err := en.GetByIDContractTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDContractTypeResponse).C, response.(GetByIDContractTypeResponse).Error
}

// GetEmployeeRoleRequest collects the request parameters for the GetEmployeeRole method.
type GetEmployeeRoleRequest struct{}

// GetEmployeeRoleResponse collects the response parameters for the GetEmployeeRole method.
type GetEmployeeRoleResponse struct {
	R     []io.EmployeeRole `json:"r"`
	Error error             `json:"error"`
}

// MakeGetEmployeeRoleEndpoint returns an endpoint that invokes GetEmployeeRole on the service.
func MakeGetEmployeeRoleEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, error := s.GetEmployeeRole(ctx)
		return GetEmployeeRoleResponse{
			Error: error,
			R:     r,
		}, nil
	}
}

// Failed implements Failer.
func (r GetEmployeeRoleResponse) Failed() error {
	return r.Error
}

// AddEmployeeRoleRequest collects the request parameters for the AddEmployeeRole method.
type AddEmployeeRoleRequest struct {
	EmployeeRole io.EmployeeRole `json:"employee_role"`
}

// AddEmployeeRoleResponse collects the response parameters for the AddEmployeeRole method.
type AddEmployeeRoleResponse struct {
	R     io.EmployeeRole `json:"r"`
	Error error           `json:"error"`
}

// MakeAddEmployeeRoleEndpoint returns an endpoint that invokes AddEmployeeRole on the service.
func MakeAddEmployeeRoleEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddEmployeeRoleRequest)
		r, error := s.AddEmployeeRole(ctx, req.EmployeeRole)
		return AddEmployeeRoleResponse{
			Error: error,
			R:     r,
		}, nil
	}
}

// Failed implements Failer.
func (r AddEmployeeRoleResponse) Failed() error {
	return r.Error
}

// DeleteEmployeeRoleRequest collects the request parameters for the DeleteEmployeeRole method.
type DeleteEmployeeRoleRequest struct {
	Id string `json:"id"`
}

// DeleteEmployeeRoleResponse collects the response parameters for the DeleteEmployeeRole method.
type DeleteEmployeeRoleResponse struct {
	Error error `json:"error"`
}

// MakeDeleteEmployeeRoleEndpoint returns an endpoint that invokes DeleteEmployeeRole on the service.
func MakeDeleteEmployeeRoleEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteEmployeeRoleRequest)
		error := s.DeleteEmployeeRole(ctx, req.Id)
		return DeleteEmployeeRoleResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteEmployeeRoleResponse) Failed() error {
	return r.Error
}

// GetByIDEmployeeRoleRequest collects the request parameters for the GetByIDEmployeeRole method.
type GetByIDEmployeeRoleRequest struct {
	Id string `json:"id"`
}

// GetByIDEmployeeRoleResponse collects the response parameters for the GetByIDEmployeeRole method.
type GetByIDEmployeeRoleResponse struct {
	R     io.EmployeeRole `json:"r"`
	Error error           `json:"error"`
}

// MakeGetByIDEmployeeRoleEndpoint returns an endpoint that invokes GetByIDEmployeeRole on the service.
func MakeGetByIDEmployeeRoleEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDEmployeeRoleRequest)
		r, error := s.GetByIDEmployeeRole(ctx, req.Id)
		return GetByIDEmployeeRoleResponse{
			Error: error,
			R:     r,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDEmployeeRoleResponse) Failed() error {
	return r.Error
}

// GetEmployeeRole implements Service. Primarily useful in a client.
func (en Endpoints) GetEmployeeRole(ctx context.Context) (r []io.EmployeeRole, error error) {
	request := GetEmployeeRoleRequest{}
	response, err := en.GetEmployeeRoleEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetEmployeeRoleResponse).R, response.(GetEmployeeRoleResponse).Error
}

// AddEmployeeRole implements Service. Primarily useful in a client.
func (en Endpoints) AddEmployeeRole(ctx context.Context, employeeRole io.EmployeeRole) (r io.EmployeeRole, error error) {
	request := AddEmployeeRoleRequest{EmployeeRole: employeeRole}
	response, err := en.AddEmployeeRoleEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddEmployeeRoleResponse).R, response.(AddEmployeeRoleResponse).Error
}

// DeleteEmployeeRole implements Service. Primarily useful in a client.
func (en Endpoints) DeleteEmployeeRole(ctx context.Context, id string) (error error) {
	request := DeleteEmployeeRoleRequest{Id: id}
	response, err := en.DeleteEmployeeRoleEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteEmployeeRoleResponse).Error
}

// GetByIDEmployeeRole implements Service. Primarily useful in a client.
func (en Endpoints) GetByIDEmployeeRole(ctx context.Context, id string) (r io.EmployeeRole, error error) {
	request := GetByIDEmployeeRoleRequest{Id: id}
	response, err := en.GetByIDEmployeeRoleEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDEmployeeRoleResponse).R, response.(GetByIDEmployeeRoleResponse).Error
}

// GetRequestTypeRequest collects the request parameters for the GetRequestType method.
type GetRequestTypeRequest struct{}

// GetRequestTypeResponse collects the response parameters for the GetRequestType method.
type GetRequestTypeResponse struct {
	R     []io.RequestType `json:"r"`
	Error error            `json:"error"`
}

// MakeGetRequestTypeEndpoint returns an endpoint that invokes GetRequestType on the service.
func MakeGetRequestTypeEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, error := s.GetRequestType(ctx)
		return GetRequestTypeResponse{
			Error: error,
			R:     r,
		}, nil
	}
}

// Failed implements Failer.
func (r GetRequestTypeResponse) Failed() error {
	return r.Error
}

// AddRequestTypeRequest collects the request parameters for the AddRequestType method.
type AddRequestTypeRequest struct {
	RequestType io.RequestType `json:"request_type"`
}

// AddRequestTypeResponse collects the response parameters for the AddRequestType method.
type AddRequestTypeResponse struct {
	R     io.RequestType `json:"r"`
	Error error          `json:"error"`
}

// MakeAddRequestTypeEndpoint returns an endpoint that invokes AddRequestType on the service.
func MakeAddRequestTypeEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequestTypeRequest)
		r, error := s.AddRequestType(ctx, req.RequestType)
		return AddRequestTypeResponse{
			Error: error,
			R:     r,
		}, nil
	}
}

// Failed implements Failer.
func (r AddRequestTypeResponse) Failed() error {
	return r.Error
}

// DeleteRequestTypeRequest collects the request parameters for the DeleteRequestType method.
type DeleteRequestTypeRequest struct {
	Id string `json:"id"`
}

// DeleteRequestTypeResponse collects the response parameters for the DeleteRequestType method.
type DeleteRequestTypeResponse struct {
	Error error `json:"error"`
}

// MakeDeleteRequestTypeEndpoint returns an endpoint that invokes DeleteRequestType on the service.
func MakeDeleteRequestTypeEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequestTypeRequest)
		error := s.DeleteRequestType(ctx, req.Id)
		return DeleteRequestTypeResponse{Error: error}, nil
	}
}

// Failed implements Failer.
func (r DeleteRequestTypeResponse) Failed() error {
	return r.Error
}

// GetByIDRequestTypeRequest collects the request parameters for the GetByIDRequestType method.
type GetByIDRequestTypeRequest struct {
	Id string `json:"id"`
}

// GetByIDRequestTypeResponse collects the response parameters for the GetByIDRequestType method.
type GetByIDRequestTypeResponse struct {
	R     io.RequestType `json:"r"`
	Error error          `json:"error"`
}

// MakeGetByIDRequestTypeEndpoint returns an endpoint that invokes GetByIDRequestType on the service.
func MakeGetByIDRequestTypeEndpoint(s service.RhService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequestTypeRequest)
		r, error := s.GetByIDRequestType(ctx, req.Id)
		return GetByIDRequestTypeResponse{
			Error: error,
			R:     r,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIDRequestTypeResponse) Failed() error {
	return r.Error
}

// GetRequestType implements Service. Primarily useful in a client.
func (en Endpoints) GetRequestType(ctx context.Context) (r []io.RequestType, error error) {
	request := GetRequestTypeRequest{}
	response, err := en.GetRequestTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetRequestTypeResponse).R, response.(GetRequestTypeResponse).Error
}

// AddRequestType implements Service. Primarily useful in a client.
func (en Endpoints) AddRequestType(ctx context.Context, requestType io.RequestType) (r io.RequestType, error error) {
	request := AddRequestTypeRequest{RequestType: requestType}
	response, err := en.AddRequestTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddRequestTypeResponse).R, response.(AddRequestTypeResponse).Error
}

// DeleteRequestType implements Service. Primarily useful in a client.
func (en Endpoints) DeleteRequestType(ctx context.Context, id string) (error error) {
	request := DeleteRequestTypeRequest{Id: id}
	response, err := en.DeleteRequestTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteRequestTypeResponse).Error
}

// GetByIDRequestType implements Service. Primarily useful in a client.
func (en Endpoints) GetByIDRequestType(ctx context.Context, id string) (r io.RequestType, error error) {
	request := GetByIDRequestTypeRequest{Id: id}
	response, err := en.GetByIDRequestTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIDRequestTypeResponse).R, response.(GetByIDRequestTypeResponse).Error
}
