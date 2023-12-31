package models

var URL string = "https://www.dota2.com/webapi/ILeaderboard/GetDivisionLeaderboard/v0001?"

func FormURL(region string) string {
	// TODO: search for this api

	return URL + "division=" + region + "&leaderboard=0"
}
