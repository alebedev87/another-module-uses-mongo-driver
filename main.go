package main

import (
	"github.com/alebedev87/another-module-uses-mongo-driver/pkg/output"
	nodeps "github.com/alebedev87/another-module-uses-mongo-driver/pkg/outputnodeps"
)

func main() {
	output.Output()
	nodeps.Output()
}
