package methordStruct

func (p Point) MoveWrong(dx, dy int) { p.X += dx; p.Y += dy }
func (p *Point) Move(dx, dy int)     { p.X += dx; p.Y += dy }
