package structs

type Label struct {
	IdLabel   int
	LabelName string
	Tasks     []Task `"gorm:" foreignkey:ProjectID `
}
