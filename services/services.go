package services

import (
	"afford_meds_interview/helpers/tokenhelper"
	"afford_meds_interview/models"
	"errors"
)

var userMap map[string]models.User

// UserLogin - to login user and store in it memory map
func UserLogin(user interface{}) (string, error) {
	userData, ok := user.(models.User)
	if !ok {
		return "", errors.New("Data is not in Json Format")
	}
	if userData.Username == "" && userData.Password == "" {
		return "", errors.New("Please Check Username or Password")
	}
	valUser, valOK := userMap[userData.Username]
	if !valOK {
		return "", errors.New("No Data Found for the username, Please Login")
	}
	token, err := tokenhelper.GenerateToken(valUser.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

//UserRegistration - to validate user and return token for restricted routes
func UserRegistration(user interface{}) error {
	userData, ok := user.(models.User)
	if !ok {
		return errors.New("Data is not in Json Format")
	}
	if userData.Username == "" && userData.Password == "" {
		return errors.New("Please Check Username or Password")
	}
	userMap[userData.Username] = userData
	return nil
}

func EditUser(user interface{}) error {
	userData, ok := user.(models.User)
	if !ok {
		return errors.New("Data is not in Json Format")
	}
	if userData.Username != "" {

	}
	return nil
}
func ChangePassword() error {
	return nil
}
