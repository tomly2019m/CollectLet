package collector

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const statPath = "/proc/stat"

type ComputeCollector struct{}

func GetCpuUsage() (float64, error) {
	contents, err := os.ReadFile(statPath)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(contents), "\n")
	for _, l := range lines {
		fields := strings.Fields(l)
		if fields[0] == "cpu" {
			return calUsage(fields[1:])
		}
	}
	return 0, errors.New("cpu info not found")
}

func calUsage(fields []string) (float64, error) {
	var total uint64
	var idle uint64

	for i, field := range fields {
		val, err := strconv.ParseUint(field, 10, 64)
		if err != nil {
			return 0, err
		}
		total += val
		if i == 3 {
			idle = val
		}
	}
	return 100 * (float64(total-idle) / float64(total)), nil
}
