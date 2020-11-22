package main

import (
	"fmt"
	"io"
	"log"
)

type milector struct {
	contenido []byte
	indice    int
}

func nuevomilector(contenido []byte) *milector {

	return &milector{contenido: contenido, indice: 0}
}

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

func (ml *milector) Read(b []byte) (n int, err error) {

	if ml.indice >= len(ml.contenido) {

		n = 0
		err = io.EOF
		return

	}

	for i := 0; i < len(b); i++ {

		if ml.indice == len(ml.contenido) {
			break
		}

		b[i] = ml.contenido[ml.indice]
		ml.indice++
		n++

	}

	return
}
