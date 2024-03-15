package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"xinmiao/app/models"
	"xinmiao/config/database"
)

func GetUserByStudentIDAndPassword(studentID, password string) (*models.User, error) {
	h := sha256.New()
	h.Write([]byte(password))
	user := models.User{}
	pass := hex.EncodeToString(h.Sum(nil))
	result := database.DB.Where(
		&models.User{
			StudentID: studentID,
			Password:  pass,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}

func GetUserID(id string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			ID: id,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}
