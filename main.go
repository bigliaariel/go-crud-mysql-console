package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"

	_ "github.com/go-sql-driver/mysql"
)

var salir bool

func main() {
	//Go to the readme-mysql-freeservice.txt file to learn how to get a connection string from a free mysql service for testing
	db, err := sql.Open("mysql", "databaseUser:youpassword@tcp(sql10.freesqldatabase.com:3306)/databaseName")

	if err != nil {
		panic(err.Error())
	}

	// defer the close tigll after the main function has finished
	// executing
	defer db.Close()

	salir = false
	for {
		menu := `
		---------------------------
		1- Add contact
		2- Delete contact
		3- View all contacts
		4- Exit
		---------------------------
		> `
		fmt.Print(menu)
		var eleccion int
		fmt.Scanln(&eleccion)

		switch eleccion {
		case 1:
			// Add new contact to database
			var nombre string
			var telefono string
			fmt.Print("Name: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				nombre = scanner.Text()
			}
			fmt.Print("Phone: ")
			if scanner.Scan() {
				telefono = scanner.Text()
			}
			query := fmt.Sprintf("INSERT INTO test VALUES (null, '%s', '%s' )", nombre, telefono)
			insert, err := db.Query(query)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			// be careful deferring Queries if you are using transactions
			insert.Close()
		case 2:
			// Delete contact fron database by Id
			var id string
			fmt.Print("Id: ")
			fmt.Scanln(&id)
			query := fmt.Sprintf("DELETE FROM test WHERE id=%s", id)
			delete, err := db.Query(query)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			// be careful deferring Queries if you are using transactions
			delete.Close()
		case 3:
			// View all contact from console
			query := "SELECT id,name,phone FROM test ORDER BY name"
			rows, err := db.Query(query)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			// be careful deferring Queries if you are using transactions
			w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.TabIndent)
			fmt.Fprintf(w, "\n\n")
			fmt.Fprintln(w, "------------------------------------------------------------------")
			fmt.Fprintln(w, "   Id\tName\t\tPhone")
			fmt.Fprintln(w, "------------------------------------------------------------------")
			for rows.Next() {
				var id int
				var nombre string
				var telefono string
				rows.Scan(&id, &nombre, &telefono)
				fmt.Fprintf(w, "   %d\t%s\t\t%s\n", id, nombre, telefono)
				w.Flush()
			}
			fmt.Println("------------------------------------------------------------------")
			fmt.Fprintf(w, "\n\n")
			w.Flush()
			rows.Close()
		case 4:
			salir = true
		default:
			fmt.Println("**** Error! ****")
		}
		if salir {
			break
		}
	}

}
