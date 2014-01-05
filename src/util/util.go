package util

import (
	"logger"
	"os"
	"path/filepath"
	"strings"
)

/*
 * Get the process ids for a given process name, full path is required.
 *
 * Returns strings because it's easier for most of the things we'll use this
 * for.
 */

func GetPids(process_name string, log *logger.Logger) []string {
	pids := []string{}

	f := func(path string, info os.FileInfo, err error) error {
		// we don't care about errors because we shouldn't be root
		// so most of the errors here just return nil to keep walk going
		if err != nil {
			return nil
		}

		parts := strings.Split(path, "/")

		if len(parts) != 4 || parts[3] != "exe" {
			return nil
		}

		link, err := os.Readlink(path)

		if err != nil {
			return nil
		}

		if link == process_name {
			pids = append(pids, parts[2])
		}

		return nil
	}

	err := filepath.Walk("/proc", f)

	if err != nil {
		log.Log("info", "Error while trying to walk /proc: "+err.Error())
	}

	return pids
}
