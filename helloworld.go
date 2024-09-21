package main

/*
HTTP is the protocol that is used to fetch data from web servers.
HTTP is text sent by the client to a server to request a particular action, and the server replies.
The following commmands are useful with curl:
--verbose
--trace-ascii

The following APIs are implemented in this pgrm:
- GET (Client issues a GET request to the server and receives the document asked for)
- POST (Puts file specified on system)
- DELETE (Remove file off of system)
- PUT * to be done*

*/
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func xget(w http.ResponseWriter, r *http.Request){
	path:= r.URL.Path
	parts := strings.Split(path, "/")

	filename := parts[2]
	fmt.Printf(filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil{
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func xdelete(w http.ResponseWriter, r *http.Request){
	path:= r.URL.Path
	parts:= strings.Split(path, "/")

	filename := parts[2]
	e := os.Remove(filename)
	if e != nil{
		return
	}
}

func xpost(w http.ResponseWriter, r *http.Request){
	path:= r.URL.Path
	parts := strings.Split(path, "/")
	filename := parts[2]
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")

	if err != nil{
		return
	}

	defer file.Close()
	outFile, err := os.Create(filename)
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	w.WriteHeader(http.StatusCreated)
	fmt.Printf("Uploaded the file: %s", handler.Filename)
}


func main(){
	http.HandleFunc("/rkavi/", xdelete)
	fmt.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}
