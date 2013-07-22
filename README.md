Introduction
================

Very simple (Thanks to Go!) SSL terminator proxy.
Use it when you need to terminate SSL sessions transparently as a MITM interceptor to read the plain raw data.


Usage
================

1. run some http server on port 80
2. go run go-sslterminator
3. curl -k https://localhost:443/


Help
================

go run go-sslterminator --help


License
================

Licensed under the New BSD License.


Author
================

Uri Shamay (shamayuri@gmail.com)
