package webRouter

import "net/http"

type middleware func(http.Handler) http.Handler

type Router struct {
    middlewareChain [] middleware
    mux             map[string]http.Handler
}

// 路由
func NewRouter() *Router {
    return &Router{}
}

// 使用中间件
func (r *Router) Use(m middleware) {
    r.middlewareChain = append(r.middlewareChain, m)
}

// 添加路由
func (r *Router) Add(route string, h http.Handler) {
    var mergedHandler = h

    for i := len(r.middlewareChain) - 1; i >= 0; i-- {
        mergedHandler = r.middlewareChain[i](mergedHandler)
    }
    r.mux[route] = mergedHandler
}
