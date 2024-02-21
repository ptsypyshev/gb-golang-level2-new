package merge

func MergeSortedSlices(a, b []int) []int {
	var idxA, idxB int
	lenRes := len(a) + len(b)
	if lenRes == 0 {
		return nil
	}

	res := make([]int, 0, lenRes)
	for {
		if idxA == len(a) {
			res = append(res, b[idxB:]...)
			return res
		}

		if idxB == len(b) {
			res = append(res, a[idxA:]...)
			return res
		}

		if a[idxA] < b[idxB] {
			res = append(res, a[idxA])
			idxA++
		} else if b[idxB] < a[idxA] {
			res = append(res, b[idxB])
			idxB++
		} else {
			res = append(res, a[idxA])
			res = append(res, b[idxB])
			idxA++
			idxB++
		}
	}
}
