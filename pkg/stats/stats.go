package stats

import (
	"github.com/rgsgit/bank/v2/pkg/types"
)

//Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	cnt := types.Money(0)
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
			cnt++
		}
	}
	return sum / cnt
}

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}

	return sum

}


