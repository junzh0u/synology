package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/junzh0u/synology"
)

func main() {
	flag.Set("stderrthreshold", "FATAL")
	host := flag.String("host", "localhost:5000", "NAS host and port")
	username := flag.String("username", "admin", "NAS username")
	password := flag.String("password", "", "NAS password")
	flag.Parse()
	client := synology.Client{
		Host:     *host,
		Username: *username,
		Password: *password,
	}

	switch flag.Arg(0) {
	case "DSGetConfig":
		config, err := client.DownloadStationInfoGetConfig()
		if err != nil {
			glog.Fatal(err)
		}
		output, err := json.Marshal(config)
		if err != nil {
			glog.Fatal(err)
		}
		fmt.Println(string(output))

	case "DSSetConfig":
		fs := flag.NewFlagSet("DSSetConfig", flag.ContinueOnError)
		btMaxDownload := fs.String("bt_max_download", "", "Max BT download speed in KB/s (“0” means unlimited)")
		btMaxUpload := fs.String("bt_max_upload", "", "Max BT upload speed in KB/s (“0” means unlimited)")
		emuleMaxDownload := fs.String("emule_max_download", "", "Max eMule download speed in KB/s (“0” means unlimited)")
		emuleMaxUpload := fs.String("emule_max_upload", "", "Max eMule upload speed in KB/s (“0” means unlimited)")
		nzbMaxDownload := fs.String("nzb_max_download", "", "Max NZB download speed in KB/s (“0” means unlimited)")
		httpMaxDownload := fs.String("http_max_download", "", "Max HTTP download speed in KB/s (“0” means unlimited). For more info, please see Limitations.")
		ftpMaxDownload := fs.String("ftp_max_download", "", "Max FTP download speed in KB/s (“0” means unlimited). For more info, please see Limitations.")
		emuleEnabled := fs.String("emule_enabled", "", "If eMule service is enabled")
		unzipServiceEnabled := fs.String("unzip_service_enabled", "", "If Auto unzip service is enabled for users except admin or administrators group")
		defaultDestination := fs.String("default_destination", "", "Default destination")
		emuleDefaultDestination := fs.String("emule_default_destination", "", "Emule default destination")
		fs.Parse(flag.Args()[1:])

		params := map[string]string{
			"bt_max_download":           *btMaxDownload,
			"bt_max_upload":             *btMaxUpload,
			"emule_max_download":        *emuleMaxDownload,
			"emule_max_upload":          *emuleMaxUpload,
			"nzb_max_download":          *nzbMaxDownload,
			"http_max_download":         *httpMaxDownload,
			"ftp_max_download":          *ftpMaxDownload,
			"emule_enabled":             *emuleEnabled,
			"unzip_service_enabled":     *unzipServiceEnabled,
			"default_destination":       *defaultDestination,
			"emule_default_destination": *emuleDefaultDestination,
		}
		_, err := client.DownloadStationInfoSetServerConfig(params)
		if err != nil {
			glog.Fatal(err)
		}

		config, err := client.DownloadStationInfoGetConfig()
		if err != nil {
			glog.Fatal(err)
		}
		output, err := json.Marshal(config)
		if err != nil {
			glog.Fatal(err)
		}
		fmt.Println(string(output))

	default:
		flag.PrintDefaults()
	}
}
