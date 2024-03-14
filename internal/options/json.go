package options

import (
	"bytes"
	"encoding/json"
	"os"
	"text/template"
)

const (
	defaultJsonFile            = "config.json"
	defaultEnvJsonFileTemplate = "config.{{ .Environment }}.json"
)

type WithEnvJsonFile[T any] struct {
	FilePathTemplate string
}

func (opt *WithEnvJsonFile[T]) Apply(options *ConfigBuilderOptions[T]) error {
	var (
		cfg              T
		err              error
		filePathTemplate *template.Template
		templateString   = opt.FilePathTemplate
	)

	if opt.FilePathTemplate == "" {
		templateString = defaultEnvJsonFileTemplate
	}

	filePathTemplate, err = template.New("EnvJsonFile").Parse(templateString)
	if err != nil {
		return err
	}

	var buf []byte
	filePath := bytes.NewBuffer(buf)
	err = filePathTemplate.Execute(filePath, struct {
		Environment string
	}{options.Environment})
	if err != nil {
		return err
	}

	err = readJsonFile(filePath.String(), &cfg)
	if err != nil {
		return err
	}

	options.Merge(&cfg)

	return nil
}

type WithJsonFile[T any] struct {
	FilePath string
}

func (opt WithJsonFile[T]) Apply(options *ConfigBuilderOptions[T]) error {
	var (
		cfg      T
		filePath = opt.FilePath
	)

	if filePath == "" {
		filePath = defaultJsonFile
	}

	err := readJsonFile(filePath, &cfg)
	if err != nil {
		return err
	}

	options.Merge(&cfg)

	return nil
}

type WithJsonString[T any] struct {
	ConfigJson string
}

func (opt *WithJsonString[T]) Apply(options *ConfigBuilderOptions[T]) error {
	var cfg T

	err := json.Unmarshal([]byte(opt.ConfigJson), &cfg)
	if err != nil {
		return err
	}

	options.Merge(&cfg)

	return nil
}

func readJsonFile(file string, value any) error {
	_, err := os.Stat(file)
	if err != nil {
		return err
	}

	jsonContent, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonContent, value)
	if err != nil {
		return err
	}

	return err
}
