package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/4lexvav/algo/helpers"
	"github.com/4lexvav/algo/other/735.AsteroidCollision/asteriod"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var asteroids []int
	helpers.CheckError(json.Unmarshal(scanner.Bytes(), &asteroids))

	scanner.Scan()
	var expected []int
	helpers.CheckError(json.Unmarshal(scanner.Bytes(), &expected))

	result := asteriod.AsteroidCollision(asteroids)

	if !reflect.DeepEqual(result, expected) {
		data, err := json.Marshal(result)
		helpers.CheckError(err)
		fmt.Printf("Failed, got unexpected result: %s\n", string(data))
		return
	}

	fmt.Println("Success!")
}
