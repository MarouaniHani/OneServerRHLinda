// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	http1 "net/http"
	endpoint "onServicemgo/rh/pkg/endpoint"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeGetHandler(m, endpoints, options["Get"])
	makeAddHandler(m, endpoints, options["Add"])
	makeDeleteHandler(m, endpoints, options["Delete"])
	makeUpdateHandler(m, endpoints, options["Update"])
	makeGetByIDHandler(m, endpoints, options["GetByID"])
	makeGetByCreteriaHandler(m, endpoints, options["GetByCreteria"])
	makeGetByMultiCriteriaHandler(m, endpoints, options["GetByMultiCriteria"])
	makeGetDepartmentHandler(m, endpoints, options["GetDepartment"])
	makeAddDepartmentHandler(m, endpoints, options["AddDepartment"])
	makeDeleteDepartmentHandler(m, endpoints, options["DeleteDepartment"])
	makeGetByIDDepartmentHandler(m, endpoints, options["GetByIDDepartment"])
	makeGetEventHandler(m, endpoints, options["GetEvent"])
	makeAddEventHandler(m, endpoints, options["AddEvent"])
	makeDeleteEventHandler(m, endpoints, options["DeleteEvent"])
	makeGetByIDEventHandler(m, endpoints, options["GetByIDEvent"])
	makeGetEventByMultiCriteriaHandler(m, endpoints, options["GetEventByMultiCriteria"])
	makeGetAdminRequestHandler(m, endpoints, options["GetAdminRequest"])
	makeAddAdminRequestHandler(m, endpoints, options["AddAdminRequest"])
	makeDeleteAdminRequestHandler(m, endpoints, options["DeleteAdminRequest"])
	makeGetByIDAdminRequestHandler(m, endpoints, options["GetByIDAdminRequest"])
	makeGetAdminRequestByMultiCriteriaHandler(m, endpoints, options["GetAdminRequestByMultiCriteria"])
	makeGetLeaveRequestHandler(m, endpoints, options["GetLeaveRequest"])
	makeAddLeaveRequestHandler(m, endpoints, options["AddLeaveRequest"])
	makeDeleteLeaveRequestHandler(m, endpoints, options["DeleteLeaveRequest"])
	makeGetByIDLeaveRequestHandler(m, endpoints, options["GetByIDLeaveRequest"])
	makeGetLeaveRequestByMultiCriteriaHandler(m, endpoints, options["GetLeaveRequestByMultiCriteria"])
	makeGetConventionHandler(m, endpoints, options["GetConvention"])
	makeAddConventionHandler(m, endpoints, options["AddConvention"])
	makeDeleteConventionHandler(m, endpoints, options["DeleteConvention"])
	makeGetByIDConventionHandler(m, endpoints, options["GetByIDConvention"])
	makeGetConventionByMultiCriteriaHandler(m, endpoints, options["GetConventionByMultiCriteria"])
	makeGetContractTypeHandler(m, endpoints, options["GetContractType"])
	makeAddContractTypeHandler(m, endpoints, options["AddContractType"])
	makeDeleteContractTypeHandler(m, endpoints, options["DeleteContractType"])
	makeGetByIDContractTypeHandler(m, endpoints, options["GetByIDContractType"])
	makeGetEmployeeRoleHandler(m, endpoints, options["GetEmployeeRole"])
	makeAddEmployeeRoleHandler(m, endpoints, options["AddEmployeeRole"])
	makeDeleteEmployeeRoleHandler(m, endpoints, options["DeleteEmployeeRole"])
	makeGetByIDEmployeeRoleHandler(m, endpoints, options["GetByIDEmployeeRole"])
	makeGetRequestTypeHandler(m, endpoints, options["GetRequestType"])
	makeAddRequestTypeHandler(m, endpoints, options["AddRequestType"])
	makeDeleteRequestTypeHandler(m, endpoints, options["DeleteRequestType"])
	makeGetByIDRequestTypeHandler(m, endpoints, options["GetByIDRequestType"])
	makeGetDocumentTypeHandler(m, endpoints, options["GetDocumentType"])
	makeAddDocumentTypeHandler(m, endpoints, options["AddDocumentType"])
	makeDeleteDocumentTypeHandler(m, endpoints, options["DeleteDocumentType"])
	makeGetByIDDocumentTypeHandler(m, endpoints, options["GetByIDDocumentType"])
	return m
}
