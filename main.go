package main

import (
	"fmt"
	"onefootball/constants"
	"onefootball/types"
	"onefootball/utils"
	"strings"
	"sync"

	"go.uber.org/zap"
)

var globalLogger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	globalLogger = logger.Sugar()
}

type teamPlayers map[string][]types.PlayersInfo

// used as a set
var teamNames map[string]bool

func queryAPIEndpoint(start int64, end int64, jobs chan<- int, apiResponse chan<- teamPlayers) chan<- teamPlayers {

	defer func() {
		// release a slot back into the worker pool
		jobs <- 0
	}()

	for i := start; i < end; i++ {
		if utils.AreAllTeamsFound(teamNames) {
			break
		}

		url := fmt.Sprintf(constants.API_ENDPOINT, i)
		resp, err := utils.HandleRequest(url)
		if err != nil || resp == nil {
			globalLogger.Debugw("Error in fetching team data", "err", err)
		} else {
			teamName := utils.GetTeamName(resp)

			globalLogger.Debugw("Processing team", "name", teamName)
			// check if the provided name or team is present in the response
			if utils.IsRequiredTeam(teamNames, strings.ToLower(teamName)) {
				playerData := teamPlayers{}

				// Get all the players information
				playerData[teamName] = resp.Data.Team.Players

				apiResponse <- playerData

				// delete the records if the given team name is present
				utils.RemoveFromRequiredTeams(teamNames, strings.ToLower(teamName))
			}
		}
	}
	return apiResponse
}

// extractPlayerInfo fetch all the response from api and print the output
func extractPlayerInfo() []types.PlayerInformation {
	numberOfTeams := len(teamNames)

	var currentID int64 = 0
	var jobsPerGoroutine int64 = 10

	numJobs := 10

	jobs := make(chan int, numJobs)
	apiResponse := make(chan teamPlayers, numberOfTeams)
	retmap := map[string]interface{}{}
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {

		defer wg.Done()

		for resp := range apiResponse {
			for k, v := range resp {
				retmap[k] = v
			}
		}
	}()

	for {
		<-jobs
		go queryAPIEndpoint(currentID, currentID+jobsPerGoroutine, jobs, apiResponse)
		currentID += jobsPerGoroutine

		if utils.AreAllTeamsFound(teamNames) {
			break
		}
	}

	// closing api response so that our response parser goroutine can end
	close(apiResponse)

	wg.Wait()

	compiledPlayerInfo := map[string]types.PlayerInformation{}

	// extract the player info and manipulate and complied the information
	for teamName, playerInfo := range retmap {
		for _, player := range playerInfo.([]types.PlayersInfo) {
			if info, ok := compiledPlayerInfo[player.Id]; !ok {
				compiledPlayerInfo[player.Id] = types.PlayerInformation{
					PlayerId: player.Id,
					Name:     player.Name,
					Age:      player.Age,
					Teams:    teamName,
				}
			} else {
				info.Teams = info.Teams + ", " + teamName
				compiledPlayerInfo[player.Id] = info
			}
		}
	}

	// Sort the players by ID in ascending order
	playerData := utils.GetSortedPlayers(compiledPlayerInfo)

	return playerData

}

func main() {
	// we initialize our team names array
	teamNames = map[string]bool{
		"germany":           true,
		"england":           true,
		"france":            true,
		"spain":             true,
		"manchester united": true,
		"arsenal":           true,
		"chelsea":           true,
		"barcelona":         true,
		"real madrid":       true,
		"bayern munich":     true,
	}

	// calling function to extract player information
	playerData := extractPlayerInfo()

	globalLogger.Info("Players information's with their teams")

	for id, player := range playerData {
		playerInfo := utils.FormatPlayersInformation(id, player)
		fmt.Println(playerInfo)
	}
}
