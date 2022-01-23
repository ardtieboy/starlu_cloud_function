package starlucloudfunction

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ardtieboy/starlu/imageprocessing"
)

func BorderImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	outputFilename := r.FormValue("myFileName")

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("/tmp", fmt.Sprintf("pragafied-*%s", filepath.Ext(handler.Filename)))
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
	bytes_written, err := tempFile.Write(fileBytes)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Failure saving tempfile: %s!", err.Error())
	}

	fmt.Printf("Tempfile is %s and contains %d bytes", tempFile.Name(), bytes_written)

	conv, err := imageprocessing.Crop(tempFile.Name())
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Failure due to %s!", err.Error())
	} else {
		dat, err := os.ReadFile(conv)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Failure exporting %s!", err.Error())
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		if outputFilename == "" {
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(conv)))
		} else {
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-pragafied%s", outputFilename, filepath.Ext(conv)))
		}
		w.Write(dat)
	}

}
