package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Reviews struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Review      string
	Timestamp time.Time
}