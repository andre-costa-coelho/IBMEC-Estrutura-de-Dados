// 4 + n + 1 + (4 n + n + 1 + 3) n
// n + 5 + 5n^2 + 4n
// 5n^2 + 5n + 5

package main

func procuraMatriz(matriz [][]int , n int, x int) bool{
	i := 0
	for i < n {
		j := 0
		for j < n {
			if matriz[i][j] == x {
				return true // Valor encontrado
			}
		}
		j ++
	}
	i++
return false // Valor não encontrado
}