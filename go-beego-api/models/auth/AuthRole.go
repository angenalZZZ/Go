package auth

import (
	"time"
)

type Authrole struct {
	Id          string    `json:"Id" xorm:"not null pk VARCHAR(32)"`
	Code        string    `json:"Code" xorm:"not null VARCHAR(32)"`
	Sortcode    string    `json:"SortCode" xorm:"not null VARCHAR(8)"`
	Name        string    `json:"Name" xorm:"not null NVARCHAR(96)"`
	Type        string    `json:"Type" xorm:"not null VARCHAR(32)"`
	Inworkflow  string    `json:"InWorkFlow" xorm:"not null VARCHAR(1)"`
	Status      string    `json:"Status" xorm:"not null VARCHAR(32)"`
	Summary     string    `json:"Summary" xorm:"not null NVARCHAR(1024)"`
	Revision    int       `json:"Revision" xorm:"not null INT(4)"`
	Createdby   string    `json:"CreatedBy" xorm:"not null VARCHAR(32)"`
	Createdtime time.Time `json:"CreatedTime" xorm:"not null DATETIME(8)"`
	Updatedby   string    `json:"UpdatedBy" xorm:"not null VARCHAR(32)"`
	Updatedtime time.Time `json:"UpdatedTime" xorm:"not null DATETIME(8)"`
}
