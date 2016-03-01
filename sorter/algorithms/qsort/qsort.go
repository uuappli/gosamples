package qsort

func quickSort(values []int, left, reght int) {
	temp := values[left]
	p := left

	i, j := left, reght

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}

		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}

	if reght-p > 1 {
		quickSort(values, p+1, reght)
	}
}


func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
