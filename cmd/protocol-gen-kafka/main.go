package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ethanmoffat/kafka-protocol/internal/codegen"
	"github.com/ethanmoffat/kafka-protocol/internal/jsonmodel"
)

func main() {
	const (
		KafkaSubmoduleRootDefault = "kafka"
		KafkaMessageSpecPath      = "clients/src/main/resources/common/message"

		ProtocolOutputDirDefault = "pkg/protocol/messages"

		RequestSuffix  = "Request.json"
		ResponseSuffix = "Response.json"
		HeaderSuffix   = "Header.json"
	)
	var (
		inputDir  string
		outputDir string
	)

	flag.StringVar(&inputDir, "i", KafkaSubmoduleRootDefault, "The input directory for eo-protocol files.")
	flag.StringVar(&outputDir, "o", ProtocolOutputDirDefault, "The output directory for generated code.")
	flag.Parse()

	if _, err := os.Stat(inputDir); err != nil {
		fmt.Printf("error: input directory %s does not exist\n", inputDir)
		os.Exit(1)
	}

	if _, err := os.Stat(outputDir); err != nil {
		fmt.Printf("error: output directory %s does not exist\n", outputDir)
		os.Exit(1)
	}

	fmt.Printf("Using parameters:\n  inputDir:  %s\n  outputDir: %s\n", inputDir, outputDir)

	requestFiles := []string{}
	responseFiles := []string{}
	headerFiles := []string{}
	filepath.WalkDir(path.Join(inputDir, KafkaMessageSpecPath), func(currentPath string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(currentPath, RequestSuffix) {
			requestFiles = append(requestFiles, currentPath)
		} else if strings.HasSuffix(currentPath, ResponseSuffix) {
			responseFiles = append(responseFiles, currentPath)
		} else if strings.HasSuffix(currentPath, HeaderSuffix) {
			headerFiles = append(headerFiles, currentPath)
		} else {
			fmt.Printf("warning: skipping unexpected file %s\n", currentPath)
		}
		return nil
	})

	sets := [][]string{
		requestFiles,
		responseFiles,
		headerFiles,
	}

	for _, set := range sets {
		fmt.Printf("processing %d files...", len(set))

		for _, file := range set {
			fp, err := os.Open(file)
			if err != nil {
				fmt.Printf("error opening file: %v\n", err)
				os.Exit(1)
			}
			defer fp.Close()

			bytes, err := io.ReadAll(fp)
			if err != nil {
				fmt.Printf("error reading file: %v\n", err)
				os.Exit(1)
			}

			split := strings.Split(string(bytes), "\n")
			mod := make([]byte, 0, len(bytes))
			for _, s := range split {
				commentNdx := strings.Index(s, "//")
				if commentNdx >= 0 {
					mod = append(mod, []byte(s[:commentNdx])...)
				} else {
					mod = append(mod, []byte(s)...)
				}
			}

			var next jsonmodel.MessageSpec
			if err := json.Unmarshal(mod, &next); err != nil {
				fmt.Printf("error unmarshalling json: %v\n", err)
				os.Exit(1)
			}

			fileNameWithExt := filepath.Base(file)
			parts := strings.Split(fileNameWithExt, ".")
			var fileName string
			if len(parts) <= 1 {
				fileName = fileNameWithExt
			} else {
				fileName = parts[1]
			}

			err = codegen.Generate(outputDir, fileName, next)
			if err != nil {
				fmt.Printf("error generating code: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Printf("done\n")
	}
}
