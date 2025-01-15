package models

type Contact struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Linkedin string `json:"Linkedin"` // Note the capital "L" to match JSON key
}

type Experience struct {
	Company          string   `json:"company"`
	Position         string   `json:"position"`
	StartDate        string   `json:"start_date"`
	EndDate          string   `json:"end_date"`
	Responsibilities []string `json:"responsibilities"`
}

type Education struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type Resume struct {
	Name       string       `json:"name"`
	Contact    Contact      `json:"contact"`
	Summary    string       `json:"summary"`
	Experience []Experience `json:"experience"`
	Education  []Education  `json:"education"`
	Skills     []string     `json:"skills"`
}
