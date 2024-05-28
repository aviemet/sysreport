package network

import (
    "bytes"
    "fmt"
    "net"
    "net/http"
)

func GetIP() (string, error) {
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

func SendReport(url, report string) error {
    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(report)))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("failed to send report, status code: %d", resp.StatusCode)
    }

    return nil
}
