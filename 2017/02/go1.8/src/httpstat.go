package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

// START OMIT
func main() {
	req, _ := http.NewRequest("GET", "https://golang.org", nil)

	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result) // 🙆‍
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, _ := client.Do(req)

	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	result.End(time.Now())

	fmt.Printf("%+v\n", result)
}

// END OMIT
