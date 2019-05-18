package models

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Test_InitDB(t *testing.T) {

	fmt.Println("DB testing start. Test target SID: '1010066'")

	var testDB = new(Database)

	testDB.InitDB()

	songsTestFilter := bson.M{"sid": "1010066"}
	tokensTestFilter := bson.M{"token": "B290346CC8BCC4C3C668FE3A25027344B1DEA70986FECB16E22E724BFB56C72A"}

	songsTestResult := &Song{}

	tokensTestResult := &Token{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	testDB.Songs.FindOne(ctx, songsTestFilter).Decode(&songsTestResult)

	testDB.Tokens.FindOne(ctx, tokensTestFilter).Decode(&tokensTestResult)

	cancel()

	if songsTestResult.SID != "1010066" {
		t.Error("Songs DB test fails. (Not Expected SID)")
	} else if songsTestResult.Title != "前來敬拜" {
		t.Error("Songs DB test fails. (Not Expected Title)")
	} else {
		t.Log("PASS!")
	}

	if tokensTestResult.Token != "B290346CC8BCC4C3C668FE3A25027344B1DEA70986FECB16E22E724BFB56C72A" {
		t.Error("Tokens DB test fails. (Not Expected Token)")
	} else if tokensTestResult.Autho != "FORTEST" {
		t.Error("Tokens DB test fails. (Not Expected Autho)")
	} else {
		t.Log("PASS!")
	}
}
