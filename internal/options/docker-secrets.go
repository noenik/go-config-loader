package options

import (
	"github.com/noenik/go-config-loader/internal/utils"
	"os"
	"path/filepath"
	"reflect"
)

const defaultDockerSecretDir = "/var/run/secrets"

type WithDockerSecrets[T any] struct {
	Dirs []string
}

func (opt *WithDockerSecrets[T]) Apply(options *ConfigBuilderOptions[T]) error {
	var secretDirs []string

	if opt.Dirs != nil && len(opt.Dirs) > 0 {
		secretDirs = opt.Dirs
	} else {
		secretDirs = []string{defaultDockerSecretDir}
	}

	secrets := make(map[string]string)

	for _, dir := range secretDirs {
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}
			name := file.Name()
			secretPath := filepath.Join(dir, name)
			content, _ := os.ReadFile(secretPath)
			secrets[name] = string(content)
		}
	}

	utils.FindAndApplyValuesByTag(secrets, reflect.ValueOf(options.Config).Elem(), "secret", 1)

	return nil
}
