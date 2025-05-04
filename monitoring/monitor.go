package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Service struct {
	Name string
	URL  string
}

var services = []Service{
	{"Frontend", "http://nginx"},
	{"Backend API", "http://nginx/api/products"},
	{"Order Service", "http://nginx/orders"},
}

const errorLogFile = "errors.log"

func main() {
	for {
		checkServices()
		time.Sleep(5 * time.Second)
	}
}

func checkServices() {
	for _, svc := range services {
		resp, err := http.Get(svc.URL)
		if err != nil || resp.StatusCode >= 400 {
			fmt.Printf("%s: DOWN\n", svc.Name)
			logError(svc.Name, err)
			continue
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		fmt.Printf("%s: OK\n", svc.Name)
	}
	fmt.Println("---")
}

func logError(serviceName string, err error) {
	f, ferr := os.OpenFile(errorLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if ferr != nil {
		fmt.Println("Cannot open error log file:", ferr)
		return
	}
	defer f.Close()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("[%s] %s is DOWN. Error: %v\n", timestamp, serviceName, err)
	f.WriteString(msg)
}
