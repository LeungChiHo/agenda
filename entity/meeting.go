package entity

type Meeting struct {
	Sponsor       string
	StartDate     string
	EndDate       string
	Title         string
	Participators []User
}

func (m *Meeting) initMeeting(t_sponsor, t_startDate, t_endDate, t_title string, t_participator []User) {
	m.Sponsor = t_sponsor
	m.StartDate = t_startDate
	m.EndDate = t_endDate
	m.Title = t_title
	m.Participators = t_participator
	
}

func (m Meeting) getSponsor() string {
	return m.Sponsor
}

func (m *Meeting) setSponsor(s string) {
	m.Sponsor = s
}

func (m Meeting) getParticipators() []User {
	return m.Participators
}

func (m *Meeting) setParticipators(p []User) {
	m.Participators = p
}
func (m Meeting) getStartDate() string {
	return m.StartDate
}

func (m *Meeting) setStartDate(s string) {
	m.StartDate = s
}
func (m Meeting) getEndDate() string {
	return m.EndDate
}

func (m *Meeting) setEndDate(e string) {
	m.EndDate = e
}
func (m Meeting) getTitle() string {
	return m.Title
}

func (m *Meeting) setTitle(t string) {
	m.Title = t
}

func (m Meeting) isParticipator(name string) bool {
	for _, t := range m.Participators {
		if t.Name == name {
			return true
		}
	}
	return false	
}
