// Treasure-Stealing Trolls
// John Sloan

// This program generates n trolls, numbered 0 to n-1. Each troll has a bag of treasure t
// initially worth $1000000 and starts under a bridge at a position p equal to its id * 1000000.
// The trolls take turns moving according to id in ascending order. On a trolls turn, it moves to
// position p + r * t, where p is its current position, r is a random float between -2 and 2,
// and t is its current amount of treasure. After moving, a troll steals half of the treasure from
// both its new neighbors. If it only has one neighbor, it steals all that neighbors treasure.
// If a troll loses all its treasure, it is removed from the game. The trolls continue taking
// turns until only one remains.

package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Troll struct {
	idNumber, gold, turnIndex, posIndex, position int
}

// By declaring a String() function on type Troll, Troll implicitly implements the Stringer interface
func (t *Troll) String() string {
	return fmt.Sprintf("%d > t=%d  p=%d\n", t.idNumber, t.gold, t.position)
}

func (t *Troll) move(s []*Troll) {
	r := ((rand.Float32() * 4) - 2)
	distance := float32(t.gold) * r
	t.position += int(distance)

	switch {
	case r < 0:
		prev := s[t.posIndex-1]
		for t.position < prev.position {
			s[t.posIndex] = prev
			prev.posIndex++
			s[t.posIndex-1] = t
			t.posIndex--
			prev = s[t.posIndex-1]
		}
	case r > 0:
		next := s[t.posIndex+1]
		for t.position > next.position {
			s[t.posIndex] = next
			next.posIndex--
			s[t.posIndex+1] = t
			t.posIndex++
			next = s[t.posIndex+1]
		}
	default:
	}
}

func (t *Troll) steal(s, b []*Troll) ([]*Troll, []*Troll) {
	var dead *Troll

	switch t.posIndex {
	case 1:
		loot := b[t.posIndex+1].gold
		t.gold += loot
		b[t.posIndex+1].gold -= loot
		dead = b[t.posIndex+1]
	case len(b) - 2:
		loot := b[t.posIndex-1].gold
		t.gold += loot
		b[t.posIndex-1].gold -= loot
		dead = b[t.posIndex-1]
	default:
		loot := b[t.posIndex+1].gold / 2
		t.gold += loot
		b[t.posIndex+1].gold -= loot
		loot = b[t.posIndex-1].gold / 2
		t.gold += loot
		b[t.posIndex-1].gold -= loot
		return s, b
	}

	st1 := s[:dead.turnIndex]
	st2 := s[dead.turnIndex+1:]
	s = append(st1, st2...)

	for i := dead.turnIndex; i < len(s); i++ {
		s[i].turnIndex = i
	}

	sb1 := b[:dead.posIndex]
	sb2 := b[dead.posIndex+1:]
	b = append(sb1, sb2...)
	for i := dead.posIndex; i < len(b); i++ {
		b[i].posIndex = i
	}

	return s, b
}

func drawTrolls() {
	clearScreen := exec.Command("clear")
	clearScreen.Stdout = os.Stdout
	clearScreen.Run()

	fmt.Printf("\n      |\\,_,/|          |\\,_,/|")
	fmt.Printf("\n      ( 0 0 )          ( 0 0 )")
	fmt.Printf("\n ,-oOO--{_)--OOo----oOO--{_}--OOo-,")
	fmt.Printf("\n \\      ~Treasure  Stealing~      \\")
	fmt.Printf("\n /   ooO      ~Trolls~      Ooo   /")
	fmt.Printf("\n `---(_)---Ooo--+--+---ooO--(_)---`")
	fmt.Printf("\n           (_)  |  |   (_)")
	fmt.Printf("\n ^^^^^^^^^^^^^^^+--+^^^^^^^^^^^^^^^\n\n")
}

func showTrolls(x int, s []*Troll, w *bufio.Writer) {
	fmt.Fprintf(w, "\n ~~Turn %d~~\n", x)

	for i := 1; i < len(s)-1; i++ {
		fmt.Fprint(w, s[i].String())
	}

	w.Flush()
}

func main() {
	var (
		numberOfTrolls, turn int
	)

	w := bufio.NewWriterSize(os.Stdout, 16384)
	drawTrolls()
	fmt.Print("How many trolls? ")
	_, err := fmt.Scan(&numberOfTrolls)

	if err != nil {
		fmt.Printf("Error: %s.\n\n", err)
		os.Exit(1)
	}

	fmt.Printf("Alright, then. Just lemme toss Trolls 0 through %d under the bridge...\nAaanndd... ",
		numberOfTrolls-1)

	for i := 3; i > 0; i-- {
		fmt.Printf("%d... ", i)
		time.Sleep(time.Second / 2)
	}

	fmt.Println("GO!!!\n")
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	trolls := make([]*Troll, numberOfTrolls)
	bridge := make([]*Troll, numberOfTrolls+2)

	for i := 0; i < len(trolls); i++ {
		troll := Troll{idNumber: i, gold: 1000000, turnIndex: i, posIndex: i + 1, position: i * 1000000}
		trolls[i] = &troll
		bridge[i+1] = &troll
	}

	for i := 0; i < len(bridge); i += len(bridge) - 1 {
		x := 1

		if i == 0 {
			x = -1
		}

		troll := Troll{-1, 0, -1, i, x * math.MaxInt64}
		bridge[i] = &troll
	}

	showTrolls(turn, bridge, w)

	for len(trolls) > 1 {
		turn++

		for i := 0; i < len(trolls); i++ {
			t := trolls[i]
			t.move(bridge)
			trolls, bridge = t.steal(trolls, bridge)
		}

		showTrolls(turn, bridge, w)
	}

	elapsed := time.Since(start)
	t := bridge[1]
	fmt.Printf(
		"\nWinner: %d Treasure: $%d Position: %d\nElapsed time: %s Number of Trolls: %d Number of Turns: %d\n",
		t.idNumber, t.gold, t.position, elapsed, numberOfTrolls, turn)
}
