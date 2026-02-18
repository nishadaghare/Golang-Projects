package models

type Region struct {
	Code     string    `json:"code"`
	Parent   *Region   `json:"-"`
	Children []*Region `json:"children,omitempty"`
}
