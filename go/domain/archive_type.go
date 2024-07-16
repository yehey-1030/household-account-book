package domain

type ArchiveType interface {
	Id() int
	Name() string
}

type archiveType struct {
	id   int
	name string
}

func (a *archiveType) Id() int {
	return a.id
}

func (a *archiveType) Name() string {
	return a.name
}

func NewArchiveType(id int, name string) ArchiveType {
	return &archiveType{id: id, name: name}
}
