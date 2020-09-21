package stats

import (
	"fmt"

	"github.com/sulton1987/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 5000_00,
		},
		{
			Amount: 0,
		},
	}
	fmt.Println(Avg(payments))
	//Output:
	//250000
}
