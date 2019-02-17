package routine

import "sync"

var Waitgroup = sync.WaitGroup{}

func QuickSort01(array []int) {
	if len(array) == 1 || len(array) == 0 || array == nil {
		return
	}
	key := array[0]
	r := len(array) - 1
	l := 1
	for r != l {
		if array[l] > key {
			swap(&array[r], &array[l])
			r--
		} else {
			l++
		}
	}
	if array[r] > key {
		swap(&array[0], &array[r-1])
	} else {
		swap(&array[0], &array[r])
	}
	QuickSort01(array[:r])
	QuickSort01(array[r:])
	return
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func QuickSort02(array []int) {
	if len(array) == 1 || len(array) == 0 || array == nil {
		Waitgroup.Done()
		return
	}
	key := array[0]
	r := len(array) - 1
	l := 1
	for r != l {
		if array[l] > key {
			swap(&array[r], &array[l])
			r--
		} else {
			l++
		}
	}
	if array[r] > key {
		swap(&array[0], &array[r-1])
	} else {
		swap(&array[0], &array[r])
	}
	Waitgroup.Add(2)
	go QuickSort02(array[:r])
	go QuickSort02(array[r:])
	Waitgroup.Done()
}