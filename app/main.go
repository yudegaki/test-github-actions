package main

import (
	"fmt"

	"github.com/yudegaki/test-github-actions/user"
)

func main() {
	u := user.GenerateUser(1, "taro", 3)
	fmt.Println(u)
}
