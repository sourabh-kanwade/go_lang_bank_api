package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(string) error
	UpdateAccount(*Account) error
	GetAccountByID(string) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type MongoStore struct {
	db *mongo.Database
}

func NewMongoStore() (*MongoStore, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		return nil, err
	}

	database := client.Database("bank_api")
	return &MongoStore{
		db: database,
	}, nil
}

func (s *MongoStore) CreateAccount(account *Account) error {

	accountsCollection := s.db.Collection("accounts")
	insertResult, err := accountsCollection.InsertOne(context.Background(), account)
	if err != nil {
		return (err)
	}
	fmt.Println(insertResult)
	return nil
}
func (s *MongoStore) DeleteAccount(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	accountsCollection := s.db.Collection("accounts")
	_, err = accountsCollection.UpdateByID(context.Background(), objectId, bson.D{{Key: "$set", Value: bson.D{{Key: "isDeleted", Value: true}}}})
	if err != nil {
		return err
	}
	return nil
}
func (s *MongoStore) UpdateAccount(*Account) error {
	return nil
}
func (s *MongoStore) GetAccountByID(id string) (*Account, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	accountsCollection := s.db.Collection("accounts")
	var account *Account

	if err := accountsCollection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&account); err != nil {
		return nil, err
	}

	return account, nil
}
func (s *MongoStore) GetAccounts() ([]*Account, error) {

	accountsCollection := s.db.Collection("accounts")
	accounts := []*Account{}
	cursor, err := accountsCollection.Find(context.Background(), bson.M{"isDeleted": false})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &accounts); err != nil {
		return nil, err
	}
	fmt.Printf("%+v", accounts)
	return accounts, nil
}
