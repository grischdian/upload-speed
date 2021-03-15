# Just check your Network speed over http

# Run check
## On Serverside
```
go run main.go
```
## On Clientside
```
 curl -X PUT -k -F "myFile=@$testfile" https://$servername/upload
```
