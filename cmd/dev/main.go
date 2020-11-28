package main

import (
	"path"

	"sungora/lib/app"
)

func main() {

	obj := &Test1{}
	obj.Name = "/page1/page2/page3/file.txt"

	app.Dumper(path.Split(obj.Name))


}

type Test1 struct {
	Test2
	Name string
	ID   string
}

type Test2 struct {
	Name string
}
