// Package cbs loads data from a csv file specified by a flag, and sends information
// to the kochava blacklist
package cbs

import "log"

// Action is a generic handler that facilitates the loading and sending of references
func Action(logger *log.Logger, filename string, api string, debug bool, action string) error {

	list, err := LoadBlackList(logger, filename)
	if err != nil {
		logger.Println(err)
	}

	err = SendList(logger, list, api, debug, action)
	if err != nil {
		logger.Println(err)
	}

	if err == nil {
		logger.Println("List sent correctly")

	}
	return nil
}
