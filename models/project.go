package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Project model.
type Project struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);unique"`
	Tasks     []Task
	Entries   []Entry
	TotalTime time.Duration
}

// AllProjects queries the database for, and returns, all projects
// after scanning them into a slice of structs.
func AllProjects() []Project {
	p := []Project{}
	DB.Find(&p)
	return p
}

// GetProject queries the database for, and returns, one project
// after scanning it into the struct.
func GetProject(n string) Project {
	p := Project{}
	DB.Where("name = ?", n).First(&p)
	return p
}

// AddProject queries the database for one project by name (project names are unique).
// If the record exists then it is returned;
// else, it will create the record and return that one.
func AddProject(n string) Project {
	p := Project{Name: n}
	// db.FirstOrCreate(&p, Project{Name: n})
	DB.Create(&p)
	return p
}