package main

import (
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateAccountReq struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Number    int64              `bson:"number" json:"number"`
	Balance   int64              `bson:"balance" json:"balance"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	IsDeleted bool               `bson:"isDeleted" json:"isDeleted"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63n(100000),
		CreatedAt: time.Now().UTC(),
	}
}
