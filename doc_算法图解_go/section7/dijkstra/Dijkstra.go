package dijkstra

import (
	"math"
)

//DijkstraDemo-迪克斯特拉算法示例
//graph 存储节点及其邻居
//cost 存储节点开销
//parent 存储节点及其父节点
/*

S->A, 6       A
S->B, 2     / | \
B->A, 3    S  |  E
A->E, 1     \ | /
B->E, 5       B

 graph:             |  costs              |  parents:
  __________        |   __________        |   __________
 |    A| 6  |       |  |    A| 6  |       |  |    A| S  |
 | S   |    |       |  |    B| 2  |       |  |    B| S  |
 |____B|_2__|       |  |____E|_--_|       |  |____E|_--_|
 |_A__E|_1__|       |                     |
 |    A| 3  |       |                     |
 | B   |    |       |                     |
 |____E|_5__|       |                     |
 |_E___|_--_|       |                     |

*/
func DijkstraDemo(graph map[string]map[string]int, costs map[string]int, parents map[string]string) {
	var processed []string
	//所有最小开销的节点全部被处理完成后结束循环
	for node := find_lowest_cost_node(costs, processed); node != ""; node = find_lowest_cost_node(costs, processed) {
		cost := costs[node]
		neighbors := graph[node]
		for key := range neighbors { //遍历当前节点的所有邻居节点
			new_cost := cost + neighbors[key] //从该节点到邻居节点的开销
			if costs[key] > new_cost {        //当前节点到达邻居节点的开销更小
				costs[key] = new_cost //更新邻居节点的开销
				parents[key] = node   //更新邻居节点的父节点为当前节点
			}
		}
		processed = append(processed, node) //当前节点标记为已处理
	}

}

//找到邻居节点中开销最小的节点
func find_lowest_cost_node(costs map[string]int, processed []string) (node string) {
	node = ""
	low_cost := math.MaxInt
	var in bool
	for _, n := range processed {
		_, ok := costs[n]
		if ok {
			in = true
		}
	}
	for key := range costs { //遍历节点开销
		if low_cost > costs[key] && !in { //当前节点开销最低且未处理
			low_cost = costs[key] //开销最低的节点
			node = key
		}
	}
	return
}
