## Webscraper tester

`go run main.go` 

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