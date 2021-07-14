package models

type ApplicationVersion struct {
	ID      string
	Major   int32
	Minor   int32
	Patch   int32
	Build   int32
	Postfix string
}
