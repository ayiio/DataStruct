package breadthfirstsearch

import (
	"fmt"

	"github.com/eapache/queue"
)

func BreadthFirstSRCH(searchMap map[string][]string, name string) bool {
	var myqueue = queue.New()
	for _, v := range searchMap[name] {
		myqueue.Add(v) //FIFO，辅助遍历队列
	}
	var searched []string //记录已被检查的节点
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("queue empty")
		}
	}()
	for head := myqueue.Peek(); head != nil; head = myqueue.Remove() {
		headKey := head.(string)
		var in bool
		for _, sd := range searched {
			if sd == headKey {
				in = true
			}
		}
		if !in { //节点没被检查时才做检查
			if headKey[len(headKey)-1:] == "m" {
				fmt.Println("find target")
				return true
			} else {
				for _, v := range searchMap[headKey] {
					myqueue.Add(v) //未命中，追加该节点的邻居节点
				}
				searched = append(searched, headKey) //标记节点为已检查
			}
		}
	}
	return false
}
