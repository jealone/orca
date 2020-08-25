package orca

import (
	"github.com/fasthttp/router"
)


type (
	Router = router.Router
	Group = router.Group
)

func NewRouter() *Router {
	root := router.New()
	root.MethodNotAllowed = notAllowed
	root.NotFound = notFound
	root.PanicHandler = panicHandler
	return root
}
