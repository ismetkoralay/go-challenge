package entity

import (
	"database/sql"
	"time"
)

type DateType time.Time

func (t DateType) String() string {
	return time.Time(t).String()
}

type Customer struct {
	Id         int           `gorm:"primary_key;column:Id"`
	Email      string        `gorm:"column:Email"`
	Password   string        `gorm:"column:Password"`
	IsActive   []uint8       `gorm:"column:IsActive"`
	CreatedAt  time.Time     `gorm:"column:CreatedAt"`
	CreatedBy  int           `gorm:"column:CreatedBy"`
	ModifiedAt sql.NullTime  `gorm:"column:ModifiedAt"`
	ModifiedBy sql.NullInt32 `gorm:"column:ModifiedBy"`
}

func (c Customer) TableName() string {
	return "Customer"
}
