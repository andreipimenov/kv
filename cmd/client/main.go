package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/andreipimenov/kvstore/config"
)

func main() {
	configFile := flag.String("config", "", "configuration file")
	port := flag.Int("port", -1, "client port")
	server := flag.String("server", "", "server host:port")
	flag.Parse()

	c, err := NewConfig(config.New(*configFile))
	if *configFile != "" && err != nil {
		log.Println(err)
	}
	if *port >= 0 {
		c.Port = *port
	}
	if *server != "" {
		serverData := strings.Split(*server, ":")
		if len(serverData) == 2 {
			c.ServerHost = serverData[0]
			serverPort, _ := strconv.Atoi(serverData[1])
			c.ServerPort = serverPort
		}
	}

	cl := NewClient(c.ServerHost, c.ServerPort)
	http.Handle("/", cl.WebUI())
	http.Handle("/process", cl.ProcessWebUI())
	log.Printf("Start listening on port %d", c.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", c.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
