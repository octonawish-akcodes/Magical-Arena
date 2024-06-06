package player

import (
	"fmt"
	"os"
	"testing"
)

// ANSI escape codes for text color
const (
	redColor   = "\033[31m"
	greenColor = "\033[32m"
	resetColor = "\033[0m"
)

// TestGetPlayerBaseAttributes tests the GetPlayerBaseAttributes function.
func TestNewPlayer(t *testing.T) {
	//testing NewPlayer and GetPlayerBaseAttributes as a single unit
	//TEST 1: create a player name Ironman with health 100, strength 10, attack 5, and check the assignment of values is correct
	player := NewPlayer("Ironman", 100, 10, 5)
	name, health, strength, attack := GetPlayerBaseAttributes(player)
	if name != "Ironman" {
		t.Errorf(redColor+"Expected player.name to be shaleen, got %s"+resetColor, name)
	}
	if health != 100 {
		t.Errorf(redColor+"Expected player.health to be 100, got %d"+resetColor, health)
	}
	if strength != 10 {
		t.Errorf(redColor+"Expected player.strength to be 10, got %d"+resetColor, strength)
	}
	if attack != 5 {
		t.Errorf(redColor+"Expected player.attack to be 5, got %d"+resetColor, attack)
	} else {
		fmt.Println(greenColor + "TestNewPlayer: Test1 : Passed" + resetColor)
	}

	//TEST 2: create a playerA with health 50, strength 5, attack 2
	//		create a playerB with health 100, strength 10, attack 5, and check the assignment of values is correct
	playerA := NewPlayer("PlayerA", 50, 5, 2)
	playerB := NewPlayer("PlayerB", 100, 10, 5)
	nameA, healthA, strengthA, attackA := GetPlayerBaseAttributes(playerA)
	nameB, healthB, strengthB, attackB := GetPlayerBaseAttributes(playerB)
	//check playerA
	if nameA != "PlayerA" {
		t.Errorf(redColor+"Expected playerA.name to be PlayerA, got %s"+resetColor, nameA)
	}
	if healthA != 50 {
		t.Errorf(redColor+"Expected playerA.health to be 50, got %d"+resetColor, healthA)
	}
	if strengthA != 5 {
		t.Errorf(redColor+"Expected playerA.strength to be 5, got %d"+resetColor, strengthA)
	}
	if attackA != 2 {
		t.Errorf(redColor+"Expected playerA.attack to be 2, got %d"+resetColor, attackA)
	}
	//check playerB
	if nameB != "PlayerB" {
		t.Errorf(redColor+"Expected playerB.name to be PlayerB, got %s"+resetColor, nameB)
	}
	if healthB != 100 {
		t.Errorf(redColor+"Expected playerB.health to be 100, got %d"+resetColor, healthB)
	}
	if strengthB != 10 {
		t.Errorf(redColor+"Expected playerB.strength to be 10, got %d"+resetColor, strengthB)
	}
	if attackB != 5 {
		t.Errorf(redColor+"Expected playerB.attack to be 5, got %d"+resetColor, attackB)
	} else {
		fmt.Println(greenColor + "TestNewPlayer: Test2 : Passed" + resetColor)

	}

	//TEST 3: create a playerA with health 0, strength 0, attack 0
	//create a playerB with health 0, strength 0, attack 0, and check that they are equal.
	playerA = NewPlayer("0", 0, 0, 0)
	playerB = NewPlayer("0", 0, 0, 0)
	nameA, healthA, strengthA, attackA = GetPlayerBaseAttributes(playerA)
	nameB, healthB, strengthB, attackB = GetPlayerBaseAttributes(playerB)
	if nameA != nameB || healthA != healthB || strengthA != strengthB || attackA != attackB {
		t.Errorf(redColor+"Expected playerA and playerB to be equal, got %d %d %d and %d %d %d"+resetColor, healthA, strengthA, attackA, healthB, strengthB, attackB)
	} else {
		fmt.Println(greenColor + "TestNewPlayer: Test3 : Passed" + resetColor)
	}
}

// TestMain runs the main testing suite.
func TestMain(m *testing.M) {
	fmt.Println("Testing player package...")
	Result := m.Run()
	fmt.Println("Testing complete.")
	os.Exit(Result)
}
