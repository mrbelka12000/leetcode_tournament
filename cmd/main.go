package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/AlekSi/pointer"

	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/delivery"
	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/leetcode"
	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/repo"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/core"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/usecase"
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

	resp, tCount, err := repo.UsrEvent.GetUsrEvents(context.TODO(), &models.UsrGetEventsPars{
		EventListPars: models.EventListPars{
			PaginationParams: models.PaginationParams{
				Offset: 0,
				Limit:  15,
			},
			OnlyCount: false,
			IDs:       nil,
			UsrIDs:    nil,
			StatusIDs: nil,
			Condition: nil,
			EventGetPars: models.EventGetPars{
				UsrID: pointer.ToInt64(2),
			},
		},
		Active: nil,
		Winner: pointer.ToBool(true),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tCount)
	for _, v := range resp {
		fmt.Printf("%+v \n", v)
	}

	cr := core.New(repo, leetcodeAdapter)
	uc := usecase.New(cr)
	deliv := delivery.NewDeliveryHTTP(uc)
	r := deliv.InitRoutes()

	log.Println("starting on port: ", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, r); err != nil {
		log.Printf("run server error: %v \n", err)
		return
	}
}
