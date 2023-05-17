package main

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

func CreateApiKey(a ApiKey)error{
  return nil
}

// mark a key as inactive incase server is compromised or something
func InvalidateKey(srvId,key string, active bool) error{
  return nil
}

func UpdateKey(srvId, key string)error{
  return nil
}

func GetApiKey(srvId string) error{
  return nil
}

func GetApiKeys(active bool)([]ApiKey,error){
  return nil,nil
}
