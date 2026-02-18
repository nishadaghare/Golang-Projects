package models

type Distributor struct {
	Name     string
	Parent   *Distributor
	Includes map[string]bool
	Excludes map[string]bool
}
