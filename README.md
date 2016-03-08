Introduction
================

Very simple (Thanks to Go!) SSL terminator proxy.
Use it when you need to terminate SSL sessions transparently as a MITM interceptor to read the plain raw data.

Usage
================

1. run some http server on port 8000
2. go run go-sslterminator.go // assumed you have key.pem & cert.pem in `pwd`
3. curl -v -k https://localhost:44300/

If you want to do some quick test, you can use that commands:

1. create SSL key & certificate for go-sslterminator
	openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365 -nodes -subj "/C=GO/ST=Golang/L=Golang/O=Golang/OU=IT Department/CN=golang.org"
2. run dummy http server
	while true; do { echo -e 'HTTP/1.1 200 OK\r\n'; } | nc -l 8000; done
3. run go-sslterminator
	go run go-sslterminator.go
3. run dummy ssl client
	curl -v -k https://localhost:44300/

Help
================

	go run go-sslterminator.go --help
	  -b string
    		backend address (default ":8000")
  	  -c string
    		SSL certificate path (default "cert.pem")
  	  -k string
    		SSL key path (default "key.pem")
          -l string
    		local address (default ":44300")

License
================

Licensed under the New BSD License.

Author
================

Uri Shamay (shamayuri@gmail.com)
