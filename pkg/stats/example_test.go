package stats
import (
	"github.com/rgsgit/bank/pkg/types"
	"fmt"
)

func ExampleAvg() {

	payments := []types.Payment{
		{
			ID:     0,
			Amount: 10,
		},
		{
			ID:     1,
			Amount: 20,
		},
		{
			ID:     2,
			Amount: 20,
		},
		{
			ID:     3,
			Amount: 20,
		},
		{
			ID:     4,
			Amount: 30,
		},
	}
	result := Avg(payments)
	fmt.Println(result)

	//Output:
	//20
}

func ExampleTotalInCategory() {

	payments := []types.Payment{
		{
			ID:     0,
			Amount: 10,
			Category: "apteka",
		},
		{
			ID:     1,
			Amount: 20,
			Category: "apteka",
		},
		{
			ID:     2,
			Amount: 20,
			Category: "retauran",
		},
		{
			ID:     3,
			Amount: 20,
			Category: "auto",
		},
		{
			ID:     4,
			Amount: 30,
			Category: "apteka",
		},
	}
	result := TotalInCategory(payments, types.Category("apteka"))
	fmt.Println(result)

	//Output:
	//60
}