package array

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T)  {
	
	t.Run("Collection of 5 numbers", func (t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
	
		if got != want {
			t.Errorf("got '%d' want '%d' numbers '%v'", got, want, numbers)
		}
	})

	t.Run("Collection of any size", func (t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		got := Sum(numbers)
		want := 28
	
		if got != want {
			t.Errorf("got '%d' want '%d' numbers '%v'", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3}, []int{2, 5})
	want := []int{4, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func (t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}

	t.Run("Make the sums of tails of", func (t *testing.T) {
		got := SumAllTails([]int{0, 2}, []int{1, 9})
		want := []int{2, 9}

		checkSum(t, got, want)
	})

	t.Run("safely sum empty slices", func (t *testing.T) {
		got := SumAllTails([]int{}, []int{11, 1, 12})
		want := []int{0, 13}

		checkSum(t, got, want)
	})
}