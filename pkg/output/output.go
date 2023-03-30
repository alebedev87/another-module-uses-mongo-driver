package output

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

func Output() {
	var i mongo.StreamType
	i++
	fmt.Println(i)
}
