package orca

import (
	"github.com/fasthttp/router"
)

type (
	fasthttpRouter = router.Router
	fasthttpGroup  = router.Group
)

func NewRouter(options ...func(*Router)) *Router {
	root := &Router{
		router.New(),
	}
	root.MethodNotAllowed = notAllowed
	root.NotFound = notFound
	root.PanicHandler = panicHandler

	for _, option := range options {
		option(root)
	}

	return root
}

type Router struct {
	*fasthttpRouter
}

func (r *Router) Group(path string, middleware ...MiddlewareHandler) *Group {
	return &Group{
		fasthttpGroup: r.fasthttpRouter.Group(path),
		middleware:    LambdaMiddleware(nil, middleware...),
	}
}

type Group struct {
	middleware MiddlewareHandler
	*fasthttpGroup
}

func (g *Group) Group(path string, middleware ...MiddlewareHandler) *Group {
	return &Group{
		fasthttpGroup: g.fasthttpGroup.Group(path),
		middleware:    LambdaMiddleware(g.middleware, middleware...),
	}
}

// GET is a shortcut for group.Handle(fasthttp.MethodGet, path, handler)
func (g *Group) GET(path string, handler RequestHandler) {
	g.fasthttpGroup.GET(path, Middleware(g.middleware, handler))
}

// HEAD is a shortcut for group.Handle(fasthttp.MethodHead, path, handler)
func (g *Group) HEAD(path string, handler RequestHandler) {
	g.fasthttpGroup.HEAD(path, Middleware(g.middleware, handler))
}

// OPTIONS is a shortcut for group.Handle(fasthttp.MethodOptions, path, handler)
func (g *Group) OPTIONS(path string, handler RequestHandler) {
	g.fasthttpGroup.OPTIONS(path, Middleware(g.middleware, handler))
}

// POST is a shortcut for group.Handle(fasthttp.MethodPost, path, handler)
func (g *Group) POST(path string, handler RequestHandler) {
	g.fasthttpGroup.POST(path, Middleware(g.middleware, handler))
}

// PUT is a shortcut for group.Handle(fasthttp.MethodPut, path, handler)
func (g *Group) PUT(path string, handler RequestHandler) {
	g.fasthttpGroup.PUT(path, Middleware(g.middleware, handler))
}

// PATCH is a shortcut for group.Handle(fasthttp.MethodPatch, path, handler)
func (g *Group) PATCH(path string, handler RequestHandler) {
	g.fasthttpGroup.PATCH(path, Middleware(g.middleware, handler))
}

// DELETE is a shortcut for group.Handle(fasthttp.MethodDelete, path, handler)
func (g *Group) DELETE(path string, handler RequestHandler) {
	g.fasthttpGroup.DELETE(path, Middleware(g.middleware, handler))
}

// DELETE is a shortcut for group.Handle(fasthttp.MethodDelete, path, handler)
func (g *Group) ANY(path string, handler RequestHandler) {
	g.fasthttpGroup.ANY(path, Middleware(g.middleware, handler))
}
