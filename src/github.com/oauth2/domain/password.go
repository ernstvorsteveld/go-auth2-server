package domain

import "golang.org/x/crypto/bcrypt"

type PasswordHash []byte

type BcryptPassword struct {
  password string
  hash PasswordHash
  cost int
} 

type Password interface {
  HashPassword() (PasswordHash, error)
  CheckPasswordHash() bool
}

func (h BcryptPassword) HashPassword() (error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(h.password), h.cost)
    h.hash = PasswordHash(bytes)
    return err
}

func (h BcryptPassword) CheckPasswordHash() bool {
    err := bcrypt.CompareHashAndPassword([]byte(h.hash), []byte(h.password))
    return err == nil
}
