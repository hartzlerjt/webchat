package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type config struct {
	Server   bool   `json:"server"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func main() {

	myconfig, err := readConfig("config.json")

	//fmt.Printf("%v", myconfig)

	if err != nil {
		fmt.Println("Error reading config")
		panic(err)
	}

	if myconfig.Server == true {
		fmt.Println("We are a server")
		runServer(&myconfig)
	}
	if myconfig.Server == false {
		fmt.Println("We are a client")
		runClient(&myconfig)
	}
}

func readConfig(filename string) (config, error) {
	fmt.Println("Reading Config")
	var conf config

	configFile, myerr := os.Open(filename)
	defer configFile.Close()

	if myerr != nil {
		return conf, myerr
	}

	jsonParser := json.NewDecoder(configFile)

	myerr = jsonParser.Decode(&conf)

	fmt.Println("Returning from readConfig")

	return conf, myerr

}

func runServer(c *config) {
	fmt.Println("Running Server")
	//fmt.Printf("%v\n", c.Server)
	http.HandleFunc("/reg", regClient)
	serverString := fmt.Sprintf("%v:%v", c.Host, c.Port)
	log.Fatal(http.ListenAndServe(serverString, nil))

}

func runClient(c *config) {
	fmt.Println("Running Client")
}

func regClient(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You are Registering!\n")
	fmt.Println(req.Method)
	if req.Method != "POST" {
		io.WriteString(w, "Only POST allowed here\n")
	}
	//fmt.Println(req.UserAgent())
}
