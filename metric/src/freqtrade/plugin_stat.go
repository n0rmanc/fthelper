package freqtrade

import (
	"github.com/n0rmanc/fthelper/metric/v4/src/connection"
	"github.com/n0rmanc/fthelper/shared/datatype"
)

const (
	STAT_CONST     = "stats"
	StatNotANumber = -1
)

func NewStat() *stat {
	return &stat{
		Reasons: make(map[string]statObject),
		Duration: durationObject{
			Win:  StatNotANumber,
			Loss: StatNotANumber,
			Draw: StatNotANumber,
		},
	}
}

type statObject struct {
	Win  int64 `json:"wins"`
	Loss int64 `json:"losses"`
	Draw int64 `json:"draws"`
}

type durationObject struct {
	Win  float64 `json:"wins"`
	Loss float64 `json:"losses"`
	Draw float64 `json:"draws"`
}

type stat struct {
	Reasons  map[string]statObject `json:"sell_reasons"`
	Duration durationObject        `json:"durations"`
}

func (s *stat) WinDuration() float64 {
	return s.Duration.Win
}

func (s *stat) DrawDuration() float64 {
	return s.Duration.Draw
}

func (s *stat) LossDuration() float64 {
	return s.Duration.Loss
}

func (s *stat) Name() string {
	return STAT_CONST
}

func (s *stat) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	err := connection.Http.GET(s.Name(), s)
	return s, err
}

func ToStat(connector connection.Connector) (*stat, error) {
	raw, err := connector.Connect(STAT_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*stat), nil
}
