package services

type LoginService interface {
	LoginUser(id int64, email string, nome string) bool
}

type loginInformation struct {
	ID    int64
	email string
	nome  string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		ID:    1,
		email: "isaqueveras@gmail.com",
		nome:  "isaque veras",
	}
}

func (info *loginInformation) LoginUser(id int64, email string, nome string) bool {
	return info.email == email && info.nome == nome && info.ID == id
}
