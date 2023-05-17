package main

import (
  "fmt"
  "net"
  "strings"
  "errors"
  "regexp"
)
// check if the email's domain exist
func CheckEmailDomain(email string) error{
  dom := strings.Index(email,"@")
  host := email[dom+1:]
  fmt.Println("Looking up host: ",host)
  _,err := net.LookupMX(host)
  if err != nil{
    return fmt.Errorf("Could not find email's domanin server. %q",err)
  }
  /*if err := ValidateBySend(mxRecords,email); ree != nil{
    return err
  }*/
  return nil
}

// verify email syntax
func VerifyEmailSyntax(email string) error{
  emailRegex, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		return fmt.Errorf("Sorry, something went wrong while compiling regex.\n %qERROR: %q",err)
	}
	rg := emailRegex.MatchString(email)
	if rg != true {
		return errors.New("Email address is not a valid syntax, please check again")
	}
	// check email length
	if len(email) < 4 {
		return errors.New("Email length is too short")
	}
	if len(email) > 253 {
		return errors.New("Email length is too long")
	}
	return nil
}

func ValidateBySend(mxRecords []net.NS,email string)error{
  mailServer := mxRecords[0].Host
  conn,err := net.Dial("tcp",mailServer+":25")
  if err != nil {
    // update this to try reaching out to all returned mail servers
    return fmt.Errorf("Error connecting to the mail server.\nERROR: %q",err)
  }
  resp := make([]byte,1024)
  if _,err = conn.Read(resp); err != nil{
    return fmt.Errorf("Error reading from server response.\nERROR: %q",err)
  }
  cmd := "HELLO sbvalidator.com\r\n"
  conn.Write([]byte(cmd))
  if _,err = conn.Read(resp); err != nil{
    return fmt.Errorf("Error reading from server response.\nERROR: %q",err)
  }
  cmd = "MAIL FROM: validator@sbvalidator.com\r\n"
  conn.Write([]byte(cmd))
  if _,err = conn.Read(resp); err != nil{
    return fmt.Errorf("Error reading from server response.\nERROR: %q",err)
  }
  respCode := string(resp[:3])
  if respCode == "250"{
    return nil
  } else {
    return errors.New("Invalid response code. Mail probably doesn't exist.")
  }
  return nil
}
