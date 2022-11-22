package main

import (
	"container/list"
	"fmt"
)

type City struct {
	Name   string
	Routes map[string]int
}
type Graph struct {
	Cities  map[string]*City
	Visited map[string]bool
	Cost    map[string]int
}

const MaxInt = int(^uint(0) >> 1)

func (g *Graph) NewCity(name string) *City {
	g.Visited[name] = false
	g.Cost[name] = MaxInt //9223372036854775807
	g.Cities[name] = &City{
		Name:   name,
		Routes: make(map[string]int),
	}
	return g.Cities[name]
}

func (c *City) AddRoute(city string, price int) {
	c.Routes[city] = price
}

func main() {
	graph := Graph{
		Cities:  map[string]*City{},
		Visited: map[string]bool{},
		Cost:    map[string]int{},
	}

	tokyo := graph.NewCity("도쿄")
	tokyo.AddRoute("오사카", 40)
	tokyo.AddRoute("삿뽀로", 100)
	sapporo := graph.NewCity("삿뽀로")
	sapporo.AddRoute("오키나와", 180)
	osaka := graph.NewCity("오사카")
	osaka.AddRoute("후쿠오카", 80)
	osaka.AddRoute("삿뽀로", 80)
	fukuoka := graph.NewCity("후쿠오카")
	fukuoka.AddRoute("오키나와", 40)
	fukuoka.AddRoute("삿뽀로", 120)
	okinawa := graph.NewCity("오키나와")
	okinawa.AddRoute("오사카", 180)
	okinawa.AddRoute("도쿄", 160)

	cost := graph.getMinCost("도쿄", "오키나와")
	fmt.Println(cost)
}

func (g *Graph) getMinCost(start, end string) int {
	fmt.Println("***", start, "에서 출발할 때 최소 비용 계산***")
	queue := list.New()
	queue.PushBack(g.Cities[start])
	g.Cost[start] = 0
	for {
		if queue.Len() == 0 {
			break
		}
		e := queue.Remove(queue.Front())
		now, ok := e.(*City) // 최소값을 탐색할 city 를 꺼낸다.
		if ok {
			g.Visited[now.Name] = true
			for name, cost := range now.Routes { // 현재 city 에서 갈 수 있는 모든 루트를 visit 한다.
				city := g.Cities[name]
				if !g.Visited[name] { // 다음 탐색을 위해 방문된적 없는 city 를 큐에 넣는다.
					queue.PushBack(city)
				}
				g.visitCity(city, g.Cost[now.Name]+cost)
			}
		} else {
			fmt.Print("에러;")
			return 0
		}
	}

	if !g.Visited[g.Cities[end].Name] {
		fmt.Print(end, "까지의 경로 없음;")
		return 0
	}

	cost, ok := g.Cost[end]
	if ok {
		fmt.Println(start, "부터", end, "까지의 최소 비용:", cost)
		return cost
	}
	fmt.Print("에러;")
	return 0
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (g *Graph) visitCity(city *City, cost int) {
	g.Visited[city.Name] = true
	g.Cost[city.Name] = getMin(g.Cost[city.Name], cost)
	fmt.Println(city.Name, "까지 최소 비용", g.Cost[city.Name])
}
