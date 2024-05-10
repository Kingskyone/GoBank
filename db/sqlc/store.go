package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

// Store 允许查询以事务的形式而不是单个SQL
type Store struct {
	*Queries // 组合
	db       *pgx.Conn
}

// 创建一个Store
func NewStore(db *pgx.Conn) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// 创建一个执行事务的函数，根据是否出错进行提交或者回滚  保证原子性
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

// TransferTxParams 两个账号转账需要的参数
type TransferTxParams struct {
	FromAccountID int64   `json:"from_account_id"`
	ToAccountID   int64   `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

// TransferTxResult 转账结果返回
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"fromAccount"`
	ToAccount   Account  `json:"toAccount"`
	FromEntry   Ectry    `json:"fromEntry"`
	ToEntry     Ectry    `json:"toEntry"`
}

// 结构体，结合上下文传递当前go协程的信息
var txKey = struct{}{}

// TransferTx 实现转账操作
// A向B转账，按顺序：创建transfer记录、创建A与B的entry记录，更改A与B的account余额
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	// 调用execTx，传入CRUD操作组成函数
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)
		fmt.Println(txName, "create transfer")

		// 创建转账记录Transfer
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Money:         arg.Amount,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "create entry")
		// 转出账户创建Entry
		result.FromEntry, err = q.CreateEctry(ctx, CreateEctryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		// 转入账户创建Entry
		result.ToEntry, err = q.CreateEctry(ctx, CreateEctryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account")
		// 更新余额  获取->变更->更新   多线程条件下需要加锁
		// 使用addMoney函数，判断ID大小，优先更新ID小的来防止死锁
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, int64(-arg.Amount), arg.ToAccountID, int64(arg.Amount))
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, int64(arg.Amount), arg.FromAccountID, int64(-arg.Amount))
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err

}

// addMoney 对两个账户进行修改余额操作
func addMoney(ctx context.Context, q *Queries, accountID1 int64, amount1 int64, accountID2 int64, amount2 int64) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: float64(amount1),
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: float64(amount2),
	})
	return
}
