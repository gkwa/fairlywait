package main

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
)

func main() {
    fmt.Println("App Version:", versioninfo.Version)
    fmt.Println("Git Revision:", versioninfo.Revision) 
    fmt.Println("Last Commit:", versioninfo.LastCommit)
    fmt.Println("Dirty Build:", versioninfo.DirtyBuild)
    fmt.Println("Short Version Info:", versioninfo.Short())
}
