package route

type Repository interface {
	addRoute()
	getDest(scr string) string
}
