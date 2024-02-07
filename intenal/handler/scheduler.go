package handler

import (
	"time"
)

// Scheduler TODO: realize the work of the scheduler using cron lib
func (h *Handler) Scheduler() {
	for {

		regions := []string{"europe"}

		//database
		leaderboard, err := h.service.MakeRequest(regions)
		if err != nil {
			h.errorLogger.Fatalf("Error while making request to dotaapi: %s", err)
		}

		server, err := h.service.RequestServerInfo()
		if err != nil {
			h.errorLogger.Fatalf("Error while requesting server info %s", err)
		}

		if err := h.service.UpdateServerInfo(server); err != nil {
			h.errorLogger.Fatalf("Error while updating server infor %s", err)
		}

		if err := h.service.UpdateLeaderboard(leaderboard[0]); err != nil {
			h.errorLogger.Fatalf("Error while updating leaderboard: %s", err)
		}

		if err := h.service.Filler.UpdateRedis(leaderboard[0]); err != nil {
			h.errorLogger.Fatalf("error while updating redis %s", err)
		}

		sleepTime := server.NextScheduledPostTime - server.ServerTime

		//logic for sleepTime for < 60 seconds
		if sleepTime < 60 {
			sleepTime += 3540
		}

		h.infoLogger.Printf("time posted of the leaderboard is %d\n", server.TimePosted)

		h.infoLogger.Printf("next scheduled post time is %d\n", server.NextScheduledPostTime)

		h.infoLogger.Printf("server time is %d\n", server.ServerTime)

		h.infoLogger.Printf("next scheduled post time in seconds is %d\n", sleepTime)

		sleepTimeInMinute := sleepTime / 60

		h.infoLogger.Printf("Request has been handled! Next scheduled request will be in %d minute\n\n", sleepTimeInMinute)

		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
