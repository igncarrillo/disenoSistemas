package singleton

var instance *database

type database struct {
	name string
}

func GetInstance(n string) *database {
	if instance == nil {
		instance = &database{n}
	}
	return instance
}

func (db *database) Getname() string {
	return db.name
}
