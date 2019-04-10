// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "onServicemgo/rh/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetEndpoint                            endpoint.Endpoint
	AddEndpoint                            endpoint.Endpoint
	DeleteEndpoint                         endpoint.Endpoint
	UpdateEndpoint                         endpoint.Endpoint
	GetByIDEndpoint                        endpoint.Endpoint
	GetByCreteriaEndpoint                  endpoint.Endpoint
	GetByMultiCriteriaEndpoint             endpoint.Endpoint
	GetDepartmentEndpoint                  endpoint.Endpoint
	AddDepartmentEndpoint                  endpoint.Endpoint
	DeleteDepartmentEndpoint               endpoint.Endpoint
	GetByIDDepartmentEndpoint              endpoint.Endpoint
	GetEventEndpoint                       endpoint.Endpoint
	AddEventEndpoint                       endpoint.Endpoint
	DeleteEventEndpoint                    endpoint.Endpoint
	GetByIDEventEndpoint                   endpoint.Endpoint
	GetEventByMultiCriteriaEndpoint        endpoint.Endpoint
	GetAdminRequestEndpoint                endpoint.Endpoint
	AddAdminRequestEndpoint                endpoint.Endpoint
	DeleteAdminRequestEndpoint             endpoint.Endpoint
	GetByIDAdminRequestEndpoint            endpoint.Endpoint
	GetAdminRequestByMultiCriteriaEndpoint endpoint.Endpoint
	GetLeaveRequestEndpoint                endpoint.Endpoint
	AddLeaveRequestEndpoint                endpoint.Endpoint
	DeleteLeaveRequestEndpoint             endpoint.Endpoint
	GetByIDLeaveRequestEndpoint            endpoint.Endpoint
	GetLeaveRequestByMultiCriteriaEndpoint endpoint.Endpoint
	GetConventionEndpoint                  endpoint.Endpoint
	AddConventionEndpoint                  endpoint.Endpoint
	DeleteConventionEndpoint               endpoint.Endpoint
	GetByIDConventionEndpoint              endpoint.Endpoint
	GetConventionByMultiCriteriaEndpoint   endpoint.Endpoint
	GetContractTypeEndpoint                endpoint.Endpoint
	AddContractTypeEndpoint                endpoint.Endpoint
	DeleteContractTypeEndpoint             endpoint.Endpoint
	GetByIDContractTypeEndpoint            endpoint.Endpoint
	GetEmployeeRoleEndpoint                endpoint.Endpoint
	AddEmployeeRoleEndpoint                endpoint.Endpoint
	DeleteEmployeeRoleEndpoint             endpoint.Endpoint
	GetByIDEmployeeRoleEndpoint            endpoint.Endpoint
	GetRequestTypeEndpoint                 endpoint.Endpoint
	AddRequestTypeEndpoint                 endpoint.Endpoint
	DeleteRequestTypeEndpoint              endpoint.Endpoint
	GetByIDRequestTypeEndpoint             endpoint.Endpoint
	GetDocumentTypeEndpoint                endpoint.Endpoint
	AddDocumentTypeEndpoint                endpoint.Endpoint
	DeleteDocumentTypeEndpoint             endpoint.Endpoint
	GetByIDDocumentTypeEndpoint            endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.RhService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddAdminRequestEndpoint:                MakeAddAdminRequestEndpoint(s),
		AddContractTypeEndpoint:                MakeAddContractTypeEndpoint(s),
		AddConventionEndpoint:                  MakeAddConventionEndpoint(s),
		AddDepartmentEndpoint:                  MakeAddDepartmentEndpoint(s),
		AddDocumentTypeEndpoint:                MakeAddDocumentTypeEndpoint(s),
		AddEmployeeRoleEndpoint:                MakeAddEmployeeRoleEndpoint(s),
		AddEndpoint:                            MakeAddEndpoint(s),
		AddEventEndpoint:                       MakeAddEventEndpoint(s),
		AddLeaveRequestEndpoint:                MakeAddLeaveRequestEndpoint(s),
		AddRequestTypeEndpoint:                 MakeAddRequestTypeEndpoint(s),
		DeleteAdminRequestEndpoint:             MakeDeleteAdminRequestEndpoint(s),
		DeleteContractTypeEndpoint:             MakeDeleteContractTypeEndpoint(s),
		DeleteConventionEndpoint:               MakeDeleteConventionEndpoint(s),
		DeleteDepartmentEndpoint:               MakeDeleteDepartmentEndpoint(s),
		DeleteDocumentTypeEndpoint:             MakeDeleteDocumentTypeEndpoint(s),
		DeleteEmployeeRoleEndpoint:             MakeDeleteEmployeeRoleEndpoint(s),
		DeleteEndpoint:                         MakeDeleteEndpoint(s),
		DeleteEventEndpoint:                    MakeDeleteEventEndpoint(s),
		DeleteLeaveRequestEndpoint:             MakeDeleteLeaveRequestEndpoint(s),
		DeleteRequestTypeEndpoint:              MakeDeleteRequestTypeEndpoint(s),
		GetAdminRequestByMultiCriteriaEndpoint: MakeGetAdminRequestByMultiCriteriaEndpoint(s),
		GetAdminRequestEndpoint:                MakeGetAdminRequestEndpoint(s),
		GetByCreteriaEndpoint:                  MakeGetByCreteriaEndpoint(s),
		GetByIDAdminRequestEndpoint:            MakeGetByIDAdminRequestEndpoint(s),
		GetByIDContractTypeEndpoint:            MakeGetByIDContractTypeEndpoint(s),
		GetByIDConventionEndpoint:              MakeGetByIDConventionEndpoint(s),
		GetByIDDepartmentEndpoint:              MakeGetByIDDepartmentEndpoint(s),
		GetByIDDocumentTypeEndpoint:            MakeGetByIDDocumentTypeEndpoint(s),
		GetByIDEmployeeRoleEndpoint:            MakeGetByIDEmployeeRoleEndpoint(s),
		GetByIDEndpoint:                        MakeGetByIDEndpoint(s),
		GetByIDEventEndpoint:                   MakeGetByIDEventEndpoint(s),
		GetByIDLeaveRequestEndpoint:            MakeGetByIDLeaveRequestEndpoint(s),
		GetByIDRequestTypeEndpoint:             MakeGetByIDRequestTypeEndpoint(s),
		GetByMultiCriteriaEndpoint:             MakeGetByMultiCriteriaEndpoint(s),
		GetContractTypeEndpoint:                MakeGetContractTypeEndpoint(s),
		GetConventionByMultiCriteriaEndpoint:   MakeGetConventionByMultiCriteriaEndpoint(s),
		GetConventionEndpoint:                  MakeGetConventionEndpoint(s),
		GetDepartmentEndpoint:                  MakeGetDepartmentEndpoint(s),
		GetDocumentTypeEndpoint:                MakeGetDocumentTypeEndpoint(s),
		GetEmployeeRoleEndpoint:                MakeGetEmployeeRoleEndpoint(s),
		GetEndpoint:                            MakeGetEndpoint(s),
		GetEventByMultiCriteriaEndpoint:        MakeGetEventByMultiCriteriaEndpoint(s),
		GetEventEndpoint:                       MakeGetEventEndpoint(s),
		GetLeaveRequestByMultiCriteriaEndpoint: MakeGetLeaveRequestByMultiCriteriaEndpoint(s),
		GetLeaveRequestEndpoint:                MakeGetLeaveRequestEndpoint(s),
		GetRequestTypeEndpoint:                 MakeGetRequestTypeEndpoint(s),
		UpdateEndpoint:                         MakeUpdateEndpoint(s),
	}
	for _, m := range mdw["Get"] {
		eps.GetEndpoint = m(eps.GetEndpoint)
	}
	for _, m := range mdw["Add"] {
		eps.AddEndpoint = m(eps.AddEndpoint)
	}
	for _, m := range mdw["Delete"] {
		eps.DeleteEndpoint = m(eps.DeleteEndpoint)
	}
	for _, m := range mdw["Update"] {
		eps.UpdateEndpoint = m(eps.UpdateEndpoint)
	}
	for _, m := range mdw["GetByID"] {
		eps.GetByIDEndpoint = m(eps.GetByIDEndpoint)
	}
	for _, m := range mdw["GetByCreteria"] {
		eps.GetByCreteriaEndpoint = m(eps.GetByCreteriaEndpoint)
	}
	for _, m := range mdw["GetByMultiCriteria"] {
		eps.GetByMultiCriteriaEndpoint = m(eps.GetByMultiCriteriaEndpoint)
	}
	for _, m := range mdw["GetDepartment"] {
		eps.GetDepartmentEndpoint = m(eps.GetDepartmentEndpoint)
	}
	for _, m := range mdw["AddDepartment"] {
		eps.AddDepartmentEndpoint = m(eps.AddDepartmentEndpoint)
	}
	for _, m := range mdw["DeleteDepartment"] {
		eps.DeleteDepartmentEndpoint = m(eps.DeleteDepartmentEndpoint)
	}
	for _, m := range mdw["GetByIDDepartment"] {
		eps.GetByIDDepartmentEndpoint = m(eps.GetByIDDepartmentEndpoint)
	}
	for _, m := range mdw["GetEvent"] {
		eps.GetEventEndpoint = m(eps.GetEventEndpoint)
	}
	for _, m := range mdw["AddEvent"] {
		eps.AddEventEndpoint = m(eps.AddEventEndpoint)
	}
	for _, m := range mdw["DeleteEvent"] {
		eps.DeleteEventEndpoint = m(eps.DeleteEventEndpoint)
	}
	for _, m := range mdw["GetByIDEvent"] {
		eps.GetByIDEventEndpoint = m(eps.GetByIDEventEndpoint)
	}
	for _, m := range mdw["GetEventByMultiCriteria"] {
		eps.GetEventByMultiCriteriaEndpoint = m(eps.GetEventByMultiCriteriaEndpoint)
	}
	for _, m := range mdw["GetAdminRequest"] {
		eps.GetAdminRequestEndpoint = m(eps.GetAdminRequestEndpoint)
	}
	for _, m := range mdw["AddAdminRequest"] {
		eps.AddAdminRequestEndpoint = m(eps.AddAdminRequestEndpoint)
	}
	for _, m := range mdw["DeleteAdminRequest"] {
		eps.DeleteAdminRequestEndpoint = m(eps.DeleteAdminRequestEndpoint)
	}
	for _, m := range mdw["GetByIDAdminRequest"] {
		eps.GetByIDAdminRequestEndpoint = m(eps.GetByIDAdminRequestEndpoint)
	}
	for _, m := range mdw["GetAdminRequestByMultiCriteria"] {
		eps.GetAdminRequestByMultiCriteriaEndpoint = m(eps.GetAdminRequestByMultiCriteriaEndpoint)
	}
	for _, m := range mdw["GetLeaveRequest"] {
		eps.GetLeaveRequestEndpoint = m(eps.GetLeaveRequestEndpoint)
	}
	for _, m := range mdw["AddLeaveRequest"] {
		eps.AddLeaveRequestEndpoint = m(eps.AddLeaveRequestEndpoint)
	}
	for _, m := range mdw["DeleteLeaveRequest"] {
		eps.DeleteLeaveRequestEndpoint = m(eps.DeleteLeaveRequestEndpoint)
	}
	for _, m := range mdw["GetByIDLeaveRequest"] {
		eps.GetByIDLeaveRequestEndpoint = m(eps.GetByIDLeaveRequestEndpoint)
	}
	for _, m := range mdw["GetLeaveRequestByMultiCriteria"] {
		eps.GetLeaveRequestByMultiCriteriaEndpoint = m(eps.GetLeaveRequestByMultiCriteriaEndpoint)
	}
	for _, m := range mdw["GetConvention"] {
		eps.GetConventionEndpoint = m(eps.GetConventionEndpoint)
	}
	for _, m := range mdw["AddConvention"] {
		eps.AddConventionEndpoint = m(eps.AddConventionEndpoint)
	}
	for _, m := range mdw["DeleteConvention"] {
		eps.DeleteConventionEndpoint = m(eps.DeleteConventionEndpoint)
	}
	for _, m := range mdw["GetByIDConvention"] {
		eps.GetByIDConventionEndpoint = m(eps.GetByIDConventionEndpoint)
	}
	for _, m := range mdw["GetConventionByMultiCriteria"] {
		eps.GetConventionByMultiCriteriaEndpoint = m(eps.GetConventionByMultiCriteriaEndpoint)
	}
	for _, m := range mdw["GetContractType"] {
		eps.GetContractTypeEndpoint = m(eps.GetContractTypeEndpoint)
	}
	for _, m := range mdw["AddContractType"] {
		eps.AddContractTypeEndpoint = m(eps.AddContractTypeEndpoint)
	}
	for _, m := range mdw["DeleteContractType"] {
		eps.DeleteContractTypeEndpoint = m(eps.DeleteContractTypeEndpoint)
	}
	for _, m := range mdw["GetByIDContractType"] {
		eps.GetByIDContractTypeEndpoint = m(eps.GetByIDContractTypeEndpoint)
	}
	for _, m := range mdw["GetEmployeeRole"] {
		eps.GetEmployeeRoleEndpoint = m(eps.GetEmployeeRoleEndpoint)
	}
	for _, m := range mdw["AddEmployeeRole"] {
		eps.AddEmployeeRoleEndpoint = m(eps.AddEmployeeRoleEndpoint)
	}
	for _, m := range mdw["DeleteEmployeeRole"] {
		eps.DeleteEmployeeRoleEndpoint = m(eps.DeleteEmployeeRoleEndpoint)
	}
	for _, m := range mdw["GetByIDEmployeeRole"] {
		eps.GetByIDEmployeeRoleEndpoint = m(eps.GetByIDEmployeeRoleEndpoint)
	}
	for _, m := range mdw["GetRequestType"] {
		eps.GetRequestTypeEndpoint = m(eps.GetRequestTypeEndpoint)
	}
	for _, m := range mdw["AddRequestType"] {
		eps.AddRequestTypeEndpoint = m(eps.AddRequestTypeEndpoint)
	}
	for _, m := range mdw["DeleteRequestType"] {
		eps.DeleteRequestTypeEndpoint = m(eps.DeleteRequestTypeEndpoint)
	}
	for _, m := range mdw["GetByIDRequestType"] {
		eps.GetByIDRequestTypeEndpoint = m(eps.GetByIDRequestTypeEndpoint)
	}
	for _, m := range mdw["GetDocumentType"] {
		eps.GetDocumentTypeEndpoint = m(eps.GetDocumentTypeEndpoint)
	}
	for _, m := range mdw["AddDocumentType"] {
		eps.AddDocumentTypeEndpoint = m(eps.AddDocumentTypeEndpoint)
	}
	for _, m := range mdw["DeleteDocumentType"] {
		eps.DeleteDocumentTypeEndpoint = m(eps.DeleteDocumentTypeEndpoint)
	}
	for _, m := range mdw["GetByIDDocumentType"] {
		eps.GetByIDDocumentTypeEndpoint = m(eps.GetByIDDocumentTypeEndpoint)
	}
	return eps
}
