package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	collectionName = "users"
)

type mongoRepository struct {
	DB         *mgo.Database
	Collection *mgo.Collection
}

func (m *mongoRepository) Create(user *User) (id string, err error) {
	err = m.Collection.Insert(user)
	if err != nil {
		return "", err
	}
	return user.ID.String(), nil
}

func (m *mongoRepository) Read(email string) (*User, error) {
	user := &User{}
	err := m.Collection.Find(bson.M{"email": email}).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *mongoRepository) Update(user *User) error {
	err := m.Collection.Update(bson.M{"email": user.Email}, user)
	return err
}

func (m *mongoRepository) Delete(email string) error {
	err := m.Collection.Remove(bson.M{"email": email})
	return err
}

func newMongoRepository(DB *mgo.Database) *mongoRepository {
	collection := DB.C(collectionName)
	index := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := collection.EnsureIndex(index)
	if err != nil {
		log.Panicf("error creating users collection index: %v", err)
	}
	return &mongoRepository{DB: DB, Collection: collection}
}
