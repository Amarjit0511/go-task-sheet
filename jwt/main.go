package main

import (
	"net/http"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var key []byte

func main() {
	// Loading the config.env file
	err := godotenv.Load("config.env")
	if err != nil {
		panic("Error loading .env file")
	}

	key = []byte(getEnv("SECRET_KEY", ""))

	r := gin.Default()

	r.POST("/register", register)
	r.POST("/login", login)
	r.GET("/protected", authMiddleware(), protected)

	r.Run(":8080")
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var users = map[string]User{
	"user1": User{Username: "user1", Password: "password1"},
	"user2": User{Username: "user2", Password: "password2"},
}

func register(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if _, exists := users[u.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	users[u.Username] = User{
		Username: u.Username,
		Password: u.Password,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration Successfull"})
}

func login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	sUser, exists := users[u.Username]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if u.Password != sUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := generateToken(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func generateToken(username string) (string, error) {
	expTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tStr := c.GetHeader("Authorization")
		if tStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("username", claims.Username)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}

func protected(c *gin.Context) {
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + username.(string) + "! Successfully implemented JWT from task"})
}
