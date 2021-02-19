package model

import (
	"fmt"
)

type account struct {
	accountNo string
	pwd string
	balance float64
}

func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度错误...")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码长度错误...")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额不对...")
		return nil
	}

	return &account{
		accountNo : accountNo,
		pwd : pwd,
		balance : balance,
	}
}

// 存款
func (account *account) Deposite(money float64, pwd string){
	if pwd != account.pwd {
		fmt.Println("输入的密码错误")
		return
	}

	if money <= 0 {
		fmt.Println("输入的金额不正确")
		return
	}

	account.balance += money
	fmt.Println("存款成功")
}

// 取款
func (account *account) WithDraw(money float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("输入的密码错误...")
		return
	}

	if money <= 0 || money > account.balance {
		fmt.Println("输入的金额错误...")
		return
	}

	account.balance -= money
	fmt.Println("取款成功...")
}

func (account *account) Query(pwd string) {
	if pwd != account.pwd {
		fmt.Println("输入的密码错误")
		return 
	}

	fmt.Printf("账号：%v 余额：%v \n", account.accountNo, account.balance)
}