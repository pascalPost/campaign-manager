package types

type Project struct {
	id   uint
	name string
}

func NewProject(id uint, name string) *Project {
	return &Project{
		id:   id,
		name: name,
	}
}
