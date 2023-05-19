package main

/*
  This package contains helper functions for running the server or do small tasks
*/

import(
  "os"
  "io"
  "fmt"
  "log"
  "time"
  "bytes"
  "strings"
  "io/ioutil"
  "math/rand"
  "crypto/sha256"
  "golang.org/x/crypto/bcrypt"
)

//takes in a password and hashes it into a an encrypted string
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    return string(bytes), err
}

// Takes in a password and a given hash then compares them if they rhyme.
func CheckPasswordHash(password, hash string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// hash a whole struct datatype(FOr generating api key)
func HashStruct(data interface{})string{
  h := sha256.New()
  s := fmt.Sprintf("%v",data)
  sum := h.Sum([]byte(s))
  return string(sum)
}

// Opens a file, reads it and returns the lines in a string array
func GetEmailsFromFile(fileName string)([]string,error){
  var emails []string
  buf,err := ioutil.ReadFile(fileName)
  if err != nil {
    return nil,err
  }
  lines := bytes.Split(buf,[]byte("\n"))
  for _,line := range lines {
    emails = append(emails,string(line))
  }
  return emails,nil
}

//Check if a string is empty returns True if string is a string
func CheckifStringIsEmpty(data string) bool{
  if len(strings.TrimSpace(data)) == 0{
    return false
  }
  if len(data) == 0{
    return false
  }
  return true
}

func TrueRand(len int) string{
  bytes := make([]byte,len)
  for i := 0; i < len; i++{
    bytes[i] = byte(randInt(97,122))
  }
  if !CheckifStringIsEmpty(string(bytes)){
    TrueRand(len)
  }
  return string(bytes)
}

func randInt(min int, max int) int {
  return min + rand.Intn(max-min)
}

func RandString(length int) string{
  var output strings.Builder
  rand.Seed(time.Now().Unix())
  charset := []rune("QWERTYUIOPLKJHGFDSAZXCVBNM123456789qwertyuioplkjhgfdsazxcvbnm")
  for i := 0; i < length; i++{
    random := rand.Intn(len(charset))
    randomChar := charset[random]
    output.WriteRune(randomChar)
  }
  id := output.String()
  id = strings.ToUpper(id)
  if !CheckifStringIsEmpty(id){
    RandString(length)
  }
  return id
}

//Retrns a random string with numbers and letters (caps on)
func RandNoLetter(length int) string{
  var output strings.Builder
  rand.Seed(time.Now().Unix())
  charset := []rune("QWERTYUIOPLKJHGFDSAZXCVBNM123456789")
  for i := 0; i < length; i++{
    random := rand.Intn(len(charset))
    randomChar := charset[random]
    output.WriteRune(randomChar)
  }
  id := output.String()
  id = strings.ToUpper(id)
  if !CheckifStringIsEmpty(id){
    RandNoLetter(length)
  }
  return id
}

//Returns A Random letters
func RandLetters(length int) string{
  var output strings.Builder
  rand.Seed(time.Now().Unix())
  charset := []rune("qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBBNM")
  for i := 0; i < length; i++{
    random := rand.Intn(len(charset))
    randomChar := charset[random]
    output.WriteRune(randomChar)
  }
  id := output.String()
  if !CheckifStringIsEmpty(id){
    RandLetters(length)
  }
  return id
}

//Returns a random number in string format
func RandNo(length int) string{
  var output strings.Builder
  rand.Seed(time.Now().Unix())
  charset := []rune("1234567890")
  for i := 0; i < length; i++{
    random := rand.Intn(len(charset))
    randomChar := charset[random]
    output.WriteRune(randomChar)
  }
  id := output.String()
  if !CheckifStringIsEmpty(id){
    RandNo(length)
  }
  return id
}

//Log and error to file Allows format string input
func LogErrorToFile(name string,text ...interface{}) error{
  name = "./.data/logs/"+name+".log"
  f,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
  if err != nil{
    return err
  }
  defer f.Close()
  writer := io.MultiWriter(os.Stdout,f)
  log.SetOutput(writer)
  log.Println(text)
  return nil
}

func Logerror(e error){
  if e != nil{
    log.Println(e)
  } else {
    log.Println("Can not log a nil error.")
  }
}
