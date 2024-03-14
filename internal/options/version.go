package options

import (
	"os"
	"reflect"
	"strings"
)

const defaultVersionFile = "version"

type WithVersionFile[T any] struct {
	FilePath string
}

func (opt WithVersionFile[T]) Apply(options *ConfigBuilderOptions[T]) error {
	var fp = opt.FilePath

	if fp == "" {
		fp = defaultVersionFile
	}

	if _, err := os.Stat(fp); err != nil {
		return err
	}

	versionFileContent, err := os.ReadFile(fp)
	if err != nil {
		return err
	}

	return loadVersionInfo(options.Config, string(versionFileContent))
}

type WithVersionString[T any] struct {
	Version string
}

func (opt WithVersionString[T]) Apply(options *ConfigBuilderOptions[T]) error {
	return loadVersionInfo(options.Config, opt.Version)
}

func loadVersionInfo(dst any, serializedVersionString string) error {
	var (
		build   string
		version string
	)

	lines := strings.Split(serializedVersionString, "\n")
	for _, v := range lines {
		if !strings.Contains(v, "=") {
			continue
		}
		kvp := strings.Split(v, "=")
		key := kvp[0]
		value := kvp[1]

		switch key {
		case "version":
			version = value
			break
		case "build":
			build = value
		}
	}

	ps := reflect.ValueOf(dst).Elem()
	versionField := ps.FieldByName("Version")
	if versionField.IsNil() {
		versionField.Set(reflect.New(versionField.Type().Elem()))
		versionField = versionField.Elem()
	}

	if versionField.IsValid() && versionField.CanSet() {
		bf := versionField.FieldByName("Build")
		if bf.IsValid() && bf.CanSet() && bf.Kind() == reflect.String {
			bf.SetString(build)
		}

		vf := versionField.FieldByName("Version")
		if vf.IsValid() && vf.CanSet() && vf.Kind() == reflect.String {
			vf.SetString(version)
		}
	}

	return nil
}
