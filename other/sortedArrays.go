package other

//
//import (
//	"container/heap"
//	"container/list"
//	"fmt"
//	"net/http"
//)
//
//func MergeTwoSortedArrays(x []int, y []int) []int {
//	z := make([]int, len(x)+len(y))
//	i, j, k := 0, 0, 0
//	for ok := true; ok; ok = i < len(x) && j < len(y) {
//		if x[i] <= y[j] {
//			z[k] = x[i]
//			i++
//		} else {
//			z[k] = y[j]
//			j++
//		}
//		k++
//	}
//	if len(x) > len(y) {
//		copy(z[k:], x[i:])
//	} else {
//		copy(z[k:], y[j:])
//	}
//	return z
//}
//
//func ReverseString(s string) string {
//	rns := []rune(s)
//	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
//		rns[i], rns[j] = rns[j], rns[i]
//	}
//	return string(rns)
//}
//func checkLink(link string, c chan string) {
//	_, err := http.Get(link)
//	if err != nil {
//		fmt.Println(link, " might be down")
//		c <- link
//		return
//	}
//	fmt.Println(link, " is up")
//	c <- link
//}
//
//func contains(arr []string, str string) bool {
//	for _, a := range arr {
//		if a == str {
//			return true
//		}
//	}
//	return false
//}
//func runningSum(nums []int) []int {
//	p := make([]int, len(nums))
//	p[0] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		p[i] = p[i-1] + nums[i]
//	}
//	return p
//}
//func joinTwoSortedArrays() {
//	x := []int{1, 4, 6, 8, 9}
//	y := []int{2, 3, 6}
//	z := make([]int, len(x)+len(y))
//	i := 0
//	j := 0
//	k := 0
//	for ok := true; ok; ok = i < len(x) && j < len(y) {
//		if x[i] <= y[j] {
//			z[k] = x[i]
//			i++
//			k++
//		} else {
//			z[k] = y[j]
//			j++
//			k++
//		}
//	}
//	if len(x) > len(y) {
//		copy(z[k:], x[i:])
//	} else {
//		copy(z[k:], y[j:])
//	}
//	fmt.Println(z)
//}
//
//func isPalindroome(a string) bool {
//	fifo := list.New()
//	palindrome := true
//	for _, item := range a {
//		fifo.PushBack(item)
//	}
//	lifo := list.New()
//	for _, item := range a {
//		lifo.PushBack(item)
//	}
//	for i := 0; i < len(a); i++ {
//		front := lifo.Front()
//		back := fifo.Back()
//		if front.Value == back.Value {
//			fifo.Remove(back)
//			lifo.Remove(front)
//		} else {
//			palindrome = false
//		}
//
//	}
//	return palindrome
//}
//
//type Element struct {
//	Value interface{}
//	Priority int
//	index int
//}
//
//type MinHeap []*Element
//
//func (m MinHeap) Len() int {
//	return  len(m)
//}
//
//func (m MinHeap) Less(i, j int) bool {
//	panic("implement me")
//}
//
//func (m MinHeap) Swap(i, j int) {
//	panic("implement me")
//}
//
//func (m MinHeap) Push(x interface{}) {
//	panic("implement me")
//}
//
//func (m MinHeap) Pop() interface{} {
//	panic("implement me")
//}
//
//var elem int
//func NewMinHeap()*MinHeap{
//	mh:=&MinHeap{}
//	heap.Init(mh)
//	return mh
//}
//func NewIntegerHeap()*elem{
//	mh:=&elem
//	heap.Init(mh)
//	return mh
//}
//func MergeKSortedArrays(input [][]int) []int {
//	z := make([]int, 50)
//
//	return z
//}
