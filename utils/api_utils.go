package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"onefootball/types"
	"sort"
	"strconv"
	"sync"

	"go.uber.org/zap"
)

var globalLogger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	globalLogger = logger.Sugar()
}

// mutex for accessing the map
var mapMutex sync.RWMutex

// IsRequiredTeam return True if required team is found in api response
func IsRequiredTeam(teamNames map[string]bool, name string) bool {
	mapMutex.RLock()
	defer mapMutex.RUnlock()

	_, ok := teamNames[name]
	return ok
}

// RemoveFromRequiredTeams removed the team name from the map if record of the desired team is found
func RemoveFromRequiredTeams(teamNames map[string]bool, name string) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	delete(teamNames, name)
}

// AreAllTeamsFound returns True if all teams players information are extracted
func AreAllTeamsFound(teamNames map[string]bool) bool {
	mapMutex.RLock()
	defer mapMutex.RUnlock()

	return len(teamNames) == 0
}

// GetTeamName Execution should be stopped if error is not returned nil from this function.
func GetTeamName(data *types.TeamInformation) string {
	if data.Status != "ok" {
		globalLogger.Debugw("Error in JSON response status", "err", data.Status)
		return ""
	} else {
		return data.Data.Team.Name
	}
}

// HandleRequest calling api request to get the response from the api
func HandleRequest(url string) (*types.TeamInformation, error) {
	globalLogger.Debugw("Making API request", "url", url)
	res, err := http.Get(url)
	if err != nil {
		globalLogger.Debugw("Error in fetching team data", "err", err)
		return nil, err
	} else {
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()

		// If the API response isn't 200 then I will consider it as an error response and will imediately return from here.
		if err != nil {
			globalLogger.Debugw("Error in API response", "err", err)
			return nil, fmt.Errorf("non ok response received")

			// Here, the response body is being checked. If there is an error in reading the response then function will also return from here.
		} else if res.StatusCode != http.StatusOK {
			globalLogger.Debugw("Error in reading the response body of the API call")
			return nil, fmt.Errorf("non ok response received")
		}

		var data types.TeamInformation

		if err := json.Unmarshal(body, &data); err != nil {
			globalLogger.Debugw("Error in parsing recieved response", "err", err)
			return nil, err
		}
		// Only success response.
		return &data, nil
	}
}

// GetSortedPlayers sort the players by their players id
func GetSortedPlayers(compiledPlayerInfo map[string]types.PlayerInformation) []types.PlayerInformation {
	var sortedPlayers []types.PlayerInformation

	//  First getting all the keys, which are string convert to ints, then sorting those keys.
	keys := make([]int, 0, len(compiledPlayerInfo))
	for _, value := range compiledPlayerInfo {
		id, _ := strconv.Atoi(value.PlayerId)
		keys = append(keys, id)
	}

	sort.Ints(keys)

	// Creating another array of sorted players, then iterating over sorted keys and adding them to the sorted player array.
	for _, k := range keys {
		for key, value := range compiledPlayerInfo {
			key1, _ := strconv.Atoi(key)
			if key1 == k {
				sortedPlayers = append(sortedPlayers, value)
			}
		}
	}
	return sortedPlayers
}

// FormatPlayersInformation prettify/format the players information
func FormatPlayersInformation(id int, player types.PlayerInformation) string {
	playerInfo := fmt.Sprintf("%d. %s; %s; %s; %s", id, player.PlayerId, player.Name, player.Age, player.Teams)
	return playerInfo
}
