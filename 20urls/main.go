package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://search.naver.com/search.naver?sm=tab_sug.top&where=nexearch&ssc=tab.nx.all&query=%EA%B3%A0%EC%9A%A9%EB%85%B8%EB%8F%99%EB%B6%80+%EB%85%B8%EB%8F%99%ED%8F%AC%ED%84%B8&oquery=%EC%83%A4%EB%A8%BC+%EA%B7%80%EC%8B%A0%EC%A0%84&tqi=ipv55lpzLiwssmP3nswssssstSh-431517&acq=%EA%B3%A0%EC%9A%A9%EB%85%B8%EB%8F%99%EB%B6%80&acr=10&qdt=0"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["where"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
