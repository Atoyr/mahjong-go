package main

import (
  "math/rand"
  "strconv"
  "time"
  "sort"
  "fmt"
)

type Pai int

func CreateYama() []Pai {
  pais := CreatePais()
  rand.Seed(time.Now().UnixNano())
  rand.Shuffle(len(pais), func(i, j int) { pais[i], pais[j] = pais[j], pais[i]})
  return pais
}

func GetPai(pais *[]Pai) Pai {
  temp := *pais
  if len(temp) == 0 {
    return Pai(0)
  }
  p := (temp)[0]
  *pais = (temp[1:])
  return p
}

func GetPais(pais *[]Pai, count int) []Pai {
  temp := *pais
  if 0 < count && count <= len(temp) {
    ret := temp[:count - 1]
    *pais = (temp[count:])
    return ret
  }else {
    return make([]Pai,0)
  }
}

func CreatePais() []Pai {
  pais := make([]Pai,136)
  for i := 0 ; i < 136; i++ {
    pais[i] = Pai(i)
  }
  return pais
}

func SortPais(pais []Pai) {
  sort.Slice(pais, func(i,j int) bool { return int(pais[i]) < int(pais[j])})
}

func PrintPais(pais []Pai) {
  for i := 0; i < len(pais); i++ {
      fmt.Printf("- ")
  }
  fmt.Println()
  for i := 0; i < 3; i++ {
    for j := 0; j < len(pais); j++ {
      fmt.Printf("%s ",pais[j].String()[i])
    }
    fmt.Println()
  }
  for i := 0; i < len(pais); i++ {
      fmt.Printf("- ")
  }
  fmt.Println()
}

func(p Pai) String() []string {
  i := int(p)
  ret := make([]string,3)
  t := i / 36
  n := (i % 36) / 4 + 1
  switch t {
  case 0:
    ret[0] = strconv.Itoa(n)
    ret[1] = " "
    ret[2] = "M"
  case 1:
    ret[0] = strconv.Itoa(n)
    ret[1] = " "
    ret[2] = "S"
  case 2:
    ret[0] = strconv.Itoa(n)
    ret[1] = " "
    ret[2] = "P"
  case 3:
    ret[0] = " "
    ret[2] = " "
    switch n {
    case 0:
      ret[1] = "E"
    case 1:
      ret[1] = "S"
    case 2:
      ret[1] = "W"
    case 3:
      ret[1] = "N"
    case 4:
      ret[1] = " "
    case 5:
      ret[1] = "H"
    case 6:
      ret[1] = "C"
    }
  }
  return ret
}
