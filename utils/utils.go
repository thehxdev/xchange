package utils

import (
    "log"
    "io"
    "os"
    "strings"
    "net/http"
    "io/ioutil"
    "archive/zip"
    "path/filepath"
)


// check for root user permissions
func IsRoot() bool {
    return os.Geteuid() == 0
}


// get public IP address
func PublicIP() ([]byte, error) {
    // More secure domains with `https`
    // banned iran ips :)
    // for example `https://ipinfo.io/ip`
    url := "http://ifconfig.in"
    log.Printf("Sending request to %s", url)

    res, err := http.Get(url)

    if err != nil {
        return nil, err
    }

    defer res.Body.Close()
    data, err := ioutil.ReadAll(res.Body)

    if err != nil {
        return nil, err
    }

    res.Body.Close()
    return data, nil
}


//func ResolvDomain(string domain) 


func DownloadFile(filepath string, url string) (err error) {
    out, err := os.Create(filepath)
    if err != nil  {
        return err
    }
    defer out.Close()

    res, err := http.Get(url)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    _, err = io.Copy(out, res.Body)
    if err != nil  {
        return err
    }

    return nil
}


func ExtractZip(src, dest string) (err error) {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    for _, f := range r.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        fpath := filepath.Join(dest, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, f.Mode())
        } else {
            var fdir string
            if lastIndex := strings.LastIndex(fpath,string(os.PathSeparator)); lastIndex > -1 {
                fdir = fpath[:lastIndex]
            }

            err = os.MkdirAll(fdir, f.Mode())
            if err != nil {
                return err
            }
            f, err := os.OpenFile(
                fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer f.Close()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

