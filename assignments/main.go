package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("hello world")

	fmt.Println(arraySign([]int{2, 1}))                    // 1
	fmt.Println(arraySign([]int{-2, 1}))                   // -1
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1})) // 1

	fmt.Println(isAnagram("anak", "kana"))       // true
	fmt.Println(isAnagram("anak", "mana"))       // false
	fmt.Println(isAnagram("anagram", "managra")) // true

	fmt.Println(string(findTheDifference("abcd", "abcde"))) // 'e'
	fmt.Println(string(findTheDifference("abcd", "abced"))) // 'e'
	fmt.Println(string(findTheDifference("", "y")))         // 'y'

	fmt.Println(canMakeArithmeticProgression([]int{1, 5, 3}))    // true; 1, 3, 5 adalah baris aritmatik +2
	fmt.Println(canMakeArithmeticProgression([]int{5, 1, 9}))    // true; 9, 5, 1 adalah baris aritmatik -4
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4, 8})) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// write code here

	sign := 1
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			sign = sign * -1
		}
	}

	return sign // if positive
	// return -1 // if negative
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// write code here
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	for _, c := range t {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	// write code here
	result := byte(0)
	for i := 0; i < len(s); i++ {
		result ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		result ^= t[i]
	}

	return result
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// write code here
	if len(arr) < 0 {
		return true
	}

	sort.Ints(arr)

	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true
}

// Deck represents a "standard" deck consisting of 52 cards.
type Deck struct {
	cards []Card
}

// Card represents a card in a "standard" deck.
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New initializes the deck with 52 cards, sorted by symbol and then number.
func (d *Deck) New() {
	d.cards = nil            // Clear any existing cards
	for s := 0; s < 4; s++ { // symbol: 0 (spade), 1 (heart), 2 (club), 3 (diamond)
		for n := 1; n <= 13; n++ { // number: 1 (Ace) to 13 (King)
			d.cards = append(d.cards, Card{symbol: s, number: n})
		}
	}
}

// PeekTop returns n cards from the top of the deck.
func (d Deck) PeekTop(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]
}

// PeekBottom returns n cards from the bottom of the deck.
func (d Deck) PeekBottom(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex returns the card at the specified index.
func (d Deck) PeekCardAtIndex(idx int) Card {
	if idx < 0 || idx >= len(d.cards) {
		return Card{} // Return a zero-value card if index is out of range
	}
	return d.cards[idx]
}

// Shuffle randomly shuffles the deck.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// Cut performs a single "Cut" technique. Moves the top n cards to the bottom.
func (d *Deck) Cut(n int) {
	if n < 0 || n >= len(d.cards) {
		return
	}
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

// ToString returns a string representation of the card.
func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
