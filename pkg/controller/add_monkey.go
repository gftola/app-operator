package controller

import (
	"github.com/gftola/app-operator/pkg/controller/monkey"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, monkey.Add)
}
