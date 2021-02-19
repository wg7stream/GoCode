package main
import (
	"fmt"
)

func main() {
	// 声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "asdf"
		monsters[0]["age"] = "66"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "fdfds"
		monsters[1]["age"] = "400"
	}

	// append
	newMonster := map[string]string {
		"name" : "sdefds",
		"age" : "555",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)
}