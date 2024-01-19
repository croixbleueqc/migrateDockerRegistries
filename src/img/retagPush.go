// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/img/retagPush.go
// Original timestamp: 2024/01/17 15:43

package img

import (
	"bufio"
	"fmt"
	"migrateDockerRegistries/env"
	"migrateDockerRegistries/helpers"
	"os"
	"strings"
)

func retagImages(regs env.DockerRegistryCreds) error {
	if err := createShellScript(regs); err != nil {
		return err
	}
	fmt.Printf("%s\n", helpers.Blue("Command script created"))
	return nil
}

func createShellScript(regs env.DockerRegistryCreds) error {

	basename := regs.Source.Name + "-" + regs.Dest.Name
	// read the missing images file
	inputFile, err := os.Open(basename + ".txt")
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// write the shell script file
	outputFile, err := os.Create(basename + ".sh")
	if err != nil {
		return err
	}
	defer outputFile.Close()
	outputFile.WriteString("#!/usr/bin/env bash\n\n")

	// iterate tru inputfile, manipulate strings, send to output file
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		var srcStr, dstStr, delStr, pushStr string
		lineString := scanner.Text()

		outputFile.WriteString(fmt.Sprintf("docker pull %s\n", lineString))

		srcStr = lineString[strings.LastIndex(lineString, "/")+1:]
		dstStr = fmt.Sprintf("docker tag %s %s%s\n", lineString, stripProtocol(regs.Dest.URL), srcStr)
		outputFile.WriteString(dstStr)
		if DeleteOrg {
			delStr = fmt.Sprintf("docker rmi %s\n", lineString)
			outputFile.WriteString(delStr)
		}
		if Push {
			pushStr = fmt.Sprintf("docker push %s%s\n", stripProtocol(regs.Dest.URL), srcStr)
			outputFile.WriteString(pushStr)
		}
		if DeleteOrg || Push {
			outputFile.WriteString("\n")
		}
		os.Chmod(basename+".sh", 0755)
		//fmt.Printf("lineString: %s\n", lineString)
		//fmt.Printf("srcStr: %s\n", srcStr)
		//fmt.Printf("dstStr: %s", dstStr)
		//fmt.Printf("delStr: %s", delStr)
		//fmt.Printf("pushStr: %s", pushStr)
	}
	return nil
}
