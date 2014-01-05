package process_count

import (
	"logger"
	"os"
	"path/filepath"
	"strings"
)

func GetMetric(params interface{}, log *logger.Logger) interface{} {
	process_name := params.(string)
	count := 0

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
			count++
		}

		return nil
	}

	err := filepath.Walk("/proc", f)

	if err != nil {
		log.Log("info", "Error while trying to walk /proc: "+err.Error())
	}

	return count
}
