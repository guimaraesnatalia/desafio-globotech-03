package main

import (
	"fmt"
	"os"
	"github.com/jedib0t/go-pretty/v6/table"
)

type compra struct{
	comprador string
	qtd int
	total float32
}

func main(){
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	var bilheteria []compra

	var compraTemp compra
	const valor = 10.50
	var reserva string
	var novaCompra string

	fmt.Println("Bem vindo ao sistema de reserva de tickets")
	
	for{
		fmt.Printf("Qual seu nome? ")
		fmt.Scanln(&compraTemp.comprador)

		fmt.Printf("Quantos tickets você quer reservar? ")
		fmt.Scan(&compraTemp.qtd)
		compraTemp.total = valor * float32(compraTemp.qtd)

		fmt.Printf("%s, você confirma a reserva de %d tickets no valor de %.2f reais? (s/n)\n", compraTemp.comprador, compraTemp.qtd, compraTemp.total )
		fmt.Scan(&reserva)

		if (reserva == "s") {
			bilheteria = append(bilheteria, compraTemp)
			fmt.Println("Wow.. seus tickets foram reservados!!!")
		}else if (reserva == "n"){
			continue;
		}

		fmt.Println("Quer reservar mais tickets? (s/n)")
		fmt.Scan(&novaCompra)

		if (novaCompra == "n"){
			fmt.Println("Tickets reservados: ")
			// fmt.Println( "Nome: Qtd: Total: ")
			t.AppendHeader(table.Row{ "Nome", "Quantidade", "Total"})
			for i := 0; i<len(bilheteria); i++ {
				t.AppendRow(table.Row{bilheteria[i].comprador, bilheteria[i].qtd, bilheteria[i].total})				
				// fmt.Println(bilheteria[i].comprador, " - ", bilheteria[i].qtd, " - ", bilheteria[i].total)
			}
			t.Render()
			break;
		}

	}
}