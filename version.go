package gover

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const VersionChars = "1234567890."

var InvalidVersion = errors.New("invalid version string")

type Version struct {
	Major int
	Minor int
	Micro int
	raw   string
}

type GitHubVersion struct {
	Name        string `json:"name"`
	TagName     string `json:"tag_name"`
	Body        string `json:"body"`
	HtmlUrl     string `json:"html_url"`
	CreatedAt   string `json:"created_at"`
	PublishedAt string `json:"published_at"`
}

func NewVersion(vStr string) (*Version, error) {
	if strings.Trim(vStr, "") == "" {
		return nil, InvalidVersion
	}

	trimmed := ""
	for i := 0; i < len(vStr); i++ {
		if strings.Index(VersionChars, string(vStr[i])) != -1 {
			trimmed += string(vStr[i])
		}
	}

	nums := strings.Split(trimmed, ".")
	if len(nums) == 0 {
		return nil, InvalidVersion
	}

	ver := &Version{
		raw: trimmed,
	}

	if len(nums) >= 1 {
		ver.Major, _ = strconv.Atoi(nums[0])
	}
	if len(nums) >= 2 {
		ver.Minor, _ = strconv.Atoi(nums[1])
	}
	if len(nums) >= 3 {
		ver.Micro, _ = strconv.Atoi(nums[2])
	}

	return ver, nil
}

func (v Version) NewerThan(ver Version) bool {
	if v.Major > ver.Major {
		return true
	} else if v.Major == ver.Major && v.Minor > ver.Minor {
		return true
	} else if v.Major == ver.Major && v.Minor == ver.Minor && v.Micro > ver.Micro {
		return true
	} else {
		return false
	}
}

func (v Version) NewerThanStr(ver string) bool {
	vv, err := NewVersion(ver)
	if err != nil {
		return false
	}
	return v.NewerThan(*vv)
}

func (v Version) EqualOrHigher(ver Version) bool {
	if v.Major > ver.Major {
		return true
	} else if v.Major == ver.Major && v.Minor >= ver.Minor && v.Micro >= ver.Micro {
		return true
	} else {
		return false
	}
}

func (v Version) EqualOrHigherStr(ver string) bool {
	vv, err := NewVersion(ver)
	if err != nil {
		return false
	}
	return v.EqualOrHigher(*vv)
}

func (v Version) UpTo(ver Version) bool {
	return !v.NewerThan(ver)
}

func (v Version) UpToStr(ver string) bool {
	vv, err := NewVersion(ver)
	if err != nil {
		return false
	}
	return v.UpTo(*vv)
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Micro)
}

func GetGithubVersion(user, repo string) (*GitHubVersion, error) {
	client := &http.Client{}
	client.Timeout = time.Second * 30
	resp, err := client.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", user, repo))
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	version := &GitHubVersion{}
	err = json.Unmarshal(bytes, version)
	if err != nil {
		return nil, err
	}

	return version, nil
}

func (v GitHubVersion) GetVersion() (*Version, error) {
	return NewVersion(v.TagName)
}
