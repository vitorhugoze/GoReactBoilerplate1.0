package models

type UserIdentifier interface {
	GetUserIdentifier() string
}

type User struct {
	Id       int
	Username string `binding:"required,min=5,max=18" json:"user"`
	PessName string `binding:"required,min=5,max=30" json:"pess"`
	UserMail string `binding:"required,min=5,max=30" json:"mail"`
	Password string `binding:"required,min=5,max=60" json:"pass"`
}

/*
Identifier can be either Username or Email
*/
type UserLogin struct {
	Identifier string `binding:"required,min=5,max=18" json:"user"`
	Password   string `binding:"required,min=5,max=60" json:"pass"`
}

func (user User) GetUserIdentifier() string {

	if len(user.Username) > 0 {
		return user.Username
	} else {
		return user.UserMail
	}

}

func (user UserLogin) GetUserIdentifier() string {
	return user.Identifier
}
