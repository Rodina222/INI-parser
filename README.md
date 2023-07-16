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
```ini

import "github.com/codescalersinternships/INIParser-Rodina"
```

2. Load the configuration data as follows:
```ini

//from a file:

fileData, error := LoadFromFile(/path/to/file.ini)

//from a string:

stringData := LoadFromString([section1]\nkey=value\n[section2]\nkey=value\n..etc)
```

5. Create a new ini Parser object using "NewINIParser" function:
```ini
iniParser := iniparser.NewINIParser()
```

5. Obtain all sections in the form of a map of key-value pairs as follows:
```ini
iniMap := GetSections()
```

6. Obtain sections names only in the form of a slice as follows:
```ini
sectionNames := GetSectionNames()
```

7. Get configuration values by section and key as follows:
```ini
value, error := GetValue("section" , "key")
```

5. Update key values or set values for new keys as follows:
```ini
error := SetValue("section" ,"key" , "value")
```

6. Convert the INI map to string using the ToString function as follows:
```ini
strData := ToString()
```

7. Finally saving all changes back into an INI file as follows:
```ini
error := SaveToFile("/path/to/file.ini")
```

### __Here is a guide of error messages returned by the parser when an error occurred:__
```ini
iniparser.ErrInvalidExtension
iniparser.ErrInvalidSyntax
iniparser.ErrKeyNotExist
iniparser.ErrSectionAlreadyExists
iniparser.ErrKeyNotFound
iniparser.ErrEmptyKey 
iniparser.ErrValuesEmpty
```

### __Here is an example of a valid INI File:__

```ini
[NETWORK]
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
user
```


### __Here is a guide for testing:__

1. Install the needed dependencies as follows:
```ini
go get -d ./....
```

2. Run all the tests as follows: 
```ini
go test ./....
```
If all tests pass on, the result should show that the tests were successful as follows:
```ini
PASS
ok      github.com/codescalersinternships/INIParser-Rodina       0.002s
```
If any tests fail, the output will indicate which tests failed.

3. Run the test for a method as follows:
```ini
go test -run TestGetSections()
```
