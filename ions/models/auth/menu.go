package auth

type MenuTree struct {
	Id int
	AuthName string
	UrlFor string
	Weight int
	Children []*MenuTree
}
