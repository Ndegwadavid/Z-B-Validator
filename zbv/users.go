package main

type UServer struct{
  Name string
  ServerId string
  Email string //incase admin wants to make coomunication to the admin of the server in question
  CreatedAt string
  UpdatedAt string
}

func CreateServer(s UServer)error{
  return nil
}

func UpddateServer(s UServer) error{
  return nil
}

func GetServer(id string)(*UServer,error){
  return nil,nil
}

func GetServers()([]UServer,error){
  return nil,nil
}
