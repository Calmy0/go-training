package main

import (
	"fmt";
	"collection/list";
)


func main() {
	c := list.New();

	c.Print()

	fmt.Println(c.Remove(4))
	c.Print()

	fmt.Println(c.Get(3))
	fmt.Println(c.First())
	fmt.Println(c.Last())
	fmt.Println(c.Length())

	print("****** after initialization: ******\n")
	for i := 0; i < 10; i++ {
		c.Add(i)
	}

	c.Print()

	fmt.Println(c.Remove(4))
	c.Print()

	fmt.Println(c.Get(3))
	fmt.Println(c.First())
	fmt.Println(c.Last())
	fmt.Println(c.Length())

	b, _:= c.Get(0)
	fmt.Println(b.Value())

	b = b.Prev()
	fmt.Println(b)
}
