package stats

import(
	"github.com/sulton1987/bank/pkg/bank/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	avg := types.Money(0)
	for _, payment := range payments {
		avg += payment.Amount;
	}
	avg = avg/types.Money(len(payments))
	return avg
}