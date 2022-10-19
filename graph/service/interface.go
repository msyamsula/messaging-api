package service

type Edge struct {
	From string
	To   string
}

type Graph struct {
	G map[string][]string
	E map[Edge]int64
	N map[string]bool
}

type GraphSolver interface {
	DFS(startNode string, endNode string) ([]string, error)
}
