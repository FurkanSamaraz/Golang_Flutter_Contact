package checkError

import (
	"fmt"
	"log"
)

func ErrorContr(err error, num string) {
	if err != nil {
		log.Fatal(err)
		fmt.Println(num)
	}

}
