package main

func main() {
  taku := CreateTaku()
  taku.Start()
  PrintPais(taku.Players[0].Hand)
}
