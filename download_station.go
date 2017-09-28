package synology

// DownloadStationInfoGetConfig gets DownloadStation config.
func (c Client) DownloadStationInfoGetConfig() (interface{}, error) {
	return c.get(
		"/webapi/DownloadStation/info.cgi",
		map[string]string{
			"api":     "SYNO.DownloadStation.Info",
			"version": "1",
			"method":  "getconfig",
		},
		map[int]string{},
	)
}

// DownloadStationInfoSetServerConfig sets DownloadStation config.
func (c Client) DownloadStationInfoSetServerConfig(config map[string]string) (interface{}, error) {
	params := map[string]string{
		"api":     "SYNO.DownloadStation.Info",
		"version": "1",
		"method":  "setserverconfig",
	}
	for k, v := range config {
		if v != "" {
			params[k] = v
		}
	}
	return c.get(
		"/webapi/DownloadStation/info.cgi",
		params,
		map[int]string{},
	)
}
