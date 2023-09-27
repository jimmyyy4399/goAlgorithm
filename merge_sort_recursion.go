package main

func Merge_Sort_recursion(A []int, N int) {
	TmpA := make([]int, N)

	Msort(A, TmpA, 0, N-1)
}

func Msort(A []int, TmpA []int, L int, RightEnd int) {
	if L < RightEnd {
		Center := (L + RightEnd) / 2
		Msort(A, TmpA, L, Center)
		Msort(A, TmpA, Center+1, RightEnd)
		Merge_r(A, TmpA, L, Center+1, RightEnd)
	}
}

func Merge_r(A []int, TmpA []int, L int, R int, RightEnd int) {
	i := L
	Tmp := L
	LeftEnd := R - 1
	for L <= LeftEnd && R <= RightEnd {
		if A[L] < A[R] {
			TmpA[Tmp] = A[L]
			Tmp++
			L++
		} else {
			TmpA[Tmp] = A[R]
			Tmp++
			R++
		}
	}
	for L <= LeftEnd {
		TmpA[Tmp] = A[L]
		Tmp++
		L++
	}
	for R <= RightEnd {
		TmpA[Tmp] = A[R]
		Tmp++
		R++
	}
	for ; i <= RightEnd; i++ {
		A[i] = TmpA[i]
	}
}

/* func main() {
	matrix := [][]int{
		{4, 3, 2, 1},
		{1, 2, 3, 4},
		{},
		{5},
		{-1, -2, -3, -4},
		{10, -10, 20, -20},
		{5, 5, 5, 5, 5},
		{1, 100, 2, 200, 3, 300},
	}
	for _, arr := range matrix {
		Merge_Sort_recursion(arr, len(arr))
	}
	fmt.Println(matrix)
}
*/
