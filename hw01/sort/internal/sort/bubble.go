package sort

func BubbleSort(a []int) {
	for i := 0; i < len(a) - 1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}