type Service struct {
	db *db
}

type db struct {
	host string
}

func (d *db) GetById()    {}
func (d *db) DeleteById() {}

type Service struct {
	db repositoty
}

type repositoty interface {
	GetById()
	DeleteById()
}
