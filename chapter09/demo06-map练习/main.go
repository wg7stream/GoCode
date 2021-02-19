package main
import (
	"fmt"
)

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "88888"
	} else {
		users[name = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "efdewf"
	}
}

func main() {
	users := make(map[string]map[string]string, 10)
	
}