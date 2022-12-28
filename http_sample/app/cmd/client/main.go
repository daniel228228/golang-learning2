package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Matrixes struct {
	A Matrix `json:"a"`
	B Matrix `json:"b"`
}

type Matrix struct {
	N    int     `json:"n"`
	M    int     `json:"m"`
	Nums [][]int `json:"nums"`
}

type errorResponse struct {
	Err string `json:"error"`
}

func genMatrix(n, m int) *Matrix {
	result := &Matrix{
		N:    n,
		M:    m,
		Nums: make([][]int, n),
	}

	for i := 0; i < n; i++ {
		result.Nums[i] = make([]int, m)

		for j := 0; j < m; j++ {
			result.Nums[i][j] = i*n + j
		}
	}

	return result
}

func main() {
	client := &http.Client{}

	matrixes := &Matrixes{
		A: *genMatrix(3, 3),
		B: *genMatrix(3, 3),
	}

	res, _ := json.Marshal(matrixes)
	b := bytes.NewReader(res)

	req, _ := http.NewRequest("POST", "http://localhost:80/matrix", b)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	// result := &Matrix{}
	// errorResponse := &errorResponse{}

	// decoder := json.NewDecoder(resp.Body)
	// decoder.DisallowUnknownFields()

	// if err := decoder.Decode(errorResponse); err != nil {
	// 	if err2 := decoder.Decode(result); err2 != nil {
	// 		panic(err2)
	// 	}
	// } else {
	// 	fmt.Println(errorResponse.Err)
	// 	return
	// }

	// for i := range result.Nums {
	// 	for j := range result.Nums[i] {
	// 		fmt.Print(result.Nums[i][j], " ")
	// 	}

	// 	fmt.Println("")
	// }
}
