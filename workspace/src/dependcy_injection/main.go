package dependcy_injection

import (
	"fmt"
	"strconv"
)

type Database interface {
	FindUser(i int) string
}

type Greeter interface {
	DoGreet(name string) string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) FindUser(i int) string {
	return "User" + strconv.Itoa(i)
}

type DefaultGreeter struct{}

func (greeter DefaultGreeter) DoGreet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

type NiceGreeter struct{}

func (greeter NiceGreeter) DoGreet(name string) string {
	return fmt.Sprintf("Hello %s, nice to meet you", name)
}

type Programm struct {
	Db      Database
	Greeter Greeter
}

func (p Programm) Execute() {
	user := p.Db.FindUser(1)
	fmt.Println(p.Greeter.DoGreet(user))
}

func main() {
	// wiring
	db := DefaultDatabase{}
	greeter := NiceGreeter{}
	p := Programm{
		Db:      db,
		Greeter: greeter,
	}
	p.Execute()
}
