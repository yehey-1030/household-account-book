package response

type TagResponse struct {
	TagId         int    `json:"tag_id"`
	Name          string `json:"name"`
	ArchiveTypeId int    `json:"archivetype_id"`
	ParentId      int    `json:"parent_id,omitempty"`
}

type TagListResponse struct {
	Tags []TagResponse `json:"tags"`
}
