package utils

// PriorityQueue to struktura implementująca kolejkę priorytetową
type PriorityQueue[T any] struct {
	elements []T               // elementy w kolejce
	cmpFunc  func(a, b T) bool // funkcja porównująca elementy
}

// NewPriorityQueue tworzy nową kolejkę priorytetową
func NewPriorityQueue[T any](cmpFunc func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{cmpFunc: cmpFunc}
}

// Push dodaje element do kolejki
func (pq *PriorityQueue[T]) Push(element T) {
	pq.elements = append(pq.elements, element)
	pq.heapifyUp(len(pq.elements) - 1)
}

// Pop zdejmuje element z kolejki
func (pq *PriorityQueue[T]) Pop() T {
	if pq.IsEmpty() {
		return *new(T)
	}

	top := pq.elements[0]
	pq.elements[0] = pq.elements[len(pq.elements)-1]
	pq.elements = pq.elements[:len(pq.elements)-1]

	pq.heapifyDown(0)

	return top
}

// IsEmpty zwraca true, jeśli kolejka jest pusta
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.elements) == 0
}

// GetElements zwraca elementy w kolejce
func (pq *PriorityQueue[T]) GetElements() []T {
	return pq.elements
}

// heapifyUp naprawia kopiec w górę
func (pq *PriorityQueue[T]) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.cmpFunc(pq.elements[index], pq.elements[parentIndex]) {
			pq.elements[index], pq.elements[parentIndex] = pq.elements[parentIndex], pq.elements[index]
			index = parentIndex
		} else {
			break
		}
	}
}

// heapifyDown naprawia kopiec w dół
func (pq *PriorityQueue[T]) heapifyDown(index int) {
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallest := index

		if leftChildIndex < len(pq.elements) && pq.cmpFunc(pq.elements[leftChildIndex], pq.elements[smallest]) {
			smallest = leftChildIndex
		}

		if rightChildIndex < len(pq.elements) && pq.cmpFunc(pq.elements[rightChildIndex], pq.elements[smallest]) {
			smallest = rightChildIndex
		}

		if smallest != index {
			pq.elements[index], pq.elements[smallest] = pq.elements[smallest], pq.elements[index]
			index = smallest
		} else {
			break
		}
	}
}
