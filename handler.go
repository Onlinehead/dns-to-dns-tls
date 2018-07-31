package main

import (
	"log"
	"time"

	"github.com/miekg/dns"
)

// RequestHandlerConfig configuration for handler
type RequestHandlerConfig struct {
	upstreams  []string
	resTimeout time.Duration
	reqTimeout time.Duration
	useHTTPS   bool
}

// MakeHandler return a handler with a proper setup
func MakeHandler(config *RequestHandlerConfig) func(w dns.ResponseWriter, r *dns.Msg) {
	return func(w dns.ResponseWriter, r *dns.Msg) {
		config := config
		var net string
		if config.useHTTPS {
			net = "https"
		} else {
			net = "tcp-tls"
		}
		clientTLS := dns.Client{
			Net:          net,
			ReadTimeout:  config.resTimeout,
			WriteTimeout: config.reqTimeout,
		}
		request := new(dns.Msg)
		response := new(dns.Msg)
		response.SetReply(r)
		request.Question = r.Question
		upstreamsNum := len(config.upstreams)
		for i, upstream := range config.upstreams {
			resp, rtt, err := clientTLS.Exchange(request, upstream)
			if err != nil {
				log.Printf("Cannot process a request with err: %s", err)
				if i < upstreamsNum {
					continue
				} else {
					return
				}
			}
			response.Answer = resp.Answer
			response.Ns = resp.Ns
			response.Extra = resp.Extra
			log.Printf("Request rtt to %s = %s", upstream, rtt)
			break
		}
		w.WriteMsg(response)
	}
}
