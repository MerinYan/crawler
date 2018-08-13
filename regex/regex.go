package main

import (
	"fmt"
	"regexp"
)

const text = `My email is aaabb@163.com.com
email1 is   qqq.fa@qq.com
cc@mm.com.hha23

`

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-z0-9.]+[a-zA-z0-9]+`)
	//match := re.FindString(text)
	match := re.FindAllString(text, -1)
	fmt.Printf("%s", match)
}
