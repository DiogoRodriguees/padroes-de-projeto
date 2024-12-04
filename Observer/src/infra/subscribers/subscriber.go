package subscribers

import "log"

type Subscriber struct {
	ID           int
	InterestedIn []string
}

func New(id int, interestedList []string) *Subscriber {
	return &Subscriber{
		ID:           id,
		InterestedIn: interestedList,
	}
}

func (o *Subscriber) GetInterest() []string {
	return o.InterestedIn
}
func (o *Subscriber) Update() {
	log.Println("Executing observer function update for id: ", o.ID)
}

func (o *Subscriber) GetID() int {
	return o.ID
}
