package auth

import (
	"time"
)

type Authuserbehavior struct {
	Id          string    `json:"Id" xorm:"not null pk VARCHAR(32)"`
	Userid      string    `json:"UserId" xorm:"not null VARCHAR(32)"`
	Objectid    string    `json:"ObjectId" xorm:"not null NVARCHAR(64)"`
	Objecttype  string    `json:"ObjectType" xorm:"not null NVARCHAR(64)"`
	Type        string    `json:"Type" xorm:"not null VARCHAR(32)"`
	Value       string    `json:"Value" xorm:"not null VARCHAR(32)"`
	Memo        string    `json:"Memo" xorm:"not null VARCHAR(512)"`
	Revision    int       `json:"Revision" xorm:"not null INT(4)"`
	Createdby   string    `json:"CreatedBy" xorm:"not null VARCHAR(32)"`
	Createdtime time.Time `json:"CreatedTime" xorm:"not null DATETIME(8)"`
	Updatedby   string    `json:"UpdatedBy" xorm:"not null VARCHAR(32)"`
	Updatedtime time.Time `json:"UpdatedTime" xorm:"not null DATETIME(8)"`
}
