package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
*Consommer une API
*/

func main(){
  url :="https://wikipedia.org"

  start_time :=time.Now()
  resp,err :=http.Get(url)
  end_time :=time.Since(start_time)

  if err != nil{
    fmt.Println(err)
    return
  }
  body,err := ioutil.ReadAll(resp.Body)
  if err !=nil{
	  fmt.Println(err)
	  return
  }
  defer resp.Body.Close()
  fmt.Print("%Site %s length %d Processing time %s \n",url, len(body), end_time)
}