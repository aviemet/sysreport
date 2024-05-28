package main

import (
    "fmt"
    "net"
    "time"

    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/host"
    "github.com/shirou/gopsutil/v3/mem"
)

func getIP() (string, error) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }

    for _, iface := range interfaces {
        addrs, err := iface.Addrs()
        if err != nil {
            return "", err
        }
        for _, addr := range addrs {
            var ip net.IP
            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }
            if ip != nil && ip.IsGlobalUnicast() {
                return ip.String(), nil
            }
        }
    }
    return "", fmt.Errorf("no global unicast address found")
}

func main() {
    for {
        ip, err := getIP()
        if err != nil {
            fmt.Println("Error getting IP:", err)
            continue
        }

        cpus, err := cpu.Percent(0, false)
        if err != nil {
            fmt.Println("Error getting CPU usage:", err)
            continue
        }

        memStat, err := mem.VirtualMemory()
        if err != nil {
            fmt.Println("Error getting memory stats:", err)
            continue
        }

        hostInfo, err := host.Info()
        if err != nil {
            fmt.Println("Error getting host info:", err)
            continue
        }

        report := fmt.Sprintf(
            "IP: %s, CPU Usage: %.2f%%, Memory Usage: %.2f%%, Serial Number: %s",
            ip, cpus[0], memStat.UsedPercent, hostInfo.HostID,
        )
        fmt.Println(report)

        // Here you would send the report to your server
        // For example, using an HTTP POST request

        // Sleep for a specified interval before repeating
        time.Sleep(60 * time.Second)
    }
}
