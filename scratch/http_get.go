package main

import (
   "fmt"
   "io/ioutil"
   "net/http"
   "time"
)

func count() {  //add a timer to get a refresh of the page
	for i := 0; i < 30; i++ {
	  fmt.Println(i)
	  time.Sleep(time.Second * 1)
	}
}

func main() {  //scrap a URL  TODo: add URL as var
	count()   //to run concurrently add go prior to fnction
	resp, _ := http.Get("http://www.securitytracker.com")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()

}