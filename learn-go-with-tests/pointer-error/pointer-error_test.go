package pointer_error

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	// 這個測試的目的是要：
	// 1. 使用者提款超過帳戶的數量時，不會進行操作。
	// 2. 回覆 error 給使用者
	t.Run("Withdraw insufficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}
func assertBalance (t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError (t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertNoError (t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

