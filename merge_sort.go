package main

func Merge_Sort(A []int, N int) {
	length := 1 /* 初始化子序列长度 */
	TmpA := make([]int, N)
	for length <= N {
		Merge_Pass(A, TmpA, N, length)
		length *= 2
		Merge_Pass(TmpA, A, N, length)
		length *= 2
	}

}

func Merge_Pass(A []int, TmpA []int, N int, length int) {
	/* 两两归并有序子序列 */
	var i, j int

	for i = 0; i <= N-2*length; i += 2 * length {
		Merge(A, TmpA, i, i+length, i+2*length-1)
	}
	if i+length < N { /* 归并最后两个子序列 */
		Merge(A, TmpA, i, i+length, N-1)
	} else { /* 最后只剩一个子序列 */
		for j = i; j < N; j++ {
			TmpA[j] = A[j]
		}
	}
}
func Merge(A []int, TmpA []int, L int, R int, RightEnd int) {
	/* 将有序的A[L]~A[R-1]和A[R]~A[RightEnd]归并成一个有序序列 */
	LeftEnd := R - 1
	Tmp := L
	//NumElements := RightEnd - L + 1

	for L <= LeftEnd && R <= RightEnd {
		if A[L] <= A[R] {
			TmpA[Tmp] = A[L] /* 将左边元素复制到TmpA */
			Tmp++
			L++
		} else {
			TmpA[Tmp] = A[R] /* 将右边元素复制到TmpA */
			Tmp++
			R++
		}
	}
	for L <= LeftEnd {
		TmpA[Tmp] = A[L] /* 直接复制左边剩下的 */
		Tmp++
		L++
	}
	for R <= RightEnd {
		TmpA[Tmp] = A[R] /* 直接复制右边剩下的 */
		Tmp++
		R++
	}
	// for i := 0; i < NumElements; i++ {
	//  A[RightEnd] = TmpA[RightEnd]    /* 将有序的TmpA[]复制回A[] */
	//  RightEnd--
	// }
}
