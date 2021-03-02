package structs

type Project struct {
	IdProject int    `"gorm:" primary_key`
	Tasks     []Task `"gorm:" foreignkey:ProjectID `
}
