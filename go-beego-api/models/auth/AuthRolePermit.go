package auth

import (
	"time"
)

type Authrolepermit struct {
	Id          string    `json:"Id" xorm:"not null pk VARCHAR(32)"`
	Permitcode  string    `json:"PermitCode" xorm:"not null VARCHAR(128)"`
	Roleid      string    `json:"RoleId" xorm:"not null VARCHAR(32)"`
	Revision    int       `json:"Revision" xorm:"not null INT(4)"`
	Createdby   string    `json:"CreatedBy" xorm:"not null VARCHAR(32)"`
	Createdtime time.Time `json:"CreatedTime" xorm:"not null DATETIME(8)"`
	Updatedby   string    `json:"UpdatedBy" xorm:"not null VARCHAR(32)"`
	Updatedtime time.Time `json:"UpdatedTime" xorm:"not null DATETIME(8)"`
}
