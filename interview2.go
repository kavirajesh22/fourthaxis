package main
import(
	"io/ioutil"
    "log"
    "net/http"
)

func main() {
	filePath := "C:\\Users\\rkavi\\hello.txt"
	resp, err := http.Get(filePath)
	if err != nil{
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln(err)
	}

	sb := string(body)
	log.printf(sb)
}