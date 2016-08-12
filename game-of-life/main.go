package main

import "./gol"

func main() {
	for i := 0; i < 100; i++ {
		gol.Gol(40, 100)
	}
}
