package main

import (
	"github.com/miekg/dns"
	"log"
)

func handleQuerie(w dns.ResponseWriter, r *dns.Msg) {
	q := r.Question[0]
	log.Printf("got question: %v\n", q)

	m := new(dns.Msg)
	m.SetReply(r)

	switch q.Qtype {
	case dns.TypeA:
		log.Println("A record !!")
	default:
		log.Println("Uhandled qtype")
	}

}

func main() {
	dns.HandleFunc(".", handleQuerie)

	server := &dns.Server{Addr: ":8053", Net: "udp"}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
