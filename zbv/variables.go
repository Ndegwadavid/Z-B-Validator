package main

/*
   This package contains global variables and any plausible helper function
*/
import (
  "time"
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
