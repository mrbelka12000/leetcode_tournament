package main

import (
	"log"
	"net/http"

	"github.com/mrbelka12000/leetcode_tournament/internal/client/leetcode"
	"github.com/mrbelka12000/leetcode_tournament/internal/delivery"
	"github.com/mrbelka12000/leetcode_tournament/internal/repo"
	"github.com/mrbelka12000/leetcode_tournament/internal/service"
	"github.com/mrbelka12000/leetcode_tournament/internal/usecase"
	"github.com/mrbelka12000/leetcode_tournament/pkg/config"
	"github.com/mrbelka12000/leetcode_tournament/pkg/database/postgres"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("get config: %v", err)
	}

	db, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatalf("connect to postgres: %v", err)
	}
	defer db.Close()

	leetcodeAdapter := leetcode.New(cfg.LeetCodeApiURL)

	repo := repo.New(db)
	cr := service.New(repo, leetcodeAdapter)
	uc := usecase.New(cr)
	deliv := delivery.NewDeliveryHTTP(uc)
	r := deliv.InitRoutes()

	log.Println("starting on port: ", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, r); err != nil {
		log.Printf("run server error: %v \n", err)
		return
	}
}
