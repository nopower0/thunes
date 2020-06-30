package http

import (
	"os"
	"testing"
	"thunes/components"
	"thunes/objects/business"
	"thunes/objects/models"
	"thunes/settings"
	"thunes/tools"
	"time"
)

const (
	TestTokenLogin    = "___a_login_token___"
	TestTokenNotLogin = "___a_not_login_token___"
	TestUID           = 10
	TestUsername      = "___test_username___"
	TestPassword      = "___test_password___"
	TestSGD           = 233
	TestTransferToUID = 15
)

type MockAuthService struct{}

func (*MockAuthService) GetTokenInfo(token string) (*business.TokenInfo, error) {
	var info *business.TokenInfo
	if token == TestTokenNotLogin {
		info = &business.TokenInfo{}
	} else if token == TestTokenLogin {
		info = &business.TokenInfo{
			Username: TestUsername,
			UID:      TestUID,
		}
	}
	return info, nil
}
func (*MockAuthService) CreateTokenInfo(info *business.TokenInfo, expireAt time.Time) (string, error) {
	if info.UID != 0 && len(info.Username) != 0 {
		return TestTokenLogin, nil
	} else {
		return TestTokenNotLogin, nil
	}
}
func (*MockAuthService) UpdateTokenInfo(token string, info *business.TokenInfo, expireAt time.Time) error {
	return nil
}

type MockUserManager struct{}

func (m *MockUserManager) Get(uid int) (*models.User, error) {
	if users, err := m.GetMany([]int{uid}); err != nil {
		return nil, err
	} else {
		return users[uid], nil
	}
}
func (*MockUserManager) GetMany(uids []int) (map[int]*models.User, error) {
	users := make(map[int]*models.User)
	for _, id := range uids {
		switch id {
		case TestUID:
			users[id] = &models.User{
				UID:      TestUID,
				Username: TestUsername,
				Password: TestPassword,
			}
		case TestTransferToUID:
			users[id] = &models.User{
				UID:      TestTransferToUID,
				Username: "",
				Password: "",
			}
		}
	}
	return users, nil
}
func (*MockUserManager) GetByCredential(username, password string) (*models.User, error) {
	if username == TestUsername && password == tools.PasswordHash(TestPassword) {
		return &models.User{
			UID:      TestUID,
			Username: TestUsername,
			Password: TestPassword,
		}, nil
	}
	return nil, nil
}
func (*MockUserManager) Create(username, password string) (*models.User, error) {
	return &models.User{
		UID:      0,
		Username: username,
		Password: password,
	}, nil
}

type MockWalletManager struct{}

func (*MockWalletManager) Get(uid int) (*models.Wallet, error) {
	if uid == TestUID {
		return &models.Wallet{
			UID: TestUID,
			SGD: TestSGD,
		}, nil
	}
	return nil, nil
}
func (*MockWalletManager) Create(uid, sgd int) (*models.Wallet, error) {
	return &models.Wallet{
		UID: uid,
		SGD: sgd,
	}, nil
}
func (*MockWalletManager) Transfer(from, to, amount int) (*models.TransferReceipt, error) {
	return &models.TransferReceipt{
		From: &models.Wallet{
			UID: TestUID,
			SGD: TestSGD - amount,
		},
		To: &models.Wallet{
			UID: TestTransferToUID,
			SGD: amount,
		},
		Time: time.Now(),
	}, nil
}

type MockTransactionHistoryManager struct{}

func (*MockTransactionHistoryManager) Get(uid, start, length int) ([]*models.TransferHistory, error) {
	return []*models.TransferHistory{
		{
			ID:      0,
			FromUID: TestUID,
			ToUID:   TestTransferToUID,
			Amount:  10,
		},
		{
			ID:      1,
			FromUID: TestTransferToUID,
			ToUID:   TestUID,
			Amount:  15,
		},
	}, nil
}

func TestMain(m *testing.M) {
	_ = os.Setenv("CONF", "conf/unit_test.json")
	settings.Init()

	components.DefaultAuthService = new(MockAuthService)

	models.DefaultUserManager = new(MockUserManager)
	models.DefaultWalletManager = new(MockWalletManager)
	models.DefaultTransferHistoryManager = new(MockTransactionHistoryManager)

	os.Exit(m.Run())
}
