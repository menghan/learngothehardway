package main

import "fmt"

func main() {
	type Foo struct {
		bar int
		baz string
	}

	foo := Foo{1, "foo"}
	fmt.Printf("foo: %+v\n", foo)

	var bar Foo
	bar.bar = 2
	bar.baz = "bar"
	fmt.Printf("bar: %+v\n", bar)

	baz := Foo{1, "foo"}
	fmt.Println("foo is equivalent to baz?", baz == foo)

	fmt.Printf("foo: bar is %d and baz is %s\n", foo.bar, foo.baz)
}
