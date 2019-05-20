package auth

import (
	"github.com/xormplus/xorm"
	"golang.org/x/gofrontend/libgo/go/regexp"
)

type AuthuserXorm struct {
	DB *xorm.Engine
}

func (o *AuthuserXorm) Get() (list []Authuser, err error) {
	list = make([]Authuser, 0)
	err = o.DB.Desc("CreatedTime").Find(&list)
	return
}

func (o *AuthuserXorm) GetById(id string, columns ...string) (user *Authuser, err error) {
	user = &Authuser{}
	_, err = o.DB.Cols(columns...).ID(id).Get(user)
	return
}

func (o *AuthuserXorm) GetByEmailOrPhone(emailOrPhone string) (list []Authuser, err error) {
	list = make([]Authuser, 0)
	if OK, _ := regexp.Match(`/^\d+$/`, []byte(emailOrPhone)); OK {
		err = o.DB.Find(&list, &Authuser{Phone: emailOrPhone})
	} else {
		err = o.DB.Find(&list, &Authuser{Email: emailOrPhone})
	}
	return
}

func (o *AuthuserXorm) Create(model *Authuser) (err error) {
	_, err = o.DB.Insert(model)
	return
}

func (o *AuthuserXorm) Update(model *Authuser, id string, columns ...string) (err error) {
	_, err = o.DB.ID(id).MustCols(columns...).Update(model)
	return
}

func (o *AuthuserXorm) Delete(id string) (err error) {
	_, err = o.DB.Delete(&Authuser{Id: id})
	return
}
