package postgresStorage

type Postgres struct{}

func NewPostgresStorage() *Postgres {
	return &Postgres{}
}
