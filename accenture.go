package main

import (
	"fmt"
	"strings"
)

type Parametro struct {
	File string
	Sql  string
}

func main() {
	var p Parametro
	p.File =
		`Pais;id_operacao
Brasil;1
Argentina;2
Uruguai;3`

	p.Sql = `select * from operacao where id_operacao = <<id_operacao>>;`

	splitLinha := "\n"

	splitColuna := ";"

	linhas := strings.Split(p.File, splitLinha)

	colunas := [][]string{}

	for _, s := range linhas {
		l := strings.Split(s, splitColuna)
		var coluna []string
		for _, c := range l {
			coluna = append(coluna, c)
		}
		colunas = append(colunas, coluna)
	}

	cabecalho := colunas[0]

	aux1 := strings.Index(p.Sql, "<<") + 2

	aux2 := strings.Index(p.Sql, ">>")

	variavel := p.Sql[aux1:aux2]

	var index int

	for i, s := range cabecalho {
		if variavel == s {
			index = i
		}
	}

	for i, s := range colunas {
		if i != 0 {
			fmt.Println(strings.ReplaceAll(p.Sql, "<<"+variavel+">>", s[index]))
		}
	}
}
