package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your_secret_key") // Replace this with a secure secret key for production use.

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// In-memory user storage
var users = map[string]User{
	"user1": User{Username: "user1", Password: "password1"},
	"user2": User{Username: "user2", Password: "password2"},
	// Add more users as needed.
}

func main() {
	r := gin.Default()

	// User registration endpoint
	r.POST("/register", register)

	// User login endpoint
	r.POST("/login", login)

	// Protected route example
	r.GET("/protected", authMiddleware(), protected)

	r.Run(":8080")
}

func register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the username already exists
	if _, exists := users[user.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Replace this with proper password hashing in production.
	users[user.Username] = User{
		Username: user.Username,
		Password: user.Password,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the username exists in the user map.
	storedUser, exists := users[user.Username]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Replace this with proper password hashing and comparison in production.
	if user.Password != storedUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// For a successful login, generate a JWT token and send it in the response.
	token, err := generateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute) // Token will be valid for 15 minutes
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
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
	// Extract the username from the context.
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + username.(string) + "! This is a protected route."})
}
