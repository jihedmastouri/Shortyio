package service

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func (s *Service) GetKV(key string) (string, error) {
	kv, _, err := s.api.KV().Get(key, nil)
	if err != nil {
		return "", err
	}

	if kv == nil {
		return "", fmt.Errorf("key %s not found", key)
	}

	return string(kv.Value), nil
}

func (s *Service) SetKV(key, value string) error {
	p := &api.KVPair{
		Key:   key,
		Value: []byte(value),
	}

	_, err := s.api.KV().Put(p, nil)
	if err != nil {
		return err
	}

	return nil
}
