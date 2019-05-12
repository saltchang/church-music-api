package models

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Test_InitDB(t *testing.T) {

	fmt.Println("DB testing start. Test target SID: 1010066")

	var testDB = new(Database)

	testDB.InitDB()

	testFilter := bson.M{"sid": "1010066"}

	testResult := &Song{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	testDB.Songs.FindOne(ctx, testFilter).Decode(&testResult)

	cancel()

	if testResult.SID != "1010066" {
		t.Error("DB test fails. (Not Expected SID)")
	} else if testResult.Title != "前來敬拜" {
		t.Error("DB test fails. (Not Expected Title)")
	} else {
		t.Log("PASS!")
	}
}
