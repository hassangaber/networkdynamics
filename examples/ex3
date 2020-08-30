## Embedded file utility and uses

The embeddedfile.go file aims to serve a custom file written by the programmer to the server. The file
is written inside the go file and must be specified.

It is served based on the directory specified in the main function, in this case the html file will be
served upon a '/hello.html' request from the client. 

This is just an example, and I encourage the users of GoServe edit it by example and host the files they need
with the directories needed. I promote this naming scheme for the function so you do not get confused with the
directory names.

```go
func _CustomeFileN(...)
```
Where N is a positive integer, it may also help to comment the directory when writing the function out.

```go
http.HandleFunc("/example.html", _CustomFileN) 
//example.html maps to the _CustomFileN function.
```

This is how to run it

```
$ cd networkdynamics/server
```

```
server$ go run embeddedfile.go
```
