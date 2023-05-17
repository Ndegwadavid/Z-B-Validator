package main

/*
  This package contains all the handlers for the email validator service as declared in router
*/

import(
  "net/http"
)

func Requestkey(res http.ResponseWriter, req *http.Request){
}

func Validatekey(res http.ResponseWriter, req *http.Request){}


func Validatemail(res http.ResponseWriter, req *http.Request){}

func Updateapikey(res http.ResponseWriter, req *http.Request){}


func Getapikey(res http.ResponseWriter, req *http.Request){}


func Dashboard(res http.ResponseWriter, req *http.Request){}

func Login(res http.ResponseWriter, req *http.Request){}

func Logout(res http.ResponseWriter, req *http.Request){}

func Validatekeyrequest(res http.ResponseWriter, req *http.Request){}

func Listapikeys(res http.ResponseWriter, req *http.Request){}

func Listservers(res http.ResponseWriter, req *http.Request){}

func Createadmin(res http.ResponseWriter, req *http.Request){}

func Changepassword(res http.ResponseWriter, req *http.Request){}
