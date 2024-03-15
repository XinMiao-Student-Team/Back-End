package userServices

import (
	"xinmiao/app/apiException"
	"xinmiao/config/database"
)

func ResetPass(studentID, oldpass, newpass string) error {
	user, err := GetUserByStudentIDAndPassword(studentID, oldpass)
	if err != nil {
		return apiException.UserNotFind
	}
	user.Password = newpass
	if err := database.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
