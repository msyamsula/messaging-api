package handler

import "github.com/gin-gonic/gin"

type Node struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Edge struct {
	To   string `json:"to"`
	From string `json:"from"`
}

type GraphStructure struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Handler interface {
	SolveDFS(c *gin.Context)
}
