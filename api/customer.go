package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	responseModels "goChallenge/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type insertCustomerRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func (s *Server) InsertNewCustomer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		JsonResponse(w, 400, false, nil, "Invalid request body")
		return
	}

	payload := &insertCustomerRequest{}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		JsonResponse(w, 400, false, nil, err.Error())
		return
	}

	if len(payload.Email) <= 0 || len(payload.Password) <= 0 {
		JsonResponse(w, 400, false, nil, "Email and password cannot be empty")
		return
	}

	customer, err := s.customerService.InsertNewCustomer(payload.Email, payload.Password)
	if err != nil {
		JsonResponse(w, 400, false, nil, err.Error())
		return
	}

	customerModel := responseModels.CustomerModel{
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		IsActive:  len(customer.IsActive) != 0 && customer.IsActive[0] == 1,
	}
	JsonResponse(w, 200, true, &customerModel, "")
}

func (s *Server) GetCustomer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idparam, err := url.QueryUnescape(params.ByName("id"))
	if err != nil {
		JsonResponse(w, 400, false, nil, err.Error())
		return
	}

	id, err := strconv.Atoi(idparam)
	if err != nil {
		JsonResponse(w, 400, false, nil, err.Error())
		return
	}

	if id <= 0 {
		JsonResponse(w, 400, false, nil, "id cannot be 0")
		return
	}

	customer, err := s.customerService.GetById(id)
	if err != nil {
		JsonResponse(w, 400, false, nil, err.Error())
		return
	}
	customerModel := responseModels.CustomerModel{
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		IsActive:  len(customer.IsActive) != 0 && customer.IsActive[0] == 1,
	}
	JsonResponse(w, 200, true, &customerModel, "")
}
