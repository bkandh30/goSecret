package main

import (
	"fmt"

	secret "github.com/bkandh30/goSecret"
)

func main() {
	v := secret.File("my-key", ".secrets")
	err := v.Set("demo_key1", "123 value")
	if err != nil {
		panic(err)
	}

	err = v.Set("demo_key2", "456 value")
	if err != nil {
		panic(err)
	}

	err = v.Set("demo_key3", "789 value")
	if err != nil {
		panic(err)
	}

	plain, err := v.Get("demo_key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)

	plain, err = v.Get("demo_key2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)

	plain, err = v.Get("demo_key3")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
