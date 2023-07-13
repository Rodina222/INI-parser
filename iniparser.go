package iniparser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (

	// ErrInvalidExtension is returned when the extension of the input file is not ".ini"
	ErrInvalidExtension = errors.New("the file should have an ini extension (.ini)")

	// ErrInvalidSyntax is returned when there is an invalid ini syntax in a line with specifing its number
	ErrInvalidSyntax = errors.New("invalid ini syntax")

	//ErrSectionNotFound is returned when trying to get the value for a key in a non existent section
	ErrSectionNotFound = errors.New("key you entered doesn't exist")

	//ErrSectionAlreadyExists is returned when there are two sections with the same name
	ErrSectionAlreadyExists = errors.New("section you entered already exists")

	//ErrKeyNotFound is returned when trying to get the value for a key that doesn't exist
	ErrKeyNotFound = errors.New("key you entered doesn't exist")

	// ErrEmptyKey is returned when a line has a value without key before the equal sign
	ErrEmptyKey = errors.New("key is empty")

	//ErrKeyEmptyValue is returned when not entering the section name or key
	ErrValuesEmpty = errors.New("section name or key can't be empty")
)

// INISection represents a section in the INI config file
type INISection map[string]string

// INIParser represents a struct of all the INI config file sections
type INIParser struct {
	sections map[string]INISection
}

// NewINIParser is used to create a new INI parser
func NewINIParser() INIParser {

	return INIParser{
		sections: make(map[string]INISection),
	}
}

// loadFromReader is used to create the map from either a file or a string
func (parser *INIParser) loadFromReader(reader io.Reader) error {

	parser.sections = make(map[string]INISection)

	// Create a scanner to read the data line by line
	scanner := bufio.NewScanner(reader)

	index := 0
	section := ""

	// Read the file line by line
	for scanner.Scan() {

		index++

		line := strings.TrimSpace(scanner.Text())

		// skipping empty lines and comments
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}

		if line[0] == '[' && line[len(line)-1] == ']' && strings.Count(line, "[") == 1 && strings.Count(line, "]") == 1 {

		}

		// INIsection name starts with "["and ends with "]" and repeated only once
		if line[0] == '[' && line[len(line)-1] == ']' && strings.Count(line, "[") == 1 && strings.Count(line, "]") == 1 {

			section = strings.TrimSpace(line[1 : len(line)-1])

			// check section is not empty
			if len(section) == 0 {
				return fmt.Errorf("%w: invalid section at line %d", ErrInvalidSyntax, index)
			}

			// check section does not exist
			_, ok := parser.sections[section]
			if ok {

				return ErrSectionAlreadyExists
			}
			// make a new section
			parser.sections[section] = make(INISection)
			continue
		}

		// valid key-pair line
		if strings.Contains(line, "=") && section != "" {

			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])

			// check whether the key is empty or not
			if key != "" {

				value := strings.TrimSpace(parts[1])
				parser.sections[section][key] = value
				continue
			}

			return ErrEmptyKey
		}

		//invalid syntax with specifing the number of line that has the error
		wrappedErr := fmt.Errorf("%w line %d", ErrInvalidSyntax, index)
		return wrappedErr

	}
	return nil

}

// LoadFromFile is responsible for loading the INI file data into a reader and sends it to loadFromReader method
func (parser *INIParser) LoadFromFile(path string) error {

	// check the extension of the input file
	ext := filepath.Ext(path)
	if ext != ".ini" {
		return ErrInvalidExtension
	}

	//reading the file
	file, err := os.ReadFile(path)
	if err != nil {

		return fmt.Errorf("failed to read the file %s %w", path, err)

	}

	reader := strings.NewReader(string(file))

	return parser.loadFromReader(reader)
}

// LoadFromFile is responsible for loading the INI string data into a reader and sends it to loadFromReader method
func (parser *INIParser) LoadFromString(data string) error {

	reader := strings.NewReader(data)
	return parser.loadFromReader(reader)

}

// GetSections returns the whole INI sections data stored in the map
func (parser *INIParser) GetSections() map[string]INISection {

	return parser.sections
}

// GetSectionNames returns the INI sections names stored in the map
func (parser *INIParser) GetSectionNames() []string {

	sectionNames := make([]string, 0, len(parser.sections))
	for key := range parser.sections {
		sectionNames = append(sectionNames, key)
	}

	return sectionNames

}

// GetValue returns the value given its section name and key
func (parser *INIParser) GetValue(sectionName, key string) (string, error) {

	if sectionName == "" || key == "" {
		return "", ErrValuesEmpty
	}

	section, ok := parser.sections[sectionName]
	if !ok {
		return "", ErrSectionNotFound
	}

	value, ok := section[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

// SetValue sets a value for a key in a section whether this section exists or not
func (parser *INIParser) SetValue(SectionName, key, value string) error {

	if SectionName == "" || key == "" {
		return ErrValuesEmpty
	}

	section, ok := parser.sections[SectionName]

	// create a new section if section does not exist
	if !ok {
		section = make(INISection)
		parser.sections[SectionName] = section
	}

	section[key] = value

	return nil
}

// String converts the INIParser into string
func (parser *INIParser) String() string {

	var sb strings.Builder

	for key, value := range parser.sections {
		sectionName := fmt.Sprintf("[%s]\n", key)
		sb.WriteString(sectionName)

		// Convert the section to a string
		for k, v := range value {
			pair := fmt.Sprintf("%s = %s\n", k, v)
			sb.WriteString(pair)
		}
		//sb.WriteString("\n")

	}
	return sb.String()
}

// SaveToFile saves the ini data to a file
func (parser *INIParser) SaveToFile(path string) error {

	// check the file extension
	fileExt := filepath.Ext(path)
	if fileExt != ".ini" {
		return ErrInvalidExtension
	}

	// Convert the ini data to string
	ini_data := parser.String()

	// Write the string to a file
	// 0644 is the file permissions which means that the file is readable and writable by the owner, and readable by everyone else
	return ioutil.WriteFile(path, []byte(ini_data), 0644)

}
