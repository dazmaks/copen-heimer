package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Changing log output
	logfile, err := os.Create("output.log")
	ErrorHandler(err)
	defer logfile.Close()

	log.SetOutput(logfile)
	log.Print("Info [main]: Setted log output to output.log")

	var delay int
	var masscan string
	var database string

	log.Print("Info [main]: Getting a delay")
	fmt.Print("Input a delay after every IP check (35 is recommended): ")
	if _, err := fmt.Scan(&delay); err != nil {
		ErrorHandler(err)
	}

	log.Print("Info [main]: Getting a masscan filename")
	fmt.Print("Input a masscan filename: ")
	if _, err := fmt.Scan(&masscan); err != nil {
		ErrorHandler(err)
	}

	log.Print("Info [main]: Getting a database filename")
	fmt.Print("Input a database filename (.db extension): ")
	if _, err := fmt.Scan(&database); err != nil {
		ErrorHandler(err)
	}

	log.Printf("Info [database]: Opening %s\n", database)
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	ErrorHandler(err)

	// Migrate the schema
	db.AutoMigrate(&Server{})

	// Opening masscan file
	log.Printf("Info [main]: Opening %s\n", masscan)
	file, err := os.Open("masscan.txt")
	ErrorHandler(err)
	defer file.Close()

	// Scanning file line by line
	scanner := bufio.NewScanner(file)
	//time.Sleep(10 * time.Millisecond) // Sleep for database creation
	for scanner.Scan() {
		// Getting server IP
		ip := strings.Split(scanner.Text(), " ")[3]
		//log.Printf("Info [main]: Current ip is %s\n", ip)
		//go IPHandler(db, ip)
		go IPHandler(db, ip)
		time.Sleep(time.Duration(delay) * time.Millisecond)

	}
	if err := scanner.Err(); err != nil {
		ErrorHandler(err)
	}
}
