package services

type ServiceForNoteApp struct{}

type StudentNote struct {
	CoureseName string `json:"course_name"`
	CourseID    string `json:"course_id"`
	Note        string `json:"note"`
}

func (n *ServiceForNoteApp) GetNotes() []StudentNote {
	notes := []StudentNote{
		{CoureseName: "Math", CourseID: "M101", Note: "Algebra notes"},
		{CoureseName: "Science", CourseID: "S101", Note: "Physics notes"},
	}
	return notes
}

func (n *ServiceForNoteApp) CreateNewNote() []StudentNote {
	notes := []StudentNote{
		{CoureseName: "German", CourseID: "DE101", Note: "Deutsch notes"},
	}
	return notes
}
