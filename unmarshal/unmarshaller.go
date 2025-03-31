package unmarshal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Type   string `json:"type"`
	Social Social `json:"social"`
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Type: %s, Social: {Facebook: %s, VK: %s}",
		u.Name, u.Age, u.Type, u.Social.Facebook, u.Social.VK)
}

type Social struct {
	Facebook string `json:"fb"`
	VK       string `json:"vk"`
}

func ShowUnmarshalExample() {
	jsonFile, err := os.Open("advanced_users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// Unmarshal the JSON data into a slice of User structs
	var users Users
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users.Users {
		fmt.Println(user)
	}
}
