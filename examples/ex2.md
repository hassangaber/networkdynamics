## Utilization of listening tools

The main.go file serves a purpose of sending an http request to the server and if the connection is successful, it will return the string written in the file.
Upon returning the string it will also show you what specific URL you have directed to. This is only a listening procedure for network tests. 
You can change the string returned, or you can return something bigger by also using the statfile.go file, however, that is better used a hosting tool. 

Upon running main.go 

```
$ cd networkdynamics/server
```
```
server$ go run main.go
```
And then opening a web-browser and going to http://localhost:8081 you should get 
```
Hello, "/"
```
Displayed.

If there are no errors, then the server has a working TCP protocol setup with your computer.
The reason I am using port 81, it to avoid any exit errors with port 80 if the user decides to run both tests simultaneously. 
