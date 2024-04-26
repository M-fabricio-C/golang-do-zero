package main

import "fmt"

func main() {
//	idade := map[string]int{}
//	idade["levi"] = 28
//	idade["fabricio"] = 4
//	fmt.Println(idade)
//	fmt.Println(idade["levi"])
//	fmt.Println(idade["fabricio"])

    anoNasc := map[string]int{
	    "levi": 2010,
	    "fabricio": 1988,
    }

    fmt.Println(anoNasc)
    fmt.Println(anoNasc["levi"])
    fmt.Println(anoNasc["fabricio"])
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)

}