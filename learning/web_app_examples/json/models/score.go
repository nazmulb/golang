package models

import (
	"gopkg.in/mgo.v2"
)

var mongo *mgo.Session

// Score of providers
type Score struct {
	ProviderUserID               int     `json:"provider_user_id" bson:"provider_user_id"`
	TypeOfWorkID                 string  `json:"type_of_work_id" bson:"type_of_work_id"`
	ProviderScoreRank            float32 `json:"provider_score_rank" bson:"provider_score_rank"`
	ProviderQualityScoreRank     float32 `json:"provider_quality_score_rank" bson:"provider_quality_score_rank"`
	ProviderPerformanceScoreRank float32 `json:"provider_performance_score_rank" bson:"provider_performance_score_rank"`
}

// CalculateScoreRank method calculates rank
func (s *Score) CalculateScoreRank() {
	s.ProviderScoreRank = (s.ProviderQualityScoreRank + s.ProviderPerformanceScoreRank) / 2
}

// SetTypeOfWork method ser type of work
func (s *Score) SetTypeOfWork(tow string) {
	s.TypeOfWorkID = tow
}

// GetDB method provides MongoDB connection
func GetDB() *mgo.Session {

	if mongo != nil {
		return mongo
	}

	session, err := mgo.Dial("127.0.0.1")
	CheckErr(err)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	mongo = session

	return mongo
}

// CheckErr method checks errors and throws
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
