package match

import (
	"fmt"
	"os"
	"proj/pkg/player"
	"testing"
)

// ANSI escape codes for text color
const (
	redColor   = "\033[31m"
	greenColor = "\033[32m"
	resetColor = "\033[0m"
)

// TestGetDeterminStartingPlayer tests the GetDeterminStartingPlayer function,
// which determines the starting player for a match based on the health attributes
// of the players.
//
// Test scenarios:
//  1. Create a match with PlayerA's health 100 and PlayerB's health 50. Check that
//     PlayerB is the starting player.
//  2. Create a match with PlayerA's health 50 and PlayerB's health 100. Check that
//     PlayerA is the starting player.
//  3. Create a match with PlayerA's health 50 and PlayerB's health 100. Check that
//     PlayerA is the starting player.
//
// Note: The tests cover different scenarios to ensure that the starting player
// determination is based on health attributes as expected.
func TestGetDeterminStartingPlayer(t *testing.T) {
	//TEST 1: create a match with playerA health 100, playerB health 50
	//		check that playerB is the starting player
	playerA := player.NewPlayer("PlayerA", 100, 10, 5)
	playerB := player.NewPlayer("PlayerB", 50, 5, 2)
	match := NewMatch(playerA, playerB)
	startingPlayer := GetDeterminStartingPlayer(match)
	startingPlayerName, _, _, _ := player.GetPlayerBaseAttributes(startingPlayer)
	if startingPlayerName != "PlayerB" {
		t.Errorf(redColor+"Expected startingPlayerName to be PlayerB, got %s"+resetColor, startingPlayerName)
	} else {
		fmt.Println(greenColor + "TestGetDeterminStartingPlayer : Test1 : Passed" + resetColor)
	}

	//TEST 2: create a match with playerA health 50, playerB health 100
	//		check that playerB is the starting player
	playerA = player.NewPlayer("PlayerA", 100, 5, 2)
	playerB = player.NewPlayer("PlayerB", 100, 10, 5)
	match = NewMatch(playerA, playerB)
	startingPlayer = GetDeterminStartingPlayer(match)
	startingPlayerName, _, _, _ = player.GetPlayerBaseAttributes(startingPlayer)
	if startingPlayerName != "PlayerA" {
		t.Errorf(redColor+"Expected startingPlayerName to be PlayerA, got %s"+resetColor, startingPlayerName)
	} else {
		fmt.Println(greenColor + "TestGetDeterminStartingPlayer : Test2 : Passed" + resetColor)
	}

	//TEST 3: create a match with playerA health 50, playerB health 100
	// 		check that playerA is the starting player
	playerA = player.NewPlayer("PlayerA", 50, 5, 2)
	playerB = player.NewPlayer("PlayerB", 100, 10, 5)
	match = NewMatch(playerA, playerB)
	startingPlayer = GetDeterminStartingPlayer(match)
	startingPlayerName, _, _, _ = player.GetPlayerBaseAttributes(startingPlayer)
	if startingPlayerName != "PlayerA" {
		t.Errorf(redColor+"Expected startingPlayerName to be PlayerA, got %s"+resetColor, startingPlayerName)
	} else {
		fmt.Println(greenColor + "TestGetDeterminStartingPlayer : Test3 : Passed" + resetColor)
	}
}

