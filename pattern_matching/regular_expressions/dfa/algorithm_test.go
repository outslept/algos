package dfa

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
)

// Test states and symbols for even-zeroes language DFA:
// Accepts strings with even number of zeros
var (
    s0 = State("S0") // Even zeros seen (accepting)
    s1 = State("S1") // Odd zeroes seen (non-accepting)

    s2 = State("S2")
    
    zero = Symbol('0')
    one  = Symbol('1')
)

// createTestDFA constructs DFA for even-zeros language:
// - States: {s0, s1}
// - Alphabet: {0, 1}
// - Initial: s0
// - Accepting: {s0}
// - Transitions preserve even/odd parity of zeros
func createTestDFA(t *testing.T) *DFA {
    states := []State{s0, s1}
    alphabet := []Symbol{zero, one}
    acceptStates := []State{s0}
    
    transitions := func(s State, symbol Symbol) State {
        switch s {
        case s0:
            if symbol == zero {
                return s1
            }
            return s0
        case s1:
            if symbol == zero {
                return s0
            }
            return s1
        default:
            t.Fatalf("Invalid state: %v", s)
            return ""
        }
    }

    dfa, err := NewDFA(states, alphabet, transitions, s0, acceptStates)
    assert.NoError(t, err)
    return dfa
}

// TestDFA_Creation verifies DFA constructor validation:
// - Accepts valid configurations
// - Rejects empty states/alphabet
// - Rejects invalid initial/accept states
func TestDFA_Creation(t *testing.T) {
    tests := []struct {
        name          string
        states        []State
        alphabet      []Symbol
        transitions   TransitionFunc
        initialState  State
        acceptStates  []State
        expectedError bool
    }{
        {
            name:          "Valid DFA",
            states:        []State{s0, s1},
            alphabet:      []Symbol{zero, one},
            transitions:   func(s State, sym Symbol) State { return s0 },
            initialState:  s0,
            acceptStates:  []State{s0},
            expectedError: false,
        },
        {
            name:          "Empty States",
            states:        []State{},
            alphabet:      []Symbol{zero, one},
            transitions:   func(s State, sym Symbol) State { return s0 },
            initialState:  s0,
            acceptStates:  []State{s0},
            expectedError: true,
        },
        {
            name:          "Invalid Initial State",
            states:        []State{s0, s1},
            alphabet:      []Symbol{zero, one},
            transitions:   func(s State, sym Symbol) State { return s0 },
            initialState:  s2,
            acceptStates:  []State{s0},
            expectedError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            dfa, err := NewDFA(
                tt.states,
                tt.alphabet,
                tt.transitions,
                tt.initialState,
                tt.acceptStates,
            )

            if tt.expectedError {
                assert.Error(t, err)
                assert.Nil(t, dfa)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, dfa)
            }
        })
    }
}

// TestDFA_Accept verifies language recognition:
// - Empty string (even zeros, accept)
// - Single zero (odd, reject)
// - Two zeros (even, accept)
// - Only ones (even, accept)
// - Mixed input validation
func TestDFA_Accept(t *testing.T) {
    dfa := createTestDFA(t)

    tests := []struct {
        name           string
        input         string
        shouldAccept  bool
        expectedError bool
    }{
        {
            name:          "Empty string",
            input:         "",
            shouldAccept:  true,
            expectedError: false,
        },
        {
            name:          "Single zero",
            input:         "0",
            shouldAccept:  false,
            expectedError: false,
        },
        {
            name:          "Two zeros",
            input:         "00",
			shouldAccept:  true,
            expectedError: false,
        },
        {
            name:          "Only ones",
            input:         "111",
            shouldAccept:  true,
            expectedError: false,
        },
        {
            name:          "Complex valid input",
            input:         "1011001",
            shouldAccept:  false,
            expectedError: false,
        },
        {
            name:          "Invalid symbol",
            input:         "10x01",
            shouldAccept:  false,
            expectedError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            accepted, err := dfa.Accept(tt.input)

            if tt.expectedError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.shouldAccept, accepted)
            }
        })
    }
}

func TestDFA_Run(t *testing.T) {
    dfa := createTestDFA(t)

    tests := []struct {
        name           string
        input         string
        expectedStates []State
        expectedError  bool
    }{
        {
            name:           "Empty string",
            input:         "",
            expectedStates: []State{s0},
            expectedError:  false,
        },
        {
            name:           "Single transition",
            input:         "0",
            expectedStates: []State{s0, s1},
            expectedError:  false,
        },
        {
            name:           "Multiple transitions",
            input:         "001",
            expectedStates: []State{s0, s1, s0, s0},
            expectedError:  false,
        },
        {
            name:           "Invalid input",
            input:         "0x1",
            expectedStates: nil,
            expectedError:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            states, err := dfa.Run(tt.input)

            if tt.expectedError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.expectedStates, states)
            }
        })
    }
}

