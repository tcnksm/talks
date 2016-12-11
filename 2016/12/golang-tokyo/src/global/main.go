package main

import "log"

// START OMIT
const defaultConfigPath = "./golang-tokyo.json" // デフォルト値

func main() {
	// ...
	if len(cfgPath) == 0 {
		cfgPath = defaultConfigPath // mainでローカル変数に代入！
	}

	config, err := LoadConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	// ...
}

// END OMIT
