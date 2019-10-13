package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T)  {
	got := Repeat("g", 5)
	want := "ggggg"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestLetterEqual(t *testing.T) {
	got := LetterEqual("Gabriel Sabadini Facina", "a")
	want := 5

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestLowerCase(t *testing.T) {
	got := LowerCase("GABRIEL SABADINI FACINA")
	want := "gabriel sabadini facina"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("g", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}

func ExampleTestLetterEqual() {
	result := LetterEqual("Jenkins Continuous Integration/Continuous Delivery", "i")
	fmt.Println(result)
	// Output: 6
}