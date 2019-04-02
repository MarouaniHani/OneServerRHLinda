package service

import (
	"context"
	io "onServicemgo/rh/pkg/io"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(RhService) RhService

type loggingMiddleware struct {
	logger log.Logger
	next   RhService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a RhService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next RhService) RhService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (e []io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "Get", "e", e, "error", error)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) Add(ctx context.Context, employee io.Employee) (e io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "Add", "employee", employee, "e", e, "error", error)
	}()
	return l.next.Add(ctx, employee)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "error", error)
	}()
	return l.next.Delete(ctx, id)
}
func (l loggingMiddleware) GetByID(ctx context.Context, id string) (e io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "GetByID", "id", id, "e", e, "error", error)
	}()
	return l.next.GetByID(ctx, id)
}
func (l loggingMiddleware) GetByCreteria(ctx context.Context, creteria string) (e []io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "GetByCreteria", "creteria", creteria, "e", e, "error", error)
	}()
	return l.next.GetByCreteria(ctx, creteria)
}
func (l loggingMiddleware) GetByMultiCriteria(ctx context.Context, urlMap string) (e []io.Employee, error error) {
	defer func() {
		l.logger.Log("method", "GetByMultiCriteria", "urlMap", urlMap, "e", e, "error", error)
	}()
	return l.next.GetByMultiCriteria(ctx, urlMap)
}

func (l loggingMiddleware) GetDepartment(ctx context.Context) (d []io.Department, error error) {
	defer func() {
		l.logger.Log("method", "GetDepartment", "d", d, "error", error)
	}()
	return l.next.GetDepartment(ctx)
}
func (l loggingMiddleware) AddDepartment(ctx context.Context, department io.Department) (d io.Department, error error) {
	defer func() {
		l.logger.Log("method", "AddDepartment", "department", department, "d", d, "error", error)
	}()
	return l.next.AddDepartment(ctx, department)
}
func (l loggingMiddleware) DeleteDepartment(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "DeleteDepartment", "id", id, "error", error)
	}()
	return l.next.DeleteDepartment(ctx, id)
}
func (l loggingMiddleware) GetByIDDepartment(ctx context.Context, id string) (d io.Department, error error) {
	defer func() {
		l.logger.Log("method", "GetByIDDepartment", "id", id, "d", d, "error", error)
	}()
	return l.next.GetByIDDepartment(ctx, id)
}

func (l loggingMiddleware) GetEvent(ctx context.Context) (d []io.Event, error error) {
	defer func() {
		l.logger.Log("method", "GetEvent", "d", d, "error", error)
	}()
	return l.next.GetEvent(ctx)
}
func (l loggingMiddleware) AddEvent(ctx context.Context, event io.Event) (d io.Event, error error) {
	defer func() {
		l.logger.Log("method", "AddEvent", "event", event, "d", d, "error", error)
	}()
	return l.next.AddEvent(ctx, event)
}
func (l loggingMiddleware) DeleteEvent(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "DeleteEvent", "id", id, "error", error)
	}()
	return l.next.DeleteEvent(ctx, id)
}
func (l loggingMiddleware) GetByIDEvent(ctx context.Context, id string) (d io.Event, error error) {
	defer func() {
		l.logger.Log("method", "GetByIDEvent", "id", id, "d", d, "error", error)
	}()
	return l.next.GetByIDEvent(ctx, id)
}

func (l loggingMiddleware) GetEventByMultiCriteria(ctx context.Context, urlMap string) (e []io.Event, error error) {
	defer func() {
		l.logger.Log("method", "GetEventByMultiCriteria", "urlMap", urlMap, "e", e, "error", error)
	}()
	return l.next.GetEventByMultiCriteria(ctx, urlMap)
}

func (l loggingMiddleware) GetAdminRequest(ctx context.Context) (a []io.AdminRequest, error error) {
	defer func() {
		l.logger.Log("method", "GetAdminRequest", "a", a, "error", error)
	}()
	return l.next.GetAdminRequest(ctx)
}
func (l loggingMiddleware) AddAdminRequest(ctx context.Context, adminRequest io.AdminRequest) (a io.AdminRequest, error error) {
	defer func() {
		l.logger.Log("method", "AddAdminRequest", "adminRequest", adminRequest, "a", a, "error", error)
	}()
	return l.next.AddAdminRequest(ctx, adminRequest)
}
func (l loggingMiddleware) DeleteAdminRequest(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "DeleteAdminRequest", "id", id, "error", error)
	}()
	return l.next.DeleteAdminRequest(ctx, id)
}
func (l loggingMiddleware) GetByIDAdminRequest(ctx context.Context, id string) (a io.AdminRequest, error error) {
	defer func() {
		l.logger.Log("method", "GetByIDAdminRequest", "id", id, "a", a, "error", error)
	}()
	return l.next.GetByIDAdminRequest(ctx, id)
}
func (l loggingMiddleware) GetAdminRequestByMultiCriteria(ctx context.Context, urlMap string) (a []io.AdminRequest, error error) {
	defer func() {
		l.logger.Log("method", "GetAdminRequestByMultiCriteria", "urlMap", urlMap, "a", a, "error", error)
	}()
	return l.next.GetAdminRequestByMultiCriteria(ctx, urlMap)
}

func (lo loggingMiddleware) GetLeaveRequest(ctx context.Context) (l []io.LeaveRequest, error error) {
	defer func() {
		lo.logger.Log("method", "GetLeaveRequest", "l", l, "error", error)
	}()
	return lo.next.GetLeaveRequest(ctx)
}
func (lo loggingMiddleware) AddLeaveRequest(ctx context.Context, leaveRequest io.LeaveRequest) (l io.LeaveRequest, error error) {
	defer func() {
		lo.logger.Log("method", "AddLeaveRequest", "leaveRequest", leaveRequest, "l", l, "error", error)
	}()
	return lo.next.AddLeaveRequest(ctx, leaveRequest)
}
func (lo loggingMiddleware) DeleteLeaveRequest(ctx context.Context, id string) (error error) {
	defer func() {
		lo.logger.Log("method", "DeleteLeaveRequest", "id", id, "error", error)
	}()
	return lo.next.DeleteLeaveRequest(ctx, id)
}
func (lo loggingMiddleware) GetByIDLeaveRequest(ctx context.Context, id string) (l io.LeaveRequest, error error) {
	defer func() {
		lo.logger.Log("method", "GetByIDLeaveRequest", "id", id, "l", l, "error", error)
	}()
	return lo.next.GetByIDLeaveRequest(ctx, id)
}
func (lo loggingMiddleware) GetLeaveRequestByMultiCriteria(ctx context.Context, urlMap string) (l []io.LeaveRequest, error error) {
	defer func() {
		lo.logger.Log("method", "GetLeaveRequestByMultiCriteria", "urlMap", urlMap, "l", l, "error", error)
	}()
	return lo.next.GetLeaveRequestByMultiCriteria(ctx, urlMap)
}
