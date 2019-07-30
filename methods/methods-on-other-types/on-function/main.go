package main

import "fmt"

type By func(person User) bool

type User struct {
	id         int
	name       string
	age        int
	visitDays  int
	totalLikes int
	followers  int
}

func (by By) Filter(s1 []User) []User {
	filtered := make([]User, 0)
	for _, ele := range s1 {
		if by(ele) {
			filtered = append(filtered, ele)
		}
	}
	return filtered
}

var persons []User = []User{
	User{101, "A", 18, 20, 645, 2342},
	User{102, "B", 23, 110, 323, 110},
	User{103, "C", 40, 125, 1120, 4577},
	User{104, "D", 36, 45, 323, 1201},
	User{105, "D", 42, 45, 323, 1201},
}

func main() {

	//Frequent User
	frequent := func(user User) bool {
		return user.visitDays > 100
	}

	//Liked User
	appreciated := func(user User) bool {
		return user.totalLikes > 500
	}

	// Large followers
	respected := func(user User) bool {
		return user.followers > 1000
	}

	// Matured User
	matured := func(user User) bool {
		return user.age > 35
	}
	frequestUsers := By(frequent).Filter(persons)
	fmt.Println(frequestUsers)

	appreciatedUsers := By(appreciated).Filter(persons)
	fmt.Println(appreciatedUsers)

	respectedUsers := By(respected).Filter(persons)
	fmt.Println(respectedUsers)

	maturedUsers := By(matured).Filter(persons)
	fmt.Println(maturedUsers)
}
