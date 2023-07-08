package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type (
	INI_section map[string]string
)

type INI_Parser map[string]INI_section

func LoadFromFile(filename string) string {

	// Open the file for reading
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error while opening the file!")
	}

	// Scheduling closing the file after the function finish reading its content
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Concatenate the lines into a single variable
	var builder strings.Builder

	// Read the file line by line

	for scanner.Scan() {
		line := scanner.Text()
		builder.WriteString(line + "\n")

		// skipping empty lines and comments
		if len(line) == 0 || line[0] == ';' {
			continue
		}

	}
	// Check for errors during scanning

	if scan_err := scanner.Err(); scan_err != nil {
		fmt.Println("Error while scanning the file!")
	}

	// Get the concatenated string
	file_data := builder.String()

	return file_data

}

func LoadFromString(input_data string) string {

	// parsing the string to lines according to commas
	parsed_values := strings.Split(input_data, ",")

	// concatenate the parsed values into a string separated by lines
	var builder strings.Builder

	for _, value := range parsed_values {

		builder.WriteString(value + "\n")
	}
	// Get the concatenated string
	string_data := builder.String()

	return string_data

}

func GetSections(ini_data string) (INI_Parser, error) {

	parser_map := make(INI_Parser)

	// Create a scanner to read the string line by line
	scanner := bufio.NewScanner(strings.NewReader(ini_data))
	var section string = ""

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		// skipping empty lines and comments
		if len(line) == 0 || line[0] == ';' {
			continue
		}

		// we know it is a section name if it starts with "["and ends with "]"
		if line[0] == '[' && line[len(line)-1] == ']' {
			section = strings.TrimSpace(line[1 : len(line)-1])
			parser_map[section] = make(INI_section)

			// Missing one of the 2 brackets case
		} else if (line[0] == '[' && line[len(line)-1] != ']') || (line[0] != '[' && line[len(line)-1] == ']') {

			return parser_map, errors.New("There is a Section name that misses one of the 2 brackets! please, fix the issue.")

			// the final case is that it is a normal line so we will append it to the INI parser map but we should check it first
		} else {

			if strings.Contains(line, "=") {

				line_tokens := strings.SplitN(line, "=", 2) //The 2 argument specifies that we want to split the string into a maximum of two substrings
				key := line_tokens[0]
				value := line_tokens[1]

				parser_map[section][key] = value

			} else {
				return parser_map, errors.New("There is a line without '=' which means there is a key without a value! please, fix the issue.")
			}
		}

	}

	return parser_map, nil

}

func GetSectionNames(parser_map INI_Parser) []string {

	section_names := make([]string, 0, len(parser_map))

	for key := range parser_map {
		section_names = append(section_names, key)
	}
	return section_names
}

func GetValue(parser_map INI_Parser, section_name string, key string) (string, error) {

	if section_name == "" || key == "" {
		return "", errors.New("Section name or key can't be empty! please, try again.")
	}

	value := parser_map[section_name][key]

	if value != "" {
		return value, nil

	} else {
		return value, errors.New("The key you entered isn't correct or doesn't exist! please, try again.")

	}

}

func SetValue(parser_map *INI_Parser, section_name string, KEY string, value string) error {

	if section_name == "" || KEY == "" || value == "" {

		return errors.New("Section name or key or value can't be empty! please, enter a value.")
	}

	for key := range *parser_map {

		if key == section_name {

			(*parser_map)[section_name][KEY] = value
			return nil

		}

	}

	// Create a new section in the existing map if the input section doesn't exist
	(*parser_map)[section_name] = make(INI_section)
	(*parser_map)[section_name][KEY] = value
	return nil
}

func ToString(parser_map INI_Parser) string {

	str_data := ""

	for key, value := range parser_map {

		section_name := "[" + key + "]"
		str_data += section_name
		str_data += "\n"

		// Convert the section to a string
		section_str := ""
		for K, V := range value {
			section_str += K + "=" + V + "\n"
		}

		str_data += section_str
	}
	return str_data

}

func SaveToFile(ini_data string) error {

	// Write the INI string to a file
	// 0644  is the file permissions which means that the file is readable and writable by the owner, and readable by everyone else
	err := ioutil.WriteFile("written_file.ini", []byte(ini_data), 0644)
	if err != nil {
		return errors.New("An error occurred while writing the file!")
	}

	fmt.Println("INI file saved successfully")
	return nil
}

func main() {

	// test a working file
	file_data := LoadFromFile("config.ini")

	//fmt.Println(file_data)

	config_data := "[NETWORK], host = example.com, port = 7878, [database], host = localhost, port = 5432,username = postgres,password = password, [Email], username= host_email.com,password=12345"

	string_data := LoadFromString(config_data)

	INI_sections, error := GetSections(file_data)

	if error != nil {

		fmt.Println("Error:", error)

	} else {

		fmt.Println(string_data)
		fmt.Println(INI_sections)
		section_names := GetSectionNames(INI_sections)
		fmt.Println("Hello from Section_names!", section_names)

		value, error := GetValue(INI_sections, "Email", "password")

		if error == nil {
			fmt.Println(value)
		}

		err := SetValue(&INI_sections, "NETWORK", "host", "ex.com")

		if err == nil {
			fmt.Println(INI_sections)
		}

		str_data := ToString(INI_sections)
		fmt.Println("string map!", str_data)

		e := SaveToFile(str_data)

		if e != nil {
			fmt.Println(e)
		}

	}

	// test a failed file
	failed_file := LoadFromFile("failed_file1.ini")
	sections, error := GetSections(failed_file)

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Map:", sections)
	}

}
