package service

import (
	"context"
	"fmt"
	"onServicemgo/rh/pkg/db"
	io "onServicemgo/rh/pkg/io"
	"onServicemgo/utils"
	"reflect"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

// RhService describes the service.
type RhService interface {
	//Employee services
	Get(ctx context.Context) (e []io.Employee, error error)
	Add(ctx context.Context, employee io.Employee) (e io.Employee, error error)
	Delete(ctx context.Context, id string) (error error)
	GetByID(ctx context.Context, id string) (e io.Employee, error error)
	GetByCreteria(ctx context.Context, creteria string) (e []io.Employee, error error)
	GetByMultiCriteria(ctx context.Context, urlMap string) (e []io.Employee, error error)
	//Department services
	GetDepartment(ctx context.Context) (d []io.Department, error error)
	AddDepartment(ctx context.Context, department io.Department) (d io.Department, error error)
	DeleteDepartment(ctx context.Context, id string) (error error)
	GetByIDDepartment(ctx context.Context, id string) (d io.Department, error error)
	//Event service
	GetEvent(ctx context.Context) (d []io.Event, error error)
	AddEvent(ctx context.Context, event io.Event) (d io.Event, error error)
	DeleteEvent(ctx context.Context, id string) (error error)
	GetByIDEvent(ctx context.Context, id string) (d io.Event, error error)
	GetEventByMultiCriteria(ctx context.Context, urlMap string) (e []io.Event, error error)
	//AdminRequest services
	GetAdminRequest(ctx context.Context) (a []io.AdminRequest, error error)
	AddAdminRequest(ctx context.Context, adminRequest io.AdminRequest) (a io.AdminRequest, error error)
	DeleteAdminRequest(ctx context.Context, id string) (error error)
	GetByIDAdminRequest(ctx context.Context, id string) (a io.AdminRequest, error error)
	GetAdminRequestByMultiCriteria(ctx context.Context, urlMap string) (a []io.AdminRequest, error error)
	//Leave Request services
	GetLeaveRequest(ctx context.Context) (l []io.LeaveRequest, error error)
	AddLeaveRequest(ctx context.Context, leaveRequest io.LeaveRequest) (l io.LeaveRequest, error error)
	DeleteLeaveRequest(ctx context.Context, id string) (error error)
	GetByIDLeaveRequest(ctx context.Context, id string) (l io.LeaveRequest, error error)
	GetLeaveRequestByMultiCriteria(ctx context.Context, urlMap string) (l []io.LeaveRequest, error error)
	//Convention services
	GetConvention(ctx context.Context) (c []io.Convention, error error)
	AddConvention(ctx context.Context, convention io.Convention) (c io.Convention, error error)
	DeleteConvention(ctx context.Context, id string) (error error)
	GetByIDConvention(ctx context.Context, id string) (c io.Convention, error error)
	GetConventionByMultiCriteria(ctx context.Context, urlMap string) (c []io.Convention, error error)
	//Cantarct Type services
	GetContractType(ctx context.Context) (c []io.ContractType, error error)
	AddContractType(ctx context.Context, contractType io.ContractType) (c io.ContractType, error error)
	DeleteContractType(ctx context.Context, id string) (error error)
	GetByIDContractType(ctx context.Context, id string) (c io.ContractType, error error)
	//Employee Role services
	GetEmployeeRole(ctx context.Context) (r []io.EmployeeRole, error error)
	AddEmployeeRole(ctx context.Context, employeeRole io.EmployeeRole) (r io.EmployeeRole, error error)
	DeleteEmployeeRole(ctx context.Context, id string) (error error)
	GetByIDEmployeeRole(ctx context.Context, id string) (r io.EmployeeRole, error error)
	//Request Type services
	GetRequestType(ctx context.Context) (r []io.RequestType, error error)
	AddRequestType(ctx context.Context, requestType io.RequestType) (r io.RequestType, error error)
	DeleteRequestType(ctx context.Context, id string) (error error)
	GetByIDRequestType(ctx context.Context, id string) (r io.RequestType, error error)
}

type basicRhService struct{}

func (b *basicRhService) Get(ctx context.Context) (e []io.Employee, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	error = c.Find(nil).All(&e)
	return e, error
}
func (b *basicRhService) Add(ctx context.Context, employee io.Employee) (e io.Employee, error error) {
	employee.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	error = c.Insert(&employee)
	return employee, error
}
func (b *basicRhService) Delete(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicRhService) GetByID(ctx context.Context, id string) (e io.Employee, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	error = c.FindId(bson.ObjectIdHex(id)).One(&e)
	return e, error
}

func (b *basicRhService) GetByCreteria(ctx context.Context, creteria string) (e []io.Employee, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	error1 := c.Find(bson.M{"EmployeeName": creteria}).All(&e)
	if error1 == nil {
		error = error1
	}
	return e, error
}

func (b *basicRhService) GetByMultiCriteria(ctx context.Context, urlMap string) (e []io.Employee, error error) {

	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employees")
	qm := utils.QlSeparator(urlMap)
	if qm["EmployeeName"] != "" {
		error1 := c.Find(bson.M{"EmployeeName": qm["EmployeeName"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["ZipCode"] != "" {
		zipcode, _ := strconv.Atoi(qm["ZipCode"])
		error1 := c.Find(bson.M{"ZipCode": zipcode}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeEmail"] != "" {
		error1 := c.Find(bson.M{"EmployeeEmail": qm["EmployeeEmail"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["Address"] != "" {
		error1 := c.Find(bson.M{"Address": qm["Address"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeBirthDate"] != "" {
		error1 := c.Find(bson.M{"EmployeeBirthDate": qm["EmployeeBirthDate"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeNumTel"] != "" {
		numTel, _ := strconv.Atoi(qm["EmployeeNumTel"])
		error1 := c.Find(bson.M{"EmployeeNumTel": numTel}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmergencyContactName"] != "" {
		error1 := c.Find(bson.M{"EmergencyContactName": qm["EmergencyContactName"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmergencyContactTel"] != "" {
		contactTel, _ := strconv.Atoi(qm["EmergencyContactTel"])
		error1 := c.Find(bson.M{"EmergencyContactTel": contactTel}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeStartDate"] != "" {
		error1 := c.Find(bson.M{"EmployeeStartDate": qm["EmployeeStartDate"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeSalary"] != "" {
		salary, _ := strconv.Atoi(qm["EmployeeSalary"])
		error1 := c.Find(bson.M{"EmployeeSalary": salary}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeIban"] != "" {
		iban, _ := strconv.Atoi(qm["EmployeeIban"])
		error1 := c.Find(bson.M{"EmployeeIban": iban}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EmployeeBic"] != "" {
		bic, _ := strconv.Atoi(qm["EmployeeBic"])
		error1 := c.Find(bson.M{"EmployeeBic": bic}).All(&e)
		if error1 == nil {
			error = error1
		}
	}

	return e, error
}
func (b *basicRhService) GetDepartment(ctx context.Context) (d []io.Department, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return d, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("departments")
	error = c.Find(nil).All(&d)
	return d, error
}
func (b *basicRhService) AddDepartment(ctx context.Context, department io.Department) (d io.Department, error error) {
	department.ID = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return d, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("departments")
	error = c.Insert(&department)
	return department, error
}
func (b *basicRhService) DeleteDepartment(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("departments")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicRhService) GetByIDDepartment(ctx context.Context, id string) (d io.Department, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return d, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("departments")
	error = c.FindId(bson.ObjectIdHex(id)).One(&d)
	return d, error
}

func (b *basicRhService) GetEvent(ctx context.Context) (d []io.Event, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return d, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("events")
	error = c.Find(nil).All(&d)
	return d, error
}
func (b *basicRhService) AddEvent(ctx context.Context, event io.Event) (d io.Event, error error) {
	event.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return d, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("events")
	error = c.Insert(&event)
	return event, error
}
func (b *basicRhService) DeleteEvent(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("events")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicRhService) GetByIDEvent(ctx context.Context, id string) (d io.Event, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return d, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("events")
	error = c.FindId(bson.ObjectIdHex(id)).One(&d)
	return d, error
}
func (b *basicRhService) GetEventByMultiCriteria(ctx context.Context, urlMap string) (e []io.Event, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return e, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("events")
	qm := utils.QlSeparator(urlMap)
	if qm["EventName"] != "" {
		error1 := c.Find(bson.M{"EventName": qm["EventName"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	if qm["EventStartDate"] != "" {
		error1 := c.Find(bson.M{"EventStartDate": qm["EventStartDate"]}).All(&e)
		if error1 == nil {
			error = error1
		}
	}
	return e, error
}
func (b *basicRhService) GetAdminRequest(ctx context.Context) (a []io.AdminRequest, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return a, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("adminRequests")
	error = c.Find(nil).All(&a)
	return a, error
}
func (b *basicRhService) AddAdminRequest(ctx context.Context, adminRequest io.AdminRequest) (a io.AdminRequest, error error) {
	adminRequest.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return a, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("adminRequests")
	error = c.Insert(&adminRequest)
	return adminRequest, error
}
func (b *basicRhService) DeleteAdminRequest(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("adminRequests")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicRhService) GetByIDAdminRequest(ctx context.Context, id string) (a io.AdminRequest, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return a, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("adminRequests")
	error = c.FindId(bson.ObjectIdHex(id)).One(&a)
	return a, error
}
func (b *basicRhService) GetAdminRequestByMultiCriteria(ctx context.Context, urlMap string) (a []io.AdminRequest, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return a, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("adminRequests")
	qm := utils.QlSeparator(urlMap)

	/* if qm["ApliedBy"] != "" {
		error1 := c.Find(bson.M{"ApliedBy": qm["ApliedBy"]}).All(&a)
		if error1 == nil {
			error = error1
		}
	} */
	if qm["RequestStatus"] != "" {
		reqStatus, _ := strconv.ParseBool(qm["RequestStatus"])
		fmt.Println(reflect.TypeOf(reqStatus))
		error1 := c.Find(bson.M{"RequestStatus": reqStatus}).All(&a)
		if error1 == nil {
			error = error1
		}
	}
	/* if qm["RequestType"] != "" {
		error1 := c.Find(bson.M{"RequestType": qm["RequestType"]}).All(&a)
		if error1 == nil {
			error = error1
		}
	} */
	return a, error
}
func (b *basicRhService) GetLeaveRequest(ctx context.Context) (l []io.LeaveRequest, error error) {
	fmt.Println("entered in leave reason GET")
	session, err := db.GetMongoSession()
	if err != nil {
		return l, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("leaveRequests")
	error = c.Find(nil).All(&l)
	return l, error
}
func (b *basicRhService) AddLeaveRequest(ctx context.Context, leaveRequest io.LeaveRequest) (l io.LeaveRequest, error error) {
	fmt.Println("entered in leave reason ADD")
	leaveRequest.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return l, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("leaveRequests")
	error = c.Insert(&leaveRequest)
	return leaveRequest, error
}
func (b *basicRhService) DeleteLeaveRequest(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("leaveRequests")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicRhService) GetByIDLeaveRequest(ctx context.Context, id string) (l io.LeaveRequest, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return l, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("leaveRequests")
	error = c.FindId(bson.ObjectIdHex(id)).One(&l)
	return l, error
}
func (b *basicRhService) GetLeaveRequestByMultiCriteria(ctx context.Context, urlMap string) (l []io.LeaveRequest, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return l, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("leaveRequests")
	qm := utils.QlSeparator(urlMap)
	// fmt.Println("leave start date :", qm["LeaveStartDate"])
	/* if qm["ApliedBy"] != "" {
		error1 := c.Find(bson.M{"ApliedBy": qm["ApliedBy"]}).All(&l)
		if error1 == nil {
			error = error1
		}
	} */
	/* if qm["ApprouvedBy"] != "" {
		error1 := c.Find(bson.M{"ApprouvedBy": qm["ApprouvedBy"]}).All(&l)
		if error1 == nil {
			error = error1
		}
	} */
	if qm["LeaveStartDate"] != "" {

		error1 := c.Find(bson.M{"LeaveStartDate": qm["LeaveStartDate"]}).All(&l)
		if error1 == nil {
			error = error1
		}
	}
	if qm["LeaveEndDate"] != "" {
		error1 := c.Find(bson.M{"LeaveEndDate": qm["LeaveEndDate"]}).All(&l)
		if error1 == nil {
			error = error1
		}
	}
	if qm["RequestStatus"] != "" {
		LeaveReqStatus, _ := strconv.ParseBool(qm["RequestStatus"])
		error1 := c.Find(bson.M{"RequestStatus": LeaveReqStatus}).All(&l)
		if error1 == nil {
			error = error1
		}
	}
	/* if qm["RequestType"] != "" {
		error1 := c.Find(bson.M{"RequestType": qm["RequestType"]}).All(&l)
		if error1 == nil {
			error = error1
		}
	} */
	return l, error
}

func (b *basicRhService) GetConvention(ctx context.Context) (c []io.Convention, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return c, err
	}
	defer session.Close()
	d := session.DB("Linda_app").C("conventions")
	error = d.Find(nil).All(&c)
	return c, error
}
func (b *basicRhService) AddConvention(ctx context.Context, convention io.Convention) (c io.Convention, error error) {
	convention.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return c, err
	}
	defer session.Close()
	d := session.DB("Linda_app").C("conventions")
	error = d.Insert(&convention)
	return convention, error
}
func (b *basicRhService) DeleteConvention(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("conventions")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicRhService) GetByIDConvention(ctx context.Context, id string) (c io.Convention, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return c, err
	}
	defer session.Close()
	d := session.DB("Linda_app").C("conventions")
	error = d.FindId(bson.ObjectIdHex(id)).One(&c)
	return c, error
}
func (b *basicRhService) GetConventionByMultiCriteria(ctx context.Context, urlMap string) (c []io.Convention, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return c, err
	}
	defer session.Close()
	d := session.DB("Linda_app").C("conventions")
	qm := utils.QlSeparator(urlMap)
	if qm["ConventionName"] != "" {
		error1 := d.Find(bson.M{"ConventionName": qm["ConventionName"]}).All(&c)
		fmt.Println("services conventions :", error1)
		if error1 == nil {
			error = error1
		}
	}
	return c, error
}
func (b *basicRhService) GetContractType(ctx context.Context) (c []io.ContractType, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return c, err
	}
	defer session.Close()
	d := session.DB("Linda_app").C("contractTypes")
	error = d.Find(nil).All(&c)
	return c, error
}
func (b *basicRhService) AddContractType(ctx context.Context, contractType io.ContractType) (c io.ContractType, error error) {
	contractType.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return c, err
	}
	defer session.Close()
	d := session.DB("Linda_app").C("contractTypes")
	error = d.Insert(&contractType)
	return contractType, error
}
func (b *basicRhService) DeleteContractType(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("contractTypes")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicRhService) GetByIDContractType(ctx context.Context, id string) (c io.ContractType, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return c, err
	}
	defer session.Close()
	d := session.DB("Linda_app").C("contractTypes")
	error = d.FindId(bson.ObjectIdHex(id)).One(&c)
	return c, error
}

func (b *basicRhService) GetEmployeeRole(ctx context.Context) (r []io.EmployeeRole, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return r, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employeeRoles")
	error = c.Find(nil).All(&r)
	return r, error
}
func (b *basicRhService) AddEmployeeRole(ctx context.Context, employeeRole io.EmployeeRole) (r io.EmployeeRole, error error) {
	employeeRole.ID = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return r, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employeeRoles")
	error = c.Insert(&employeeRole)
	return employeeRole, error
}
func (b *basicRhService) DeleteEmployeeRole(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employeeRoles")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicRhService) GetByIDEmployeeRole(ctx context.Context, id string) (r io.EmployeeRole, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return r, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("employeeRoles")
	error = c.FindId(bson.ObjectIdHex(id)).One(&r)
	return r, error
}

func (b *basicRhService) GetRequestType(ctx context.Context) (r []io.RequestType, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return r, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("requestTypes")
	error = c.Find(nil).All(&r)
	return r, error
}
func (b *basicRhService) AddRequestType(ctx context.Context, requestType io.RequestType) (r io.RequestType, error error) {
	requestType.ID = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return r, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("requestTypes")
	error = c.Insert(&requestType)
	return requestType, error
}
func (b *basicRhService) DeleteRequestType(ctx context.Context, id string) (error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("requestTypes")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicRhService) GetByIDRequestType(ctx context.Context, id string) (r io.RequestType, error error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return r, err
	}
	defer session.Close()
	c := session.DB("Linda_app").C("requestTypes")
	error = c.FindId(bson.ObjectIdHex(id)).One(&r)
	return r, error
}

// NewBasicRhService returns a naive, stateless implementation of RhService.
func NewBasicRhService() RhService {
	return &basicRhService{}
}

// New returns a RhService with all of the expected middleware wired in.
func New(middleware []Middleware) RhService {
	var svc RhService = NewBasicRhService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