// TestGetIsMatchOver tests the GetIsMatchOver function, which determines whether
// a match is over based on the health attributes of the players.
//
// Test scenarios:
//  1. Create a match with PlayerA's health 100 and PlayerB's health 50. Check that
//     the match is not over.
//  2. Create a match with PlayerA's health 0 and PlayerB's health 50. Check that
//     the match is over.
//  3. Create a match with PlayerA's health 50 and PlayerB's health 0. Check that
//     the match is over.
//  4. Create a match with PlayerA's health 0 and PlayerB's health 0. Check that
//     the match is over.
func TestGetIsMatchOver(t *testing.T) {
	//TEST 1: create a match with playerA health 100, playerB health 50
	// 		GetIsMatchOver should return false as both players are alive (i.e. health > 0)
	healthA := 100
	healthB := 50
	matchOver := GetIsMatchOver(healthA, healthB)
	if matchOver != false {
		t.Errorf(redColor+"Expected matchOver to be false, got %t"+resetColor, matchOver)
	} else {
		fmt.Println(greenColor + "TestGetIsMatchOver : Test1 : Passed" + resetColor)
	}

	//TEST 2: create a match with playerA health 0, playerB health 50
	// 		GetIsMatchOver should return true as playerA is dead (i.e. health <= 0)
	healthA = 0
	healthB = 50
	matchOver = GetIsMatchOver(healthA, healthB)
	if matchOver != true {
		t.Errorf(redColor+"Expected matchOver to be true, got %t"+resetColor, matchOver)
	} else {
		fmt.Println(greenColor + "TestGetIsMatchOver : Test2 : Passed" + resetColor)
	}

	//TEST 3: create a match with playerA health 50, playerB health 0
	// 		GetIsMatchOver should return true as playerB is dead (i.e. health <= 0)
	healthA = 50
	healthB = 0
	matchOver = GetIsMatchOver(healthA, healthB)
	if matchOver != true {
		t.Errorf(redColor+"Expected matchOver to be true, got %t"+resetColor, matchOver)
	} else {
		fmt.Println(greenColor + "TestGetIsMatchOver : Test3 : Passed" + resetColor)
	}

	//TEST 4: create a match with playerA health 0, playerB health 0
	// 		GetIsMatchOver should return true as both players are dead (i.e. health <= 0)
	healthA = 0
	healthB = 0
	matchOver = GetIsMatchOver(healthA, healthB)
	if matchOver != true {
		t.Errorf(redColor+"Expected matchOver to be true, got %t"+resetColor, matchOver)
	} else {
		fmt.Println(greenColor + "TestGetIsMatchOver : Test4 : Passed" + greenColor)
	}
}

