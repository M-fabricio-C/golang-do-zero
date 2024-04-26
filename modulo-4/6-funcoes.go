package main

import "fmt"

func main() {
    somaDosValores := soma(42, 13)
    fmt.Println(somaDosValores)

    nome, sobrenome, idade := printaNomeCompleto("JOÃO LEVI", "GADELHA", "14")
    fmt.Println(nome)
    fmt.Println(sobrenome)
    fmt.Println(idade)

}

// Função começando com letra minúscula:
// Função é PRIVADA
// Função ela só pode ser utilizada no próprio pacote
func printaNomeCompleto(nome, sobrenome, idade string) (string, string, string) {
    return nome, sobrenome, idade
}

// Função começando com letra maiuscula:
// Função é PÚBLICA
// Função ela só pode ser utilizada fora do próprio pacote
// Como utilizaria ela fora do pacote: main.PrintaNome(nome)
//func printaNomeCompleto(nome string) string {
//    return nome
//}

func soma(x int, y int) int {
    return x + y
}