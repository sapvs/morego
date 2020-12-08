package user

import "doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() {
	u.Doer.Do(1, "abc")
}
