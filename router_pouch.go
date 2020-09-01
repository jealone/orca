package orca

import (
	"github.com/fasthttp/router"
)

type (
	Router = router.Router
	Group  = router.Group
)

func NewRouter(options ...func(*Router)) *Router {
	root := router.New()
	root.MethodNotAllowed = notAllowed
	root.NotFound = notFound
	root.PanicHandler = panicHandler

	for _, option := range options {
		option(root)
	}

	return root
}
