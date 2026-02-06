package main

import (
	"fmt"
	"jsonStorage/bins"
)


func main() {
    binList := bins.BinList(&[]bins.Bin{})
    bins.NewBin("1", false, "My Bin", binList)
    fmt.Println(*binList)
}
