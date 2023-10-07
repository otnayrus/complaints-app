package model

type FieldType string

const (
	FieldTypeSingleFileImage   FieldType = "single_file_image"
	FieldTypeMultipleFileImage FieldType = "multiple_file_image"
	FieldTypeDropdownSelection FieldType = "dropdown_selection"
	FieldTypeText              FieldType = "text"
	FieldTypeTextArea          FieldType = "text_area"
	FieldTypeNumber            FieldType = "number"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

const (
	ComplaintStatusPending int = iota + 1
	ComplaintStatusResolved
)
