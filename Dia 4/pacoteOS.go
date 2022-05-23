package main

import (
	"fmt"
	"os"
)

func main() {
	// Variaveis de ambiente
	err := os.Setenv("Name", "Rodney")

	if err != nil {
		fmt.Println("Ocorreu um erro")
	}
	fmt.Println(os.Getenv("Name"))

	if value, ok := os.LookupEnv("Name2"); ok == true {
		fmt.Println(value)
	} else {
		fmt.Println("Nao existe")
	}
	// Arquivos
	files, err := os.ReadDir("./")
	if err != nil {
		fmt.Println("Deu erro")
	} else {
		for _, v := range files {
			fmt.Println(v.Name())
		}
	}

	data, err2 := os.ReadFile("pacoteOS.go")

	if err2 != nil {
		fmt.Println("Erro:", err2)
	}
	fmt.Println(string(data))

	minhaString := []byte("Meu novo arquivo")
	err3 := os.WriteFile("teste.txt", minhaString, 0777)

	if err3 != nil {
		fmt.Println("Erro na escrita")
	}

	// O_RDONLY	It opens the file read-only
	// O_WRONLY	It opens the file write-only
	// O_RDWR	It opens the file read-write
	// O_APPEND	It appends data to the file when writing
	// O_CREATE	It creates a new file if none exists
	// O_EXCL	It is used with O_CREATE, and the file must not exist
	// O_SYNC	It opens for synchronous I/O
	// O_TRUNC	if possible, truncate the file when opened

	minhaString2 := []byte("Meu novo arquivo 2")
	meuArquivo, err4 := os.OpenFile("teste 2.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer meuArquivo.Close()
	if err4 != nil {
		fmt.Println("Erro na escrita")
	}

	meuArquivo.Write(minhaString2)
	meuArquivo.WriteString("\nAdd segunda linha")
}
