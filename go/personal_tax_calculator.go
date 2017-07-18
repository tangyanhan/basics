package main

import "fmt"

type TaxRateItem struct{
	start float64  // 起征点
	rate float64  // 税率
	sub float64  // 速算扣除数
}

func main(){
	salary := []float64{3000, 4000, 8000, 12000, 20000, 40000, 80000, 90000, 5000000}
	taxList :=[]TaxRateItem{
		TaxRateItem{-99999, 0.0, 0},  // 凑数用，用于处理低于3500的情况
		TaxRateItem{start: 0, rate: 0.03, sub: 0},
		TaxRateItem{start: 1500, rate:0.1, sub: 105},
		TaxRateItem{4500, 0.2, 555},
		TaxRateItem{9000, 0.25, 1005},
		TaxRateItem{35000,0.3, 2755},
		TaxRateItem{55000, 0.35, 5505},
		TaxRateItem{80000, 0.45, 13505},
	}

	const start = 3500.0
	for i, s := range salary{
		originalSalary := s
		s -= start  // 减去起征点
		iSuitableItem := 0
		for iTax, taxItem := range taxList{
			// Similar to Perl, if statement should always wrapped with curly brackets
			if s >= taxItem.start{
				iSuitableItem = iTax
			}else{
				break
			}
		}

		suitableItem := taxList[iSuitableItem]
		tax := suitableItem.rate * s - suitableItem.sub
		remainSalary := originalSalary - tax
		fmt.Println("#", i, "Salary=", originalSalary, "Tax=", tax, "Remain=", remainSalary, "SuitableItem:", suitableItem)
	}
}