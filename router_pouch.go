package orca

import (
	"github.com/fasthttp/router"
)

type FasthttpRoute struct {
	*router.Router
}

func NewRouter() *FasthttpRoute {
	root := router.New()
	return &FasthttpRoute{
		root,
	}
}

func (r *FasthttpRoute) Group(path string) Router {
	return &FasthttpGroup{
		r.Router.Group(path),
	}

}

// 伪装匿名类
type shadeGroup = router.Group

type FasthttpGroup struct {
	*shadeGroup
}

func (g *FasthttpGroup) Group(path string) Router {
	return &FasthttpGroup{
		// 伪装匿名类的作用
		g.shadeGroup.Group(path),
	}
}
