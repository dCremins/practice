package main


type Move struct {
    angle   string
    amount	int
}

func BMain(directions []Move)int {
	position:=0
	depth:=0
	aim:=0

	for _,move:= range directions {
		dir:=move.angle
		amount:=move.amount
		if dir=="forward" {
			position+=amount
			depth+=(aim*amount)
		} else if dir=="down"{
			aim+=amount
		} else {
			aim-=amount
		}
	}

	return position*depth
}