func TestGetConductRound(t *testing.T) {
	// TEST 1: creating a current player with name "testA".
	// playerA is the current player, playerB is the opponent
	// playerA attributes: health 100, strength 10, attack 10
	// playerB attributes: health 50, strength 5, attack 2
	//expected health of PlayerB after round = 50 - max(0, 10*4 - 5*4) = 30
	playerA := player.NewPlayer("testA", 100, 10, 10)
	//playerB := player.NewPlayer("PlayerB", 50, 5, 2)
	currentPlayer := playerA
	roundResult, healthA, healthB := conductRound(currentPlayer, "testA", 100, 10, 10, "PlayerB", 50, 5, 2)
	if healthB != 30 {
		t.Errorf(redColor+"Expected healthB to be 30, got %d"+resetColor, healthB)
	}
	if roundResult != "testA attacked PlayerB for 20 damage" {
		t.Errorf(redColor+"Expected roundResult to be 'testA attacked PlayerB for 20 damage', got %s"+resetColor, roundResult)
	}
	if healthA != 100 {
		t.Errorf(redColor+"Expected healthA to be 100, got %d"+resetColor, healthA)
	} else {
		fmt.Println(greenColor + "TestGetConductRound : Test1 : Passed" + resetColor)
	}

	// TEST 2: creating a current player with name "testA".
	// playerA is the current player, playerB is the opponent
	// playerA attributes: health 100, strength 10, attack 4
	// playerB attributes: health 50, strength 5, attack 2
	//expected health of PlayerB after round = 50 - max(0, 4*4 - 5*4) = 50
	playerA = player.NewPlayer("testA", 100, 10, 4)
	//playerB := player.NewPlayer("PlayerB", 50, 5, 2)
	currentPlayer = playerA
	roundResult, healthA, healthB = conductRound(currentPlayer, "testA", 100, 10, 4, "PlayerB", 50, 5, 2)
	if healthB != 50 {
		t.Errorf(redColor+"Expected healthB to be 50, got %d"+resetColor, healthB)
	}
	if roundResult != "testA attacked PlayerB for 0 damage" {
		t.Errorf(redColor+"Expected roundResult to be 'testA attacked PlayerB for 0 damage', got %s"+resetColor, roundResult)
	}
	if healthA != 100 {
		t.Errorf(redColor+"Expected healthA to be 100, got %d"+resetColor, healthA)
	} else {
		fmt.Println(greenColor + "TestGetConductRound : Test2 : Passed" + resetColor)
	}

	// TEST 3: creating a current player with name "testB".
	// playerB is the current player, playerA is the opponent
	// playerA attributes:  health 50, strength 5, attack 2
	// playerB attributes: health 100, strength 10, attack 10
	// expected health of PlayerA after round = 50 - max(0, 10*4 - 5*4) = 30
	playerB := player.NewPlayer("testB", 100, 10, 10)
	currentPlayer = playerB
	roundResult, healthA, healthB = conductRound(currentPlayer, "PlayerA", 50, 5, 2, "testB", 100, 10, 10)
	if healthA != 30 {
		t.Errorf(redColor+"Expected healthA to be 30, got %d"+resetColor, healthA)
	}
	if roundResult != "testB attacked PlayerA for 20 damage" {
		t.Errorf(redColor+"Expected roundResult to be 'testB attacked PlayerA for 20 damage', got %s"+resetColor, roundResult)
	}
	if healthB != 100 {
		t.Errorf(redColor+"Expected healthB to be 100, got %d"+resetColor, healthB)
	} else {
		fmt.Println(greenColor + "TestGetConductRound : Test3 : Passed" + resetColor)
	}

	// TEST 4: creating a current player with name "testB".
	// playerB is the current player, playerA is the opponent
	// playerA attributes:  health 50, strength 5, attack 2
	// playerB attributes: health 100, strength 10, attack 4
	// expected health of PlayerA after round = 50 - max(0, 4*4 - 5*4) = 50
	playerB = player.NewPlayer("testB", 100, 10, 4)
	currentPlayer = playerB
	roundResult, healthA, healthB = conductRound(currentPlayer, "PlayerA", 50, 5, 2, "testB", 100, 10, 4)
	if healthA != 50 {
		t.Errorf(redColor+"Expected healthA to be 50, got %d"+resetColor, healthA)
	}
	if roundResult != "testB attacked PlayerA for 0 damage" {
		t.Errorf(redColor+"Expected roundResult to be 'testB attacked PlayerA for 0 damage', got %s"+resetColor, roundResult)
	}
	if healthB != 100 {
		t.Errorf(redColor+"Expected healthB to be 100, got %d"+resetColor, healthB)
	} else {
		fmt.Println(greenColor + "TestGetConductRound : Test4 : Passed" + resetColor)
	}
}

// TestGetSwitchCurrentPlayer tests the switching of the current player between playerA and playerB.
//
// TEST 1: If the current player is playerA, it should switch to playerB.
// - Create playerA and playerB with different attributes.
// - Set currentPlayer to playerA and call GetSwitchCurrentPlayer.
// - Check if the currentPlayerName is now "PlayerB".
//
// TEST 2: If the current player is playerB, it should switch to playerA.
// - Create playerA and playerB with different attributes.
// - Set currentPlayer to playerB and call GetSwitchCurrentPlayer.
// - Check if the currentPlayerName is now "PlayerA".
func TestGetSwitchCurrentPlayer(t *testing.T) {
	// TEST 1: if current player is playerA, switch to playerB
	playerA := player.NewPlayer("PlayerA", 100, 10, 10)
	currentPlayer := playerA
	playerB := player.NewPlayer("PlayerB", 50, 5, 2)
	GetSwitchCurrentPlayer(&currentPlayer, playerA, playerB)
	currentPlayerName, _, _, _ := player.GetPlayerBaseAttributes(currentPlayer)
	if currentPlayerName != "PlayerB" {
		t.Errorf(redColor+"Expected currentPlayerName to be PlayerB, got %s"+resetColor, currentPlayerName)
	} else {
		fmt.Println(greenColor + "TestGetSwitchCurrentPlayer : Test1 : Passed" + resetColor)
	}

	// TEST 2: if current player is playerB, switch to playerA
	playerA = player.NewPlayer("PlayerA", 100, 10, 10)
	playerB = player.NewPlayer("PlayerB", 50, 5, 2)
	currentPlayer = playerB
	GetSwitchCurrentPlayer(&currentPlayer, playerA, playerB)
	currentPlayerName, _, _, _ = player.GetPlayerBaseAttributes(currentPlayer)
	if currentPlayerName != "PlayerA" {
		t.Errorf(redColor+"Expected currentPlayerName to be PlayerA, got %s"+resetColor, currentPlayerName)
	} else {
		fmt.Println(greenColor + "TestGetSwitchCurrentPlayer : Test2 : Passed" + resetColor)
	}
}

