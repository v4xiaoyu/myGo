package entities

type BaseEntity struct {
	Id int64
}

func (this *BaseEntity) GetId() int64 {
	return this.Id
}
