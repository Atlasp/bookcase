package server

func (rt *Router) addRoutes() {
	rt.Router.Get("/book", rt.AddBook())
}
