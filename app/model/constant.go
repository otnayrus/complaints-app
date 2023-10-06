package model

type Field string

const (
	FieldSingleFileImage   Field = "single_file_image"
	FieldMultipleFileImage Field = "multiple_file_image"
	FieldDropdownSelection Field = "dropdown_selection"
	FieldText              Field = "text"
	FieldTextArea          Field = "text_area"
	FieldNumber            Field = "number"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)
