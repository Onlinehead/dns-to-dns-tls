package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"github.com/miekg/dns"
)

type server struct {
	config *Config
}

func (s *server) serverAddr() string {
	return net.JoinHostPort(s.config.Server.Host,
		strconv.Itoa(s.config.Server.Port))
}

// Serve prepare configuration and start listner
func (s *server) Serve() {
	// Register handler for "."
	handlerConf := RequestHandlerConfig{
		upstreams:  s.config.Upstreams,
		resTimeout: s.config.ResTimeout,
		reqTimeout: s.config.ReqTimeout,
	}
	requestHandler := MakeHandler(&handlerConf)
	dns.HandleFunc(".", requestHandler)
	// Prepare TCP and UDP servers
	if s.config.Server.TCP {
		tcpServer := &dns.Server{Addr: s.serverAddr(),
			Net:          "tcp",
			ReadTimeout:  s.config.ReqTimeout,
			WriteTimeout: s.config.ResTimeout}
		go s.start(tcpServer)
	}
	if s.config.Server.UDP {
		udpServer := &dns.Server{Addr: s.serverAddr(),
			Net:          "udp",
			UDPSize:      65535,
			ReadTimeout:  s.config.ReqTimeout,
			WriteTimeout: s.config.ResTimeout}
		go s.start(udpServer)
	}
}

func (s *server) start(srv *dns.Server) {
	log.Printf("Start %s listener on %s\n", srv.Net, s.serverAddr())
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Start %s listener on %s failed:%s", srv.Net, s.serverAddr(), err.Error())
	}
}

func main() {
	// Get configruation file location
	var configFile string
	flag.StringVar(&configFile, "config", "./config.yaml",
		"Configuration file in YAML format")
	flag.Parse()
	// Load configuration file
	var config Config
	config.readConfig(configFile)

	server := new(server)
	server.config = &config

	// Start serving
	server.Serve()
	log.Println("DNS to DNS-over-TLS proxy started")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	// Wait for a syscall to exit
waiting:
	for {
		select {
		case <-sig:
			log.Println("Interrupt signal received, stopping")
			break waiting
		}
	}
}
