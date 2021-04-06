package stats

import (
	"reflect"
	"testing"

	"github.com/rgsgit/bank/v2/pkg/types"
)

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100_00,
			Category: "auto",
			Status:   types.StatusFail,
		},
		{
			ID:       2,
			Amount:   5_00,
			Category: "mobile",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   10_00,
			Category: "auto",
			Status:   types.StatusFail,
		},
	}
	expected := map[types.Category]types.Money{
		"auto":   110_00,
		"mobile": 5_00,
	}
	result := CategoriesTotal(payments)
	if !reflect.DeepEqual(expected, result) {
		if len(result) != 0 {
			t.Errorf("Invalid Result, expected %v, actual %v", expected, result)
		}
	}
}
func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100_00,
			Category: "auto",
			Status:   types.StatusFail,
		},
		{
			ID:       2,
			Amount:   5_00,
			Category: "auto",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   10_00,
			Category: "mobile",
			Status:   types.StatusFail,
		},
		{
			ID:       4,
			Amount:   10_00,
			Category: "mobile",
			Status:   types.StatusInProgress,
		},
	}
	expected := map[types.Category]types.Money{
		"auto":   5250,
		"mobile": 10_00,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		if len(result) != 0 {
			t.Errorf("Invalid Result, expected %v, actual %v", expected, result)
		}
	}
}