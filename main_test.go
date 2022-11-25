package main

import (
	"fmt"
	"onefootball/constants"
	"onefootball/types"
	"onefootball/utils"
	"testing"
)

func TestALLTeamsFound(t *testing.T) {
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
	isAllTeamFound := utils.AreAllTeamsFound(teamNames)
	if isAllTeamFound != false {
		t.Fatalf("failed to find all teams")
	}

	// Test Case 2
	teamNames = map[string]bool{
		"germany": true,
	}
	isAllTeamFound = utils.AreAllTeamsFound(teamNames)
	if isAllTeamFound != false {
		t.Fatalf("failed to find all teams")
	}

	// Test Case 3
	teamNames = map[string]bool{}
	isAllTeamFound = utils.AreAllTeamsFound(teamNames)
	if isAllTeamFound != true {
		t.Fatalf("failed to find all teams")
	}
}

func TestGetTeamName(t *testing.T) {

	url := fmt.Sprintf(constants.API_ENDPOINT, 1)
	resp, err := utils.HandleRequest(url)
	if err != nil {
		globalLogger.Debugw("Error in fetching team data", "err", err)
	}
	teamName := utils.GetTeamName(resp)
	if teamName != "APOEL" {
		t.Fatalf("Expected result %s got %s", "APOEL", teamName)
	}

	// Test Case 2
	url = fmt.Sprintf(constants.API_ENDPOINT, 5)
	resp, err = utils.HandleRequest(url)
	if err != nil {
		globalLogger.Debugw("Error in fetching team data", "err", err)
	}

	teamName = utils.GetTeamName(resp)
	if teamName != "Barcelona" {
		t.Fatalf("Expected result %s got %s", "Barcelona", teamName)
	}
}

func TestRemoveFromRequiredTeams(t *testing.T) {
	teamNames = map[string]bool{
		"germany": true,
	}
	type args struct {
		teamNames map[string]bool
		name      string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Remove Required Team",
			args: args{
				teamNames: map[string]bool{
					"germany":   true,
					"barcelona": true,
				},
				name: "barcelona",
			},
		},
		{
			name: "Test Remove Empty Team",
			args: args{
				teamNames: map[string]bool{
					"germany":   true,
					"barcelona": true,
				},
				name: "dummy",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.RemoveFromRequiredTeams(tt.args.teamNames, tt.args.name)
		})
	}
}

func TestIsRequiredTeam(t *testing.T) {
	type args struct {
		teamNames map[string]bool
		name      string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test is required team",
			args: args{
				teamNames: map[string]bool{
					"germany":   true,
					"barcelona": true,
				},
				name: "germany",
			},
			want: true,
		},
		{
			name: "Test is required team",
			args: args{
				teamNames: map[string]bool{
					"germany":   true,
					"barcelona": true,
				},
				name: "japan",
			},
			want: false,
		},
		{
			name: "Test is required team",
			args: args{
				teamNames: map[string]bool{
					"germany":   true,
					"barcelona": true,
				},
				name: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsRequiredTeam(tt.args.teamNames, tt.args.name); got != tt.want {
				t.Errorf("IsRequiredTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractALLTeamsPlayerInfo(t *testing.T) {
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
	var tests = []struct {
		want types.PlayerInformation
	}{
		{
			types.PlayerInformation{
				PlayerId: "93",
				Name:     "David De Gea",
				Age:      "32",
				Teams:    "Manchester United",
			},
		},
	}

	for _, tt := range tests {
		testname := "Test Extract All Teams Player Information"
		t.Run(testname, func(t *testing.T) {
			ans := extractPlayerInfo()
			if len(ans) != 244 {
				t.Errorf("failed to discover the complete id's")

			}
			if ans[0] != tt.want {
				t.Errorf("got %s, want %s", ans[0].PlayerId, tt.want.PlayerId)

			}
			if ans[0].PlayerId != tt.want.PlayerId && ans[0].Name != tt.want.Name && ans[0].Age != tt.want.Age && ans[0].Teams != tt.want.Teams {
				t.Errorf("got %s, want %s", ans[0].PlayerId, tt.want.PlayerId)
			}
		})
	}
}

func TestExtractTeamPlayerInfo(t *testing.T) {
	teamNames = map[string]bool{
		"germany": true,
	}
	var tests = []struct {
		want types.PlayerInformation
	}{
		{
			types.PlayerInformation{
				PlayerId: "173",
				Name:     "Thomas Müller",
				Age:      "33",
				Teams:    "Germany",
			},
		},
	}

	for _, tt := range tests {
		testname := "Test Extract Player Information"
		t.Run(testname, func(t *testing.T) {
			ans := extractPlayerInfo()
			if ans[0] != tt.want {
				t.Errorf("got %s, want %s", ans[0].PlayerId, tt.want.PlayerId)
			}
			if ans[0].PlayerId != tt.want.PlayerId && ans[0].Name != tt.want.Name && ans[0].Age != tt.want.Age && ans[0].Teams != tt.want.Teams {
				t.Errorf("got %s, want %s", ans[0].PlayerId, tt.want.PlayerId)
			}
		})
	}
}
