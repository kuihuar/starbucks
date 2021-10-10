package iterator

import "fmt"

type iterator interface {
	hasNext() bool
	getNext() *user
}
type collection interface {
	createIterator()iterator
}

type userCollection struct {
	users []*user
}
func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}
type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}
func (u *userIterator) getNext() *user{
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
type user struct {
	name string
	age int
}

func IntertorUser()  {
	user1 := &user{
		name: "a",
		age:  10,
	}
	user2 := &user{
		name: "b",
		age:  20,
	}
	user3 := &user{
		name: "c",
		age:  30,
	}

	userCollection := &userCollection{users: []*user{user1,user2, user3}}

	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("user is %+v\n", user)
	}
}