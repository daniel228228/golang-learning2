package dto

type Matrixes struct {
	A Matrix `json:"a"`
	B Matrix `json:"b"`
}

type Matrix struct {
	N    int     `json:"n"`
	M    int     `json:"m"`
	Nums [][]int `json:"nums"`
}
