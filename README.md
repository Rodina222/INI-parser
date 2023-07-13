# __INIParser Package__

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

2. Load the configuration data as follows:<br>
from a file:<br>
`fileData, error := LoadFromFile(/path/to/file.ini)`

 or from a string :<br>
 `stringData := LoadFromString([section1]\nkey=value\n[section2]\nkey=value\n..etc)`  <br>


5. Create a new ini Parser object using "NewINIParser" function:<br>

`iniParser := iniparser.NewINIParser()`

5. Obtain all sections in the form of a map of key-value pairs as follows:<br>
 `iniMap := GetSections()`

6. Obtain sections names only in the form of a slice as follows:<br>
 `sectionNames := GetSectionNames()`

7. Get configuration values by section and key as follows:<br>
 `value, error := GetValue("section" , "key")`.

5. Update key values or set values for new keys as follows: <br>
`error := SetValue("section" ,"key" , "value")`.

6. Convert the INI map to string using the ToString function as follows:<br>
 `strData := ToString()`.

7. Finally saving all changes back into an INI file as follows: <br>
`error := SaveToFile("/path/to/file.ini")`

### __Here is a guide of error messages returned by the parser when an error occurred:__

`iniparser.ErrInvalidExtension` <br>
`iniparser.ErrInvalidSyntax` <br>
`iniparser.ErrKeyNotExist` <br>
`iniparser.ErrSectionAlreadyExists` <br>
`iniparser.ErrKeyNotFound` <br>
`iniparser.ErrEmptyKey` <br>
`iniparser.ErrValuesEmpty` <br>

### __Here is an example of a valid INI File:__

[NETWORK] <br>
host = example.com <br>
port = 7878<br>

[database]<br>
host = localhost<br>
port = 5432<br>
username = postgres<br>
password = password<br>

[Email]<br>
username = host_email.com<br>
password = 12345<br>

[LOCAL]<br>
user = terry<br>


### __Here is a guide for testing:__

1. Install the needed dependencies as follows:<br>
`go get -d ./....` 

2. Run all the tests as follows:  <br>
`go test ./....`
<br>If all tests pass on, the result should show that the tests were successful. If any tests fail, the output will indicate which tests failed.

3. Run the test for a method as follows: <br>
`go test -run TestMyFunction`












