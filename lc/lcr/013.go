package main

type NumMatrix struct {
	nums [][]int
	s    [][]int
}

//func Constructor(matrix [][]int) NumMatrix {
//	if matrix == nil {
//		return NumMatrix{}
//	}
//	n, m := len(matrix), len(matrix[0])
//	s := make([][]int, n+1)
//	for i := 0; i <= n; i++ {
//		s[i] = make([]int, m+1)
//	}
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= m; j++ {
//			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + matrix[i-1][j-1]
//		}
//	}
//	for i := 1; i <= n; i++ {
//		fmt.Println(s[i])
//	}
//
//	nm := NumMatrix{
//		nums: matrix,
//		s:    s,
//	}
//	nm.SumRegion(1, 1, 1, 1)
//	return nm
//}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.s[row2][col2] - this.s[row1-1][row2] - this.s[row2][col1-1] + this.s[row1-1][col1-1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
