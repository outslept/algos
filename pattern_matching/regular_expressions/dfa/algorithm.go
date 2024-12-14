package dfa

import (
	"errors"
	"fmt"
)

// State represents a state in the DFA
type State string

// Symbol represents an input symbol
type Symbol rune

// TransitionFunc represents the transition function δ
type TransitionFunc func(State, Symbol) State

// DFA implements a Deterministic Finite Automaton based on the formal definition:
// M = (Q, Σ, δ, q0, F) where:
// - Q is a finite set of states
// - Σ is a finite input alphabet
// - δ: Q × Σ → Q is the transition function
// - q0 ∈ Q is the initial state
// - F ⊆ Q is the set of accept states
type DFA struct {
	states       map[State]bool  // Q: Set of states
	alphabet     map[Symbol]bool // Σ: Input alphabet
	transitions  TransitionFunc  // δ: Transition function
	initialState State           // q0: Initial state
	acceptStates map[State]bool  // F: Set of accept states
}

// NewDFA creates and validates a new DFA instance according to formal definition.
// Returns error if the automaton structure voilates DFA properties:
// - States and alphabet must be non-empty
// - Initial state must exist in states set
// - Accecpt states must be subset of states set
func NewDFA(
	states []State,
	alphabet []Symbol,
	transitions TransitionFunc,
	initialState State,
	acceptStates []State,
) (*DFA, error) {
	if len(states) == 0 || len(alphabet) == 0 {
		return nil, errors.New("states and alphabet cannot be empty")
	}

	dfa := &DFA{
		states:       make(map[State]bool),
		alphabet:     make(map[Symbol]bool),
		transitions:  transitions,
		initialState: initialState,
		acceptStates: make(map[State]bool),
	}

	for _, state := range states {
		dfa.states[state] = true
	}

	for _, symbol := range alphabet {
		dfa.alphabet[symbol] = true
	}

	for _, state := range acceptStates {
		if !dfa.states[state] {
			return nil, fmt.Errorf("accept state %v not in states set", state)
		}
		dfa.acceptStates[state] = true
	}

	if !dfa.states[initialState] {
		return nil, fmt.Errorf("initial state %v not in states set", initialState)
	}

	return dfa, nil
}

// Accept implements the DFA acceptance algorithm:
// 1. Start from initial state
// 2. For each input symbol, follow transition
// 3. Accept if final state is accepting
// Returns errors if input contains invalid symbols
func (d *DFA) Accept(input string) (bool, error) {
	currentState := d.initialState

	for _, symbol := range input {
		if !d.alphabet[Symbol(symbol)] {
			return false, fmt.Errorf("symbol %c not in alphabet", symbol)
		}

		currentState = d.transitions(currentState, Symbol(symbol))

		if !d.states[currentState] {
			return false, fmt.Errorf("transition to non-existent state %v", currentState)
		}
	}

	return d.acceptStates[currentState], nil
}

// GetStates returns all states of the DFA
func (d *DFA) GetStates() []State {
	states := make([]State, 0, len(d.states))
	for state := range d.states {
		states = append(states, state)
	}
	return states
}

// GetAlphabet returns the alphabet of the DFA
func (d *DFA) GetAlphabet() []Symbol {
	symbols := make([]Symbol, 0, len(d.alphabet))
	for symbol := range d.alphabet {
		symbols = append(symbols, symbol)
	}
	return symbols
}

// GetAcceptStates returns all accept states of the DFA
func (d *DFA) GetAcceptStates() []State {
	states := make([]State, 0, len(d.acceptStates))
	for state := range d.acceptStates {
		states = append(states, state)
	}
	return states
}

// IsAcceptState checks if the given state is an accept state
func (d *DFA) IsAcceptState(state State) bool {
	return d.acceptStates[state]
}

// ValidateTransitions verifies that transition function is total:
// For each state and symbol combination, transition must lead to valid state.
func (d *DFA) ValidateTransitions() error {
	for state := range d.states {
		for symbol := range d.alphabet {
			nextState := d.transitions(state, symbol)
			if !d.states[nextState] {
				return fmt.Errorf("invalid transition: state %v with symbol %c leads to non-existent state %v",
					state, symbol, nextState)
			}
		}
	}
	return nil
}

// Clone creates a deep copy of the DFA
func (d *DFA) Clone() *DFA {
	newDFA := &DFA{
		states:       make(map[State]bool),
		alphabet:     make(map[Symbol]bool),
		transitions:  d.transitions,
		initialState: d.initialState,
		acceptStates: make(map[State]bool),
	}

	for state := range d.states {
		newDFA.states[state] = true
	}

	for symbol := range d.alphabet {
		newDFA.alphabet[symbol] = true
	}

	for state := range d.acceptStates {
		newDFA.acceptStates[state] = true
	}

	return newDFA
}

// Complement returns DFA recognizing complement language.
// Implemented by inverting accept states while preserving structure,
// based on DFA closure under complement operation.
func (d *DFA) Complement() *DFA {
	complementDFA := d.Clone()

	complementDFA.acceptStates = make(map[State]bool)
	for state := range d.states {
		if !d.acceptStates[state] {
			complementDFA.acceptStates[state] = true
		}
	}

	return complementDFA
}

// Run simulates DFA execution step by step, returning state sequence.
// Returns error on invalid input symbols or transitions.
func (d *DFA) Run(input string) ([]State, error) {
	states := make([]State, 0, len(input)+1)
	currentState := d.initialState
	states = append(states, currentState)

	for _, symbol := range input {
		if !d.alphabet[Symbol(symbol)] {
			return nil, fmt.Errorf("symbol %c not in alphabet", symbol)
		}

		currentState = d.transitions(currentState, Symbol(symbol))
		if !d.states[currentState] {
			return nil, fmt.Errorf("transition to non-existent state %v", currentState)
		}
		states = append(states, currentState)
	}

	return states, nil
}
