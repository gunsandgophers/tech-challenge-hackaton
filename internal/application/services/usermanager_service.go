package services

type UserManagerServiceInterface interface {
	Login(username string, password string) (string, error)
}
