package cbs

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// LoadBlackList loads a csv file to be sent to the Kochava server
func LoadBlackList(logger *log.Logger, filename string) (BlackList, error) {

	var list BlackList

	file, err := os.Open(filename)
	if err != nil {
		logger.Println(err)
		return list, err
	}

	c := csv.NewReader(file)
	records, err := c.ReadAll()

	lum := make(map[string]int)
	for i := range records[0] {
		lum[records[0][i]] = i
	}

	for i := range records {
		record := records[i]
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Println(err)
			break
		}

		// switch keys off of type entered in the CSV record
		switch record[lum["type"]] {
		case "site_id":
			var site siteID
			site.Type = "siteId"
			site.BlackListSiteID.AccountID = record[lum["account_id"]]
			site.BlackListSiteID.NetworkID = record[lum["network_id"]]
			site.BlackListSiteID.SiteID = record[lum["site_id"]]
			site.BlackListSiteID.Reason = record[lum["reason"]]

			score, _ := strconv.Atoi(record[lum["score"]])

			site.BlackListSiteID.Score = score
			site.BlackListSiteID.Source = 2

			list.BlackListSiteIDs = append(list.BlackListSiteIDs, site)
		case "device_id":
			var device deviceID
			device.Type = "device"
			device.BlackListDevice.AccountID = record[lum["account_id"]]
			device.BlackListDevice.DeviceIDValue = record[lum["device_id_value"]]
			device.BlackListDevice.DeviceIDType = record[lum["device_id_type"]]
			device.BlackListDevice.Reason = record[lum["reason"]]

			score, _ := strconv.Atoi(record[lum["score"]])

			device.BlackListDevice.Score = score

			device.BlackListDevice.Source = 2
			list.BlackListDevices = append(list.BlackListDevices, device)
		case "ip_address":
			var ip ipAddress
			ip.Type = "ip"
			ip.BlackListIP.AccountID = record[lum["account_id"]]
			ip.BlackListIP.IPAddress = record[lum["ip_address"]]
			ip.BlackListIP.Reason = record[lum["reason"]]

			score, _ := strconv.Atoi(record[lum["score"]])

			ip.BlackListIP.Score = score
			ip.BlackListIP.Source = 2

			list.BlackListIPs = append(list.BlackListIPs, ip)
		}

	}

	return list, nil
}
