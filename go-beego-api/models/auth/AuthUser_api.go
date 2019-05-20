package auth

type AuthuserApi interface {
	Get() (list []Authuser, err error)
	GetById(id string, columns ...string) (user *Authuser, err error)
	GetByEmailOrPhone(email string) (list []Authuser, err error)

	Create(model *Authuser) (err error)

	Update(model *Authuser, id string, columns ...string) (err error)

	Delete(id string) (err error)
}
