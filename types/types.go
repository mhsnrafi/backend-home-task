package types

type TeamInformation struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   struct {
		Team struct {
			ID              int                `json:"id"`
			OptID           int                `json:"optID"`
			Country         string             `json:"country"`
			CountryName     string             `json:"countryName"`
			Name            string             `json:"name"`
			IsNational      bool               `json:"isNational"`
			HasOfficialPage bool               `json:"hasOfficialPage"`
			Players         []PlayersInfo      `json:"players"`
			LogoURLS        []LogoURLSInfo     `json:"logo_urls"`
			Competitions    []CompetitionsInfo `json:"competitions"`
			Officials       []OfficialsInfo    `json:"officials"`
			Colors          ColorsInfo         `json:"colors"`
		}
	}
	Message string `json:"message"`
}

// LogoURLSInfo LogoUrls info struct
type LogoURLSInfo struct {
	Size string `json:"size"`
	URL  string `json:"url"`
}

// PlayersInfo player info struct
type PlayersInfo struct {
	Id           string            `json:"id"`
	Country      string            `json:"country"`
	FirstName    string            `json:"firstName"`
	LastName     string            `json:"lastName"`
	Name         string            `json:"name"`
	Position     string            `json:"position"`
	Number       int               `json:"number"`
	Birthdate    string            `json:"birthdate"`
	Age          string            `json:"age"`
	Height       int               `json:"height"`
	Weight       int               `json:"weight"`
	ThumbnailSrc string            `json:"thumbnailSrc"`
	Affiliation  map[string]string `json:"affiliation"`
}

// CompetitionsInfo competitions info struct
type CompetitionsInfo struct {
	CompetitionId   int    `json:"competitionId"`
	CompetitionName string `json:"competitionName"`
}

// OfficialsInfo officials info struct
type OfficialsInfo struct {
	CountryName  string            `json:"countryName"`
	ID           string            `json:"id"`
	FirstName    string            `json:"firstName"`
	LastName     string            `json:"lastName"`
	Country      string            `json:"country"`
	Position     string            `json:"position"`
	ThumbnailSrc string            `json:"thumbnailSrc"`
	Affiliation  map[string]string `json:"affiliation"`
}

// ColorsInfo colors info struct
type ColorsInfo struct {
	ShirtColorHome string `json:"shirtColorHome"`
	ShirtColorAway string `json:"shirtColorAway"`
	CrestMainColor string `json:"CrestMainColor"`
	MainColor      string `json:"mainColor"`
}

type PlayerInformation struct {
	PlayerId string `json:"player_id"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Teams    string `json:"teams"`
}
