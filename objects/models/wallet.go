package models

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
	"xorm.io/xorm"
)

type Wallet struct {
	AbstractTimeModel `xorm:"extends"`
	UID               int `xorm:"bigint pk 'uid'" json:"uid"`
	SGD               int `xorm:"bigint 'sgd'" json:"sgd"`
}

type TransferHistory struct {
	AbstractTimeModel `xorm:"extends"`
	ID                int `xorm:"bigint pk autoincr 'id'" json:"id"`
	FromUID           int `xorm:"bigint 'from_uid'" json:"from_uid"`
	ToUID             int `xorm:"bigint 'to_uid'" json:"to_uid"`
	Amount            int `xorm:"bigint" json:"amount"`
}

type TransferReceipt struct {
	From *Wallet
	To   *Wallet
	Time time.Time
}

type IWalletManager interface {
	Get(uid int) (*Wallet, error)
	Create(uid, sgd int) (*Wallet, error)
	Transfer(from, to, amount int) (*TransferReceipt, error)
}

type WalletManager struct {
	db *xorm.Engine
}

var (
	InsufficientBalanceError = errors.New("insufficient balance")
)

func NewWalletManager(db *xorm.Engine) *WalletManager {
	return &WalletManager{db: db}
}

func (m *WalletManager) Get(uid int) (*Wallet, error) {
	if wallets, err := m.getMany([]int{uid}); err != nil {
		return nil, err
	} else {
		return wallets[uid], nil
	}
}

func (m *WalletManager) getMany(uids []int) (map[int]*Wallet, error) {
	var wallets []*Wallet
	if err := m.db.In("uid", uids).Find(&wallets); err != nil {
		return nil, errors.Wrap(err, "error getting wallet by ID from DB")
	}
	result := make(map[int]*Wallet, len(wallets))
	for _, w := range wallets {
		result[w.UID] = w
	}
	return result, nil
}

func (m *WalletManager) Create(uid, sgd int) (*Wallet, error) {
	w := &Wallet{
		UID: uid,
		SGD: sgd,
	}
	if _, err := m.db.InsertOne(w); err != nil {
		return nil, err
	}
	return w, nil
}

func (m *WalletManager) Transfer(from, to, amount int) (*TransferReceipt, error) {
	session := m.db.NewSession().ForUpdate()
	if err := session.Begin(); err != nil {
		return nil, errors.Wrap(err, "error begin transfer session")
	}
	defer session.Close()

	// processing from's wallet
	fromWallet := &Wallet{UID: from}
	if _, err := session.Get(fromWallet); err != nil {
		return nil, err
	} else {
		if fromWallet.SGD < amount {
			return nil, InsufficientBalanceError
		}
		if _, err := session.ID(from).Incr("sgd", -amount).Update(new(Wallet)); err != nil {
			return nil, errors.Wrap(err, "error deducing from's sgd")
		}
		fromWallet.SGD -= amount
	}

	// processing to's wallet
	toWallet := &Wallet{UID: to}
	if exist, err := session.Get(toWallet); err != nil {
		return nil, err
	} else if !exist {
		toWallet.SGD = amount
		if _, err := session.InsertOne(toWallet); err != nil {
			return nil, errors.Wrap(err, "error creating to's wallet")
		}
	} else {
		if _, err := session.ID(to).Incr("sgd", amount).Update(new(Wallet)); err != nil {
			return nil, errors.Wrap(err, "error increase to's sgd")
		}
		toWallet.SGD += amount
	}

	// creating history
	history := &TransferHistory{
		FromUID: from,
		ToUID:   to,
		Amount:  amount,
	}
	if _, err := session.InsertOne(history); err != nil {
		return nil, errors.Wrap(err, "error creating transfer history")
	}

	if err := session.Commit(); err != nil {
		return nil, errors.Wrap(err, "error committing transfer transaction")
	}

	return &TransferReceipt{
		From: fromWallet,
		To:   toWallet,
		Time: history.AddTime,
	}, nil
}

type ITransferHistoryManager interface {
	Get(uid, start, length int) ([]*TransferHistory, error)
}

type TransferHistoryManager struct {
	db *xorm.Engine
}

func NewTransferHistoryManager(db *xorm.Engine) *TransferHistoryManager {
	return &TransferHistoryManager{db: db}
}

func (m *TransferHistoryManager) Get(uid, start, length int) ([]*TransferHistory, error) {
	var histories []*TransferHistory
	query := m.db.Where("from_uid = ?", uid)
	if start != 0 {
		query = query.And("id < ?", start)
	}
	if err := query.Desc("id").Limit(length).Find(&histories); err != nil {
		return nil, errors.Wrap(err, "error getting transfer histories from DB")
	}
	return histories, nil
}

type WalletAnalysisManager struct {
	db *xorm.Engine
}

type TransactionSummary struct {
	Count    int `json:"count"`
	TotalSGD int `json:"total_sgd"`
}

type WalletSummary struct {
	TotalUser int `json:"total_user"`
	TotalSGD  int `xorm:"'total_sgd'" json:"total_sgd"`
}

func NewWalletAnalysisManager(db *xorm.Engine) *WalletAnalysisManager {
	return &WalletAnalysisManager{db: db}
}

func (m *WalletAnalysisManager) GetWalletSummary() (*WalletSummary, error) {
	sql := `select count(1) as total_user, sum(sgd) as total_sgd from wallet`
	summary := new(WalletSummary)
	if _, err := m.db.SQL(sql).Get(summary); err != nil {
		return nil, errors.Wrap(err, "error getting total cash from DB")
	} else {
		return summary, nil
	}
}

func (m *WalletAnalysisManager) GetTransactionSummary(start, end time.Time) (*TransactionSummary, error) {
	sql := fmt.Sprintf(`select count(1), sum(amount) from transfer_history where add_time > '%s' and add_time <= '%s'`, start.Format("2006-01-02"), end.Format("2006-01-02"))
	summary := new(TransactionSummary)
	if _, err := m.db.SQL(sql).Get(summary); err != nil {
		return nil, err
	}
	return summary, nil
}

func (m *WalletAnalysisManager) GetAllWallets() ([]*Wallet, error) {
	var wallets []*Wallet
	if err := m.db.Asc("uid").Find(&wallets); err != nil {
		return nil, err
	}
	return wallets, nil
}
