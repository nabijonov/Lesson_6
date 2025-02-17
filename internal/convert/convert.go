package convert

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Shemistan/Lesson_6/internal/models"
)

func ApiAuthConvertToService(juser string)(user models.SAuth) {
	js := []byte(juser)
	user1 := models.SAuth{}

	if json.Valid(js) {
		err := json.Unmarshal(js, &user1)
		if err != nil{
			fmt.Println(err)
		}
		return user1
	}

		return user1
}

func ApiUserConvertToService(juser string)(models.SUser, error) {
	js := []byte(juser)
	user1 := models.SUser{}
	if json.Valid(js) {
		err := json.Unmarshal(js, &user1)
		if err != nil{
			fmt.Println(err)
		}
		return user1, nil
	}

		return user1, errors.New("Error Cant Convert json")
}

func ApiUserConvertFromoService(user models.SUser) (string) {
	js,_ := json.Marshal(user)
	juser := string(js)

	return juser
}

func ApiIdConvertToService(juser string)(models.IdGenerate, error) {
	js := []byte(juser)
	id := models.IdGenerate{}
	if json.Valid(js) {
		err := json.Unmarshal(js, &id)
		if err != nil{
			fmt.Println(err)
		}
		return id, nil
	}

		return id, errors.New("Error Cant Convert json")
}

func ApiIdConvertFromoService(id models.IdGenerate) (string) {
	js,_ := json.Marshal(id)
	juser := string(js)

	return juser
}
