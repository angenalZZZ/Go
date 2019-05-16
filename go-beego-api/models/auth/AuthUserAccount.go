package auth

import (
	"time"
)

type Authuseraccount struct {
	Id          string    `json:"Id" xorm:"not null pk VARCHAR(32)"`
	Userid      string    `json:"UserId" xorm:"not null VARCHAR(32)"`
	Accountcode string    `json:"AccountCode" xorm:"not null NVARCHAR(64)"`
	Accounttype string    `json:"AccountType" xorm:"not null NVARCHAR(64)"`
	Password    string    `json:"Password" xorm:"not null VARCHAR(32)"`
	Status      string    `json:"Status" xorm:"not null VARCHAR(32)"`
	Revision    int       `json:"Revision" xorm:"not null INT(4)"`
	Createdby   string    `json:"CreatedBy" xorm:"not null VARCHAR(32)"`
	Createdtime time.Time `json:"CreatedTime" xorm:"not null DATETIME(8)"`
	Updatedby   string    `json:"UpdatedBy" xorm:"not null VARCHAR(32)"`
	Updatedtime time.Time `json:"UpdatedTime" xorm:"not null DATETIME(8)"`
}
