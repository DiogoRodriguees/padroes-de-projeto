package observers

import "log"

type Observer struct {
	ID           int
	InterestedIn []string
}

func New(id int, interestedList []string) *Observer {
	return &Observer{
		ID:           id,
		InterestedIn: interestedList,
	}
}

func (o *Observer) GetInterest() []string {
	return o.InterestedIn
}
func (o *Observer) Update() {
	log.Println("Executing observer function update for id: ", o.ID)
}

func (o *Observer) GetID() int {
	return o.ID
}
