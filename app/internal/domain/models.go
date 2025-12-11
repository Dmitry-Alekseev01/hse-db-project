package domain

import "time"

type ErrorResponse struct {
	Message string `json:"message"`
}

type CoachSpecificationName string

const (
	CoachHead        CoachSpecificationName = "Head"
	CoachGoalkeeping CoachSpecificationName = "Goalkeeping"
	CoachDefense     CoachSpecificationName = "Defense"
	CoachForward     CoachSpecificationName = "Forward"
	CoachStandard    CoachSpecificationName = "Standard"
	CoachFitness     CoachSpecificationName = "Fitness"
)

type CoachSpecification struct {
	ID   int64                  `db:"specification_id" json:"specification_id"`
	Name CoachSpecificationName `db:"specification_name" json:"specification_name"`
}

type StaffSpecificationName string

const (
	StaffCleaner   StaffSpecificationName = "Cleaner"
	StaffSecurity  StaffSpecificationName = "Security"
	StaffLawnMower StaffSpecificationName = "LawnMower"
	StaffDoctor    StaffSpecificationName = "Doctor"
	StaffDirector  StaffSpecificationName = "Director"
	StaffPresident StaffSpecificationName = "President"
)

type StaffSpecification struct {
	ID   int64                  `db:"specification_id" json:"specification_id"`
	Type StaffSpecificationName `db:"specification_type" json:"specification_type"`
}

type Club struct {
	ID           int64     `db:"club_id" json:"team_id"`
	Name         string    `db:"club_name" json:"club_name"`
	CreationDate time.Time `db:"creation_date" json:"creation_date"`
	Website      *string   `db:"website,omitempty" json:"website,omitempty"`
}

type Team struct {
	ID     int64   `db:"team_id" json:"team_id"`
	Name   string  `db:"team_name" json:"team_name"`
	Budget float64 `db:"budget" json:"budget"`
	ClubID int64   `db:"club_id" json:"club_id"`
}

type Staff struct {
	ID              int64   `db:"staff_id" json:"staff_id"`
	Name            string  `db:"staff_name" json:"staff_name"`
	Surname         string  `db:"staff_surname" json:"staff_surname"`
	Salary          float64 `db:"salary" json:"salary"`
	SpecificationID int64   `db:"specification_id" json:"specification_id"`
	ClubID          int64   `db:"club_id" json:"club_id"`
}

type Stadium struct {
	ID        int64     `db:"stadium_id" json:"stadium_id"`
	Capacity  int64     `db:"capacity" json:"capacity"`
	Location  string    `db:"location" json:"location"`
	BuildDate time.Time `db:"build_date" json:"build_date"`
}

type StadiumAndClub struct {
	ClubID    int64 `db:"club_id" json:"club_id"`
	StadiumID int64 `db:"stadium_id" json:"stadium_id"`
}

type Game struct {
	ID        int64     `db:"match_id" json:"match_id"`
	StadiumID int64     `db:"stadium_id" json:"stadium_id"`
	Team1ID   int64     `db:"team_1_id" json:"team_1_id"`
	Team2ID   int64     `db:"team_2_id" json:"team_2_id"`
	MatchDate time.Time `db:"match_date" json:"match_date"`
}

type PlayerStatusType string

const (
	StatusLoan       PlayerStatusType = "load"
	StatusOnContract PlayerStatusType = "on contract"
)

type PlayerStatus struct {
	ID             int64            `db:"player_id" json:"player_id"`
	Type           PlayerStatusType `db:"status_id" json:"status_id"`
	StartDate      time.Time        `db:"start_date" json:"start_date"`
	ExpirationDate time.Time        `db:"expiration_date" json:"expiration_date"`
}

type Player struct {
	ID        int64     `db:"player_id" json:"player_id"`
	Name      string    `db:"player_name" json:"player_name"`
	Surname   string    `db:"player_surname" json:"player_surname"`
	Number    int64     `db:"player_number" json:"player_number"`
	Salary    float64   `db:"salary" json:"salary" `
	Phone     string    `db:"phone" json:"phone"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
	TeamID    int64     `db:"team_id" json:"team_id"`
	StatusID  int64     `db:"status_id" json:"status_id"`
}

type Coach struct {
	ID      int64   `db:"coach_id" json:"coach_id"`
	Name    string  `db:"coach_name" json:"coach_name"`
	Surname string  `db:"coach_surname" json:"coach_surname"`
	Salary  float64 `db:"salary" json:"salary"`
	Phone   string  `db:"phone" json:"phone"`
	TeamID  int64   `db:"team_id" json:"team_id"`
}

type CoachAndSpecification struct {
	SpecificationID int64 `db:"specification_id" json:"specification_id"`
	CoachID         int64 `db:"coach_id" json:"coach_id"`
}
