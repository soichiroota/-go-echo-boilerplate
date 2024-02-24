package routes

import (
	"sample/db"
	"sample/models"

	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) (err error) {
	// Bind
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	user, err := models.CreateUser(u)
	if err != nil {
		return
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) (err error) {
	// Bind
	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	// Find user
	var user models.User
	ctx := context.Background()
	err = db.DB.NewSelect().Model(&user).Where("email = ? and password = ?", u.Email, u.Password).Scan(ctx)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Don't send password
	return c.JSON(http.StatusOK, u)
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func GetUsers(c echo.Context) error {
		users, err := models.GetAllUsers()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Userを取得できませんでした。")
		}
		return c.JSON(http.StatusOK, users)
}