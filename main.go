package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner *bufio.Scanner

type Object struct {
	name   string
	active bool
}

type Choice struct {
	Cmd      string
	Desc     string
	NextNode *StoryNode
	Next     *Choice
}

type StoryNode struct {
	Text   string
	Choice *Choice
}

func NewNode(text string) *StoryNode {
	return &StoryNode{
		Text: text,
	}
}

func NewObject(name string) *Object {
	return &Object{name: name, active: false}
}

func (node *StoryNode) AddChoice(cmd string, desc string, storyNode *StoryNode) {
	choice := &Choice{
		Cmd:      cmd,
		Desc:     desc,
		NextNode: storyNode,
	}

	if node.Choice == nil {
		node.Choice = choice
		return
	}

	currentChoice := node.Choice
	for currentChoice.Next != nil {
		currentChoice = currentChoice.Next
	}

	currentChoice.Next = choice
}

func (node *StoryNode) render() {
	fmt.Println(node.Text)

	currentChoice := node.Choice
	for currentChoice != nil {
		fmt.Println(currentChoice.Cmd, ":", currentChoice.Desc)
		currentChoice = currentChoice.Next
	}
}

func (node *StoryNode) executeCmd(cmd string) *StoryNode {
	currentChoice := node.Choice
	for currentChoice != nil {
		if strings.EqualFold(currentChoice.Cmd, cmd) {
			return currentChoice.NextNode
		}
		currentChoice = currentChoice.Next
	}

	fmt.Println("Sorry, I didn't understand that")
	return node
}

func (node *StoryNode) play() {
	node.render()
	if node.Choice != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := NewNode("This is the beginning of your adventure, there are two choices in front of you")
	northDungeon := NewNode("Your are in the North section of the Dungeon, there are two more rooms in front of you, one is lit up the other is pitch black")
	darkRoom := NewNode("You entered a dark room, you can't see anything")
	darkRoomLit := NewNode("The dark room, is not as dark anymore. on the wall there is a text that says, take a number 2 to win!")
	toilette := NewNode("You took a big shit and clogged the toilette, the water first comes out and then sucks you in, you escaped")
	lightRoom := NewNode("You entered the light room, where an Ogre eats your brain and you die.")
	southDungeon := NewNode("Your are in the South section of the Dungeon")
	hole := NewNode("While walking you fall into a hole in the ground and you die.")
	theEnd := NewNode("Game over")
	youWin := NewNode("You Won!")

	start.AddChoice("N", "Go North", northDungeon)
	northDungeon.AddChoice("D", "Enter the dark room", darkRoom)
	darkRoom.AddChoice("O", "Turn your lamp on", darkRoomLit)
	darkRoomLit.AddChoice("S", "Take a shit in the toilette", toilette)
	toilette.AddChoice("E", "Press E to exit", youWin)

	darkRoomLit.AddChoice("B", "Go back and start again", start)

	northDungeon.AddChoice("L", "Enter the light room", lightRoom)
	lightRoom.AddChoice("E", "This is the end of your adventure press E to escape", theEnd)

	start.AddChoice("S", "Go South", southDungeon)
	southDungeon.AddChoice("K", "Keep walking until you get to the end of this corridor", hole)
	hole.AddChoice("E", "This is the end of your adventure press E to escape", theEnd)

	start.play()

}
