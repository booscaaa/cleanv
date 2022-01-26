/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

type Github struct {
	Assets []Asset `json:"assets"`
}

type Asset struct {
	URL                string      `json:"url"`
	ID                 int64       `json:"id"`
	NodeID             string      `json:"node_id"`
	Name               string      `json:"name"`
	Label              interface{} `json:"label"`
	Uploader           Uploader    `json:"uploader"`
	ContentType        string      `json:"content_type"`
	State              string      `json:"state"`
	Size               int64       `json:"size"`
	DownloadCount      int64       `json:"download_count"`
	CreatedAt          string      `json:"created_at"`
	UpdatedAt          string      `json:"updated_at"`
	BrowserDownloadURL string      `json:"browser_download_url"`
}

type Uploader struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		downloadBinary(path)
	},
}

func downloadBinary(path string) {
	fmt.Println("Updating SDK...")
	http.DefaultTransport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	var github Github

	resp, err := http.Get("https://api.github.com/repos/booscaaa/cleanv/releases/latest")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&github); err != nil {
		fmt.Println("There was a problem with the SDK update")
		fmt.Println(err)
		return
	}

	for _, asset := range github.Assets {
		if runtime.GOOS == "windows" {
			if asset.Name == "cleanv.exe" {
				downloadToWindows(path, fmt.Sprint(asset.ID))
			}
		} else {
			if asset.Name == "cleanv" {
				downloadToLinux(path, fmt.Sprint(asset.ID))
			}
		}

	}

	fmt.Println("Cleanv command line successfully updated!")
	// }

}

func downloadToLinux(path, asset string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/booscaaa/cleanv/releases/assets/"+asset, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header = http.Header{
		"Accept": []string{"application/octet-stream"},
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	out, _ := os.Create(path + "/clenvnew")
	defer out.Close()

	io.Copy(out, resp.Body)

	os.Rename(path+"/clenvnew", path+"/cleanv")

	os.Chmod(path+"/cleanv", 0777)

	exec.Command("cleanv", "completion", "bash", ">", "/tmp/completion").Run()
}

func downloadToWindows(path, asset string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/booscaaa/cleanv/releases/assets/"+asset, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header = http.Header{
		"Accept": []string{"application/octet-stream"},
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	out, _ := os.Create(path + "/clenvnew.exe")
	defer out.Close()

	io.Copy(out, resp.Body)

	fmt.Println()
	fmt.Println("Run the command: rename " + path + "/clenvnew.exe " + path + "/cleanv.exe")
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	var path string
	if runtime.GOOS == "windows" {
		path = "C:/cleanv"
	} else {
		path = "/usr/local/cleanv"
	}
	updateCmd.Flags().StringP("path", "p", path, "Cleanv sdk path")
}
