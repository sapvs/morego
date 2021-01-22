package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type email string

type user struct {
	fname string
	email email
}

func getUsers(users []user, filter func(u *user) bool, sort bool) []user {

}
func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	numUser, _ := strconv.ParseInt(scan.Text(), 10, 32)
	var users []user
	for i := numUser; i > 0; i-- {
		scan.Scan()
		spli := strings.Split(scan.Text(), " ")
		users = append(users, user{fname: spli[0], email: email(spli[1])})
	}

}
