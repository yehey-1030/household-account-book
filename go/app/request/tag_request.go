package request

type CreateTagRequest struct {
	ArchiveTypeId int    `json:"archivetype_id"`
	ParentId      int    `json:"parent_id"`
	Name          string `json:"name"`
}
