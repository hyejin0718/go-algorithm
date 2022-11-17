package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	Depth   int
	Visited bool
	Name    string
	Friends []*Graph
}

func main() {
	var g1, g2, g3, g4, g5, g6, g7, g8 Graph

	g1 = Graph{
		Name:    "상인",
		Friends: []*Graph{&g2, &g3, &g4},
	}
	g2 = Graph{
		Name:    "은지",
		Friends: []*Graph{&g1},
	}
	g3 = Graph{
		Name:    "범기",
		Friends: []*Graph{&g1, &g5},
	}
	g4 = Graph{
		Name:    "혜진",
		Friends: []*Graph{&g1, &g6},
	}
	g5 = Graph{
		Name:    "영욱",
		Friends: []*Graph{&g2, &g7},
	}
	g6 = Graph{
		Name:    "윤희",
		Friends: []*Graph{&g4},
	}
	g7 = Graph{
		Name:    "상미",
		Friends: []*Graph{&g5, &g8},
	}
	g8 = Graph{
		Name:    "수현",
		Friends: []*Graph{&g7},
	}
	bfs(&g4)
}

func bfs(g *Graph) {
	queue := list.New()
	depth := 0
	g.visit(depth)
	depth++
	//  정점과 인접한 각 정점을 방문하고 방문한 정점이면 표시하고 큐에 쌓는다.
	for _, friend := range g.Friends {
		if !friend.Visited {
			queue.PushBack(friend)
			friend.visit(1)
		}
	}
	//	인접 정점을 모두 방문한 경우, 큐로 부터 값을 꺼내 1번을 반복한다.
	//	모든 정점에 방문 표시를 했고 큐가 비었다면 종료한다.
	for {
		if queue.Len() == 0 {
			break
		}
		e := queue.Remove(queue.Front())
		now, ok := e.(*Graph)
		if ok {
			now.Visited = true
			for _, friend := range now.Friends {
				if !friend.Visited {
					queue.PushBack(friend)
					friend.visit(now.Depth + 1)
				}
			}
		} else {
			fmt.Print("에러")
			return
		}
	}
}

func (g *Graph) visit(depth int) {
	g.Depth = depth
	g.Visited = true
	fmt.Println(g.Name, g.Depth)
}
