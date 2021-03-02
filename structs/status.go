package structs

type Status struct {
	IdStatus   int
	StatusName string
	Tasks      []Task `"gorm:" foreignkey:Status_id `
}
