package main


type Taku struct {
  Yama []Pai
  Players []Player
  Kaze int
}

func CreateTaku() *Taku {
  t := new(Taku)
  t.Players = make([]Player,4)

  for i := 0; i < 4; i++ {
    t.Players[i] = Player{}
    t.Players[i].Point = 2500
    t.Players[i].Kaze = i
  }
  return t
}

func (t *Taku) Start(){
  t.haiPai()
}

func (t *Taku) haiPai(){
  t.Yama = CreateYama()
  for i := 0; i < 4; i++ {
    t.Players[i].Hand = GetPais(&(t.Yama), 13)
  }
}
