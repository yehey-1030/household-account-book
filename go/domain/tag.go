package domain

type Tag interface {
	Id() int
	Name() string
	ParentId() int
	ArchiveTypeId() int
}

type tag struct {
	tagId         int
	name          string
	parentId      int
	archiveTypeId int
}

func (t *tag) Id() int {
	return t.tagId
}

func (t *tag) Name() string {
	return t.name
}

func (t *tag) ParentId() int {
	return t.parentId
}

func (t *tag) ArchiveTypeId() int {
	return t.archiveTypeId
}

func NewTag(id int, name string, parentId, archiveTypeId int) Tag {
	return &tag{tagId: id, name: name, parentId: parentId, archiveTypeId: archiveTypeId}
}
