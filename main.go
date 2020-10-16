package main

func main() {
  yama := CreateYama()
  pais := make([]Pai,0)

  for i := 0 ; i < 3; i++ {
    for j := 0; j < 4; j++ {
      pais = append(pais,GetPai(&yama))
    }
    PrintPais(pais)
  }
    pais = append(pais,GetPai(&yama))
  PrintPais(pais)
  SortPais(pais)
  PrintPais(pais)

}
