package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/db"
	"github.com/krinnnout/reserve-get-served/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func InsertTestUser(t *testing.T, userStore db.UserStore) *models.User {
	c := context.Background()
	user, err := models.NewUserFromParams(models.UserParams{FirstName: "James", LastName: "Bebra", Email: "james@foo.com", Password: "supersecuredpassword"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = userStore.InsertUser(c, user)
	if err != nil {
		t.Fatal(err)
	}
	return user
}

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	insertedUser := InsertTestUser(t, tdb.UserStore)
	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/auth", authHandler.HandleAuthenticate)
	params := AuthParams{
		Email:    "james@foo.com",
		Password: "supersecuredpassword",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected http status of 200 but got %d", resp.StatusCode)
	}
	var response AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Error(err)
	}
	if response.Token == "" {
		t.Fatal("expected the JWT token to be present in the auth response")
	}
	insertedUser.EncryptedPassword = ""

	if !reflect.DeepEqual(insertedUser, response.User) {
		t.Fatal("expected the user to match the inserted user")
	}
}

func TestAuthenticateWithWrongPasswordFailure(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/auth", authHandler.HandleAuthenticate)
	params := AuthParams{
		Email:    "james@foo.com",
		Password: "supersecuredpasswordnotcorrect",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected http status of 400 but got %d", resp.StatusCode)
	}
	var genResp GenericResponse
	if err := json.NewDecoder(resp.Body).Decode(&genResp); err != nil {
		t.Fatal(err)
	}
	if genResp.Type != "error" {
		t.Fatalf("expected genResp Type to be error but got %s", genResp.Type)
	}
	if genResp.Msg != "invalid credentials" {
		t.Fatalf("expected genResp Msg to be <invalid credentials> but got %s", genResp.Msg)
	}
}
