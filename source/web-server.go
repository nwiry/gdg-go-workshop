package main // Pacote principal do programa - Aqui é onde toda a mágica acontece

import (
	"fmt"      // É a partir deste pacote que utilizamos a função para imprimir nosso texto na tela
	"net/http" // Este pacote contém tudo que é necessário para trabalharmos com servidores da internet
)

// handler - Retorna o conteudo que será exibido na pagina inicial do nosso servidor web
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!") // Escrevemos hello world em nossa pagina da web
}

// main - Tudo que estiver aqui dentro será utilizado em nosso código durante a sua execução
func main() {

	http.HandleFunc("/", handler) // Associa a função handler ao inicio do nosso servidor web (pagina inicial)

	// Inicia o servidor web na porta 8080. Logo, nosso servidor é acessivel em qualquer navegador web da nossa maquina
	// através do endereço: http://localhost:8080
	http.ListenAndServe(":8080", nil)
}
