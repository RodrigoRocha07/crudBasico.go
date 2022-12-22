package main

import "fmt"

func main() {
	nomes := []string{}
	for {
		fmt.Println(" 1 - Listar \n 2 - Acrescentar \n 3 - Atualizar \n 4 - Deletar \n 0 - Finalizar")
		comando := 0
		fmt.Scan(&comando)
		if comando == 1 {
			if len(nomes) == 0 {
				fmt.Println("Nenhum nome na lista")
			} else {
				fmt.Println(nomes)
			}
		} else if comando == 2 {
			fmt.Println("Digite o nome: ")
			var novoNome string
			fmt.Scan(&novoNome)
			nomes = append(nomes, novoNome)
		} else if comando == 3 {
			if len(nomes) == 0 {
				fmt.Println("Nenhum nome para atualizar!")
			} else {
				fmt.Println("Qual nome voce deseja mudar?")
				for i := 0; i < len(nomes); i++ {
					fmt.Println(i, "-", nomes[i])
				}
				var atualizar int
				fmt.Scan(&atualizar)
				fmt.Println("Digite o novo nome: ")
				var novoNome string
				fmt.Scan(&novoNome)
				nomes[atualizar] = (novoNome)

			}

		} else if comando == 4 {
			if len(nomes) == 0 {
				fmt.Println("Nenhum nome para deletar")
			} else {
				for i := 0; i < len(nomes); i++ {
					fmt.Println(i, "-", nomes[i])
				}
				var deletar int
				fmt.Scan(&deletar)

				nomes = append(nomes[:deletar], nomes[deletar+1:]...)
			}
		} else {
			break
		}
	}
}
