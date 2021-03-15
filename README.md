# Just check your Network speed over http

# Run check
## On Serverside
```
git clone https://github.com/grischdian/upload-speed.git
go run main.go
```
### You may use correct SSL Certs
Just put your Cert into `./cert/cert.crt` and key into `./cert/cert.key` and your are fine

### Known Problem:
You need `root` privileges to open Port 443. Or edit the Port in main.go before start

```
http.ListenAndServeTLS(":443", "cert/cert.crt", "cert/cert.key", nil)

```

## On Clientside
```
 curl -X PUT -k -F "myFile=@$testfile" https://$servername/upload
```
