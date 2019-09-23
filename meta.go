package tboxcore

type Meta struct {
	Size                 int64 `json:"Size"`
	IsDirectory          bool `json:"IsDirectory"`
	IsRegularFile        bool `json:"IsRegularFile"`
	IsSimbolicLink       bool `json:"IsSimbolicLink"`
	IsHidden             bool `json:"IsHidden"`
	LastModifyTime       int64 `json:"LastModifyTime"`
	Owner                string `json:"Owner"`
	PosixFilePermissions string `json:"PosixFilePermissions"`
	MimeType             string `json:"MimeType"`
	Notes                string `json:"Notes"`
}
