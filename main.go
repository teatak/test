package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("hello test")
	// r := net.Resolver{}
	// ips, _ := r.LookupTXT(context.TODO(), "_dnsauth.dev.x-t.top")
	// fmt.Println(ips)
	certPEMBlock, err := os.ReadFile("/Users/yanggang/Documents/cert/teatak/teatak.pem")

	if err != nil {
		fmt.Println(err)
	}
	block, _ := pem.Decode(certPEMBlock)

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	data, _ := json.Marshal(cert)
	fmt.Println(string(data))

}
