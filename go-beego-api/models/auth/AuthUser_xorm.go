package auth

import "github.com/xormplus/xorm"

type AuthuserXorm struct {
	DB *xorm.Engine
}

func (o *AuthuserXorm) Get() (list []Authuser, err error) {
	list = make([]Authuser, 0)
	err = o.DB.Desc("CreatedTime").Find(&list)
	return
}

func (o *AuthuserXorm) GetById(id string) (user *Authuser, err error) {
	user = &Authuser{}
	_, err = o.DB.Id(id).Get(user)
	return
}

func (o *AuthuserXorm) GetByEmailOrPhone(emailOrPhone string) (list []Authuser, err error) {
	list = make([]Authuser, 0)
	err = o.DB.Where("Email=? or Phone=?", emailOrPhone).Desc("CreatedTime").Find(&list)
	return
}

func (o *AuthuserXorm) Create(model *Authuser) (err error) {
	_, err = o.DB.Insert(model)
	return
}

func (o *AuthuserXorm) Update(model *Authuser, columns ...string) (err error) {
	_, err = o.DB.Id(model.Id).MustCols(columns...).Update(model)
	return
}

func (o *AuthuserXorm) Delete(id string) (err error) {
	_, err = o.DB.Delete(&Authuser{Id: id})
	return
}
