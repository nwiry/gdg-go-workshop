# 💻 Oficina de Código em Go

Bem-vindo à oficina de código em Go! Nesta oficina, você aprenderá o básico da linguagem de programação Go, como configurá-la no seu ambiente de desenvolvimento, e explorar alguns dos recursos mais legais que ela oferece.

## 📖 Resumo Básico de Go

Go, também conhecido como Golang, é uma linguagem de programação desenvolvida pelo Google. Ela foi projetada para ser simples, eficiente e segura, com suporte a concorrência nativa. Alguns dos principais benefícios de usar Go incluem:

- **Simplicidade**: Sintaxe clara e concisa.
- **Desempenho**: Código compilado com alta performance.
- **Concorrência**: Suporte robusto para concorrência com goroutines e canais.
- **Ferramentas**: Conjunto de ferramentas integrado para facilitar o desenvolvimento.

## 🛠️ Instalação do Go

### Linux

1. Baixe o arquivo tarball da versão mais recente do Go do [site oficial](https://golang.org/dl/):
    ```sh
    wget https://golang.org/dl/go{version}.linux-amd64.tar.gz
    ```

2. Extraia o arquivo baixado:
    ```sh
    sudo tar -C /usr/local -xzf go{version}.linux-amd64.tar.gz
    ```

3. Adicione o Go ao PATH:
    ```sh
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    source ~/.profile
    ```

### MacOS

1. Baixe o pacote .pkg da versão mais recente do Go do [site oficial](https://golang.org/dl/).

2. Execute o pacote baixado e siga as instruções do instalador.

3. Verifique a instalação abrindo o terminal e digitando:
    ```sh
    go version
    ```

### Windows

1. Baixe o instalador .msi da versão mais recente do Go do [site oficial](https://golang.org/dl/).

2. Execute o instalador e siga as instruções.

3. Verifique a instalação abrindo o Command Prompt (cmd) e digitando:
    ```sh
    go version
    ```

## 🖥️ Instalando o Visual Studio Code

1. Baixe o instalador do VS Code do [site oficial](https://code.visualstudio.com/).

2. Siga as instruções para instalar o VS Code no seu sistema operacional.

3. Após a instalação, abra o VS Code e instale a extensão Go:
    - Clique em **Extensions** no menu lateral.
    - Pesquise por "Go" e instale a extensão oficial da **Go Team at Google**.

## ✨ Por que escolher Golang?

- **Simplicidade e Clareza**: Go é conhecido por sua sintaxe simples e clara, o que torna a leitura e a escrita do código mais fáceis.
- **Desempenho**: Por ser uma linguagem compilada, Go oferece um desempenho próximo ao de linguagens como C e C++.
- **Concorrência Facilitada**: O suporte nativo a goroutines e canais facilita a criação de programas concorrentes e paralelos.
- **Bibliotecas Padrão Poderosas**: A biblioteca padrão do Go é robusta e inclui pacotes para praticamente todas as necessidades básicas de um desenvolvedor.
- **Ferramentas de Desenvolvimento**: Ferramentas integradas como `go fmt`, `go test`, e `go build` facilitam o desenvolvimento, teste e implantação de aplicações.

## 🚀 Atividade Prática

Durante a oficina, apresentaremos um código simples em Go. Começaremos com o clássico "Hello World" e avançaremos até o desenvolvimento de um servidor web simples, com uma pequena introdução aos pacotes utilizados. Os códigos estarão disponíveis na pasta `source` do repositório.

---

Esta oficina é organizada pelo GDG Ilhéus. Esperamos que você aproveite e aprenda bastante com a linguagem Go!

Para mais informações e recursos adicionais, visite: 
- [Site Oficial do Go](https://go.dev).
- [Site Oficial do GDG Ilhéus](https://gdg.community.dev/gdg-ilheus/).
