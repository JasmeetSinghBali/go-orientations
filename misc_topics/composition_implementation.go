package main

import "fmt"

/*Has-a-Relationship i.e Composition implementation , in contrast Inheritance has (In-Relationship)*/
// avoid inheritance bcz its creates complex heirarchy and is bad for complex permutations&combinations maintainance

type alive struct {
}

type walkable struct {
}

type swimmable struct {
}

/*functions of type alive struct*/
func (alive) eat() {
	fmt.Println("Alive and eating")
}
func (alive) sleep() {
	fmt.Println("Alive and sleeping")
}

/*function of type walkable struct*/
func (walkable) walk() {
	fmt.Println("Walking the miles..")
}

/*function of type swimmable struct*/
func (swimmable) swim() {
	fmt.Println("Swimming in the high tides of ocean...")
}

/*a new struct that has three type inference to the alive,walkable & simmable structs*/
type plutoCreature struct {
	a alive
	w walkable
	s swimmable
}

func (pc plutoCreature) eat() {
	pc.a.eat()
}

func (pc plutoCreature) sleep() {
	pc.a.sleep()
}

func (pc plutoCreature) walk() {
	pc.w.walk()
}

func (pc plutoCreature) swim() {
	pc.s.swim()
}

/*a new creatur that can fly to along with a,w,s*/
type goose struct {
	a alive
	w walkable
	s swimmable
	f flyable
}

// struct that support all flying creatures
type flyable struct {
}

func (flyable) fly() {
	fmt.Println("I beleive I can fly....tadaaa!")
}

func main() {
	fmt.Println("FRANKESTEIN MONSTER-----")
	frankestine := plutoCreature{}
	frankestine.eat()
	frankestine.sleep()
	frankestine.walk()
	frankestine.swim()
	fmt.Println("GOOSE-----")
	gooseCreature := goose{}
	gooseCreature.a.eat()
	gooseCreature.a.sleep()
	gooseCreature.s.swim()
	gooseCreature.w.walk()
	gooseCreature.f.fly()
}
