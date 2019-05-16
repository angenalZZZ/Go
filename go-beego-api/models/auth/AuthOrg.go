package auth

import (
	"time"
)

type Authorg struct {
	Id          string    `json:"Id" xorm:"not null pk VARCHAR(32)"`
	Code        string    `json:"Code" xorm:"not null VARCHAR(32)"`
	Name        string    `json:"Name" xorm:"not null NVARCHAR(256)"`
	Fullname    string    `json:"FullName" xorm:"not null NVARCHAR(256)"`
	Shortname   string    `json:"ShortName" xorm:"not null NVARCHAR(256)"`
	Sortcode    string    `json:"SortCode" xorm:"not null NVARCHAR(64)"`
	Parentid    string    `json:"ParentId" xorm:"not null VARCHAR(32)"`
	Level       string    `json:"Level" xorm:"not null VARCHAR(32)"`
	Orgtype     string    `json:"OrgType" xorm:"not null VARCHAR(32)"`
	Leader      string    `json:"Leader" xorm:"not null VARCHAR(32)"`
	Remark      string    `json:"Remark" xorm:"not null NVARCHAR(1024)"`
	Revision    int       `json:"Revision" xorm:"not null INT(4)"`
	Createdby   string    `json:"CreatedBy" xorm:"not null VARCHAR(32)"`
	Createdtime time.Time `json:"CreatedTime" xorm:"not null DATETIME(8)"`
	Updatedby   string    `json:"UpdatedBy" xorm:"not null VARCHAR(32)"`
	Updatedtime time.Time `json:"UpdatedTime" xorm:"not null DATETIME(8)"`
}
