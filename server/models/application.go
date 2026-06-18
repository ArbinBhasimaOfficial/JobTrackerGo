package models

import (
	"time"
)

type JobType string

const (
	Internship JobType = "Internship"
	FullTime   JobType = "Full-time"
	PartTime   JobType = "Part-time"
)

type Status string

const (
	Applied      Status = "Applied"
	Interviewing Status = "Interviewing"
	Offer        Status = "Offer"
	Rejected     Status = "Rejected"
)

type Application struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CompanyName string    `json:"company_name" gorm:"type:varchar(255);not null"`
	JobTitle    string    `json:"job_title" gorm:"type:varchar(255);not null"`
	JobType     JobType   `json:"job_type" gorm:"type:varchar(50);not null"`
	Status      Status    `json:"status" gorm:"type:varchar(50);not null"`
	AppliedDate time.Time `json:"applied_date" gorm:"type:date;not null"`
	Notes       string    `json:"notes" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
