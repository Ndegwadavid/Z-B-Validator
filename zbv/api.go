package main

import (
  "fmt"
  "errors"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

type ApiKey struct {
  ServerID string
  Key string
  Comment string // can be anything but mostly a description of the server and what it does
  Active bool
  CreatedAt string
  UpdatedAt string
}

var GenerateApiKey = func()string{
  //should go into the db and validate that theres no such key already in existance
  //generate key
  //get all api keys
  // ensure key matches none
  // use a mutex to make the checkking faster
  return ""
}

 	//serverid 	key 	comment 	active 	created_at 	updated_at
func CreateApiKey(a ApiKey)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `zbv`.`apikey` (serverid,key,comment,active,created_at,updated_at) VALUES(?,?,?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert wpi key.\nERROR %s",err))
    Logerror(e)
    return errors.New("Server encountered an error while preparing to create apikey. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&a.ServerID,&a.Key,&a.Comment,&a.Active,&a.CreatedAt,&a.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error inserting api key.\nERROR: %s",err))
    Logerror(e)
    return fmt.Errorf("Server encountered an error while creating API Key. \nERROR: %v",err)
  }
  return nil
}

// mark a key as inactive incase server is compromised or something
func InvalidateKey(srvId,key string, active bool) error{
  upStmt := "UPDATE `zbv`.`apikey` SET (`active` = ? AND `updated_at` = ?) WHERE (`serverid` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to invalidate/validate key.\nERROR: %s",err))
    Logerror(e)
    return errors.New("Server encountered an error while preparing to invalidate/validate API Key.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(key,currentTime,srvId)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := LogErrorToFile("sql",fmt.Sprintf("Error Executing invalidate/validate key.\nERROR %s",err))
    Logerror(e)
    return errors.New("Server encountered an error while executing invalidate/validate apikey.")
  }
  return nil
}

func UpdateKey(srvId, key string)(string,error){
  upStmt := "UPDATE `zbv`.`apikey` SET (`apikey` = ? AND `updated_at` = ?) WHERE (`serverid` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to update key.\nERROR: %s",err))
    Logerror(e)
    return "",errors.New("Server encountered an error while preparing to update API Key.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(key,currentTime,srvId)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := LogErrorToFile("sql",fmt.Sprintf("Error Executing update key.\nERROR %s",err))
    Logerror(e)
    return "",errors.New("Server encountered an error while executing update apikey.")
  }
  return key,nil
}

func GetApiKey(srvId,key string) (*ApiKey,error){
  var a ApiKey
  row := db.QueryRow("SELECT * FROM `zbv`.`apikey` WHERE `key` = ? AND `serverid` = ?;",key,srvId)
  err := row.Scan(&a.ServerID,&a.Key,&a.Comment,&a.Active,&a.CreatedAt,&a.UpdatedAt)
  if err != nil{
    if err == sql.ErrNoRows {
      e := LogErrorToFile("danger",fmt.Sprintf("Non existatnt key: %s",err))
      Logerror(e)
      return nil,ErrorNonExistantKey
    }
    e := LogErrorToFile("sql",fmt.Sprintf("Error Scanning for api key %s.\nERROR %s",key,err))
    Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing apikey of %s",key))
  }
  return &a,nil
}

func GetApiKeys(active bool)([]ApiKey,error){
  stmt := "SELECT * FROM `zbv`.`apikey` WHERE (`active` = ? ) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,active)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error querying for api keys.\nERROR: %s",err))
    Logerror(e)
    return nil,errors.New("Server encountered an error while listing all api keys.")
  }
  defer rows.Close()
  var keys []ApiKey
  for rows.Next(){
    var a ApiKey
    err = rows.Scan(&a.ServerID,&a.Key,&a.Comment,&a.Active,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scaning for api keys.\nERROR: %s",err))
      Logerror(e)
      return nil,errors.New("Server encountered an error while listing apikeys.")
    }
    keys = append(keys,a)
  }
  return keys,nil
}
