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

func (r *Router) Group(path string, handlers ...RequestHandler) *Group {
	return &Group{
		fasthttpGroup: r.fasthttpRouter.Group(path),
		handler:       AfterFilter(nil, handlers...),
	}
}

type Group struct {
	handler RequestHandler
	*fasthttpGroup
}

func (g *Group) Group(path string, handlers ...RequestHandler) *Group {
	return &Group{
		fasthttpGroup: g.fasthttpGroup.Group(path),
		handler:       AfterFilter(g.handler, handlers...),
	}
}

// GET is a shortcut for group.Handle(fasthttp.MethodGet, path, handler)
func (g *Group) GET(path string, handler RequestHandler) {
	g.fasthttpGroup.GET(path, BeforeFilter(handler, g.handler))
}

// HEAD is a shortcut for group.Handle(fasthttp.MethodHead, path, handler)
func (g *Group) HEAD(path string, handler RequestHandler) {
	g.fasthttpGroup.HEAD(path, BeforeFilter(handler, g.handler))
}

// OPTIONS is a shortcut for group.Handle(fasthttp.MethodOptions, path, handler)
func (g *Group) OPTIONS(path string, handler RequestHandler) {
	g.fasthttpGroup.OPTIONS(path, BeforeFilter(handler, g.handler))
}

// POST is a shortcut for group.Handle(fasthttp.MethodPost, path, handler)
func (g *Group) POST(path string, handler RequestHandler) {
	g.fasthttpGroup.POST(path, BeforeFilter(handler, g.handler))
}

// PUT is a shortcut for group.Handle(fasthttp.MethodPut, path, handler)
func (g *Group) PUT(path string, handler RequestHandler) {
	g.fasthttpGroup.PUT(path, BeforeFilter(handler, g.handler))
}

// PATCH is a shortcut for group.Handle(fasthttp.MethodPatch, path, handler)
func (g *Group) PATCH(path string, handler RequestHandler) {
	g.fasthttpGroup.PATCH(path, BeforeFilter(handler, g.handler))
}

// DELETE is a shortcut for group.Handle(fasthttp.MethodDelete, path, handler)
func (g *Group) DELETE(path string, handler RequestHandler) {
	g.fasthttpGroup.DELETE(path, BeforeFilter(handler, g.handler))
}

// DELETE is a shortcut for group.Handle(fasthttp.MethodDelete, path, handler)
func (g *Group) ANY(path string, handler RequestHandler) {
	g.fasthttpGroup.ANY(path, BeforeFilter(handler, g.handler))
}
