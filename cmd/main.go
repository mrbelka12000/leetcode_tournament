package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/delivery"
	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/leetcode"
	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/repo"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/core"
	"github.com/mrbelka12000/leetcode_tournament/pkg/config"
	"github.com/mrbelka12000/leetcode_tournament/pkg/database/postgres"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("get config: %v", err)
	}

	rand.Seed(time.Now().UnixNano())

	db, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatalf("connect to postgres: %v", err)
	}
	defer db.Close()

	leetcodeAdapter := leetcode.New(cfg.LeetCodeApiURL)

	repo := repo.New(db)
	cr := core.New(repo, leetcodeAdapter)
	deliv := delivery.NewDeliveryHTTP(cr)
	r := deliv.InitRoutes()

	log.Println("starting on port: ", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, r); err != nil {
		log.Printf("run server error: %v \n", err)
		return

	}
}
