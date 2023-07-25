package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	var  err error
	databaseUrl := "postgres://postgres:puja@localhost:5432/db_b48"
	Conn, err = pgx.Connect(context.Background(),databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("connected to Database")
}