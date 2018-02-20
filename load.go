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

	for {
		record, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Println(err)
			break
		}

		// switch keys off of type entered in the CSV record
		switch record[8] {
		case "site_id":
			var site siteID
			site.Type = "siteId"
			site.BlackListSiteID.AccountID = record[0]
			site.BlackListSiteID.NetworkID = record[1]
			site.BlackListSiteID.SiteID = record[2]
			site.BlackListSiteID.Reason = record[3]

			score, _ := strconv.Atoi(record[4])

			site.BlackListSiteID.Score = score
			site.BlackListSiteID.Source = 2

			list.BlackListSiteIDs = append(list.BlackListSiteIDs, site)
		case "device_id":
			var device deviceID
			device.Type = "device"
			device.BlackListDevice.AccountID = record[0]
			device.BlackListDevice.DeviceIDValue = record[5]
			device.BlackListDevice.DeviceIDType = record[6]
			device.BlackListDevice.Reason = record[3]

			score, _ := strconv.Atoi(record[4])

			device.BlackListDevice.Score = score

			device.BlackListDevice.Source = 2
			list.BlackListDevices = append(list.BlackListDevices, device)
		case "ip_address":
			var ip ipAddress
			ip.Type = "ip"
			ip.BlackListIP.AccountID = record[0]
			ip.BlackListIP.IPAddress = record[7]
			ip.BlackListIP.Reason = record[3]

			score, _ := strconv.Atoi(record[4])

			ip.BlackListIP.Score = score
			ip.BlackListIP.Source = 2

			list.BlackListIPs = append(list.BlackListIPs, ip)
		}

	}

	return list, nil
}
