type Dequeue struct {
	prev  *Dequeue
	next  *Dequeue
	key   int
	value int
}

type LRUCache struct {
	cache      map[int]*Dequeue
	head, tail *Dequeue
	capacity   int
}

func Constructor(capacity int) LRUCache {
	head := &Dequeue{key: -1, value: -1}
	tail := &Dequeue{key: -1, value: -1}
	head.next = tail
	tail.prev = head

	return LRUCache{
		cache:    make(map[int]*Dequeue),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

// TimeComplexity: O(1)
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.bringToFront(node)
		return node.value
	}
}

// TimeComplexity: O(1)
func (this *LRUCache) Put(key int, value int) {
	// If capacity, remove last item from cache
	if node, ok := this.cache[key]; ok {
		node.key = key
		node.value = value

		this.bringToFront(node)
	} else {
		if len(this.cache) == this.capacity {
			this.removeBack()
		}
		node := &Dequeue{key: key, value: value}
		this.pushFront(node)
		this.cache[key] = node
	}
}

func (this *LRUCache) bringToFront(node *Dequeue) {
	if node.prev != this.head {
		this.remove(node)
		this.pushFront(node)
	}
}

func (this *LRUCache) remove(node *Dequeue) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

func (this *LRUCache) pushFront(node *Dequeue) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeBack() {
	prev := this.tail.prev
	this.tail.prev = prev.prev
	prev.prev.next = this.tail

	delete(this.cache, prev.key)
	prev.prev = nil
	prev.next = nil
}