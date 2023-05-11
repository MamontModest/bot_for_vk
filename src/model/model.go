package model

type User struct {
	Uid      int
	Service  string
	Login    string
	Password string
}

// map [uid] status
// if status==0 ; user neet to write service
// if status==1 ; user need to write password
type UsersStatus map[int]int

type LastUserService map[int]*User

type LastUserCommand map[int]string

type LogPas struct {
	Login    string
	Password string
}

// Конструктор ...
func NewUser(uid int) *User {
	return &User{
		Uid: uid,
	}
}
