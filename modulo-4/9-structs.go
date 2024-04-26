package main 

import "fmt"


// Structs
// Forma de criar sua própria estrutura de dados
// Pensonalizar de acordo com a sua nessessidade
// Podemos usar vários tipos diferentes

// type <nome da nossa estrutura> struct { <campos> }
type Pessoa struct {
		Nome string
		Idade int 
}

type Graduacao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"fabricio", 36})
	fmt.Println(Pessoa{Nome: "Levi", Idade: 14})
	fmt.Println(Pessoa{Nome: "Regi"})
		
	p1 := Pessoa{Nome: "Bob", Idade: 2}
//	fmt.Println(p1.Nome)
//	fmt.Println(p1.Idade)

//	p1.Idade = 3
//	fmt.Println((p1.Idade))

	p2 := Pessoa{Nome: "Regi", Idade: 36}

    pessoas := []Pessoa{}
    pessoas = append(pessoas, p1, p2)
//	fmt.Println(pessoas)

// structs + map
//    alunos := map[string] []Pessoa{}
//    alunos["programação"] = pessoas
//    fmt.Println(alunos)

//    var alunos = map[string] []Pessoa{
//        "Programação": {{Nome: "Fabricio", Idade: 36}, {Nome: "Levi", Idade: 14}},
//        "Engenharia": {{Nome: "Fabricio2", Idade: 36}, {Nome: "Levi2", Idade: 14}},
//    }
//    fmt.Println(alunos)

// struct herdando campos de outra struct
grad := Graduacao{p2, "verde"}
fmt.Println(grad)
fmt.Println(grad.Pessoa.Nome)
fmt.Println(grad.Pessoa.Idade)
fmt.Println(grad.Tipo)

}

