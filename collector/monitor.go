package collector

import (
	"CollectLet/cache"
	"CollectLet/constants"
	"CollectLet/logger"
	"strconv"
	"time"
)

var cacheFactory *cache.LakeFactory

var computeLake cache.Lake[cache.ComputeCache]

var monitorLogTag = "[monitor]"

func init() {
	cacheFactory = cache.NewLakeFactory()
	computeLakeInterface, err := cacheFactory.GetObject(constants.Compute)
	if err != nil {
		logger.GetLogger().Error("%s Error getting cache Lake: %s", monitorLogTag, err.Error())
	}
	computeLake = computeLakeInterface.(cache.Lake[cache.ComputeCache])
}

func monitor() {
	for {
		computeLake.Add(cache.ComputeCache{
			DataItem: cache.NewDataItem(constants.CPUUsage, strconv.FormatFloat(GetCPUUsage(), 'f', 2, 64)),
		})
		time.Sleep(time.Duration(config.Compute.Freq) * time.Millisecond)
	}
}

func Start() {
	go monitor()
}
