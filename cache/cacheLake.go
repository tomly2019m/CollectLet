package cache

import (
	"CollectLet/util"
	"fmt"
	"sync"
)

type Lake[T any] struct {
	cacheType  string
	cacheQueue *util.Queue[T]
}

func (l *Lake[T]) Add(item T) {
	l.cacheQueue.Push(item)
}

func (l *Lake[T]) Get() (T, error) {
	pop, err := l.cacheQueue.Pop()
	if err != nil {
		var zero T
		return zero, err
	}
	return pop, nil
}

type DataItem struct {
	Name      string
	TimeStamp int64
	Value     string
}

type ComputeCache struct {
	DataItem
}

type StorageCache struct {
	DataItem
}

type NetworkCache struct {
	DataItem
}

// LakeFactory 是创建和管理 Lake 对象的工厂
type LakeFactory struct {
	lakes map[string]interface{}
	mu    sync.Mutex
}

// NewLakeFactory 创建一个新的 LakeFactory
func NewLakeFactory() *LakeFactory {
	return &LakeFactory{
		lakes: make(map[string]interface{}),
	}
}

// GetObject 根据参数返回对应的 Lake 对象，如果不存在则创建一个新的
func (f *LakeFactory) GetObject(param string) (interface{}, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	// 检查对象是否已存在
	if lake, exists := f.lakes[param]; exists {
		return lake, nil
	}

	// 创建新的 Lake 对象
	var newLake interface{}
	switch param {
	case "compute":
		newLake = &Lake[ComputeCache]{
			cacheType:  "compute",
			cacheQueue: util.NewQueue[ComputeCache](),
		}
	case "storage":
		newLake = &Lake[StorageCache]{
			cacheType:  "storage",
			cacheQueue: util.NewQueue[StorageCache](),
		}
	case "network":
		newLake = &Lake[NetworkCache]{
			cacheType:  "network",
			cacheQueue: util.NewQueue[NetworkCache](),
		}
	default:
		return nil, fmt.Errorf("unknown lake type: %s", param)
	}

	// 存储新创建的对象
	f.lakes[param] = newLake
	return newLake, nil
}
