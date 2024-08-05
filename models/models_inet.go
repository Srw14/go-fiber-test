package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type Register struct {
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required"`
	ID       string `json:"id"`
	Phone    string `json:"Phone" validate:"required,numeric,min=6,max=20"`
	Type     string `json:"type" validate:"required"`
	Webname  string `json:"Webname" validate:"required"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data        []DogsRes `json:"data"`
	Name        string    `json:"name"`
	Count       int       `json:"count"`
	Red_Sum     int       `json:"red_sum"`
	Green_Sum   int       `json:"green_sum"`
	Pink_Sum    int       `json:"pink_sum"`
	NoColor_Sum int       `json:"nocolor_sum"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type Companys struct {
	gorm.Model
	CompanyID   int    `json:"company_id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type Users struct {
	gorm.Model
	EmployeeID int       `json:"employee_id"`
	Name       string    `json:"name"`
	LastName   string    `json:"last_name"`
	BirthDay   time.Time `json:"birth_day"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	Tel        string    `json:"tel"`
}