// TestGetMatchResult tests the GetMatchResult function, which determines the
// result of a match based on the health attributes of the players.
//
// Test scenarios:
//  1. Create a match with PlayerA's health 100 and PlayerB's health 0. Check that
//     PlayerA wins.
//  2. Create a match with PlayerA's health 0 and PlayerB's health 100. Check that
//     PlayerB wins.
func TestGetMatchResult(t *testing.T) {
	// TEST 1: playerA wins
	matchResult := GetMatchResult("PlayerA", 100, "PlayerB", 0)
	if matchResult != "PlayerA wins" {
		t.Errorf(redColor+"Expected matchResult to be 'PlayerA wins', got %s"+resetColor, matchResult)
	} else {
		fmt.Println(greenColor + "TestGetMatchResult : Test1 : Passed" + resetColor)
	}

	// TEST 2: playerB wins
	matchResult = GetMatchResult("PlayerA", 0, "PlayerB", 100)
	if matchResult != "PlayerB wins" {
		t.Errorf(redColor+"Expected matchResult to be 'PlayerB wins', got %s"+resetColor, matchResult)
	} else {
		fmt.Println(greenColor + "TestGetMatchResult : Test2 : Passed" + resetColor)
	}
}

func TestConductMatch(t *testing.T) {
	//TEST 1: create a match with playerA health 100, playerB health 60
	// attribute of playerA: name=testA, health=100, strength=20, attack=20
	// attribute of playerB: name=testB, health=60, strength=10, attack=20
	// expected match result: testA wins
	playerA := player.NewPlayer("testA", 100, 20, 20)
	playerB := player.NewPlayer("testB", 60, 10, 20)
	match := NewMatch(playerA, playerB)
	_, matchResult := ConductMatch(match)
	if matchResult != "testA wins" {
		t.Errorf(redColor+"Expected matchResult to be 'testA wins', got %s"+resetColor, matchResult)
	} else {
		fmt.Println(greenColor + "TestConductMatch : Test1 : Passed" + resetColor)
	}

	//TEST 2: create a match with playerA health 60, playerB health 100
	// attribute of playerA: name=testA, health=60, strength=10, attack=20
	// attribute of playerB: name=testB, health=100, strength=20, attack=20
	// expected match result: testB wins
	playerA = player.NewPlayer("testA", 60, 10, 20)
	playerB = player.NewPlayer("testB", 100, 20, 20)
	match = NewMatch(playerA, playerB)
	_, matchResult = ConductMatch(match)
	if matchResult != "testB wins" {
		t.Errorf(redColor+"Expected matchResult to be 'testB wins', got %s"+resetColor, matchResult)
	} else {
		fmt.Println(greenColor + "TestConductMatch : Test2 : Passed" + resetColor)
	}
}

// TestMain runs the main testing suite.
func TestMain(m *testing.M) {
	fmt.Println("Testing Match package...")
	Result := m.Run()
	fmt.Println("Testing complete.")
	os.Exit(Result)
}
