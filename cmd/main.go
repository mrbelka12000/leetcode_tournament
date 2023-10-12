package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"leetcode_tournament/internal/adapters/delivery"
	"leetcode_tournament/internal/adapters/leetcode"
	"leetcode_tournament/internal/adapters/repo"
	"leetcode_tournament/internal/domain/core"
	"leetcode_tournament/pkg/sqlite"
)

const (
	withMigrates = true
)

func main() {
	os.Setenv("PORT", "8080")
	rand.Seed(time.Now().UnixNano())

	db, err := sqlite.Connect()
	if err != nil {
		log.Fatal(fmt.Errorf("connect to db: %w", err))
	}
	defer db.Close()

	if withMigrates {
		err = doMigrates(db)
		if err != nil {
			log.Println("do migrates: ", err)
			return
		}
	}

	leetcode := leetcode.New("https://leetcode.com/graphql")

	repo := repo.NewRepo(db)
	cr := core.NewCore(repo, leetcode)
	deliv := delivery.NewDeliveryHTTP(cr)

	r := deliv.InitRoutes()

	fmt.Println("starting on port: ", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
		log.Printf("run server error: %v \n", err)
		return
	}
}

func doMigrates(db *sql.DB) error {
	dir, err := os.ReadDir("ddl")
	if err != nil {
		return fmt.Errorf("read dir: %w", err)
	}

	for _, file := range dir {
		body, err := os.ReadFile("ddl/" + file.Name())
		if err != nil {
			return fmt.Errorf("os read file %v: %w", file.Name(), err)
		}

		_, err = db.Exec(string(body))
		if err != nil {
			return fmt.Errorf("sql exec %v: %w", file.Name(), err)
		}
	}
	return nil
}
