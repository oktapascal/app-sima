package domain

type Payload struct {
	Search *string
	Limit  int
	Order  string
	Start  int
}
