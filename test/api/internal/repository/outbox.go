package repository

type Outbox struct {
	ID        string
	TableID   string
	TableName string
	Status    string
}
