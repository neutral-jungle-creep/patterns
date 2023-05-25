package structural

import "fmt"

// users example
// this example uses single interface Member for structs Users and Category
// composite’s great feature is the ability to run methods recursively over the whole tree structure and sum up the results

type Member interface {
	getName() string
	getTotalPurchases() int
}

type User struct {
	name           string
	totalPurchases int
}

func (u *User) getName() string {
	return u.name
}

func (u *User) getTotalPurchases() int {
	return u.totalPurchases
}

type Category struct {
	users []Member
	name  string
}

func (c *Category) add(m Member) {
	c.users = append(c.users, m)
}

func (c *Category) getName() string {
	return c.name
}

func (c *Category) getTotalPurchases() int {
	var total = 0
	for _, u := range c.users {
		total += u.getTotalPurchases()
	}
	return total
}

func RunComposite() {
	user1 := &User{name: "Света", totalPurchases: 5000}
	user2 := &User{name: "Карина", totalPurchases: 1500}
	user3 := &User{name: "Артур", totalPurchases: 3500}

	category1 := &Category{name: "women"}
	category2 := &Category{name: "men"}
	category3 := &Category{name: "all"}

	category1.add(user1)
	category1.add(user2)
	category2.add(user3)

	category3.add(category1)
	category3.add(category2)

	fmt.Printf("category %s have total purchases %d\n", category1.name, category1.getTotalPurchases())
	fmt.Printf("category %s have total purchases %d\n", category2.name, category2.getTotalPurchases())
	fmt.Printf("category %s have total purchases %d\n", category3.name, category3.getTotalPurchases())
}
