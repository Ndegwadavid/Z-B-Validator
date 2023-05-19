package main

import (
  "sync"
)

type Runner struct{
  Servers int
  Requests int
  Perfomance int
  Reqs *Requests
}

type Requests struct{
  Servers []*UServer
}

type RMan struct{
  Rnr *Runner
  mu sync.Mutex
}

func InitializeRunner()*Runner{
  reqs,err := GetRequests()
  if err != nil{
    reqs = &Requests{}
  }
  return &Runner{
    Servers: GetServerNumber(),
    Requests: 0,
    Perfomance: 100,
    Reqs: reqs,
  }
}

func GetServerNumber() int{
  return 0
}

func GetRequests()(*Requests,error){
  return nil,nil
}
