/*
Juego 1
Juego termina cuando jugador tenga 21.
Cada turno:
- Jugador debe elegir un numero entre el 1 y el 10 x ronda.
- Lider debe elegir al azar entre 6 y 10
- Verificicar que jugadores eligieron el mismo nro que el lider.
    - En ese caso mueren
El juego no puede durar mas de 4 rondas
Todos los jugadores que no logren llegar a 21 antes de las 4 rondas, seran eliminados.

*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func juego1_jugador(jugador bool, ronda int) {
	fmt.Println("Hola!")
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	suma := 0
	numero := 0

	for ronda < 4 {
		
		// Seleccionar numero para jugar
		if jugador == true {
			fmt.Println("Elija un numero del 1 al 10")
			fmt.Print("-> ")
			char, _, err := reader.ReadRune()
	
			if err != nil {
				fmt.Println(err)
			}
	
			numero = int(char - '0')
	
		} else {
			numero = rand.Intn(10)
	
		}

		// Chequear si el numero es igual o mayor al del lider


		// sumar al numero anterior

		suma += numero

		fmt.Println(numero)
		ronda = ronda + 1
	}

	
    
}

func main() {

	// Enviar peticion para unirse al juego 
	// Recibir 

	juego1_jugador(true, 0)
}
