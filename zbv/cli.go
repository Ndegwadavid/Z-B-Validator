package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	myFigure := figure.NewFigure("Z--B--Validator", "univers", true)
	myFigure.Print()
	go RunServer()
	time.Sleep(3 * time.Second)
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/zbv")
	if err != nil {
		fmt.Println("[-] Failed to connect to the email validator's db")
		log.Fatal(err.Error())
	}
	fmt.Println("CLI For Z-B-Validator............")
	var arg string
	for {
		fmt.Printf("[+ ZBV +] Enter command:   ")
		fmt.Scanln(&arg)
		switch arg {
		case "ca", "createaccount", "createadmin":
			var pass, name, email string
			fmt.Printf("[+] Enter username: ")
			fmt.Scanln(&name)
			fmt.Printf("[+]  Enter email address: ")
			fmt.Scanln(&email)
			fmt.Printf("[+]  Enter password: ")
			fmt.Scanln(&pass)
			u := User{
				Name:      name,
				Email:     email,
				Password:  pass,
				Admin:     true,
				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			}
			if err := CreateUser(u); err != nil {
				fmt.Println("[-]  %s", err)
				continue
			}
			fmt.Println("[+]  Successfully created user. Login at 127.0.0.1:3000/login")
			fmt.Println("")
		case "cs", "create-server":
			var name, email string
			fmt.Printf("[+] Enter server name:  ")
			fmt.Scanln(&name)
			fmt.Printf("[+] Enter server email: ")
			fmt.Scanln(&email)
			s := UServer{
				ServerId:  RandString(10),
				Name:      name,
				Email:     email,
				Active:    false,
				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			}
			if err := CreateServer(s); err != nil {
				fmt.Println("[-]  ERROR: ", err)
				continue
			}
			fmt.Println("[+]  Successfully created server.")
			fmt.Println("")
		case "gk", "ak":
			srvrs, err := GetServers(false)
			if err != nil {
				fmt.Println("[-]  ERROR: ", err)
				continue
			}
			for _, s := range srvrs {
				fmt.Println("------------------------------------------------")
				fmt.Printf("Name: ", s.Name)
				fmt.Printf("    Email: ", s.Email)
				fmt.Println("")
				fmt.Println("   Server Id:    ", s.ServerId)
				fmt.Println("")
			}
			var srvId, ys string
			var accept bool
			fmt.Printf("[+] Enter server Id to validate: ")
			fmt.Scanln(&srvId)
			fmt.Printf("[+] Accept or deny. (enter YES or No):  ")
			fmt.Scanln(&ys)
			if ys == "YES" {
				accept = true
			} else {
				accept = false
			}
			api := ApiKey{
				ServerID:  srvId,
				Key:       RandString(10),
				Comment:   "Just some random server",
				Active:    accept,
				CreatedAt: currentTime,
				UpdatedAt: currentTime,
			}
			if err := CreateApiKey(api); err != nil {
				fmt.Println("[-]  ERROR: ", err)
				continue
			}
			fmt.Println("Successfully created api key.")
		case "validate-email", "ve":
			var email string
			fmt.Printf("Enter email to validate: ")
			fmt.Scanln(&email)
			if err := EmailValidate(email); err != nil {
				fmt.Println("[-]    Error validating your email.")
				fmt.Println("[-]    ERROR: ", err)
				fmt.Println("")
				continue
			}
			fmt.Printf("[+]  %s is a valid email.", email)
			fmt.Println("")
		case "validate-emails", "vas":
			var fname string
			fmt.Printf("[+] Enter file name: ")
			fmt.Scanln(&fname)
			emails, err := GetEmailsFromFile(fname)
			if err != nil {
				fmt.Printf("[-]    %s\n", err)
				continue
			}
			var valid []string
			for _, email := range emails {
				if !CheckifStringIsEmpty(email) {
					continue
				}
				if err := EmailValidate(email); err != nil {
					fmt.Printf("[-]    %s: %s\n", err, email)
					continue
				}
				valid = append(valid, email)
			}
			for _, v := range valid {
				fmt.Println("[+]    This is a valid email: ", v)
			}
			fmt.Println("")
		case "help":
			fmt.Println("Help message commming soon")
			fmt.Println("")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Whyyyy.............")
			fmt.Println("")
		}
	}
}
