package main

func quickSort2(sortArr []int, left, right int) {

	if left < right {
		p := (left + right) / 2
		key := sortArr[p]
		i, j := left, right

		for {
			// fmt.Println("111")
			in := false
			for sortArr[i] < key {
				// fmt.Println("111111111111")
				in = true
				i++
			}
			// fmt.Println("222")
			for sortArr[j] > key {
				// fmt.Println("222222222222")
				in = true
				j--
			}
			// fmt.Println("333")
			// fmt.Println(i)
			// fmt.Println(j)
			if i >= j || !in {
				break
			}
			// fmt.Println("444")
			sortArr[i], sortArr[j] = sortArr[j], sortArr[i]
		}

		quickSort2(sortArr, left, p-1)
		quickSort2(sortArr, p+1, right)
	}
}

// QuickSort2 lalala
func QuickSort2(val []int) {
	quickSort2(val, 0, len(val)-1)
}
