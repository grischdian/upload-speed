# Just check your Network speed over http
HTTPS still missing

# Run check
## On Serverside
```
go run main.go
```
## On Clientside
```
 curl -X PUT -F "myFile=@$testfile" $servername:8080/upload
```
