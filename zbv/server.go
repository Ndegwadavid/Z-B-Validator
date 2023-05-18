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

func Getapikey(res http.ResponseWriter, req *http.Request){
  if req.Method != "POST"{
    //write an invalid request back
    tpl.ExecuteTemplate(res,"blank.html",nil)
    return
  }
  //read from body and write to db
}

func RequestForapikey(res http.ResponseWriter, req *http.Request){
  if req.Method != "POST"{
    tpl.ExecuteTemplate(res,"blank.html",nil)
    return
  }
  //read from body and write to db
}


func Blank(res http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(res,"blank.html",nil)
  return
}

func Dashboard(res http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(res,"index.html",nil)
  return
}

func Login(res http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(res,"login.html",nil)
  return
}

func Logout(res http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(res,"login.html",nil)
  return
}

func Validatekeyrequest(res http.ResponseWriter, req *http.Request){}

func Listapikeys(res http.ResponseWriter, req *http.Request){}

func Listservers(res http.ResponseWriter, req *http.Request){}

func Createadmin(res http.ResponseWriter, req *http.Request){}

func Changepassword(res http.ResponseWriter, req *http.Request){}
