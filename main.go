package main

import (
	"bufio"
	"fmt"
	"os"
)
type Question struct{
	Text string
	Options []string
	Answer int
}
type GameState struct{
	Name string
	Points string
	Questions []Question
}

func (g *GameState) init(){
	fmt.Println("Seja bem vindo ao jogo")
	fmt.Println("Escreva o seu nome")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil{
		panic("Erro ao ler a string")
	}
	g.Name = name
}
func main(){
	game1 :=&GameState{}
	game1.init()
}