package memory

import (
	"reflect"
	"testing"
)

func TestSimpleTokenize(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		want     []string
	}{
		{
			name:     "Test 1: Basic sentence",
			sentence: "The quick brown fox jumps over the lazy dog.",
			want:     []string{"quick", "brown", "fox", "jumps", "over", "lazy", "dog"},
		},
		{
			name:     "Test 2: Sentence with non-subject words",
			sentence: "A cat is sitting on the mat.",
			want:     []string{"cat", "sitting", "on", "mat"},
		},
		{
			name:     "Test 3: Sentence with no non-subject words",
			sentence: "John and Mary went to the market.",
			want:     []string{"John", "and", "Mary", "went", "to", "market"},
		},
		{
			name:     "Test 4: Empty sentence",
			sentence: "",
			want:     []string{},
		},
		{
			name:     "Test 5: Sentence with no subject",
			sentence: "is a test sentence.",
			want:     []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simpleTokenize(tt.sentence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simpleTokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
