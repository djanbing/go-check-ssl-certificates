package main

import (
	"flag"
	"fmt"

	certs "go-check-ssl-certificates/tools"
)

func main() {

	var (
		certFile      string
		addr          string
		timeoutSecond int
	)
	flag.StringVar(&certFile, "file", "", "The certificates.")
	flag.StringVar(&addr, "connect", "", "The remote addr. The format should be 'example.com:ssl_port'.")
	flag.IntVar(&timeoutSecond, "timeout", 10, "The timeout in sec.")
	flag.Parse()

	// certFile = "test_certificates/my-server.crt"
	// addr = "jd.com:443"
	// timeoutSecond = 1

	var (
		c   *certs.Cert
		err error
	)

	if addr != "" {
		c, err = certs.ParseRemoteCertificate(addr, timeoutSecond)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if certFile != "" {
		c, err = certs.ParseCertificateFile(certFile)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return
	}

	fmt.Println(c.Jsonify())
}
