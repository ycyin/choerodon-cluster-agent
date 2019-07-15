package controller

import (
	"github.com/choerodon/choerodon-cluster-agent/pkg/controller/deployment"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, deployment.Add)
}
