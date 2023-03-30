package outputnodeps

import (
	"fmt"
	"github.com/alebedev87/another-module-uses-mongo-driver/pkg/output"
)

func Output() {
	fmt.Println(2)
}

func OutputiWithDep() {
	output.Output()
}
