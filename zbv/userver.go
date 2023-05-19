package main

import (
  "fmt"
  "errors"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

type UServer struct{
  Name string
  ServerId string
  Email string //incase admin wants to make coomunication to the admin of the server in question
  Active bool
  CreatedAt string
  UpdatedAt string
}

// name 	email 	serverid 	created_at 	updated_at
// create a server that should be asigned an api key
func CreateServer(s UServer)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `zbv`.`servers` (name,serverid,email,active,created_at,updated_at) VALUES(?,?,?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert server.\nERROR %s",err))
    Logerror(e)
    return errors.New("Server encountered an error while preparing to create server. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&s.Name,&s.ServerId,&s.Email,&s.Active,&s.CreatedAt,&s.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error inserting server.\nERROR: %s",err))
    Logerror(e)
    return fmt.Errorf("Server encountered an error while creating server. \nERROR: %v",err)
  }
  return nil
}

func UpddateServer(s UServer) error{
  return nil
}

//return the requested server from the db
func GetServer(sid string)(*UServer,error){
  var s UServer
  row := db.QueryRow("SELECT * FROM `zbv`.`servers` WHERE `serverid` = ?",sid)
  err := row.Scan(&s.Name,&s.ServerId,&s.Email,&s.Active,&s.CreatedAt,&s.UpdatedAt)
  if err != nil{
    if err == sql.ErrNoRows {
      e := LogErrorToFile("danger",fmt.Sprintf("Non existatnt server id: %s\nERROR: ",sid,err))
      Logerror(e)
      return nil,errors.New("No server with such a key.")
    }
    e := LogErrorToFile("sql",fmt.Sprintf("Error Scanning for server %s.\nERROR %s",sid,err))
    Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing serverwith id %s",sid))
  }
  return &s,nil
}

// get all servers from DB, Active or inactive
func GetServers(active bool)([]UServer,error){
  stmt := "SELECT * FROM `zbv`.`servers` WHERE  `active` = ? ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,active)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error querying for servers.\nERROR: %s",err))
    Logerror(e)
    return nil,errors.New("Server encountered an error while listing all servers.")
  }
  defer rows.Close()
  var srvrs []UServer
  for rows.Next(){
    var s UServer
    err = rows.Scan(&s.Name,&s.ServerId,&s.Email,&s.Active,&s.CreatedAt,&s.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scaning for servers.\nERROR: %s",err))
      Logerror(e)
      return nil,errors.New("Server encountered an error while listing servers.")
    }
    srvrs = append(srvrs,s)
  }
  return srvrs,nil
}
