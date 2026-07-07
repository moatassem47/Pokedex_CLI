package main

import (
	"testing"

)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input: "hello world",
		expected: []string{
			"hello",
			"world",
		},
	},
	}
	for _,Cases:= range cases{
		actual:= cleanInput(Cases.input)
		if len(actual)!= len(Cases.expected){
			t.Errorf("lengths aren't equal %v,%v:",len(actual),len(Cases.expected))
			continue
		}
		for i:= range actual{
			actualWord:= actual[i]
			expectedWord:= Cases.expected[i]
			if actualWord!=expectedWord{
				t.Errorf("words aren't the same:%v vs %v",actualWord,expectedWord)
			}
		}
	}
}

