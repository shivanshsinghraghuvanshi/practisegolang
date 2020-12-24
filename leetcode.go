package main

import "container/list"

//MinPartitions
func MinPartitions(n string) int {
	runes := []rune(n)
	max := 0
	for _, e := range runes {
		if max < int(e-'0') {
			max = int(e - '0')
		}
	}
	return max
}

func NumIslands(grid [][]byte) int {
	nr := len(grid)
	nc := len(grid[0])
	numIslands := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				numIslands += 1
				dfs(grid, r, c)
			}
		}
	}
	return numIslands
}

func dfs(grid [][]byte, r int, c int) {
	nr := len(grid)
	nc := len(grid[0])
	if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

type LRUCache struct {
	capacity int
	m        map[int]*list.Element
	l        *list.List
}
type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*list.Element, capacity),
		l:        new(list.List),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		value := node.Value.(*list.Element).Value.(Pair).value
		this.l.MoveToFront(node)
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.l.MoveToFront(node)
		node.Value.(*list.Element).Value = Pair{
			key:   key,
			value: value,
		}
	} else {
		if this.l.Len() == this.capacity {
			delKey := this.l.Back().Value.(*list.Element).Value.(Pair).key
			delete(this.m, delKey)
			this.l.Remove(this.l.Back())
		}
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		ptr := this.l.PushFront(node)
		this.m[key] = ptr
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
