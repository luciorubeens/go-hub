package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
}

func main() {
	user_1 := User{"User 1"}
	user_bytes, err := json.Marshal(user_1)

	fmt.Println(err, user_bytes, string(user_bytes))

	json_string := `{"Name": "Testando"}`
	user_2 := User{}

	err_2 := json.Unmarshal([]byte(json_string), &user_2)

	fmt.Println(err_2, user_2)
}
