package infra

import "fmt"

type Observer struct {
	ID int
}

func (o *Observer) Update() {
	fmt.Println("Executing observer function update for id: ", o.ID)
}

func (o *Observer) GetID() int {
	return o.ID
}
