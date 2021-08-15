package graphs

import (
	"fmt"
	"sync"
)

func NewSocialGraph(num int) *SocialGraph {
	return &SocialGraph{
		Size:  num,
		M:     &sync.RWMutex{},
		Links: make([][]Link, num),
	}
}

func (sg *SocialGraph) AddLink(v1 int, v2 int, w int) {
	sg.M.Lock()
	defer sg.M.Unlock()
	sg.Links[v1] = append(sg.Links[v1], Link{
		V1:    v1,
		V2:    v2,
		LinkW: w,
	})

	return
}

func (sg *SocialGraph) PrintEntireGraph() {
	for _, adj := range sg.Links {
		for _, link := range adj {
			fmt.Printf("Link : %d -> %d (%d)\n", link.V1, link.V2, link.LinkW)
		}
	}
}
