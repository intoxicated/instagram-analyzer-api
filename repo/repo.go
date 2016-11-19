package repo

import (
	"time"

	"github.com/intoxicated/instagram-analyzer-api/config"
	"github.com/intoxicated/instagram-analyzer-api/dto"
)

//UpdateTablesFromData ...
func UpdateTablesFromData(db *config.DB, data interface{}) error {
	//update relevant tables from media array

	//update hashtag_trend_count
	return nil
}

//GetHashtagTrends ...
func GetHashtagTrends(db *config.DB,
	startDate time.Time,
	endDate time.Time) (*dto.InstagramHashtags, error) {
	//query hashtag_trend_count
	//something like..
	/* [{
			name: "iteawon",
			changes: [{count:3, date:"11/04/2016"}, //start date
							{count:8, date:"11/05/2016"},
							{count:14, date:"11/06/2016"},
							{count:39, date:"11/07/2016"}] //end date
	 	 },
		 {
		 	 name: "yeoksam",
			 changes: [{count:42, date:"11/04/2016"},
			 					{count:31, date:"11/06/2016"}]
		 }]
	*/
	return nil, nil
}

//UpdateHashtags ...
func UpdateHashtags(db *config.DB, tags dto.InstagramHashtags) *dto.InstagramHashtags {
	//for each instagramHashtag
	//get hashtag id (insert if not exist)
	//insert count in row hashtag_count with hashtag id
	return nil
}

//GetHashtagDiff ...
func GetHashtagDiff(db *config.DB, tag string) *dto.InstagramHashtag {
	//get today
	//get yesterdays count
	//substract
	return nil
}

//GetWatchingHashtags ...
//return a set of hashtag that a user is watching with +/- changes
//compare to day before
func GetWatchingHashtags(db *config.DB) *dto.InstagramHashtags {
	return nil
}
