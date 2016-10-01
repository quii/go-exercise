package main

import "fmt"

type person struct{
	name string
	friends []*person
}

func makeFriends(a, b *person){
	a.friends = append(a.friends, b)
	b.friends = append(b.friends, a)
}

func listFriends(p *person) (friends []string) {
	for _, friend := range p.friends{
		friends = append(friends, friend.name)
	}
	return
}

func main() {
	chris := &person{"Chris", nil}
	ruth := &person{"Ruth", nil}
	dan := &person{"Dan", nil}

	makeFriends(chris, ruth)
	makeFriends(dan, chris)

	fmt.Println(listFriends(chris))
	fmt.Println(listFriends(ruth))
}
