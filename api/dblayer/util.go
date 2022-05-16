package dblayer

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	//convert password string to byte slice so that we can use it with the bcrypt package
	sBytes := []byte(*s)
	//Obtain hashed password via the GenerateFromPassword() method
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingPass string) bool {
	//this method will return an error if the hash does not match the provided password string
	return bcrypt.CompareHashAndPassword([]byte(existingHash),
		[]byte(incomingPass)) == nil
}
