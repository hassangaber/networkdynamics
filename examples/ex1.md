## Static File server example and utility 

Open a terminal and direct into the GoServer repo, specfically the Server file then run statfile.go

```
$ cd GoServe/server
```
```
server $ go run statfile.go
```
Accept any requests your network manager may ask you.

After that, open http://localhost:8080 in your web-browser. You should get a 404 error, that is fine. 
To serve a static file in the SAME FOLDER as the statfile.go file, go to http://localhost:8080/[FILENAME]
What this means is, when you run the go file, the file you want to host should be in the same folder as the go file.

Here, we are using a simple html file, so I would do http://localhost:8080/index.html
If I then visit that address, I should have it displayed!
