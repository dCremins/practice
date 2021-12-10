package main

type Move struct {
    angle   string
    amount	int
}

func AMain(directions []Move)int {
	position:=0
	depth:=0

	for _,move:= range directions {
		dir:=move.angle
		amount:=move.amount
		if dir=="forward" {
			position+=amount
		} else if dir=="down"{
			depth+=amount
		} else {
			depth-=amount
		}
	}

	return position*depth
}