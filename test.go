package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp, _ := http.Get("https://www.random.org/integers/?num=32&min=1&max=255&col=1&base=10&format=plain&rnd=new")
	body,_ :=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}