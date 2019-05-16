package auth

import (
	"time"
)

type Authpermit struct {
	Code        string    `json:"Code" xorm:"not null pk VARCHAR(128)"`
	Name        string    `json:"Name" xorm:"not null VARCHAR(128)"`
	Revision    int       `json:"Revision" xorm:"not null INT(4)"`
	Createdby   string    `json:"CreatedBy" xorm:"not null VARCHAR(32)"`
	Createdtime time.Time `json:"CreatedTime" xorm:"not null DATETIME(8)"`
	Updatedby   string    `json:"UpdatedBy" xorm:"not null VARCHAR(32)"`
	Updatedtime time.Time `json:"UpdatedTime" xorm:"not null DATETIME(8)"`
}
