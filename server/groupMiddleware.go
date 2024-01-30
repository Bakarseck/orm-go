package server

type RouteGroup struct {
    Prefix      string
    Middlewares []Middleware
    Router      *Router
}

func (r *Router) Group(prefix string) *RouteGroup {
    return &RouteGroup{
        Prefix: prefix,
        Router: r,
    }
}

func (g *RouteGroup) Use(middleware Middleware) {
    g.Middlewares = append(g.Middlewares, middleware)
}

func (g *RouteGroup) AddRoute(method, pattern string, handler RouteHandler) {
    wrappedHandler := handler
    for i := len(g.Middlewares) - 1; i >= 0; i-- {
        wrappedHandler = g.Middlewares[i](wrappedHandler)
    }
    fullPattern := g.Prefix + pattern
    g.Router.AddRoute(method, fullPattern, wrappedHandler)
}
