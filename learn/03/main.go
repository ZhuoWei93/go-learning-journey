package main

import (
	"fmt"

	"net/http"

	"net/http/httputil"
)

func main() {
	url := "http://www.baidu.com"
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent",

		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{

		CheckRedirect: func(req *http.Request, via []*http.Request) error {

			fmt.Println("Redirect:", req)

			return nil

		},
	}

	resp, err := client.Do(request)

	//resp, err := http.DefaultClient.Do(request)

	//resp, err := http.Get("http://www.imooc.com")

	if err != nil {

		panic(err)

	}

	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)

	if err != nil {

		panic(err)

	}

	fmt.Println(string(s))

}
