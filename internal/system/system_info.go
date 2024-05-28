package system

import (
    "fmt"
    "runtime"

    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/host"
    "github.com/shirou/gopsutil/v3/mem"
    "sysreport/internal/network"
)

func CollectSystemInfo() (string, error) {
    ip, err := network.GetIP()
    if err != nil {
        return "", err
    }

    cpus, err := cpu.Percent(0, false)
    if err != nil {
        return "", err
    }

    memStat, err := mem.VirtualMemory()
    if err != nil {
        return "", err
    }

    hostInfo, err := host.Info()
    if err != nil {
        return "", err
    }

    report := fmt.Sprintf(
        "IP: %s, CPU Usage: %.2f%%, Memory Usage: %.2f%%, Serial Number: %s, OS: %s",
        ip, cpus[0], memStat.UsedPercent, hostInfo.HostID, runtime.GOOS,
    )
		
    return report, nil
}
