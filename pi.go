// https://stackoverflow.com/questions/11677369/how-to-calculate-pi-to-n-number-of-places-in-c-sharp-using-loops

package main

import (
	"net/http"
	"strconv"
)

func calc(response http.ResponseWriter, request *http.Request) {
	digits := request.URL.Query().Get("digits")
	if digits == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Digits not specified."))
		return
	}
	d, err := strconv.ParseInt(digits, 10, 64)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid number of digits."))
		return
	}
	x := make([]int, d*10/3+1)
	r := make([]int, d*10/3+1)
	pi := make([]int, d)
	for i := 0; i < len(x); i++ {
		x[i] = 20
	}
	for i := int64(0); i < d; i++ {
		carry := 0
		for j := 0; j < len(x); j++ {
			num := len(x) - j - 1
			dem := num*2 + 1
			x[j] += carry
			q := x[j] / dem
			r[j] = x[j] % dem
			carry = q * num
		}
		pi[i] = x[len(x)-1] / 10
		r[len(x)-1] = x[len(x)-1] % 10
		for j := 0; j < len(x); j++ {
			x[j] = r[j] * 10
		}
	}
	result := ""
	c := 0
	for i := len(pi) - 1; i >= 0; i-- {
		pi[i] += c
		c = pi[i] / 10
		result = strconv.Itoa(pi[i]%10) + result
	}
	response.Write([]byte(result))
}

func main() {
	http.HandleFunc("/", calc)
	http.ListenAndServe(":80", nil)
}
