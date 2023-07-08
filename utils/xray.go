package utils

import (
    "os"
    "fmt"
    "errors"
    "os/exec"
    "path/filepath"
)

var ConfigPool      = PathExpandUser("~/.local/xconfs")
var XrayConfigFile  = PathExpandUser("/etc/xray/config.json")


func ChangeConfigFile() error {
    _, err := os.Stat(XrayConfigFile)
    if err != nil {
        return err
    }

    _ = MakeDir(ConfigPool)

    config_files, err := readConfigsDirectory(ConfigPool)
    if err != nil {
        return err
    }
    var idx int

    fmt.Println("Local configs:")
    for i, entry := range config_files {
        if !entry.IsDir() {
            fmt.Println(" ", i, entry.Name())
        }
    }

    fmt.Print("\nChoose one config file: ")
    _, err = fmt.Scan(&idx)
    if err != nil {
        return err
    }

    err = checkIdx(idx, len(config_files))
    if err != nil {
        return err
    }

    chosenConfig := filepath.Join(ConfigPool, config_files[idx].Name())
    err = CopyFileContent(XrayConfigFile, chosenConfig)
    if err != nil {
        return err
    }


    cmd := exec.Command("systemctl", "restart", "xray")
    _, err = cmd.Output()
    if err != nil {
        return err
    }

    fmt.Println("Change Config to", chosenConfig)
    return nil
}


func checkIdx(idx, length int) error {
    if (idx < 0 || idx > length) {
        err := errors.New("Index out of range")
        return err
    }
    return nil
}

