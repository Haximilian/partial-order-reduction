package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Vertex represents a state transition
type Vertex struct {
	Identifier  int               `json:"identifier"`
	Transitions []string          `json:"transitions"`
	Hash        int               `json:"stateHash"`
	Children    map[string]Vertex `json:"children"`
}

// Transition is a wrapper over Vertex that maintain a back tracking set
type StateWrapper struct {
	State        Vertex
	BacktrackSet []Vertex
}

func main() {
	f := "schedule.json"
	graphBinary, _ := ioutil.ReadFile(f)

	var root Vertex = Vertex{}
	json.Unmarshal(graphBinary, &root)

	fmt.Println(root)
}

func Mazurkiewicz() {
	var stack []StateWrapper

	// whenever a new state is reached discover is called

}

func MazurkiewiczDiscover(stack []StateWrapper) {
	currentState := stack[len(stack)]
	for nextTransition := range currentState.State.Children {
		for i := len(stack) - 1; i >= 0; i-- {
			length := len(stack[i].State.Transitions)
			previousTransition := stack[i].State.Transitions[length-1]
			if !stack[i].State.independant(nextTransition, previousTransition) {

			}
			// co enabled
		}
	}
}

func (v Vertex) independant(a, b string) bool {
	result := true

	// if a is enabled in v
	_, aInCurrent := v.Children[a]
	if aInCurrent {
		// either b is enabled in v and v->[a] or b is enabled in neither
		_, bInCurrent := v.Children[b]
		_, bInForward := v.Children[a].Children[b]

		result = (!bInCurrent && !bInForward)

		// if b is enabled in both, the application of a, b and b, a
		// results in the same state
		if bInCurrent && bInForward {
			result = v.Children[a].Children[b].Hash == v.Children[b].Children[a].Hash
		}
	}

	return result
}
