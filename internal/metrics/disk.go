package metrics

import "github.com/shirou/gopsutil/disk"

func GetDiskUsage() (float64, error) {
	usageStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return usageStat.UsedPercent, nil
}
