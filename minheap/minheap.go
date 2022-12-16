package minheap

import "fmt"

type StringMinHeap struct {
    heapArray []string
    size      int
    maxsize   int
}

func NewMinHeap(maxsize int) *StringMinHeap {
    minheap := &StringMinHeap{
        heapArray: []string{},
        size:      0,
        maxsize:   maxsize,
    }
    return minheap
}

func (m *StringMinHeap) leaf(index int) bool {
    if index >= (m.size/2) && index <= m.size {
        return true
    }
    return false
}

func (m *StringMinHeap) parent(index int) int {
    return (index - 1) / 2
}

func (m *StringMinHeap) leftchild(index int) int {
    return 2*index + 1
}

func (m *StringMinHeap) rightchild(index int) int {
    return 2*index + 2
}

func (m *StringMinHeap) Insert(item string) error {
    if m.size >= m.maxsize {
        return fmt.Errorf("Heap is full")
    }
    m.heapArray = append(m.heapArray, item)
    m.size++
    m.upHeapify(m.size - 1)
    return nil
}

func (m *StringMinHeap) swap(first, second int) {
    temp := m.heapArray[first]
    m.heapArray[first] = m.heapArray[second]
    m.heapArray[second] = temp
}

func (m *StringMinHeap) upHeapify(index int) {
    for m.heapArray[index] < m.heapArray[m.parent(index)] {
        m.swap(index, m.parent(index))
    }
}

func (m *StringMinHeap) downHeapify(current int) {
    if m.leaf(current) {
        return
    }
    smallest := current
    leftChildIndex := m.leftchild(current)
    rightRightIndex := m.rightchild(current)
    //If current is smallest then return
    if leftChildIndex < m.size && m.heapArray[leftChildIndex] < m.heapArray[smallest] {
        smallest = leftChildIndex
    }
    if rightRightIndex < m.size && m.heapArray[rightRightIndex] < m.heapArray[smallest] {
        smallest = rightRightIndex
    }
    if smallest != current {
        m.swap(current, smallest)
        m.downHeapify(smallest)
    }
    return
}
func (m *StringMinHeap) buildMinHeap() {
    for index := ((m.size / 2) - 1); index >= 0; index-- {
        m.downHeapify(index)
    }
}

func (m *StringMinHeap) Remove() string {
    top := m.heapArray[0]
    m.heapArray[0] = m.heapArray[m.size-1]
    m.heapArray = m.heapArray[:(m.size)-1]
    m.size--
    m.downHeapify(0)
    return top
}

func (m *StringMinHeap) Init(inputArray []string){
	for i := 0; i < len(inputArray); i++ {
		m.Insert(inputArray[i])
	}
	m.buildMinHeap()
}

func (m *StringMinHeap) GetHeapSize() int{
   return m.size
}