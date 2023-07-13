# __INI_Parser_Library__

This is a useful package that contains Golang functions for parsing simple INI configuration files.

## __Features:__
- Parse an INI file/string and obtain configuration values.
- Store the INI data in a map after double-check its syntax.
- Get the value of a key after specifing its section.
- Set the value for a key after specifing its section.
- Converting the INI data to a string.
- Save the updated INI data to an INI file.


## __Manual:__

### __Here is a guide of how you can use the package methods:__


1. Load the package to your project as follows:
`import "github.com/codescalersinternships/INIParser-Rodina"`.

2. Load the configuration data as follows:
from a file:
`fileData, error := LoadFromFile(/path/to/file.ini)` 
or from a string :
 `stringData := LoadFromString([section1]\nkey=value\n[section2]\nkey=value\n..etc)`  
 where data is the INI configuration data in string format.

3. Create a new ini Parser object using "NewINIParser" function:
iniParser := iniparser.NewINIParser()

4. Obtain all sections in the form of a map of key-value pairs as follows:
 `iniMap := GetSections()`

5. Obtain sections names only in the form of a slice as follows:
 `sectionNames := GetSectionNames()`

6. Get configuration values by section and key as follows:
 `value, error := GetValue("section" , "key")`.

5. Update key values or set values for new keys as follows: 
`error := SetValue("section" ,"key" , "value")`.

6. Convert the INI map to string using the ToString function as follows:
 `strData := ToString()`.

7. Finally saving all changes back into an INI file as follows: 
error := `SaveToFile("/path/to/file.ini")`

### __Here is a guide of error messages returned by the parser when an error occurred:__

iniparser.ErrInvalidExtension
iniparser.ErrInvalidSyntax
iniparser.ErrKeyNotExist
iniparser.ErrSectionAlreadyExists
iniparser.ErrKeyNotFound
iniparser.ErrEmptyKey
iniparser.ErrValuesEmpty

### __Here is an example for a valid INI File:__

`[NETWORK]
host = example.com
port = 7878

[database]
host = localhost
port = 5432
username = postgres
password = password

[Email]
username = host_email.com
password = 12345

[LOCAL]
user = terry`

### __Here is a guide for testing:__

1. Install the needed dependencies as follows:
`go get -d ./....` 

2. Run all the tests as follows:  
`go test ./....`
If all tests pass on, the result should show that the tests were successful. If any tests fail, the output will indicate which tests failed.

3. Run the test for a method as follows: 
`go test -run TestMyFunction`












