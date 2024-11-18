package main

import (
	"bytes"
	"encoding/json"
	"go/projcet-Adv/internal/auth"
	"go/projcet-Adv/internal/users"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&users.User{
		Email:    "trixey@mail.com",
		Password: "$2a$10$gUTLUowXY/OSP5zSeQvOFOpOBIKvllQdBIDW6s.nHs/XF7l.LTdDa",
		Name:     "Vasya",
	})
}

func RemoveData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "trixey@mail.com").
		Delete(&users.User{})
}

func TestLoginSuccess(t *testing.T) {
	//Prepare
	db := initDb()
	initData(db)
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "trixey@mail.com",
		Password: "1",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatalf("Token empty")
	}
	RemoveData(db)
}

func TestLoginFail(t *testing.T) {
	db := initDb()
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "2",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, res.StatusCode)
	}
	RemoveData(db)
}
