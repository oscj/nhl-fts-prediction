// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    game, err := UnmarshalGameInfo(bytes)
//    bytes, err = game.Marshal()

package models

import "encoding/json"

func UnmarshalGameInfo(data []byte) (GameInfo, error) {
	var r GameInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GameInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GameInfo struct {
	Copyright string   `json:"copyright"`
	GamePk    int64    `json:"gamePk"`
	Link      string   `json:"link"`
	MetaData  MetaData `json:"metaData"`
	GameData  GameData `json:"gameData"`
	LiveData  LiveData `json:"liveData"`
}

type GameData struct {
	Game     Game                   `json:"game"`
	Datetime Datetime               `json:"datetime"`
	Status   Status                 `json:"status"`
	Teams    GameDataTeams          `json:"teams"`
	Players  map[string]PlayerValue `json:"players"`
	Venue    CurrentTeamClass       `json:"venue"`
}

type Datetime struct {
	DateTime    string `json:"dateTime"`
	EndDateTime string `json:"endDateTime"`
}

type PlayerValue struct {
	ID                 int64            `json:"id"`
	FullName           string           `json:"fullName"`
	Link               string           `json:"link"`
	FirstName          string           `json:"firstName"`
	LastName           string           `json:"lastName"`
	PrimaryNumber      string           `json:"primaryNumber"`
	BirthDate          string           `json:"birthDate"`
	CurrentAge         int64            `json:"currentAge"`
	BirthCity          string           `json:"birthCity"`
	BirthStateProvince *string          `json:"birthStateProvince,omitempty"`
	BirthCountry       string           `json:"birthCountry"`
	Nationality        string           `json:"nationality"`
	Height             string           `json:"height"`
	Weight             int64            `json:"weight"`
	Active             bool             `json:"active"`
	AlternateCaptain   bool             `json:"alternateCaptain"`
	Captain            bool             `json:"captain"`
	Rookie             bool             `json:"rookie"`
	ShootsCatches      Type             `json:"shootsCatches"`
	RosterStatus       RosterStatus     `json:"rosterStatus"`
	CurrentTeam        CurrentTeamClass `json:"currentTeam"`
	PrimaryPosition    Position         `json:"primaryPosition"`
}

type CurrentTeamClass struct {
	ID           int64        `json:"id"`
	Name         VenueName    `json:"name"`
	Link         Link         `json:"link"`
	TriCode      *TriCodeEnum `json:"triCode,omitempty"`
	Abbreviation *TriCodeEnum `json:"abbreviation,omitempty"`
}

type Position struct {
	Code         Type                        `json:"code"`
	Name         TypeEnum                    `json:"name"`
	Type         TypeEnum                    `json:"type"`
	Abbreviation PrimaryPositionAbbreviation `json:"abbreviation"`
}

type GameDataTeams struct {
	Away PurpleAway `json:"away"`
	Home PurpleAway `json:"home"`
}

type PurpleAway struct {
	ID              int64            `json:"id"`
	Name            VenueName        `json:"name"`
	Link            Link             `json:"link"`
	Venue           PurpleVenue      `json:"venue"`
	Abbreviation    TriCodeEnum      `json:"abbreviation"`
	TriCode         TriCodeEnum      `json:"triCode"`
	TeamName        string           `json:"teamName"`
	LocationName    string           `json:"locationName"`
	FirstYearOfPlay string           `json:"firstYearOfPlay"`
	Division        CurrentTeamClass `json:"division"`
	Conference      CurrentTeamClass `json:"conference"`
	Franchise       Franchise        `json:"franchise"`
	ShortName       string           `json:"shortName"`
	OfficialSiteURL string           `json:"officialSiteUrl"`
	FranchiseID     int64            `json:"franchiseId"`
	Active          bool             `json:"active"`
}

type Franchise struct {
	FranchiseID int64  `json:"franchiseId"`
	TeamName    string `json:"teamName"`
	Link        string `json:"link"`
}

type PurpleVenue struct {
	Name     string   `json:"name"`
	Link     string   `json:"link"`
	City     string   `json:"city"`
	TimeZone TimeZone `json:"timeZone"`
	ID       *int64   `json:"id,omitempty"`
}

type TimeZone struct {
	ID     string `json:"id"`
	Offset int64  `json:"offset"`
	Tz     string `json:"tz"`
}

type LiveData struct {
	Plays     Plays     `json:"plays"`
	Linescore Linescore `json:"linescore"`
	Boxscore  Boxscore  `json:"boxscore"`
	Decisions Decisions `json:"decisions"`
}

type Boxscore struct {
	Teams     BoxscoreTeams `json:"teams"`
	Officials []Official    `json:"officials"`
}

type Official struct {
	Official     FirstStar `json:"official"`
	OfficialType string    `json:"officialType"`
}

type FirstStar struct {
	ID       int64  `json:"id"`
	FullName string `json:"fullName"`
	Link     string `json:"link"`
}

type BoxscoreTeams struct {
	Away FluffyAway `json:"away"`
	Home Home       `json:"home"`
}

type FluffyAway struct {
	Team       CurrentTeamClass `json:"team"`
	TeamStats  TeamStats        `json:"teamStats"`
	Goalies    []int64          `json:"goalies"`
	Skaters    []int64          `json:"skaters"`
	OnIce      []int64          `json:"onIce"`
	OnIcePlus  []OnIcePlus      `json:"onIcePlus"`
	Scratches  []int64          `json:"scratches"`
	PenaltyBox []interface{}    `json:"penaltyBox"`
	Coaches    []Coach          `json:"coaches"`
}

type Coach struct {
	Person   CoachPerson `json:"person"`
	Position Position    `json:"position"`
}

type CoachPerson struct {
	FullName string `json:"fullName"`
	Link     string `json:"link"`
}

type OnIcePlus struct {
	PlayerID      int64 `json:"playerId"`
	ShiftDuration int64 `json:"shiftDuration"`
	Stamina       int64 `json:"stamina"`
}

type SkaterStats struct {
	TimeOnIce            string   `json:"timeOnIce"`
	Assists              int64    `json:"assists"`
	Goals                int64    `json:"goals"`
	Shots                int64    `json:"shots"`
	Hits                 int64    `json:"hits"`
	PowerPlayGoals       int64    `json:"powerPlayGoals"`
	PowerPlayAssists     int64    `json:"powerPlayAssists"`
	PenaltyMinutes       int64    `json:"penaltyMinutes"`
	FaceOffPct           *float64 `json:"faceOffPct,omitempty"`
	FaceOffWINS          int64    `json:"faceOffWins"`
	FaceoffTaken         int64    `json:"faceoffTaken"`
	Takeaways            int64    `json:"takeaways"`
	Giveaways            int64    `json:"giveaways"`
	ShortHandedGoals     int64    `json:"shortHandedGoals"`
	ShortHandedAssists   int64    `json:"shortHandedAssists"`
	Blocked              int64    `json:"blocked"`
	PlusMinus            int64    `json:"plusMinus"`
	EvenTimeOnIce        string   `json:"evenTimeOnIce"`
	PowerPlayTimeOnIce   string   `json:"powerPlayTimeOnIce"`
	ShortHandedTimeOnIce string   `json:"shortHandedTimeOnIce"`
}

type StatsClass struct {
}

type ID8473503Stats struct {
	GoalieStats GoalieStats `json:"goalieStats"`
}

type GoalieStats struct {
	TimeOnIce                  string  `json:"timeOnIce"`
	Assists                    int64   `json:"assists"`
	Goals                      int64   `json:"goals"`
	PIM                        int64   `json:"pim"`
	Shots                      int64   `json:"shots"`
	Saves                      int64   `json:"saves"`
	PowerPlaySaves             int64   `json:"powerPlaySaves"`
	ShortHandedSaves           int64   `json:"shortHandedSaves"`
	EvenSaves                  int64   `json:"evenSaves"`
	ShortHandedShotsAgainst    int64   `json:"shortHandedShotsAgainst"`
	EvenShotsAgainst           int64   `json:"evenShotsAgainst"`
	PowerPlayShotsAgainst      int64   `json:"powerPlayShotsAgainst"`
	Decision                   string  `json:"decision"`
	SavePercentage             float64 `json:"savePercentage"`
	PowerPlaySavePercentage    *int64  `json:"powerPlaySavePercentage,omitempty"`
	EvenStrengthSavePercentage float64 `json:"evenStrengthSavePercentage"`
}

type TeamStats struct {
	TeamSkaterStats TeamSkaterStats `json:"teamSkaterStats"`
}

type TeamSkaterStats struct {
	Goals                  int64  `json:"goals"`
	PIM                    int64  `json:"pim"`
	Shots                  int64  `json:"shots"`
	PowerPlayPercentage    string `json:"powerPlayPercentage"`
	PowerPlayGoals         int64  `json:"powerPlayGoals"`
	PowerPlayOpportunities int64  `json:"powerPlayOpportunities"`
	FaceOffWinPercentage   string `json:"faceOffWinPercentage"`
	Blocked                int64  `json:"blocked"`
	Takeaways              int64  `json:"takeaways"`
	Giveaways              int64  `json:"giveaways"`
	Hits                   int64  `json:"hits"`
}

type Home struct {
	Team       CurrentTeamClass `json:"team"`
	TeamStats  TeamStats        `json:"teamStats"`
	Goalies    []int64          `json:"goalies"`
	Skaters    []int64          `json:"skaters"`
	OnIce      []int64          `json:"onIce"`
	OnIcePlus  []OnIcePlus      `json:"onIcePlus"`
	Scratches  []int64          `json:"scratches"`
	PenaltyBox []interface{}    `json:"penaltyBox"`
	Coaches    []Coach          `json:"coaches"`
}

type Decisions struct {
	Winner     FirstStar `json:"winner"`
	Loser      FirstStar `json:"loser"`
	FirstStar  FirstStar `json:"firstStar"`
	SecondStar FirstStar `json:"secondStar"`
	ThirdStar  FirstStar `json:"thirdStar"`
}

type Linescore struct {
	CurrentPeriod              int64                `json:"currentPeriod"`
	CurrentPeriodOrdinal       CurrentPeriodOrdinal `json:"currentPeriodOrdinal"`
	CurrentPeriodTimeRemaining string               `json:"currentPeriodTimeRemaining"`
	Periods                    []Period             `json:"periods"`
	ShootoutInfo               ShootoutInfo         `json:"shootoutInfo"`
	Teams                      LinescoreTeams       `json:"teams"`
	PowerPlayStrength          string               `json:"powerPlayStrength"`
	HasShootout                bool                 `json:"hasShootout"`
	IntermissionInfo           IntermissionInfo     `json:"intermissionInfo"`
	PowerPlayInfo              PowerPlayInfo        `json:"powerPlayInfo"`
}

type IntermissionInfo struct {
	IntermissionTimeRemaining int64 `json:"intermissionTimeRemaining"`
	IntermissionTimeElapsed   int64 `json:"intermissionTimeElapsed"`
	InIntermission            bool  `json:"inIntermission"`
}

type Period struct {
	PeriodType PeriodType           `json:"periodType"`
	StartTime  string               `json:"startTime"`
	EndTime    string               `json:"endTime"`
	Num        int64                `json:"num"`
	OrdinalNum CurrentPeriodOrdinal `json:"ordinalNum"`
	Home       PeriodAway           `json:"home"`
	Away       PeriodAway           `json:"away"`
}

type PeriodAway struct {
	Goals       int64  `json:"goals"`
	ShotsOnGoal int64  `json:"shotsOnGoal"`
	RinkSide    string `json:"rinkSide"`
}

type PowerPlayInfo struct {
	SituationTimeRemaining int64 `json:"situationTimeRemaining"`
	SituationTimeElapsed   int64 `json:"situationTimeElapsed"`
	InSituation            bool  `json:"inSituation"`
}

type ShootoutInfo struct {
	Away ShootoutInfoAway `json:"away"`
	Home ShootoutInfoAway `json:"home"`
}

type ShootoutInfoAway struct {
	Scores   int64 `json:"scores"`
	Attempts int64 `json:"attempts"`
}

type LinescoreTeams struct {
	Home TentacledAway `json:"home"`
	Away TentacledAway `json:"away"`
}

type TentacledAway struct {
	Team         CurrentTeamClass `json:"team"`
	Goals        int64            `json:"goals"`
	ShotsOnGoal  int64            `json:"shotsOnGoal"`
	GoaliePulled bool             `json:"goaliePulled"`
	NumSkaters   int64            `json:"numSkaters"`
	PowerPlay    bool             `json:"powerPlay"`
}

type Plays struct {
	AllPlays      []AllPlay       `json:"allPlays"`
	ScoringPlays  []int64         `json:"scoringPlays"`
	PenaltyPlays  []int64         `json:"penaltyPlays"`
	PlaysByPeriod []PlaysByPeriod `json:"playsByPeriod"`
	CurrentPlay   CurrentPlay     `json:"currentPlay"`
}

type AllPlay struct {
	Result      AllPlayResult      `json:"result"`
	About       About              `json:"about"`
	Coordinates AllPlayCoordinates `json:"coordinates"`
	Players     []PlayerElement    `json:"players,omitempty"`
	Team        *CurrentTeamClass  `json:"team,omitempty"`
}

type About struct {
	EventIdx            int64                `json:"eventIdx"`
	EventID             int64                `json:"eventId"`
	Period              int64                `json:"period"`
	PeriodType          PeriodType           `json:"periodType"`
	OrdinalNum          CurrentPeriodOrdinal `json:"ordinalNum"`
	PeriodTime          string               `json:"periodTime"`
	PeriodTimeRemaining string               `json:"periodTimeRemaining"`
	DateTime            string               `json:"dateTime"`
	Goals               Goals                `json:"goals"`
}

type Goals struct {
	Away int64 `json:"away"`
	Home int64 `json:"home"`
}

type AllPlayCoordinates struct {
	X *int64 `json:"x,omitempty"`
	Y *int64 `json:"y,omitempty"`
}

type PlayerElement struct {
	Player      FirstStar  `json:"player"`
	PlayerType  PlayerType `json:"playerType"`
	SeasonTotal *int64     `json:"seasonTotal,omitempty"`
}

type AllPlayResult struct {
	Event           string    `json:"event"`
	EventCode       string    `json:"eventCode"`
	EventTypeID     string    `json:"eventTypeId"`
	Description     string    `json:"description"`
	SecondaryType   *string   `json:"secondaryType,omitempty"`
	PenaltySeverity *string   `json:"penaltySeverity,omitempty"`
	PenaltyMinutes  *int64    `json:"penaltyMinutes,omitempty"`
	Strength        *Strength `json:"strength,omitempty"`
	GameWinningGoal *bool     `json:"gameWinningGoal,omitempty"`
	EmptyNet        *bool     `json:"emptyNet,omitempty"`
}

type Strength struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CurrentPlay struct {
	Result      CurrentPlayResult `json:"result"`
	About       About             `json:"about"`
	Coordinates StatsClass        `json:"coordinates"`
}

type CurrentPlayResult struct {
	Event       string `json:"event"`
	EventCode   string `json:"eventCode"`
	EventTypeID string `json:"eventTypeId"`
	Description string `json:"description"`
}

type PlaysByPeriod struct {
	StartIndex int64   `json:"startIndex"`
	Plays      []int64 `json:"plays"`
	EndIndex   int64   `json:"endIndex"`
}

const (
	Hc     Type = "HC"
	L      Type = "L"
	TypeC  Type = "C"
	TypeD  Type = "D"
	TypeG  Type = "G"
	TypeNA Type = "N/A"
)

type TriCodeEnum string

const (
	Fla TriCodeEnum = "FLA"
	Min TriCodeEnum = "MIN"
)

type Link string

const (
	APIV1Conferences5 Link = "/api/v1/conferences/5"
	APIV1Conferences6 Link = "/api/v1/conferences/6"
	APIV1Divisions16  Link = "/api/v1/divisions/16"
	APIV1Divisions17  Link = "/api/v1/divisions/17"
	APIV1Teams13      Link = "/api/v1/teams/13"
	APIV1Teams30      Link = "/api/v1/teams/30"
	APIV1Venues5098   Link = "/api/v1/venues/5098"
)

type VenueName string

const (
	Atlantic         VenueName = "Atlantic"
	Central          VenueName = "Central"
	Eastern          VenueName = "Eastern"
	FloridaPanthers  VenueName = "Florida Panthers"
	MinnesotaWild    VenueName = "Minnesota Wild"
	Western          VenueName = "Western"
	XcelEnergyCenter VenueName = "Xcel Energy Center"
)

type PrimaryPositionAbbreviation string

const (
	AbbreviationC         PrimaryPositionAbbreviation = "C"
	AbbreviationD         PrimaryPositionAbbreviation = "D"
	AbbreviationG         PrimaryPositionAbbreviation = "G"
	AbbreviationHeadCoach PrimaryPositionAbbreviation = "Head Coach"
	AbbreviationNA        PrimaryPositionAbbreviation = "N/A"
	Lw                    PrimaryPositionAbbreviation = "LW"
	Rw                    PrimaryPositionAbbreviation = "RW"
)

type TypeEnum string

const (
	Center        TypeEnum = "Center"
	Defenseman    TypeEnum = "Defenseman"
	Forward       TypeEnum = "Forward"
	LeftWing      TypeEnum = "Left Wing"
	NameGoalie    TypeEnum = "Goalie"
	NameHeadCoach TypeEnum = "Head Coach"
	RightWing     TypeEnum = "Right Wing"
	Unknown       TypeEnum = "Unknown"
)

type RosterStatus string

const (
	Y RosterStatus = "Y"
)

type CurrentPeriodOrdinal string

const (
	The1St CurrentPeriodOrdinal = "1st"
	The2Nd CurrentPeriodOrdinal = "2nd"
	The3RD CurrentPeriodOrdinal = "3rd"
)

type PeriodType string

const (
	Regular PeriodType = "REGULAR"
)

type PlayerType string

const (
	Assist           PlayerType = "Assist"
	Blocker          PlayerType = "Blocker"
	DrewBy           PlayerType = "DrewBy"
	Hittee           PlayerType = "Hittee"
	Hitter           PlayerType = "Hitter"
	Loser            PlayerType = "Loser"
	PenaltyOn        PlayerType = "PenaltyOn"
	PlayerID         PlayerType = "PlayerID"
	PlayerTypeGoalie PlayerType = "Goalie"
	Scorer           PlayerType = "Scorer"
	Shooter          PlayerType = "Shooter"
	Winner           PlayerType = "Winner"
)
