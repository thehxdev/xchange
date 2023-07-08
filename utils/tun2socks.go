package utils


import (
    "os"
    "log"
    "path/filepath"
)


func InstallTun2Socks(dest string) error {
    url := "https://github.com/xjasonlyu/tun2socks/releases/latest/download/tun2socks-linux-amd64.zip"
    zipFile := PathExpandUser("~/tun2socks-linux-amd64.zip")

    log.Printf("Downloading tun2socks-linux-amd64.zip to %s", zipFile)
    err := DownloadFile(zipFile, url)
    if err != nil {
        return err
    }
    defer os.Remove(zipFile)

    installDir := PathExpandUser(dest)
    _ = MakeDir(installDir)

    log.Printf("Extracting %s to %s", zipFile, installDir)
    err = ExtractZip(zipFile, installDir)
    if err != nil {
        return err
    }

    originalName := filepath.Join(installDir, "tun2socks-linux-amd64")
    newName := filepath.Join(installDir, "tun2socks")

    // remove old tun2socks
    if _, err := os.Stat(newName); err == nil {
        log.Println("Remove Old tun2socks executable")
        os.Remove(newName)
    }

    os.Rename(originalName, newName)
    os.Chmod(newName, 0777)
    log.Printf("tun2socks installed to %s", newName)
    return nil
}

