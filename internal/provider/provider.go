package provider

import (
	"auth-go/internal/config"
	"auth-go/internal/models"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Provider struct {
	db *gorm.DB
}

func New(cfg *config.Config) *Provider {
	return &Provider{
		db: GetConnection(cfg),
	}
}

func GetConnection(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbSSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	return db
}

func (p *Provider) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := p.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}

	return &user, nil
}

func (p *Provider) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := p.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with username %s not found", username)
		}
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}
	return &user, nil
}

func (p *Provider) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := p.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}
	return &user, nil
}

func (p *Provider) CreateUser(user *models.User) error {
	if err := p.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("user with email %s already exists", user.Email)
		}

		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}

func (p *Provider) UpdateUser(user *models.User) error {
	if err := p.db.Save(&user).Error; err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

func (p *Provider) DeleteUserByID(id uint) error {
	var user models.User
	if err := p.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID %d not found", id)
		}
		return fmt.Errorf("error retrieving user: %v", err)
	}
	return nil
}
