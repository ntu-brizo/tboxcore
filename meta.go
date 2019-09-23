package tboxcore

type Meta struct {
	size                 int64 `json:"size"`
	isDirectory          bool `json:"isDirectory"`
	isRegularFile        bool `json:"isRegularFile"`
	isSimbolicLink       bool `json:"isSimbolicLink"`
	isHidden             bool `json:"isHidden"`
	lastModifyTime       int64 `json:"lastModifyTime"`
	owner                string `json:"owner"`
	posixFilePermissions string `json:"posixFilePermissions"`
	mimeType             string `json:"mimeType"`
	notes                string `json:"notes"`
}
