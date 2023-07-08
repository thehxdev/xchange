package utils

import (
	"os"
	"io"
    "fmt"
    "errors"
	"os/user"
	"path/filepath"
)


func MakeDir(path string) (err error) {
    err = os.Mkdir(path, 0755)

    if err != nil {
        return err
    }

    return nil
}


func PathExpandUser(path string) string {
    usr, _ := user.Current()
    homeDir := usr.HomeDir
    var firstSlashIdx int

    for i := 0; i < len(path); i++ {
        if path[i] == '/' {
            firstSlashIdx = i
            break
        }
    }


    if path[0] == '~' {
        return (filepath.Join(homeDir, path[firstSlashIdx:]))
    }

    return path
}


func readConfigsDirectory(path string) ([]os.DirEntry, error) {
    var files []os.DirEntry

    dirData, err := os.ReadDir(ConfigPool)
    if err != nil {
        return nil, err
    }

    if len(dirData) == 0 {
        err = errors.New(fmt.Sprintf("Cannot found config files in %s", ConfigPool))
        return nil, err
    }

    for _, entry := range dirData {
        if !entry.IsDir() {
            files = append(files, entry)
        }
    }
    return files, nil
}


func CopyFileContent(dest, src string) (err error) {
    srcFile, err := os.Open(src)
    if err != nil {
        return
    }
    defer srcFile.Close()

    destFile, err := os.Create(dest)
    if err != nil {
        return
    }
    defer func() {
        tmpErr := destFile.Close()
        if tmpErr == nil {
            err = tmpErr
        }
    }()

    if _, err = io.Copy(destFile, srcFile); err != nil {
        return
    }
    err = destFile.Sync()
    return
}

