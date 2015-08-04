package authentication

import(
	"strings" 
	"encoding/json"
	"fmt"
	"github.com/neilcwilkinson/authentication"
)

type User struct {
  Username string
  Password string
}

var Lu = make(chan int)

func IsAuthorized(message messaging.Message) (bool, User) {
	var user User

	if strings.EqualFold(message.Subject, "Authorization") {	
    	byt2 := []byte(message.Body)
    	json.Unmarshal(byt2, &user)
		fmt.Println("User: ", user.Username)
        
        if strings.EqualFold(user.Username, "neil") {
        	return true, user
        } else {
        	return false, user
        }
    }
    return false, user
}