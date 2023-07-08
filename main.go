package main

import (
    "fmt"
    "flag"
    "github.com/thehxdev/xchange/utils"
)


func main() {
    getPubIp := flag.Bool("myip", false, "Get your public IP address")
    upTun    := flag.Bool("tun2socks", false, "Install/Update tun2socks")
    flag.StringVar(&utils.ConfigPool, "dir", utils.PathExpandUser("~/.local/xconfs"), "Config files directory")
    flag.StringVar(&utils.XrayConfigFile, "xconf", utils.PathExpandUser("/etc/xray/config.json"), "Path to Xray's Config file")

    flag.Parse()

    if *getPubIp {
        publicIP, err := utils.PublicIP()
        check(err)
        fmt.Println(string(publicIP))
        return
    }

    if *upTun {
        err := utils.InstallTun2Socks("~/.local/bin")
        check(err)
        return
    }

    checkRoot()
    err := utils.ChangeConfigFile()
    check(err)
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}


func checkRoot() {
    if utils.IsRoot() == false {
        panic(fmt.Errorf("To change xray config file run app as root user"))
    }
}

