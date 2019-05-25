package main

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func main() {
	uuid, err := uuid.FromString("00000000-0000-0000-0000-111111111111")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uuid.String())
	}
}
