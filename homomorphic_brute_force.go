package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func main() {

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	cookie := "aWc2UWpSZE0zSVBNMmNaaVpjc0xGR08xdy9ack1kVlIxMDd3MFVmdkNTdEVjV0lNT1Vya0xmL0RyS2QwQzZVTVltalVPTmpsemZJS0lzMUo2OHlRRGFkNEYra3dvT2hsYXZpL2pHbzhRVkpXUXdHc1AvaVBPOE1ncksxcThsR3o="
	decodedCookie, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	//PostData := strings.NewReader("useId=5&age=12")

	bar := progressbar.Default(int64(len(alphabet) * (len(decodedCookie))))

	for index := 0; index < len(cookie)-1; index++ {
		for _, j := range alphabet {
			newCookie := append(decodedCookie[:index], byte(j))
			newCookie = append(newCookie, decodedCookie[index+1:]...)
			finalCookie := base64.URLEncoding.EncodeToString([]byte(newCookie))
			send(string(finalCookie))
			bar.Add(1)
		}
	}
}

func send(cookie string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://mercury.picoctf.net:10868/search", nil)
	req.Header.Set("Cookie", "auth_name="+cookie)
	resp, err := client.Do(req)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}
	if strings.Contains(string(data), "pico") {
		fmt.Printf("New Cookie: %s\n", cookie)
		fmt.Printf("Response = %s", string(data))
		os.Exit(0)
	}
}
