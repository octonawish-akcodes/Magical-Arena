package player

// Player represents a player in the game, encapsulating their name, health, strength, and attack attributes.
type Player struct {
	name     string // The name of the player.
	health   int    // The health attribute of the player.
	strength int    // The strength attribute of the player.
	attack   int    // The attack attribute of the player.
}

// NewPlayer creates and initializes a new Player instance with the specified attributes.
//
// Parameters:
//   - name: The name of the player.
//   - health: The health attribute of the player.
//   - strength: The strength attribute of the player.
//   - attack: The attack attribute of the player.
//
// Returns:
//   - *Player: A pointer to the newly created Player instance.
//
// Example:
//
//	player := NewPlayer("Hero", 100, 10, 5)
//	fmt.Printf("%s has %d health, %d strength, and %d attack\n", player.name, player.health, player.strength, player.attack)
//
// Note: This example assumes direct access to the Player struct fields.
// If the fields are unexported (as in this case), accessor methods should be used.
func NewPlayer(name string, health, strength, attack int) *Player {
	return &Player{name, health, strength, attack}
}

// GetPlayerBaseAttributes returns the fundamental attributes of a player, including their name, health, strength, and attack.
//
// Parameters:
//   - p: A pointer to the Player whose basic attributes are to be retrieved.
//
// Returns:
//   - string: The name of the player.
//   - int: The health attribute of the player.
//   - int: The strength attribute of the player.
//   - int: The attack attribute of the player.
//
// Example:
//
//	name, health, strength, attack := GetPlayerBaseAttributes(player)
//	fmt.Printf("%s has %d health, %d strength, and %d attack\n", name, health, strength, attack)
//
// Note: This function is designed to provide easy access to a player's core attributes for use in game logic and display.
func GetPlayerBaseAttributes(p *Player) (string, int, int, int) {
	return p.name, p.health, p.strength, p.attack
}
