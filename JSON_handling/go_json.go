package main

import (
	"encoding/json"
	"fmt"
)

type Body struct {
	ImagePath string   `json:"img_path"`
	Actions   []string `json:"actions"`
}

func main() {

	payload := "{\n    \"img_path\": \"http://localhost:8090/public/apc/people/32/image.jpg\",\n    \"actions\": [\"age\", \"gender\", \"emotion\", \"race\"]\n}"

	body := Body{}

	// converting raw json to struct object
	err := json.Unmarshal([]byte(payload), &body)

	if err != nil {
		return
	}

	//fmt.Println("Image path: ", body.ImagePath)
	//fmt.Println("Actions: ", body.Actions)

	MarshalJSONIndent(&body)

}

func MarshalJSON(obj *Body) {

	// converting struct object to raw json
	decodedBytes, err := json.Marshal(obj)

	if err != nil {
		return
	}

	fmt.Printf("%+v", string(decodedBytes))
}

func MarshalJSONIndent(obj *Body) {

	// converting struct object to raw json with indentation
	decodedBytes, err := json.MarshalIndent(obj, "", "	")

	if err != nil {
		return
	}

	fmt.Printf("%+v", string(decodedBytes))
}
