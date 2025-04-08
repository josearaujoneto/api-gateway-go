package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mutex     sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateAPIKey() string {
	byteSlice := make([]byte, 16)
	rand.Read(byteSlice)
	return hex.EncodeToString((byteSlice))
}

func NewAccount(name, email string) *Account {

	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		APIKey:    generateAPIKey(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (account *Account) AddBalance(amount float64) {
	account.mutex.Lock()
	//defer account.mutex.Lock() "defer" aguarda a execução de todo código para executar esta linha
	account.Balance += amount
	account.UpdatedAt = time.Now()
	account.mutex.Unlock()
}
