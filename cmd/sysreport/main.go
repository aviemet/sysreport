package main

import (
    "log"
    "time"

    "sysreport/internal/config"
    "sysreport/internal/network"
    "sysreport/internal/system"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    for {
        report, err := system.CollectSystemInfo()
        if err != nil {
            log.Printf("Error collecting system info: %v", err)
            continue
        }

        err = network.SendReport(cfg.ServerURL, report)
        if err != nil {
            log.Printf("Error sending report: %v", err)
        }

        time.Sleep(cfg.Interval)
    }
}
