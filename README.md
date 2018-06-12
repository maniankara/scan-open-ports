# Scan open ports on a host
A simple golang application which dials through all the ports (concurrently) via telnet to list the open ports.
The concurrency is limited by a throttle which is derived from the number of CPUs of the machine where this application is launched.


### Usage
```
$ go run scan.go [<host>|-h|--help]
```
or 
```
$ go build scan.go 
$ ./scan [<host>|-h|--help]
```

### Help
```
$ ./scan -h

	Usage:
	go run scan.go [<hostname>|-h|--help]
	./scan [<hostname>|-h|--help]

	E.g.
	./scan # Scans localhost for all open ports

	./scan server1.zyx.com # Scan host server1.zyx.com for all open ports


exit status 255
```

### Test and compile
```
$ make 
go test -v ./...

?   	_/Users/maniankara/Work/github/scan-open-ports	[no test files]
go build -o scan -v
_/Users/maniankara/Work/github/scan-open-ports
```


### Results
```
$ ./scan
Scanning localhost ports:32000/32000
Open ports:  localhost : []

$ ./scan intra-server.mydomain.com
Scanning intra-server.mydomain.com ports:32000/32000
Open ports:  intra-server.mydomain.com : [80 443]

$ ./scan goole.com
Scanning google.com ports:32000/32000
Open ports:  google.com : [80 443]

$ ./scan yahoo.com
 Scanning yahoo.com ports:32000/32000
Open ports:  yahoo.com : [80 443]

$ ./scan facebook.com
Scanning facebook.com ports:32000/32000
Open ports:  facebook.com : [80 443 843]

```

#### Simple speed tests

```
$ time ./scan
Scanning localhost ports:32000/32000
Open ports:  localhost : []

real	1m14.215s
user	0m2.142s
sys	0m2.959s
$ time ./scan intra-server.domain.com
Scanning intra-server.domain.com ports:32000/32000
Open ports:  intra-server.domain.com : [80 443]

real	0m6.725s
user	0m2.244s
sys	0m3.004s
$ time ./scan google.com
Scanning google.com ports:32000/32000
Open ports:  google.com : [80 443]

real	1m21.749s
user	0m2.054s
sys	0m2.964s
$ time ./scan yahoo.com
Scanning yahoo.com ports:32000/32000
Open ports:  yahoo.com : [80 443]

real	1m21.540s
user	0m2.469s
sys	0m3.135s
$ time ./scan facebook.com
Scanning facebook.com ports:32000/32000
Open ports:  facebook.com : [80 443 843]

real	1m21.235s
user	0m2.244s
sys	0m3.306s
```
One can tune the speeds by modifying the `limit` variable