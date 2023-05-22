package main

/*
   This package contains global variables and any plausible helper function
*/
import (
  "fmt"
  "time"
  "errors"
  "database/sql"
  "html/template"

  "github.com/gorilla/sessions"
  _ "github.com/go-sql-driver/mysql"
)

//Runtime variables for the server
var db *sql.DB
var now = time.Now()
var currentTime = now.Format("2006-01-02 15:04:05")
var tpl *template.Template
// Generate a random mnonce to be used during cookie creation for sessions
var store = sessions.NewCookieStore([]byte("ZBV" + TrueRand(20)))
// Universal errors to keep track of and log requests if nescescary
var ErrorNonExistantKey = errors.New("Api key doen't exist")
var ErrorNonExistantEmail = errors.New("user with supplied email does not exist.")
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

// defines an admin user for the UI
type User struct{
  Name string
  Email string
  Password string
  Admin bool
  CreatedAt string
  UpdatedAt string
}

//name 	email 	password 	active 	created_at 	updated_at
// writes and admin into the db
func CreateUser(u User)error{
  passHash,err := HashPassword(u.Password)
  if err != nil{
    return fmt.Errorf("Error hashing password./n ERROR: %s")
  }
  var ins *sql.Stmt
  ins,err = db.Prepare("INSERT INTO `zbv`.`users` (name,email,password,admin,created_at,updated_at) VALUES(?,?,?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert user.\nERROR %s",err))
    Logerror(e)
    return errors.New("Server encountered an error while preparing to create user. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&u.Name,&u.Email,passHash,&u.Admin,&u.CreatedAt,&u.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error inserting user.\nERROR: %s",err))
    Logerror(e)
    return fmt.Errorf("Server encountered an error while creating user. \nERROR: %v",err)
  }
  return nil
}

//  Authenticates an admin through the UI
// Check the supplied username and password if they match else returns an error
func Authenticate(mail,password string) error{
  var userEmail,hash string
  stmt := "SELECT email,password FROM `zbv`.`users` WHERE email = ?;"
  row := db.QueryRow(stmt,mail)
  err := row.Scan(&userEmail,&hash)
  if err != nil{
    if err == sql.ErrNoRows {
      e := LogErrorToFile("danger",fmt.Sprintf("A none existant user with email %s tried accessing the admin portal",mail))
      Logerror(e)
      return ErrorNonExistantEmail
    }
    e := LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for authentication.\nERROR: %s",err))
    Logerror(e)
    return errors.New("Wrong username or password supplied. Try again :).")
  }
  err = CheckPasswordHash(password,hash)
  if err != nil{
    return errors.New("Wrong username or password supplied. Try again :).")
  }
  return nil
}
