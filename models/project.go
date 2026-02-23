package models

type Project struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binfing:"optional"`
	Leader      User   `json:"leader" binding:"required"`
	Workers     []User `json:"workers" binding:"required"`
}

func NewProject(name string, description string, leader User) Project {
	return Project{
		Name:        name,
		Description: description,
		Leader:      leader,
		Workers:     make([]User, 0),
	}
}

func (p *Project) AddNewWorker(w User) bool {
	if p.workerExists(w) {
		return false
	}

	p.Workers = append(p.Workers, w)
	return true
}

func (p *Project) workerExists(w User) bool {
	for _, worker := range p.Workers {
		if worker.isEqual(w) {
			return true
		}
	}
	return false
}
