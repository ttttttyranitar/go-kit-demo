package routers

import (
	arithmaticServiceHandler "gin-operator/arithmatic/handler"
	metricServiceHandler "gin-operator/metricService/handler"
	mux "github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Methods("GET").Path("/config/add").Handler(arithmaticServiceHandler.KubeAddServiceHandler())
	r.Methods("GET").Path("/config/mul").Handler(arithmaticServiceHandler.KubeMultiServiceHandler())
	r.Methods("GET").Path("/metric/meta").Handler(metricServiceHandler.MetricServiceGetMetaHandler())

	return r
}
