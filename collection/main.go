package main

import (
	"fmt";
	"collection/listpack";
)


func main() {
	c := listpack.New();

	c.Print()
	print("\n")

	fmt.Println(c.Remove(4))
	c.Print()
	print("\n")

	fmt.Println(c.Get(3))
	fmt.Println(c.First())
	fmt.Println(c.Last())
	fmt.Println(c.Length())

	print("****** after initialization: ******\n")
	for i := 0; i < 10; i++ {
		c.Add(i)
	}

	c.Print()
	print("\n")

	fmt.Println(c.Remove(4))
	c.Print()
	print("\n")

	fmt.Println(c.Get(3))
	fmt.Println(c.First())
	fmt.Println(c.Last())
	fmt.Println(c.Length())

	_, b:= c.Get(3)
	fmt.Println(b.Value())
}
