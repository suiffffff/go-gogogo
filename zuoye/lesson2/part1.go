package main

import "fmt"

type Produce struct {
	Name  string
	Price float64
	Stock int
}

func TotalValue(p *Produce) float64 {
	return float64(p.Price) * float64(p.Stock)
}
func IsInstock(p *Produce) bool {
	if p.Stock == 0 {
		return false
	}
	return true
}
func Info(p *Produce) string {
	return fmt.Sprintf("商品：%s,价格：%f,库存：%d", p.Name, p.Price, p.Stock)
}
func (p *Produce) Restock(amount int) {
	p.Stock += amount
}
func (p *Produce) Sell(amount int) (success bool, message string) {
	if p.Stock < amount {
		return false, "库存不足" // 库存不足时返回指定结果
	}
	p.Stock -= amount
	return true, fmt.Sprintf("售卖成功")
}

func main() {
	var gobook = Produce{"go编程书", 89.5, 10}
	success, msg := gobook.Sell(5)
	fmt.Println("售卖结果：", success, msg)
	gobook.Restock(20)
	fmt.Printf("%d\n", gobook.Stock)
	success, msg = gobook.Sell(30)
	fmt.Println("售卖结果：", success, msg)

}
