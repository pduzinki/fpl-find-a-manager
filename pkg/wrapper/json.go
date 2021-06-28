package wrapper

type Manager struct {
	ID        int    `json:"id"`
	FirstName string `json:"player_first_name"`
	LastName  string `json:"player_last_name"`
}

type totalPlayers struct {
	Count int `json:"total_players"`
}
