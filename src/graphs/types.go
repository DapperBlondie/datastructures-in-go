package graphs

import "sync"

// Link holding Vertexes data and LinkWeight
type Link struct {
	V1    int
	V2    int
	LinkW int
}

type SocialGraph struct {
	Size  int
	Links [][]Link
	M     *sync.RWMutex
}
