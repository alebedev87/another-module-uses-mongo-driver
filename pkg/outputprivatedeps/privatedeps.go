package outputprivatedeps

import (
	"fmt"
	"github.com/alebedev87/another-module-uses-mongo-driver/pkg/output"
)

func Output() {
	fmt.Println(3)
}

func outputiWithDep() {
	output.Output()
}
