package controllers

import (
	"auth-go/internal/api/types"
	"auth-go/internal/config"
	"auth-go/internal/models"
	"auth-go/internal/provider"
	"auth-go/internal/statistics"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	Provider   *provider.Postgres
	Statistics *statistics.Statistics
	Config     *config.Config
}

func New(provider *provider.Postgres, statistics *statistics.Statistics, cfg *config.Config) *Controller {
	return &Controller{
		Provider:   provider,
		Statistics: statistics,
		Config:     cfg,
	}
}

func (c *Controller) GetInfo(ctx *gin.Context) {
	stats := c.Statistics.GetInfo()
	data, err := json.Marshal(stats)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error while marshaling stats"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (c *Controller) Login(ctx *gin.Context) {
	c.Statistics.IncrementRequests()
	if ctx.Request.Method != http.MethodPost {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Method not allowed!"})
		c.Statistics.AddError("Method not allowed", "Invalid method for login request")
		c.Statistics.SetStatusOK(false)
		return
	}

	var loginReq types.LoginRequest
	if err := ctx.ShouldBindBodyWithJSON(loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Statistics.AddError("Error", err.Error())
		c.Statistics.SetStatusOK(false)
		return
	}

	dbUser, err := c.Provider.GetUserByEmail(loginReq.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		c.Statistics.AddError("Error", err.Error())
		c.Statistics.SetStatusOK(false)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginReq.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password"})
		c.Statistics.AddError("Invalid password", "Password does not match with hash")
		c.Statistics.SetStatusOK(false)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": dbUser.ID,
		"jit":    uuid.New(),
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":    time.Now(),
	})

	tokenStr, err := token.SignedString(c.Config.Secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Statistics.AddError("Token signing error", err.Error())
		c.Statistics.SetStatusOK(false)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func (c *Controller) Register(ctx *gin.Context) {
	c.Statistics.IncrementRequests()
	if ctx.Request.Method != http.MethodPost {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Method not allowed!"})
		c.Statistics.AddError("Method not allowed", "Invalid method for registration request")
		c.Statistics.SetStatusOK(false)
		return
	}

	var registerReq types.RegisterRequest

	if err := ctx.ShouldBindBodyWithJSON(registerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Statistics.AddError("Error", err.Error())
		c.Statistics.SetStatusOK(false)
		return
	}

	_, err := c.Provider.GetUserByEmail(registerReq.Email)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		c.Statistics.AddError("User already exists", "Attempt to register with existing email")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Statistics.AddError("Password hashing error", err.Error())
		return
	}

	user := &models.User{
		Username: registerReq.Username,
		Email:    registerReq.Email,
		Password: string(hash),
	}

	err = c.Provider.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Statistics.AddError("User creation error", err.Error())
		c.Statistics.SetStatusOK(false)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully!"})
}
