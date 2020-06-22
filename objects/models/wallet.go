package models

import (
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
		return nil, err
	}
	result := make(map[int]*Wallet, len(wallets))
	for _, w := range wallets {
		result[w.UID] = w
	}
	return result, nil
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

	return &TransferReceipt{
		From: fromWallet,
		To:   toWallet,
		Time: history.AddTime,
	}, nil
}
