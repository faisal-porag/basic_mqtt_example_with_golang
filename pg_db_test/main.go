package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MenuType struct {
	ID       int    `db:"id" json:"id"`
	TypeName string `db:"type_name" json:"type_name"`
}

type Menu struct {
	ID         int      `db:"id" json:"id"`
	Name       string   `db:"menu_name" json:"name"`
	MenuTypeID int      `db:"menu_type_id" json:"menu_type_id"`
	MenuType   MenuType `db:"menu_type" json:"menu_type"`
}

func main() {
	// Database connection
	databaseURL := "postgres://postgres:admin@0.0.0.0:5432/testDB?sslmode=disable"
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query
	query := `
        SELECT
            m.id,
            m.menu_name,
            m.menu_type_id,
            mt.id AS "menu_type.id",
            mt.menu_type_name AS "menu_type.type_name"
        FROM
            menu AS m
        JOIN
            menu_type AS mt ON m.menu_type_id = mt.id
        WHERE
            m.id = $1
    `

	// Execute the query and scan the result
	var menu Menu
	menuID := 16
	err = db.Get(&menu, query, menuID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No menu found with ID:", menuID)
		} else {
			log.Fatal(err)
		}
	}

	// Format response as JSON
	responseData := struct {
		Data Menu `json:"data"`
	}{
		Data: menu,
	}

	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonResponse))
}
