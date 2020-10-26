package main

type Player struct {
  Hand []Pai
  Point int
  Kaze int
}

func(p *Player) ShiPai(){
  SortPais(p.Hand)
}

