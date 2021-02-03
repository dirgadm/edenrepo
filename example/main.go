package main

import (
	"edenrepo"
	"fmt"
)

func main() {
	var c = new(edenrepo.Auth)
	resp := *c
	resp = edenrepo.Auth{}
	fmt.Println(resp)
}
