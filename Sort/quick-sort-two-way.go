package main

var isPrint = true

// QSortTwoWay balbalbal.....
// func QSortTwoWay(arr []int) {

// 	arrLen := len(arr)

// 	if arrLen <= 1 {
// 		return
// 	}

// 	randNum := arrLen / 2
// 	if isPrint {
// 		fmt.Println("--------------begin---------------")
// 		fmt.Println(arr)
// 		fmt.Println("randNum: ", randNum)
// 		fmt.Println("randNum val: ", arr[randNum])
// 		fmt.Println("arr 0: ", arr[0])
// 	}

// 	// arr[randNum], arr[0] = arr[0], arr[randNum]
// 	mid := arr[randNum]

// 	head, tail := 0, arrLen-1

// 	for {
// 		in1 := false
// 		in2 := false

// 		for head <= tail && arr[head] < mid {
// 			in1 = true
// 			head++
// 		}

// 		for head <= tail && arr[tail] > mid {
// 			in2 = true
// 			tail--
// 		}

// 		if isPrint {
// 			fmt.Println("in1: ", in1)
// 			fmt.Println("in2: ", in2)
// 		}
// 		if head >= tail {
// 			break
// 		}

// 		arr[head], arr[tail] = arr[tail], arr[head]
// 		head++
// 		tail--
// 	}

// 	if isPrint {
// 		fmt.Println(arr)
// 		fmt.Println("head: ", head)
// 		// fmt.Println("head val: ", arr[head])
// 		fmt.Println("--------------end---------------")
// 	}

// 	QSortTwoWay(arr[:head])
// 	QSortTwoWay(arr[head+1:])
// }

func partitionTwoWay(arr []int, l, r int) int {
	mid := arr[l]

	left := l + 1
	right := r

	for {
		switch {
		case left > right:
			arr[l], arr[right] = arr[right], arr[l]
			return right
		case arr[left] < mid:
			left++
		case arr[right] > mid:
			right--
		default:
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
}

func qSortTwoWay(arr []int, l, r int) {
	if l > r {
		return
	}

	p := partitionTwoWay(arr, l, r)

	qSortTwoWay(arr, l, p-1)
	qSortTwoWay(arr, p+1, r)
}

// QSortTwoWay bbbbb...
func QSortTwoWay(arr []int) {
	qSortTwoWay(arr, 0, len(arr)-1)
}
