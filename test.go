package main

import (
	//"github.com/kayz666/api-app/model"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)


func main(){
	//model.DB_init()
	//re,err:= model.RegUserOfEmail("761218@qq.com","123123123")
	//
	////ss:= model.MakePasswd("762370895")
	//fmt.Println(re)
	//fmt.Println(err)
	client := &http.Client{}
	url := "http://localhost:2345/v1/user"
	req,err := http.NewRequest("POST",url,strings.NewReader("email=60900111112140@qq.com"))
	if err!=nil{return }
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response,_ := client.Do(req)
	defer response.Body.Close()
	body,err := ioutil.ReadAll(response.Body)
	if err !=nil {return }
	fmt.Println(string(body))


}