package service

import (
	"errors"
	"goChallenge/data"
	"goChallenge/entity"
	"time"
)

type ICustomerService interface {
	InsertNewCustomer(email string, password string) (*entity.Customer, error)
	GetById(id int) (*entity.Customer, error)
}

type CustomerService struct {
	dbConnector *data.DbConnector
}

func CustomerServiceConst(dbConnector *data.DbConnector) *CustomerService {
	return &CustomerService{
		dbConnector,
	}
}

func (c *CustomerService) InsertNewCustomer(email string, password string) (*entity.Customer, error) {
	myTime := time.Now()
	customer := entity.Customer{
		IsActive:  []uint8{1},
		Email:     email,
		Password:  password,
		CreatedAt: myTime,
		CreatedBy: 1,
	}
	result := c.dbConnector.Connect().Create(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (c *CustomerService) GetById(id int) (*entity.Customer, error) {
	if id == 0 {
		return nil, errors.New("Invalid Id")
	}

	customers := []entity.Customer{}
	result := c.dbConnector.Connect().Where(entity.Customer{Id: id}).Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(customers) == 0 {
		return nil, errors.New("Customer not found.")
	}
	return &customers[0], nil
}
