package iniparser

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Exported errors to the user
var ErrInvalidExtension = errors.New("the file should have an ini extension (.ini)")

var ErrFileNotExist = errors.New("the file you entered doesn't exist")

var ErrInvalidFileFormat = errors.New("invalid ini file format")

var ErrInvalidSyntax = errors.New("this line misses an equal sign")

var ErrSectionNotExist = errors.New("section you entered doesn't exist")

// Used types
type (
	INISection map[string]string
)

type INIParser struct {
	sections map[string]INISection
}

// this function is used to create a new INI parser
func New() *INIParser {

	return &INIParser{
		sections: make(map[string]INISection),
	}
}

// this function is used to create the map from either a file or a string

func LoadFromReader(reader io.Reader) (map[string]INISection, error) {

	parser := New()

	// Create a scanner to read the data line by line
	scanner := bufio.NewScanner(reader)

	// Read the file line by line

	var section string = ""

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		// skipping empty lines and comments
		if len(line) == 0 || line[0] == ';' {
			continue
		}

		// we know it is a section name if it starts with "["and ends with "]" and repeated only once
		if line[0] == '[' && line[len(line)-1] == ']' && strings.Count(line, "[]") == 1 {
			section = strings.TrimSpace(line[1 : len(line)-1])
			parser.sections[section] = INISection{}

			// Missing one of the 2 brackets case
		}
		if (line[0] == '[' && line[len(line)-1] != ']') || (line[0] != '[' && line[len(line)-1] == ']') {

			wrappedErr := errors.Wrapf(ErrInvalidFileFormat, line[1:len(line)-1])
			return parser.sections, wrappedErr

		}

		// if it is a normal line so we will append it to the INI parser map but we should check it first
		if strings.Contains(line, "=") {

			parts := strings.SplitN(line, "=", 1) //The argument specifies that we want to split the string based on the 1st occurrence of the "="
			key := parts[0]
			value := parts[1]
			parser.sections[section][key] = value

		} else {
			wrappedErr := errors.Wrapf(ErrInvalidSyntax, "line: %q", line)
			return parser.sections, wrappedErr
		}
	}
	return parser.sections, nil

}

func (parser *INIParser) LoadFromFile(path string) error {

	// checks on the input file
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return ErrFileNotExist
	}
	ext := filepath.Ext(path)
	if ext != ".ini" || len(ext) == 0 {
		return ErrInvalidExtension
	}

	//reading the file
	file, err := os.ReadFile(path)
	if err != nil {
		return errors.New("error while opening the file")
	}

	reader := strings.NewReader(string(file))
	inimap, err := LoadFromReader(reader)

	if err != nil {
		return err
	}

	// Set the sections in the INIParser instance
	parser.sections = inimap
	return nil
}

func (parser *INIParser) LoadFromString(data string) error {

	reader := strings.NewReader(data)
	inimap, err := LoadFromReader(reader)

	if err != nil {
		return err
	}
	parser.sections = inimap
	return nil
}

func (parser *INIParser) GetSections() map[string]INISection {

	return parser.sections
}

func (parser *INIParser) GetSectionNames() []string {

	section_names := make([]string, 0, len(parser.sections))
	for key := range parser.sections {
		section_names = append(section_names, key)
	}
	return section_names
}

func (parser *INIParser) GetValue(SectionName, key string) (string, error) {

	if SectionName == "" || key == "" {
		return "", errors.New("section name or key can't be empty")
	}

	section, ok := parser.sections[SectionName]
	if !ok {
		return "", ErrSectionNotExist
	}

	value, ok := section[key]
	if !ok {
		return "value is empty", nil
	}

	return value, nil
}

func (parser *INIParser) SetValue(SectionName, key, value string) error {

	if SectionName == "" || key == "" || value == "" {
		return errors.New("section name or key or value can't be empty")
	}

	section, ok := parser.sections[SectionName]
	if !ok {
		section = make(INISection)
		parser.sections[SectionName] = section
	}

	section[key] = value

	return nil
}

func (parser *INIParser) String() string {

	var sb strings.Builder

	for key, value := range parser.sections {
		section_name := "[" + key + "]\n"
		sb.WriteString(section_name)

		// Convert the section to a string
		for k, v := range value {
			pair := k + "=" + v + "\n"
			sb.WriteString(pair)
		}
	}
	return sb.String()
}

func (parser *INIParser) SaveToFile(path string) error {

	if path == "" {
		return errors.New("file path can't be empty")
	}

	// Add .ini extension to the file path if it doesn't already have one
	if !strings.HasSuffix(path, ".ini") {
		path += ".ini"
	}

	// Convert the parser to an INI string
	ini_data := parser.String()

	// Write the INI string to a file
	// 0644 is the file permissions which means that the file is readable and writable by the owner, and readable by everyone else
	err := ioutil.WriteFile(path, []byte(ini_data), 0644)
	if err != nil {
		return errors.New("error occurred while writing the file")
	}

	fmt.Println("ini file saved successfully")
	return nil
}
