package entities

type UserEntity struct {
	BaseEntity

	Name   string
	Gender int8
	Degree float32
}
