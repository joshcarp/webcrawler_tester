## Webscraper tester


There are two tests; the "Index of Sample Test-Cases" seen internally and "Custom"

# Main tests; This is a copy of "Index of Sample Test-Cases" that we use on our vms
## Setup
Edit the hostnames to include our hostnames that we'll be visiting
`sudo nano /etc/hosts`

and then add the following lines:
```
127.0.0.1 web1.comp30023
127.0.0.1 web2.comp30023
```
now you can run the program:

```bash
go run main.go
```

And you'll have a debug log of the endpoints your program hit (add -v to enable printing of the request):
```

Visited URL: /mirror/a.html
Visited URL: /
Visited URL: /robots.txt
Visited URL: /mirror//a.html
Visited URL: /1-neighbour//b.html
Visited URL: /shallow-neighbours//b.html
Visited URL: /simple-cycle//b.html
Visited URL: /triangle//b.html
Visited URL: /simple-branch-duplicate//b.html
Visited URL: /shallow-neighbours//c.html
Visited URL: /simple-cycle//a.html
Visited URL: /triangle//c.html
Visited URL: /simple-branch-duplicate//c.html
Visited URL: /shallow-neighbours//d.html
Visited URL: /triangle//a.html
```




# custom/ directory
`go run custom/main.go` 

`curl localhost/1`

see main.go to see what every endpoint returns

The terminal will output a list of endpoints that you've visited:
```bash
/1
GET http://localhost/2 HTTP/1.1
Accept: */*
Proxy-Connection: Keep-Alive
User-Agent: jcarpeggiani


/2
GET http://localhost/3 HTTP/1.1
Accept: */*
Proxy-Connection: Keep-Alive
User-Agent: jcarpeggiani


/3
GET http://localhost/4 HTTP/1.1
Accept: */*
Proxy-Connection: Keep-Alive
User-Agent: jcarpeggiani


/4
GET http://localhost/1 HTTP/1.1
Accept: */*
Proxy-Connection: Keep-Alive
User-Agent: jcarpeggiani

```