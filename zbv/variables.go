package main

/*
   This package contains global variables and any plausible helper function
*/
import (
  "time"
  "errors"
  "database/sql"
  "html/template"
)

var (
  test = false
  UniversalKey = "loiuixghjpou98y7t6txcvbiuoiugyftcvbno98igtfxcfgvbioiuyft"//use this to encrypt strings/ids
)

var db *sql.DB
var now = time.Now()
var currentTime = now.Format("2006-01-02 15:04:05")
var tpl *template.Template
//var store = sessions.NewCookieStore([]byte("ZBV"))
var ErrorNonExistantKey = errors.New("Api key doen't exist")
type DateTime struct{
  Day int
  Month string
  Year int
}

func GetDateTime()(*DateTime){
  return &DateTime{
    Day:now.Day(),
    Month:now.Month().String(),
    Year:now.Year(),
  }
}

type User struct{
  Name string
  Email string
  Password string
  Admin bool
  CreatedAt string
  UpdatedAt string
}

func CreateUser(u User)error{
  return nil
}
func Authenticate(mail,password string) error{
  return nil
}
