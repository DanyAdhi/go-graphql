package pets

type Pets struct {
	ID        int32  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Type      string `json:"type" db:"type"`
	OwnerId   int    `json:"ownerId" db:"ownerId"`
	OwnerName string `json:"ownerName" db:"ownerName"`
}

type NewPet struct {
	Name    string `json:"name"`
	Type    string `json:"type,omitempty"`
	OwnerId int    `json:"ownerId"`
}
