package response

type ArchiveTypeResponse struct {
	Id   int    `json:"archivetype_id"`
	Name string `json:"name"`
}

type ArchiveTypeListResponse struct {
	ArchiveTypes []ArchiveTypeResponse `json:"archive_types"`
}
