package main

import (
	"easy-gurl/internal/entities"
	"easy-gurl/internal/server"
	"flag"
	"fmt"
)

var options entities.Options

func main() {

	// adicionar sqllite ou csv
	// bonitim no terminal
	howToUse()

	server.Makerequest(options)
	//cmdArgs := os.Args[1:]
	//for i, opt := range cmdArgs {
	//	fmt.Println(i, opt)
	//}

}
func howToUse() {
	flag.StringVar(&options.Method, "method", "", "Specify the method (POST, GET, PUT, DELETE). Default: empty")
	flag.StringVar(&options.URL, "url", "", "Specify the URL.")
	flag.StringVar(&options.Header, "header", "", "Header options. Empty if not necessary.")
	// json is received as string
	flag.StringVar(&options.Data, "data", "", "The data/json you want to send.")
	flag.StringVar(&options.Persist, "persist", "", "Do you want to save it to a local file to use it later?(YES/NO)")

	flag.Usage = func() {
		fmt.Printf("Usage of our Program: \n")
		fmt.Println("-method = method (post, get, put, delete)")
		fmt.Println("-url = URL")
		fmt.Println("-header = Header fields options. Pattern: \"header:value\" ")
		fmt.Println("-data = Your data. Remember to scape characters with \"\\\".See the example in our docs.")
		fmt.Println("-persist = Persist your payload into a local file to reuse this server later.")
		fmt.Printf("Example: go run cmd/main.go -method POST -url https://www.google.com -data \"{\"your\": \"json\"}\" -persist YES -header X-adzerk-ApiKey:adasdasda \n")
	}
	flag.Parse()
}
