# __INI_Parser_Library__

This is a useful package that contains Golang functions for parsing simple INI configuration files.

## __Features:__
- Parse an INI file and obtain configuration values.
- double-check their syntax.
- Change the values of the keys in the INI file.
- Save the updated configuration to a new text file.

## __Enduser Manual:__
1. Load the package to your project as follows:
`import "path/to/main"`.
2. You can load the configuration data whether from a file:`file_data, error := LoadFromFile(FileName string)` 
or from a string! : `string_data := LoadFromString(config_data string)`  where config_data is the configuration data in string format.
3. You can obtain all sections in the form of a map key-value pairs as follows: `INI_data, error := GetSections(file_data string)`
4. Using the GetValue function, you can access configuration values by section and key as follows: `value, error := GetValue(INI_data INI_sections , section_name string, key string)`.
5. SetValue function can be used to update key values and set new values if you wish to add a new key with its value to an existing or new section as follows: 
`error := SetValue(INI_data &INI_sections, section_name string, key string, value string)`.
6. You can also convert the INI sections and data to string using the ToString function as follows: `str_data := ToString(INI_data INI_sections)`.
7. Finally you can save all your changes back into an INI file as follows: error := `SaveToFile(str_data string)`
- *I hope it will be help you!*




