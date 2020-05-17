package gitea

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bounteous/artixweb-packages/tools"
)

// Request http
func Request(path string) []byte {
	url := tools.Config.GITEA.ReposURL + path
	fmt.Println("Requesting -> GET %s\n", url)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Request error: %s\n", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
