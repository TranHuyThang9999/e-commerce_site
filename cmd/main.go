package main

import (
	"ecommerce_site/src/adapter"
	"ecommerce_site/src/adapter/configs"
	"fmt"
)

func main() {
	resp := adapter.NewpostgreDb(configs.Get())
	fmt.Println(resp)
}
