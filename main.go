package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "math"
    "time"
    "os"
    "path/filepath"
)

var (
	sizeInMB float64 = 999 // This is in megabytes
	suffixes [5]string
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    uploadBegin := time.Now()
    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    uploadEnd := time.Now()
    uploadDuration := uploadEnd.Sub(uploadBegin)

    fmt.Printf("Uploaded File: %+v\n", handler.Filename)

    base := math.Log(float64(handler.Size))/math.Log(1000)
    getSize := Round(math.Pow(1000, base - math.Floor(base)), .5, 2)
    getSuffix := suffixes[int(math.Floor(base))]

    var filesize string = strconv.FormatFloat(getSize, 'f', -1, 64)+" "+string(getSuffix)

    fmt.Printf("File Size: %+v\n", filesize)
    fmt.Printf("Upload duration: %+v\n", uploadDuration.Round(time.Second))

    speed := ((float64(handler.Size)/float64(uploadDuration/time.Nanosecond))*8000) //1Byte/Nanosecond = 8000Mbit/Second
    fmt.Printf("Upload speed in Mbit/s: %.2f\n", speed)
    fmt.Println("")

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("uploaded-files", handler.Filename)
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    error := ClearDir("uploaded-files")
    if error != nil {
      fmt.Println(error)
    }
    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func ClearDir(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func Round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func setupRoutes() {
    fmt.Println("Server started")
    fmt.Println("Uploading Files:")
    fmt.Println("curl -X PUT -F \"myFile=@$FILENAME\" $SERVER:8080/upload")
    fmt.Println("")
    http.HandleFunc("/upload", uploadFile)
    http.ListenAndServe(":8080", nil)
}

func main() {
    suffixes[0] = "Byte"
    suffixes[1] = "KByte"
    suffixes[2] = "MByte"
    suffixes[3] = "GByte"
    suffixes[4] = "TByte"
    setupRoutes()
}

