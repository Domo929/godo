package godo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func Load(c *cli.Context) (*List, error) {
	path, err := getGodoLocation(c.String("location"))
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return create(path)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	list := new(List)
	if err := json.NewDecoder(f).Decode(list); err != nil {
		return nil, err
	}

	return list, nil
}
func save(pathToFile string, l *List) error {
	f, err := os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(l)
}

func Save(c *cli.Context, l *List) error {
	path, err := getGodoLocation(c.String("location"))
	if err != nil {
		return err
	}

	return save(path, l)
}

func create(pathToFile string) (*List, error) {
	directory := filepath.Dir(pathToFile)
	if err := os.MkdirAll(directory, os.ModeDir|os.ModePerm); err != nil {
		return nil, err
	}
	newList := NewList()
	if err := save(pathToFile, newList); err != nil {
		return nil, err
	}
	return newList, nil
}

func getGodoLocation(cliLocation string) (string, error) {
	if cliLocation != "" {
		return validatePath(cliLocation)
	}

	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return validatePath(filepath.Join(user.HomeDir, ".godo/godo.json"))
}

func validatePath(path string) (string, error) {
	if len(path) == 0 {
		return "", fmt.Errorf("%w: path cannot be empty", ErrBadPath)
	}

	if rune(path[0]) != '~' {
		return filepath.Abs(path)
	}

	user, err := user.Current()
	if err != nil {
		return "", err
	}

	convertedFilepath := filepath.Join(user.HomeDir, path[1:])

	return filepath.Abs(convertedFilepath)
}
