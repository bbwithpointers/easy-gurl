package server

import (
	"easy-gurl/internal/entities"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Makerequest(options entities.Options) {
	data := strings.NewReader(options.Data)

	_, err := http.NewRequest(options.Method, options.URL, data)
	if err != nil {
		log.Fatalf("Error while making your request: %v", err)
	}
	if options.Header != "" {
		headers := strings.Split(options.Header, ":")
		for i, str := range headers {

			//req.Header.Set(str, str)
			fmt.Println("INDEX: ", i)
			fmt.Println("STRING: ", str)
		}
	}
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("Error while getting the response from server %v", err)
	//}
	//fmt.Println(resp)
}
