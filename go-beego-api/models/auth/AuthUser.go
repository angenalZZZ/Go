package auth

import (
	"time"
)

type Authuser struct {
	Id          string    `json:"Id" xorm:"not null pk VARCHAR(32)"`
	Code        string    `json:"Code" xorm:"not null VARCHAR(32)"`
	Name        string    `json:"Name" xorm:"not null NVARCHAR(96)"`
	Password    string    `json:"Password" xorm:"not null VARCHAR(32)"`
	Salt        string    `json:"Salt" xorm:"not null VARCHAR(24)"`
	Avatar      string    `json:"Avatar" xorm:"not null VARCHAR(64)"`
	Orgid       string    `json:"OrgId" xorm:"not null VARCHAR(32)"`
	Email       string    `json:"Email" xorm:"not null NVARCHAR(64)"`
	Phone       string    `json:"Phone" xorm:"not null VARCHAR(48)"`
	Status      string    `json:"Status" xorm:"not null VARCHAR(32)"`
	Revision    int       `json:"Revision" xorm:"not null INT(4)"`
	Createdby   string    `json:"CreatedBy" xorm:"not null VARCHAR(32)"`
	Createdtime time.Time `json:"CreatedTime" xorm:"not null DATETIME(8)"`
	Updatedby   string    `json:"UpdatedBy" xorm:"not null VARCHAR(32)"`
	Updatedtime time.Time `json:"UpdatedTime" xorm:"not null DATETIME(8)"`
}
