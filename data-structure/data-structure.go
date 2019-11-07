package main

import (
	"fmt"
)

type DataStructure interface {
	inserir(num int)
	remover()
	imprimir()
}

type pilha struct {
	content []int
}

func (p *pilha) inserir(num int) {
	fmt.Println("Inserindo ",num)
	p.content = append(p.content, num)
}

func (p *pilha) remover() {
	fmt.Println("Removendo")
	if (len(p.content) > 0) {
		p.content = p.content[:len(p.content) - 1]
	}
}

func (p *pilha) imprimir() {
	fmt.Println(p.content)
}

type fila struct {
	content []int
}

func (f *fila) inserir(num int) {
	fmt.Println("Inserindo ",num)
	f.content = append(f.content, num)
}

func (f *fila) remover() {
	fmt.Println("Removendo")
	f.content = append(f.content[:0], f.content[0+1:]...)
}

func (f *fila) imprimir() {
	fmt.Println(f.content)
}

func test(ds DataStructure) {
	ds.inserir(1)
	ds.imprimir()
	ds.inserir(2)
	ds.imprimir()
	ds.remover()
	ds.imprimir()
	ds.inserir(3)
	ds.imprimir()
	ds.remover()
	ds.imprimir()
	ds.inserir(4)
	ds.imprimir()
	ds.remover()
	ds.imprimir()
	ds.remover()
	ds.imprimir()
}

func main() {
	p := pilha{}
	test(&p)
	fmt.Println("------------------------")
	f := fila{}
	test(&f)
}
