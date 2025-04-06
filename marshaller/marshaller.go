package marshaller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type Professor struct {
	Name       string     `json:"name"`
	ScienceId  int        `json:"science_id"`
	IsWorking  bool       `json:"is_working"`
	University University `json:"university"`
}

func ShowMarshallerExample() {
	professor := Professor{
		Name:      "John Doe",
		ScienceId: 12345,
		IsWorking: true,
		University: University{
			Name: "Harvard University",
			City: "Cambridge",
		},
	}

	jsonData, err := json.MarshalIndent(professor, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
	err = os.WriteFile("professor.json", jsonData, 0664) // permissions -rw-rw-r--
	if err != nil {
		log.Fatal(err)
	}
}
