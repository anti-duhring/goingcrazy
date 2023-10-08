package config

import (
	"context"
	"encoding/json"
	"time"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Worker struct {
	DB       *gorm.DB
	Cache    *redis.Client
	ChPeople chan schema.Person
}

var Wkr *Worker

func NewWorker(DB *gorm.DB, Cache *redis.Client) *Worker {

	if Wkr == nil {
		Wkr = &Worker{
			DB:       DB,
			Cache:    Cache,
			ChPeople: make(chan schema.Person),
		}
	}

	return Wkr
}

func (w *Worker) Create(ctx context.Context, person schema.Person) error {
	if v, _ := w.Cache.Get(ctx, person.Apelido).Result(); v != "" {
		return ErrNicknameAlreadyExists
	}

	j, err := json.Marshal(person)
	if err != nil {
		return err
	}

	pipe := w.Cache.Pipeline()
	pipe.Set(ctx, person.Apelido, "t", 0)
	pipe.Set(ctx, person.ID.String(), j, 24*time.Hour)
	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("error creating person: %v", err)
		return err
	}

	w.ChPeople <- person

	return nil
}

func (w *Worker) Insert(persons []schema.Person) error {
	if err := w.DB.Create(&persons).Error; err != nil {
		logger.Errorf("error inserting person: %v", err)
		return err
	}

	return nil
}

func RunWorker(chPeople chan schema.Person, chanExit chan struct{}, batch int) {
	logger.Debug("Starting worker")
	defer logger.Debug("Finishing worker")

	i := 0
	people := make([]schema.Person, 0, batch)
	tick := time.NewTicker(1 * time.Second)

	for {
		select {
		case p, ok := <-chPeople:
			if p.ID != uuid.Nil {
				people = append(people, p)
				i++

			}
			if i == batch || !ok {
				if err := Wkr.Insert(people); err != nil {
					logger.Errorf("error inserting person: %v", err)
				}
				people = make([]schema.Person, 0, batch)
				i = 0
			}
			i++

			if !ok {
				chanExit <- struct{}{}
				return
			}
		case <-tick.C:
			if len(people) > 0 {
				if err := Wkr.Insert(people); err != nil {
					logger.Errorf("error inserting person: %v", err)
				}
				people = make([]schema.Person, 0, batch)
				i = 0
			}
		}
	}
}
