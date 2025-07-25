package main

import (
	"fmt"

	secret "github.com/bkandh30/goSecret"
)

func main() {
	v := secret.Memory("my-key")
	err := v.Set("demo_key", "some hash value")
	if err != nil {
		panic(err)
	}

	plain, err := v.Get("demo_key")
	if err != nil {
		panic(err)
	}

	fmt.Println("Plain:", plain)
}
