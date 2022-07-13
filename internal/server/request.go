package server

import (
	"bytes"
	"easy-gurl/internal/entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Makerequest(options entities.Options) {
	data := ConvertDataToBytes(options.Data)
	fmt.Println(bytes.NewBuffer(data))
	req, err := http.NewRequest(options.Method, options.URL, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error while making your request: %v", err)
	}
	if options.Header != "" {
		headers := strings.Split(options.Header, ":")
		req.Header.Set(headers[0], headers[1])
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while getting the response from server %v", err)
	}
	var response entities.Response
	json.NewDecoder(resp.Body).Decode(&response)
	fmt.Printf("\tRESPONSE: \n"+
		"STATUS: "+resp.Status+
		"\nRESPONSE: ", resp, "\n")
	defer resp.Body.Close()
}

func ConvertDataToBytes(data interface{}) []byte {
	dataStr, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error while converting data to bytes %v", err)
	}
	return []byte(dataStr)
}
