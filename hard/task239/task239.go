package task239

func MaxSlidingWindow(nums []int, k int) []int {
	var result []int
	deque := &Deque{}

	for i := 0; i < len(nums); i++ {
		for deque.Len() > 0 && deque.Front() < i-k+1 && deque.Front() != -1 {
			deque.RemoveFirst()
		}

		for deque.Len() > 0 && nums[deque.Back()] < nums[i] {
			deque.RemoveLast()
		}

		deque.PushBack(i)

		if i >= k-1 {
			result = append(result, nums[deque.Front()])
		}
	}
	return result
}

type Deque struct {
	items []int
}

func (d *Deque) PushBack(val int) {
	d.items = append(d.items, val)
}

func (d *Deque) Front() int {
	if d.Len() != 0 {
		return d.items[0]
	}
	return -1
}

func (d *Deque) Back() int {
	if d.Len() != 0 {
		return d.items[d.Len()-1]
	}
	return -1
}

func (d *Deque) Len() int {
	return len(d.items)
}

func (d *Deque) RemoveFirst() {
	if len(d.items) > 0 {
		d.items = d.items[1:]
	}
}

func (d *Deque) RemoveLast() {
	if len(d.items) > 0 {
		d.items = d.items[:len(d.items)-1]
	}
}
