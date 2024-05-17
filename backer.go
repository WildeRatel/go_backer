package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	configs, mode := loadConfig()
	fmt.Println(configs)

	copyOver(configs, mode)
}

func loadConfig() (configs [][]string, mode bool) {
	config, _ := os.Open("config.txt")
	defer config.Close()

	r := bufio.NewReader(config)
	configArgs := []string{}
	mode = false

	for i := 0; ; i++ {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			configArgs = append(configArgs, string(line))
		}
		if err != nil {
			break
		}
	}

	var dirFrom []string = configArgs[0:1]
	var dirTo []string = configArgs[1:2]
	var copyFiles []string = configArgs[2:]
	if len(configArgs) == 2 {
		mode = true
	}
	configs = append(configs, dirFrom, dirTo, copyFiles)
	return configs, mode
}

func copyOver(configs [][]string, mode bool) {
	dirTo := strings.TrimSpace(configs[1][0])
	dirFrom := strings.TrimSpace(configs[0][0])

	if !mode {
		for i := 0; i < len(configs[2]); i++ {
			tempDir := dirFrom + strings.TrimSpace(configs[2][i])
			copyCommand := exec.Command("cp", tempDir, dirTo)

			if errors.Is(copyCommand.Err, exec.ErrDot) {
				copyCommand.Err = nil
			}
			if err := copyCommand.Run(); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		copyCommand := exec.Command("cp", "-R", dirFrom, dirTo)

		if errors.Is(copyCommand.Err, exec.ErrDot) {
			copyCommand.Err = nil
		}
		if err := copyCommand.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
