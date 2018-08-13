package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

//gopm get -g -v golang.org/x/net/html 可以自动识别HTML文本编码

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Status not ok, resp.statuscode: %d", resp.StatusCode)
		return
	}
	//utf8reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8reader := transform.NewReader(resp.Body, DetermineEncoding(resp.Body).NewDecoder())
	bytes, err := ioutil.ReadAll(utf8reader)
	if err != nil {
		panic(err)
	}
	printCityListAll(bytes)
	//fmt.Printf("%s\n", bytes)
}
func printCityListAll(bytes []byte) [][][]byte {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)[^>]*>([^<]+)</a>`)
	match := re.FindAllSubmatch(bytes, -1)
	for _, m := range match {
		fmt.Printf("City:%s, \tURL:%s", m[2], m[1])
		fmt.Printf("\n")
	}
	fmt.Printf("len(match):", len(match))
	return match
}

// 根据传入的reader对象，返回对应的编码
func DetermineEncoding(r io.Reader) encoding.Encoding {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		panic(e)
	}

	encoding2, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding2
}
