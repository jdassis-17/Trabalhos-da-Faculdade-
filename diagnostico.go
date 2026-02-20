package main

import "fmt"

/*
Resposta do Diagnóstico:

	1. Na falha “A”, temos um problema de escalabilidade, pois ocorreu um aumento no tempo de resposta, passando de 2 segundos para 45 segundos um acréscimo de 43 segundos no tempo esperado para cada resposta. Dessa forma, percebemos que o sistema não conseguiu manter o desempenho conforme a carga aumentou.

	2. Na falha “B”, ocorreu um problema de confiabilidade, pois os dados de um cliente foram expostos para outro cliente. Nesse caso, os dados de pagamento estavam expostos e não devidamente protegidos pelo sistema. Percebemos que o sistema não está operando corretamente e não está funcionando de forma segura para os usuários.

	3. Na falha “C”, o problema é de manutenibilidade, pois uma manutenção de urgência, que deveria ser rápida, tornou-se uma manutenção demorada, levando vários dias. O sistema está funcionando, mas não apresenta facilidade para corrigir, modificar ou evoluir suas funcionalidades.
*/

func main() {
	fmt.Println("Ambiente configurado por: José Diego de Lima Assis")

	ftm.Println("Ambiente configurado e desafio analisado!")
}
