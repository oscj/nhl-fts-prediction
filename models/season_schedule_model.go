// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

package models

import "encoding/json"

func UnmarshalSeasonSchedule(data []byte) (SeasonSchedule, error) {
	var r SeasonSchedule
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SeasonSchedule) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SeasonSchedule struct {
	Copyright    string   `json:"copyright"`
	TotalItems   int64    `json:"totalItems"`
	TotalEvents  int64    `json:"totalEvents"`
	TotalGames   int64    `json:"totalGames"`
	TotalMatches int64    `json:"totalMatches"`
	MetaData     MetaData `json:"metaData"`
	Wait         int64    `json:"wait"`
	Dates        []Date   `json:"dates"`
}

type Date struct {
	Date         string        `json:"date"`
	TotalItems   int64         `json:"totalItems"`
	TotalEvents  int64         `json:"totalEvents"`
	TotalGames   int64         `json:"totalGames"`
	TotalMatches int64         `json:"totalMatches"`
	Games        []Game        `json:"games"`
	Events       []interface{} `json:"events"`
	Matches      []interface{} `json:"matches"`
}

type Game struct {
	GamePk   int64    `json:"gamePk"`
	Link     string   `json:"link"`
	GameType GameType `json:"gameType"`
	Season   string   `json:"season"`
	GameDate string   `json:"gameDate"`
	Status   Status   `json:"status"`
	Teams    Teams    `json:"teams"`
	Venue    Venue    `json:"venue"`
	Content  Content  `json:"content"`
}

type Content struct {
	Link string `json:"link"`
}

type Status struct {
	AbstractGameState State  `json:"abstractGameState"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     State  `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	StartTimeTBD      bool   `json:"startTimeTBD"`
}

type Teams struct {
	Away Away `json:"away"`
	Home Away `json:"home"`
}

type Away struct {
	LeagueRecord LeagueRecord `json:"leagueRecord"`
	Score        int64        `json:"score"`
	Team         Venue        `json:"team"`
}

type LeagueRecord struct {
	WINS   int64  `json:"wins"`
	Losses int64  `json:"losses"`
	Ot     *int64 `json:"ot,omitempty"`
	Type   Type   `json:"type"`
}

type Venue struct {
	ID   *int64 `json:"id,omitempty"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type MetaData struct {
	TimeStamp string `json:"timeStamp"`
}

type GameType string

const (
	A  GameType = "A"
	P  GameType = "P"
	PR GameType = "PR"
	R  GameType = "R"
)

type State string

const (
	Final State = "Final"
)

type Type string

const (
	League Type = "league"
)
