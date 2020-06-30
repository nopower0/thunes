package components

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"thunes/objects/business"
	"thunes/tools"
	"time"
)

type IAuthService interface {
	GetTokenInfo(token string) (*business.TokenInfo, error)
	CreateTokenInfo(info *business.TokenInfo, expireAt time.Time) (string, error)
	UpdateTokenInfo(token string, info *business.TokenInfo, expireAt time.Time) error
}

type AuthService struct {
	r *tools.RedisClient
}

const (
	tokenKey = "token:%s"
)

func NewAuthService(r *tools.RedisClient) *AuthService {
	return &AuthService{r: r}
}

func (s *AuthService) GetTokenInfo(token string) (*business.TokenInfo, error) {
	key := s.getKey(token)
	if content, err := s.r.RO.Get(key).Result(); err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, errors.Wrap(err, "error getting token content from redis")
	} else {
		info := new(business.TokenInfo)
		if err := json.Unmarshal([]byte(content), info); err != nil {
			return nil, errors.Wrap(err, "error unmarshalling token content")
		}
		return info, nil
	}
}

func (s *AuthService) CreateTokenInfo(info *business.TokenInfo, expireAt time.Time) (string, error) {
	token := uuid.NewV4().String()
	if err := s.setTokenInfo(token, info, expireAt); err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) setTokenInfo(token string, info *business.TokenInfo, expireAt time.Time) error {
	key := s.getKey(token)
	if content, err := json.Marshal(info); err != nil {
		return errors.Wrap(err, "error marshalling token info")
	} else {
		pipe := s.r.RW.Pipeline()
		pipe.Set(key, content, 0)
		pipe.ExpireAt(key, expireAt)
		if _, err := pipe.Exec(); err != nil {
			return errors.Wrap(err, "error creating token info")
		}
		return nil
	}
}

func (s *AuthService) UpdateTokenInfo(token string, info *business.TokenInfo, expireAt time.Time) error {
	return s.setTokenInfo(token, info, expireAt)
}

func (s *AuthService) getKey(token string) string {
	return fmt.Sprintf(tokenKey, token)
}
