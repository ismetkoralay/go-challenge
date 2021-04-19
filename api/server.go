package api

import (
	"github.com/julienschmidt/httprouter"
	"goChallenge/data"
	"goChallenge/service"
	"net/http"
)

type Server struct {
	router          *httprouter.Router
	dbConnector     *data.DbConnector
	customerService service.ICustomerService
}

func CreateServer() *Server {
	dbConnector := &data.DbConnector{}
	server := &Server{
		router:          httprouter.New(),
		dbConnector:     dbConnector,
		customerService: service.CustomerServiceConst(dbConnector),
	}
	server.registerRoutes()
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func swagger(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "swagger.json")
}

func (s *Server) registerRoutes() {
	s.router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		success := &SuccessResponse{
			StatusCode: 200,
			Message:    "Welcome!!!",
		}
		JsonResponse(w, 200, true, success, "")
	})
	s.router.POST("/api/v1/customers", s.InsertNewCustomer)
	s.router.GET("/api/v1/customers/:id", s.GetCustomer)
	s.router.GET("/swagger", swagger)
	s.router.NotFound = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			JsonResponse(w, http.StatusNotFound, false, nil, "")
		})
	s.router.PanicHandler = func(w http.ResponseWriter, r *http.Request, e interface{}) {
		error, ok := e.(ServerError)
		if ok {
			JsonResponse(w, error.GetStatusCode(), false, nil, error.GetMessage())
		} else {
			JsonResponse(w, http.StatusInternalServerError, false, nil, "")
		}
	}
}
