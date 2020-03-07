package ConfigHelper

import "log"

// Get config file
func Get() (config Config, err error) {
	config, err = Parse()
	if err != nil {
		log.Println(err)
		return
	}
	return
}
