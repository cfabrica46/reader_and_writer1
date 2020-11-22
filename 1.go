package main

import (
	"fmt"
	"log"
)

type miescritor struct {
	contenido []byte
	tamaño    int
}

func nuevomiescritor() *miescritor {
	return &miescritor{}
}

func main() {

	fmt.Println("Hola")

}

func (me *miescritor) Write(b []byte) (n int, err error) {

	me.contenido = append(me.contenido, b...)

	me.tamaño = len(me.contenido)

	n = len(me.contenido)

	if n != me.tamaño {
		log.Fatal(err)
	}

	return
}
