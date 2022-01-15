package main

import (
	"fmt"
	"github.com/Rayer/hood"
)

type Config struct {
	Host     string
	User     string
	Password string
}

type ConfigWithTag struct {
	Host     string
	User     string `confidential:"1,1"`
	Password string `confidential:"2,2"`
}

func (c ConfigWithTag) String() string {
	ret, _ := hood.PrintConfidentialData(c)
	return ret
}

func main() {
	c1 := Config{
		Host:     "https://example.host",
		User:     "admin",
		Password: "password",
	}

	c2 := ConfigWithTag(c1)

	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Printf("origin : %v\n", c1)
	fmt.Printf("concealed : %v\n", c2)
}