// TestDFA_Complement verifies language complement:
// Ensures complemented DFA accepts exactly the strings
// rejected by original and vice versa
func TestDFA_Complement(t *testing.T) {
    dfa := createTestDFA(t)
    complementDFA := dfa.Complement()

    tests := []struct {
        name          string
        input         string
        originalAccept bool
    }{
        {
            name:          "Empty string",
            input:         "",
            originalAccept: true,
        },
        {
            name:          "Single zero",
            input:         "0",
            originalAccept: false,
        },
        {
            name:          "Two zeros",
            input:         "00",
            originalAccept: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            originalResult, err := dfa.Accept(tt.input)
            assert.NoError(t, err)
            assert.Equal(t, tt.originalAccept, originalResult)

            complementResult, err := complementDFA.Accept(tt.input)
            assert.NoError(t, err)
            assert.Equal(t, !tt.originalAccept, complementResult)
        })
    }
}

func TestDFA_Validation(t *testing.T) {
    dfa := createTestDFA(t)

    t.Run("Valid transitions", func(t *testing.T) {
        err := dfa.ValidateTransitions()
        assert.NoError(t, err)
    })

    t.Run("Get states", func(t *testing.T) {
        states := dfa.GetStates()
        assert.Len(t, states, 2)
        assert.Contains(t, states, s0)
        assert.Contains(t, states, s1)
    })

    t.Run("Get alphabet", func(t *testing.T) {
        alphabet := dfa.GetAlphabet()
        assert.Len(t, alphabet, 2)
        assert.Contains(t, alphabet, zero)
        assert.Contains(t, alphabet, one)
    })

    t.Run("Get accept states", func(t *testing.T) {
        acceptStates := dfa.GetAcceptStates()
        assert.Len(t, acceptStates, 1)
        assert.Contains(t, acceptStates, s0)
    })
}

func TestDFA_Clone(t *testing.T) {
    original := createTestDFA(t)
    clone := original.Clone()

    assert.Equal(t, original.states, clone.states)
    assert.Equal(t, original.alphabet, clone.alphabet)
    assert.Equal(t, original.acceptStates, clone.acceptStates)
    assert.Equal(t, original.initialState, clone.initialState)

    testInputs := []string{"", "0", "1", "00", "01", "10", "11"}
    for _, input := range testInputs {
        originalAccept, err1 := original.Accept(input)
        cloneAccept, err2 := clone.Accept(input)

        assert.Equal(t, err1, err2)
        assert.Equal(t, originalAccept, cloneAccept)
    }
}

// TestDFA_EdgeCases verifies robustness:
// - Invalid transition functions
// - Empty configurations
// - Non-existent states
func TestDFA_EdgeCases(t *testing.T) {
    t.Run("Invalid DFA creation", func(t *testing.T) {
        invalidTransitions := func(s State, sym Symbol) State {
            return "non_existent_state"
        }

        dfa, err := NewDFA(
            []State{s0, s1},
            []Symbol{zero, one},
            invalidTransitions,
            s0,
            []State{s0},
        )

        assert.NoError(t, err) 
        err = dfa.ValidateTransitions()
        assert.Error(t, err) 
    })

    t.Run("Empty input validation", func(t *testing.T) {
        _, err := NewDFA(
            []State{},
            []Symbol{zero, one},
            func(s State, sym Symbol) State { return s0 },
            s0,
            []State{s0},
        )
        assert.Error(t, err)
    })

    t.Run("Invalid accept state", func(t *testing.T) {
        _, err := NewDFA(
            []State{s0, s1},
            []Symbol{zero, one},
            func(s State, sym Symbol) State { return s0 },
            s0,
            []State{s2}, 
        )
        assert.Error(t, err)
    })
}

// TestDFA_LongInputs verifies correct handling of
// extended input sequences, testing DFA stability
// on repeated patterns
func TestDFA_LongInputs(t *testing.T) {
    dfa := createTestDFA(t)
    
	longInput := strings.Repeat("10", 500)
    accepted, err := dfa.Accept(longInput)
    assert.NoError(t, err)
    assert.True(t, accepted)

    longInput = "0" + longInput
    accepted, err = dfa.Accept(longInput)
    assert.NoError(t, err)
    assert.False(t, accepted)
}

func TestDFA_IsAcceptState(t *testing.T) {
    dfa := createTestDFA(t)

    tests := []struct {
        name  string
        state State
        expected bool
    }{
        {
            name:     "Accept state",
            state:    s0,
            expected: true,
        },
        {
            name:     "Non-accept state",
            state:    s1,
            expected: false,
        },
        {
            name:     "Non-existent state",
            state:    s2,
            expected: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := dfa.IsAcceptState(tt.state)
            assert.Equal(t, tt.expected, result)
        })
    }
}
