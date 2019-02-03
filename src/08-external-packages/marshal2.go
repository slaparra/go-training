package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

type person2 struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bytes := marshal2(&users)
	fmt.Println(bytes)

	fmt.Println(unMarshal2(bytes))

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	people := unMarshal3([]byte(s))
	fmt.Println(people)
	fmt.Println(people[0].Sayings[1])
}

func marshal2(users *[]user) []byte {
	resultBytes, _ := json.Marshal(users)
	return resultBytes
}

func unMarshal2(otherBytes []byte) (users []user) {
	json.Unmarshal(otherBytes, &users) //the & is important: The pointer to change the content of that location
	return users
}

func unMarshal3(otherBytes []byte) (result []person2) {
	fmt.Println(otherBytes)
	json.Unmarshal(otherBytes, &result)
	return
}
