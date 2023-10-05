package main

import "fmt"

type UserManager struct {
	casher string
	useCash bool
	store string
	logger string
}

type UserBuilder struct {
	usermanager *UserManager
}

func NewUserBuilder() UserBuilder {
	return UserBuilder{
		usermanager: &UserManager{},
	}
}

func (u *UserBuilder)SetCasher(str string) Userer {
	u.usermanager.casher = str
	fmt.Printf("Using %s as cahser\n", str)
	return u
}

func (u *UserBuilder)EnableCasher() Userer {
	u.usermanager.useCash = true
	fmt.Printf("Enable cache usage\n")
	return u
}

func (u *UserBuilder)SetStoreToSave(str string) Userer {
	u.usermanager.store = str
	fmt.Printf("Using %s as store\n", str)
	return u
}

func (u *UserBuilder)SetLogger(str string) Userer {
	u.usermanager.logger = str
	fmt.Printf("Using %s as logger\n", str)
	return u
}

func (u *UserBuilder)Build() *UserManager {
	if u.usermanager.store == "" {
		u.usermanager.store = "Memory"
	}
	if u.usermanager.logger == "" {
		u.usermanager.logger = "default logger"
	}

	fmt.Println("The object was successfully created")
	return u.usermanager
}

type Userer interface {
	SetCasher(str string) Userer
	EnableCasher() Userer
	SetStoreToSave(str string) Userer
	SetLogger(str string) Userer
	Build() *UserManager
}