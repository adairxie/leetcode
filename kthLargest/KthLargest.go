package main

type heap struct {
	maxsize int
	size    int
	data    []int
}

type KthLargest struct {
	k    int
	heap *heap
}

func Constructor(k int, nums []int) KthLargest {
	kheap := KthLargest{
		k:    k,
		heap: &heap{maxsize: k, data: make([]int, 1, k+1)},
	}

	for i := 0; i < len(nums); i++ {
		kheap.Add(nums[i])
	}

	return kheap
}

func Min(heap []int, i, j int) int {
	if heap[i] < heap[j] {
		return i
	}

	return j
}

func (this *KthLargest) Add(val int) int {
	println("添加: ", val, this.heap.size)
	if this.heap.size >= this.k {
		this.heap.pop()
	}

	this.heap.push(val)
	return this.heap.data[1]
}

func (h *heap) push(val int) {
	if h.size >= h.maxsize && val <= h.data[1] {
		return
	}

	if h.size >= h.maxsize && val > h.data[1] {
		h.pop()
	}

	if h.size < h.maxsize {
		h.size++
		h.data = append(h.data, val)
	}

	println("push: ", val, " size: ", h.size, ", data: ", len(h.data))
	for i := h.size / 2; i > 0; i = i / 2 {
		if h.data[2*i] >= h.data[i] {
			break
		}

		h.data[i], h.data[2*i] = h.data[2*i], h.data[i]
	}
}

func (h *heap) pop() {
	v := h.data[1]
	h.data[1], h.data[h.size] = h.data[h.size], h.data[1]
	h.data = h.data[:h.size]
	h.size--

	for i := 1; i <= h.size; {
		var index int
		left := 2 * i
		right := 2*i + 1
		if left <= h.size && right <= h.size {
			println("left: ", left, ", right: ", right)
			index = Min(h.data, left, right)
		} else if left <= h.size {
			index = left
		} else {
			index = right
		}
		println("弹出：", v, ", size: ", h.size, ", index: ", index)

		if h.data[i] < h.data[index] {
			h.data[i], h.data[index] = h.data[index], h.data[i]
			i = index
		} else {
			break
		}
	}
}

func main() {
	kth := Constructor(3, []int{1})
	kth.Add(2)
	kth.Add(3)
	kth.Add(4)
	kth.Add(5)
	kth.Add(1)
	for i := 1; i < kth.heap.size+1; i++ {
		println(kth.heap.data[i])
	}
}
