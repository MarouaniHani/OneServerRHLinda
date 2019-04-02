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
