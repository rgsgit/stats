package stats

import (
	"github.com/rgsgit/bank/v2/pkg/types"
)

//Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}
	sum = types.Money((int(sum)) / len(payments))
	return sum
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

//CategoriesAvg считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	categoriesCnt := map[types.Category]int{}
	for _, payment := range payments {
		categoriesCnt[payment.Category]++
		categories[payment.Category] += payment.Amount
	}
	for key := range categories {
		categories[key] /= types.Money(categoriesCnt[key])
	}
	return categories
}

//CategoriesTotal все суммы покупок во всех категориях
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}
	return categories
}

//PeriodDinamics сравнивает по категориям за два периода
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	for secondId, category := range second {
		firstPeriodElem, isFind := first[secondId]
		if isFind {
			result[secondId] = category - firstPeriodElem
		} else {
			result[secondId] = category
		}
	}
	for firstId, category := range first {
		secondPeriodElem, isFind := second[firstId]
		if !isFind {
			result[firstId] = secondPeriodElem - category
		}
	}
	return result
}
