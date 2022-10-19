package object

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msyamsula/messaging-api/graph/handler"
	"github.com/msyamsula/messaging-api/graph/service"
	graphServiceO "github.com/msyamsula/messaging-api/graph/service/object"
)

type HTTPHandler struct {
	service service.GraphSolver
}

func New() handler.Handler {
	h := &HTTPHandler{}
	return h
}

func (h *HTTPHandler) SolveDFS(c *gin.Context) {
	var err error
	structure := handler.GraphStructure{}

	err = c.BindJSON(&structure)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please check your graph input",
		})
		c.Abort()
		return
	}

	nodes := map[string]bool{}
	for _, n := range structure.Nodes {
		nodes[n.ID] = true
	}

	edges := map[service.Edge]int64{}
	for _, e := range structure.Edges {
		se := service.Edge{
			To:   e.To,
			From: e.From,
		}
		edges[se] = 1
	}

	gs := map[string][]string{}
	for e, _ := range edges {
		if _, ok := gs[e.From]; !ok {
			gs[e.From] = make([]string, 0)
		}

		gs[e.From] = append(gs[e.From], e.To)
	}

	graph := service.Graph{
		N: nodes,
		E: edges,
		G: gs,
	}

	solver := graphServiceO.New(graph)

	var path []string
	path, err = solver.DFS(structure.Start, structure.End)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"path": path,
	})

}
