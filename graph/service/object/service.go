package service

import (
	"fmt"

	"github.com/msyamsula/messaging-api/graph/service"
)

type Service struct {
	Path    []string
	G       service.Graph
	Visited map[string]bool
	IsDone  bool
	Start   string
	End     string
}

func New(g service.Graph) service.GraphSolver {
	s := &Service{
		G: g,
	}

	s.Visited = map[string]bool{}

	s.clearVisited()
	s.Path = []string{}
	s.IsDone = false

	return s
}

func (s *Service) runDFS(node string) {
	s.Visited[node] = true
	s.Path = append(s.Path, node)

	if node == s.End {
		s.IsDone = true
		return
	}

	for _, v := range s.G.G[node] {
		if s.Visited[v] {
			continue
		}

		s.runDFS(v)
		if s.IsDone {
			return
		}

		s.Path = append(s.Path, fmt.Sprintf("#%v-%v", v, node))
	}
}

func (s *Service) clearVisited() {
	for n := range s.G.N {
		s.Visited[n] = false
	}
}

func (s *Service) DFS(start string, end string) ([]string, error) {
	var err error

	s.Path = []string{}
	s.clearVisited()
	s.IsDone = false
	s.Start = start
	s.End = end
	s.runDFS(start)

	fmt.Println(s.Path)

	return s.Path, err

}
