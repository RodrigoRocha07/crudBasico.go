package main

import "fmt"

func main() {
	nomes := []string{}
	for {
		fmt.Println(" 1 - Listar \n 2 - Acrescentar \n 3 - Atualizar \n 4 - Deletar \n 5 - Deletar varios \n 0 - Finalizar")
		comando := 0
		fmt.Scan(&comando)
		if comando == 1 {
			if len(nomes) == 0 {
				fmt.Println("Nenhum nome na lista")
			} else {
				fmt.Println(nomes)
			}
		} else if comando == 2 {
			fmt.Println("Quantos nomes deseja acrescentar?")
			var nAcrescentar int
			fmt.Scan(&nAcrescentar)
			for i := 0; i < nAcrescentar; i++ {
				fmt.Println("Digite o nome: ")
				var novoNome string
				fmt.Scan(&novoNome)
				nomes = append(nomes, novoNome)
			}

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
		} else if comando == 5 {
			fmt.Println("Quantos elementos deseja deletar? ")
			var nDelete int
			fmt.Scan(&nDelete)
			if nDelete > len(nomes) {
				fmt.Println("Valor maior do que a quantidade de nomes!")
			} else {
				for i := 0; i < nDelete; i++ {
					for i := 0; i < len(nomes); i++ {
						fmt.Println(i, "-", nomes[i])
					}
					var deletar int
					fmt.Scan(&deletar)

					nomes = append(nomes[:deletar], nomes[deletar+1:]...)

				}

			}

		} else {
			break
		}
	}
}
