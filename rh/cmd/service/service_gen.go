// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	endpoint "onServicemgo/rh/pkg/endpoint"
	http1 "onServicemgo/rh/pkg/http"
	service "onServicemgo/rh/pkg/service"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"Add":                            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Add", logger))},
		"AddAdminRequest":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AddAdminRequest", logger))},
		"AddDepartment":                  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AddDepartment", logger))},
		"AddEvent":                       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AddEvent", logger))},
		"AddLeaveRequest":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "AddLeaveRequest", logger))},
		"Delete":                         {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Delete", logger))},
		"DeleteAdminRequest":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteAdminRequest", logger))},
		"DeleteDepartment":               {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteDepartment", logger))},
		"DeleteEvent":                    {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteEvent", logger))},
		"DeleteLeaveRequest":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteLeaveRequest", logger))},
		"Get":                            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Get", logger))},
		"GetAdminRequest":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetAdminRequest", logger))},
		"GetAdminRequestByMultiCriteria": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetAdminRequestByMultiCriteria", logger))},
		"GetByCreteria":                  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetByCreteria", logger))},
		"GetByID":                        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetByID", logger))},
		"GetByIDAdminRequest":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetByIDAdminRequest", logger))},
		"GetByIDDepartment":              {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetByIDDepartment", logger))},
		"GetByIDEvent":                   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetByIDEvent", logger))},
		"GetByIDLeaveRequest":            {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetByIDLeaveRequest", logger))},
		"GetByMultiCriteria":             {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetByMultiCriteria", logger))},
		"GetDepartment":                  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetDepartment", logger))},
		"GetEvent":                       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetEvent", logger))},
		"GetEventByMultiCriteria":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetEventByMultiCriteria", logger))},
		"GetLeaveRequest":                {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetLeaveRequest", logger))},
		"GetLeaveRequestByMultiCriteria": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetLeaveRequestByMultiCriteria", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Get"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Get")), endpoint.InstrumentingMiddleware(duration.With("method", "Get"))}
	mw["Add"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Add")), endpoint.InstrumentingMiddleware(duration.With("method", "Add"))}
	mw["Delete"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Delete")), endpoint.InstrumentingMiddleware(duration.With("method", "Delete"))}
	mw["GetByID"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetByID")), endpoint.InstrumentingMiddleware(duration.With("method", "GetByID"))}
	mw["GetByCreteria"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetByCreteria")), endpoint.InstrumentingMiddleware(duration.With("method", "GetByCreteria"))}
	mw["GetByMultiCriteria"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetByMultiCriteria")), endpoint.InstrumentingMiddleware(duration.With("method", "GetByMultiCriteria"))}
	mw["GetDepartment"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetDepartment")), endpoint.InstrumentingMiddleware(duration.With("method", "GetDepartment"))}
	mw["AddDepartment"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "AddDepartment")), endpoint.InstrumentingMiddleware(duration.With("method", "AddDepartment"))}
	mw["DeleteDepartment"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "DeleteDepartment")), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteDepartment"))}
	mw["GetByIDDepartment"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetByIDDepartment")), endpoint.InstrumentingMiddleware(duration.With("method", "GetByIDDepartment"))}
	mw["GetEvent"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetEvent")), endpoint.InstrumentingMiddleware(duration.With("method", "GetEvent"))}
	mw["AddEvent"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "AddEvent")), endpoint.InstrumentingMiddleware(duration.With("method", "AddEvent"))}
	mw["DeleteEvent"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "DeleteEvent")), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteEvent"))}
	mw["GetByIDEvent"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetByIDEvent")), endpoint.InstrumentingMiddleware(duration.With("method", "GetByIDEvent"))}
	mw["GetEventByMultiCriteria"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetEventByMultiCriteria")), endpoint.InstrumentingMiddleware(duration.With("method", "GetEventByMultiCriteria"))}
	mw["GetAdminRequest"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetAdminRequest")), endpoint.InstrumentingMiddleware(duration.With("method", "GetAdminRequest"))}
	mw["AddAdminRequest"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "AddAdminRequest")), endpoint.InstrumentingMiddleware(duration.With("method", "AddAdminRequest"))}
	mw["DeleteAdminRequest"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "DeleteAdminRequest")), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteAdminRequest"))}
	mw["GetByIDAdminRequest"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetByIDAdminRequest")), endpoint.InstrumentingMiddleware(duration.With("method", "GetByIDAdminRequest"))}
	mw["GetAdminRequestByMultiCriteria"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetAdminRequestByMultiCriteria")), endpoint.InstrumentingMiddleware(duration.With("method", "GetAdminRequestByMultiCriteria"))}
	mw["GetLeaveRequest"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetLeaveRequest")), endpoint.InstrumentingMiddleware(duration.With("method", "GetLeaveRequest"))}
	mw["AddLeaveRequest"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "AddLeaveRequest")), endpoint.InstrumentingMiddleware(duration.With("method", "AddLeaveRequest"))}
	mw["DeleteLeaveRequest"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "DeleteLeaveRequest")), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteLeaveRequest"))}
	mw["GetByIDLeaveRequest"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetByIDLeaveRequest")), endpoint.InstrumentingMiddleware(duration.With("method", "GetByIDLeaveRequest"))}
	mw["GetLeaveRequestByMultiCriteria"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetLeaveRequestByMultiCriteria")), endpoint.InstrumentingMiddleware(duration.With("method", "GetLeaveRequestByMultiCriteria"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Get", "Add", "Delete", "GetByID", "GetByCreteria", "GetByMultiCriteria", "GetDepartment", "AddDepartment", "DeleteDepartment", "GetByIDDepartment", "GetEvent", "AddEvent", "DeleteEvent", "GetByIDEvent", "GetEventByMultiCriteria", "GetAdminRequest", "AddAdminRequest", "DeleteAdminRequest", "GetByIDAdminRequest", "GetAdminRequestByMultiCriteria", "GetLeaveRequest", "AddLeaveRequest", "DeleteLeaveRequest", "GetByIDLeaveRequest", "GetLeaveRequestByMultiCriteria"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
