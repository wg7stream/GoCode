package main

import (
	"fmt"
)

// 结构体Account
type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

// 方法
// 1. 存款
func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("输入的密码错误")
		return
	}

	if money <= 0 {
		fmt.Println("输入的金额错误")
		return
	}

	account.Balance += money
	fmt.Println("存款成功...")
}

// 取款
func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("输入密码错误")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("输入金额错误")
		return 
	}

	account.Balance -= money
	fmt.Println("取款成功...")
}


// 查询余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}

	fmt.Printf("账号为：%v 余额 = %v \n", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo : "zhangchengyu",
		Pwd : "666666",
		Balance : 100.23,
	}

	account.Query("666666")
	account.Deposite(200.3, "666666")
	account.Query("666666")


}