# GoVer

Golang version checker - the simple way to check version. GoVer can check the latest version from GitHub releases by tags with `MAJOR.MINOR.MIRCRO` pattern.

## Install

```shell
go get github.com/black40x/gover
```

## Usage

You can check version from the latest GitHub release (via tag_name).

```go
func CheckVersion() {
    currentV, _ := NewVersion("v0.0.30")
    latestV, err := GetGithubVersion("black40x", "tunl-cli")
    if err == nil {
        ver, _ := latestV.GetVersion()
        fmt.Printf("Current version: %s \n", currentV.String())
        fmt.Printf("Latest version: %s \n", ver.String())
    
        if ver.NewerThan(*currentV) {
            fmt.Println("Update available")
            fmt.Printf("Update information: \n%s\n", latestV.Body)
        }
    }
}
```
