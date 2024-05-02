package cache

import (
	"context"
	"errors"
	"frame/modules/log"
	"frame/util/optional"
	"frame/util/types"

	"go.uber.org/zap"
	"time"
)

type Operation[T any] struct {
	d Drive
}

func (o *Operation[T]) SetEx(key string, val any, exp time.Duration) bool {
	payload, SerializeErr := types.Serialize(val)
	if SerializeErr != nil {
		log.Sugar().Warn(
			"cache",
			zap.String("mes", "cache drive SerializeErr key error"),
			zap.String("key", key),
			zap.Any("val", val),
		)
		return false
	}

	state, err := o.d.SetEx(context.Background(), key, payload, exp)
	if err != nil {
		log.Sugar().Info("cache drive set key error, key: %s, val: %v", key, payload)
		return false
	}

	return state != ""
}

func (o *Operation[T]) SetNx(key string, val interface{}, exp time.Duration) bool {
	payload, SerializeErr := types.Serialize(val)
	if SerializeErr != nil {
		log.Sugar().Warn(
			"cache",
			zap.String("mes", "cache drive SerializeErr key error"),
			zap.String("key", key),
			zap.Any("val", val),
		)
		return false
	}

	state, err := o.d.SetNx(context.Background(), key, payload, exp)
	if err != nil {
		log.Sugar().Infof("cache drive set key error, key: %s, val: %v", key, payload)
		return false
	}

	return state
}

func (o *Operation[T]) Get(key string) (optional.Option[T], error) {
	payload, err := o.d.Get(context.Background(), key)
	if err != nil {
		return optional.None[T](), errors.New("data is empty")
	}

	var currentVal optional.Option[any]
	if err := types.Deserialize([]byte(payload), &currentVal); err != nil {
		log.Sugar().Infof("cache not exist, key: %s, err: %v", key, err)
		return optional.None[T](), err
	}

	if !currentVal.Has() {
		return optional.None[T](), nil
	}

	return optional.Some(currentVal.Value().(T)), nil
}

func (o *Operation[T]) Exists(key string) bool {
	return o.d.Exists(context.Background(), key)
}

func (o *Operation[T]) Refresh(key string, exp time.Duration) bool {
	return o.d.Refresh(context.Background(), key, exp)
}

func (o *Operation[T]) Del(key string) bool {
	return o.d.Del(context.Background(), key)
}

func NewOperation[T any](d Drive) *Operation[T] {
	return &Operation[T]{d: d}
}
