package main

/*
  This package contains all defined routes rof the mail validator service
*/

import (
  "os"
  "fmt"
  "log"
  "net/http"
  "html/template"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)
func main(){
  var err error
  // connect to db
  _, err = sql.Open("mysql", "root:@tcp(localhost:3306)/zbv")
  if err != nil {
    fmt.Println("[-] Failed to connect to the email validator's db")
		log.Fatal(err.Error())
	}

  // load the templates
  tpl,err = template.ParseGlob("./templates/*.html")
  if err != nil{
    log.Println("[-] This is not good like: ",err)
  }

  //Define route
  // API ROUTES\
  http.HandleFunc("/requestkey",Requestkey)
  http.HandleFunc("/validatekey",Validatekey)
  http.HandleFunc("/validatemail",Validatemail) //core functionality not reqular validation
  http.HandleFunc("/updatekey",Updateapikey)
  http.HandleFunc("/getapikey",Getapikey)

  //ADMIN ROUTES
  http.HandleFunc("/admin",Dashboard)
  http.HandleFunc("/validatekeyrequest",Validatekeyrequest)
  http.HandleFunc("/listkeys",Listapikeys)
  http.HandleFunc("listservers",Listservers)
  http.HandleFunc("/login",Login)
  http.HandleFunc("/logout",Logout)
  http.HandleFunc("/createadmin",Createadmin)
  http.HandleFunc("/changepassword",Changepassword)

  // start server
  log.Println(" Starting Odin server at: %s",currentTime)
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/static/",http.StripPrefix("/static",fs))
  err = http.ListenAndServe("0.0.0.0:3000",nil)
  if err != nil {
    log.Fatal("[+] Error starting HTTP server: ",err)
  }
}


func init(){
  emails,err := GetEmailsFromFile("../emaillist.csv")//emaillist.csv
  if err != nil{
    log.Fatal(err)
  }
  var valid []string
  for _, email := range emails {
    if !CheckifStringIsEmpty(email){
      continue
    }
    err := CheckEmailDomain(email)
    if err != nil{
      fmt.Println("[-] Error checking email domain for email: ",email)
      fmt.Sprintf("[-]  %s",err)
      fmt.Println("");continue
    }
    err = VerifyEmailSyntax(email)
    if err != nil{
      fmt.Println("[-] Error veryfing email: ",email)
      fmt.Sprintf("[-]  %s",err)
      fmt.Println("");continue
    }
    valid = append(valid,email)
  }
  for _,v := range valid{
    fmt.Println("This is a valid email: ",v)
  }
  os.Exit(0)
}
