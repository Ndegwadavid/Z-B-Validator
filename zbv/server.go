package main

/*
  This package contains all the handlers for the email validator service as declared in router
*/

import(
  "fmt"
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

// Blak page for testing
func Blank(res http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(res,"blank.html",nil)
  return
}

// Admin UI Dashboard
// add validation to ensure user has to be logged in
func Dashboard(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  _,ok := session.Values["Admin"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)
  }
  tpl.ExecuteTemplate(res,"index.html",nil)
  return
}

//log into admin panel UI and start a session for the user
func Login(res http.ResponseWriter, req *http.Request){
  if req.Method != "POST"{
    tpl.ExecuteTemplate(res,"login.html",nil)
    return
  }
  req.ParseForm()
  pass := req.FormValue("password")
  mail := req.FormValue("email")
  if err := Authenticate(mail,pass); err != nil{
    fmt.Println("[-]  ERROR: ",err)
    tpl.ExecuteTemplate(res,"login.html","Wrong username or password provided. Try again :)")
    return
  }
  session,_ := store.Get(req,"session")
  session.Values["Admin"] = "admin"
  session.Save(req,res)
  http.Redirect(res,req,"/",http.StatusSeeOther)
  return
}

// logout from admin panel and delete the user session
func Logout(res http.ResponseWriter, req *http.Request){
  if req.Method != "POST"{
    http.Redirect(res,req,"/",http.StatusSeeOther)
    return
  }
  session,_ := store.Get(req,"session")
  delete(session.Values,"Admin")
  session.Save(req,res)
  tpl.ExecuteTemplate(res,"login.html","Successfully logged out.")
  return
}

func Validatekeyrequest(res http.ResponseWriter, req *http.Request){}

func Listapikeys(res http.ResponseWriter, req *http.Request){}

func Listservers(res http.ResponseWriter, req *http.Request){}

func Createadmin(res http.ResponseWriter, req *http.Request){}

func Changepassword(res http.ResponseWriter, req *http.Request){}
