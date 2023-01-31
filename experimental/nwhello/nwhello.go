package main

import (
    "github.com/vishvananda/netlink"
    "fmt"
    "flag"
    "net"
)

func LinkList() {
  nh, err := netlink.NewHandle()

  if err != nil {
    fmt.Println("Error!", err)
    return
  }

  links, err := nh.LinkList()

  if err != nil {
    fmt.Println("Error!", err)
    return
  }

  for i := 0; i < len(links); i++ {
    fmt.Println("#", i, "type:", links[i].Type(), "name:", links[i].Attrs().Name, "hwaddr:", links[i].Attrs().HardwareAddr)
  }
}

var linkname string
var hwaddr string

func Init() {
	flag.StringVar(&linkname, "linkname", "", "link to modify")
	flag.StringVar(&hwaddr, "hwaddr", "", "hardware address to assign to the link")
  flag.Parse()
}


func DialNetlinkLib() (*netlink.Handle, error) {
  nh, err := netlink.NewHandle()
  return nh, err
}

func SetLinkMacAddress(linkname string, hwaddrstr string) (error) {
  fmt.Println("SetLinkMacAddress: req linkname:", linkname, "hwaddrstr:", hwaddrstr)

  nh, err := DialNetlinkLib()
  if err != nil {
    fmt.Println("Error: ", err)
    return err
  }

  link, err := nh.LinkByName(linkname)
  if err != nil {
    fmt.Println("Error looking for link named", linkname, ":", err)
    return err
  }

  fmt.Println("current status: type:", link.Type(), "name:", link.Attrs().Name, "hwaddr:", link.Attrs().HardwareAddr)

  hwaddr, err := net.ParseMAC(hwaddrstr)

  if err != nil {
    fmt.Println("Error parsing string version of hardware address", err)
    return err
  }

  err = nh.LinkSetHardwareAddr(link, hwaddr)
  if err != nil {
    fmt.Println("Error setting hardware address:", err)
    return err
  }

  fmt.Println("SetLinkMacAddress: resp Success linkname:", linkname, "hwaddrstr:", hwaddrstr)

  return nil
}

func main () {
  Init()

  //LinkList()
  SetLinkMacAddress(linkname, hwaddr)
}
