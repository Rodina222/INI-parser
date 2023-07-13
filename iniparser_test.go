package iniparser

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const iniValidFormat = `
[NETWORK] 
host= example.com
port = 7878

[database]
host = localhost
port = 5432
username = postgres
password = password

[Email]
username= host_email.com
password=12345

[LOCAL]
user = terry`

const iniInvalidFormat = `
[NETWORK]
host = example.com
port = 7878

[database]
host = localhost
port = 5432
username = postgres
password
= password

[Email
username= host_email.com
password=12345

;here is a comment
[LOCAL]
user = terry`

func TestLoadFromReader(t *testing.T) {

	parser := NewINIParser()

	t.Parallel()

	want := INIParser{
		sections: map[string]INISection{
			"Email": {
				"username": "host_email.com",
				"password": "12345",
			},
			"LOCAL": {
				"user": "terry",
			},
			"database": {
				"host":     "localhost",
				"port":     "5432",
				"username": "postgres",
				"password": "password",
			},
			"NETWORK": {
				"host": "example.com",
				"port": "7878",
			},
		},
	}

	t.Run("valid ini file format ", func(t *testing.T) {

		err := parser.LoadFromString(iniValidFormat)

		assert.NoError(t, err)

		got := parser.sections

		if !(reflect.DeepEqual(want.sections, got)) {

			t.Errorf("Return value is not equal to the expected value\n Expected: %+v\nActual: %+v", want.sections, got)

		}

	})

	t.Run("invalid ini file format ", func(t *testing.T) {

		err := parser.LoadFromString(iniInvalidFormat)

		assert.Error(t, err)

	})

}

func TestLoadFromFile(t *testing.T) {

	parser := NewINIParser()

	t.Parallel()

	t.Run("valid file extention", func(t *testing.T) {

		dir := t.TempDir()

		filePath := filepath.Join(dir, "config.ini")

		err := os.WriteFile(filePath, make([]byte, 0), 0644)

		assert.NoError(t, err)

		err = parser.LoadFromFile(filePath)

		assert.NoError(t, err)

	})

	t.Run("invalid file extention", func(t *testing.T) {

		dir := t.TempDir()

		filePath := filepath.Join(dir, "config.txt")

		err := os.WriteFile(filePath, make([]byte, 0), 0644)

		assert.NoError(t, err)

		err = parser.LoadFromFile(filePath)

		assert.Error(t, err)

	})

}

func TestLoadFromString(t *testing.T) {

	parser := NewINIParser()

	t.Parallel()

	t.Run("valid ini file format ", func(t *testing.T) {

		err := parser.LoadFromString(iniValidFormat)

		assert.NoError(t, err)

	})

	t.Run("invalid ini file format ", func(t *testing.T) {

		err := parser.LoadFromString(iniInvalidFormat)

		assert.Error(t, err)

	})

}

func TestGetSections(t *testing.T) {
	parser := NewINIParser()

	want := INIParser{
		sections: map[string]INISection{
			"Email": {
				"username": "host_email.com",
				"password": "12345",
			},
			"LOCAL": {
				"user": "terry",
			},
			"database": {
				"host":     "localhost",
				"port":     "5432",
				"username": "postgres",
				"password": "password",
			},
			"NETWORK": {
				"host": "example.com",
				"port": "7878",
			},
		},
	}

	err := parser.LoadFromString(iniValidFormat)

	assert.NoError(t, err)

	got := parser.GetSections()

	if !(reflect.DeepEqual(got, want.sections)) {

		t.Errorf("return value is not equal to the expected value \n Expected: %+v\nActual: %+v", want.sections, got)

	}

}

func TestGetSectionNames(t *testing.T) {

	parser := NewINIParser()

	t.Parallel()

	t.Run("section names for an empty map ", func(t *testing.T) {

		sections := parser.GetSectionNames()

		assert.Equal(t, 0, len(sections), "got %q want %q", sections, []string{})

	})

	t.Run("section names for a non empty map ", func(t *testing.T) {

		err := parser.LoadFromString(iniValidFormat)

		assert.NoError(t, err)

		want := []string{"database", "Email", "LOCAL", "NETWORK"}

		got := parser.GetSectionNames()

		sort.Strings(got)
		sort.Strings(want)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("returned value is not equal to the expected value \n Expected: %+v\nActual: %+v", want, got)
		}

	})

}

func TestGetValue(t *testing.T) {

	parser := NewINIParser()

	t.Parallel()

	t.Run("input section name  or key is empty ", func(t *testing.T) {

		_, err := parser.GetValue("", "password")

		assert.Equal(t, ErrValuesEmpty, err, "want %q but got %q", ErrValuesEmpty, err)
	})

	t.Run("get value from non existent section", func(t *testing.T) {

		_, err := parser.GetValue("company", "password")

		assert.Equal(t, ErrSectionNotFound, err, "want %q but got %q", ErrSectionNotFound, err)
	})

	t.Run("get value for a non existent key", func(t *testing.T) {

		_, err := parser.GetValue("Email", "pass")

		assert.Equal(t, ErrKeyNotFound, err, "want %q but got %q", ErrKeyNotFound, err)
	})

	t.Run("get value for existent section and key", func(t *testing.T) {

		err := parser.LoadFromString(iniValidFormat)

		assert.NoError(t, err)

		want := "12345"

		got, _ := parser.GetValue("Email", "password")

		assert.Equal(t, want, got, "got %q want %q", got, want)
	})

}

func TestSetValue(t *testing.T) {

	parser := NewINIParser()

	t.Parallel()

	t.Run("input section name  or key is empty ", func(t *testing.T) {

		err := parser.SetValue("NETWORK", "", "ex.com")

		assert.Equal(t, ErrValuesEmpty, err, "want %q but got %q", ErrValuesEmpty, err)
	})

	t.Run("set value for a key in an existent section ", func(t *testing.T) {

		err := parser.SetValue("NETWORK", "host", "ex.com")

		assert.NoError(t, err)
	})

	t.Run("set value for a key in a non existent section ", func(t *testing.T) {

		err := parser.SetValue("Information", "host", "ex.com")

		assert.NoError(t, err)
	})

}

func TestString(t *testing.T) {

	parser := NewINIParser()

	t.Parallel()

	str := `[NETWORK]
	host = example.com
	port = 7878
	`

	// Remove all occurrences of \t from the string
	want := strings.ReplaceAll(str, "\t", "")

	err := parser.LoadFromString(want)

	assert.NoError(t, err)

	got := parser.String()

	// Sort the keys of the section of the expected string
	wantLines := strings.Split(want, "\n")
	sort.Strings(wantLines[1:3])
	want = strings.Join(wantLines, "\n")

	assert.Equal(t, want, got, "got %q want %q", got, want)


}

func TestSaveToFile(t *testing.T) {

	parser := NewINIParser()

	t.Parallel()

	t.Run("save to invalid file extention", func(t *testing.T) {

		dir := t.TempDir()

		filePath := filepath.Join(dir, "config.txt")

		err := parser.SaveToFile(filePath)

		assert.Error(t, err)

	})

	t.Run("save to valid file extention", func(t *testing.T) {

		dir := t.TempDir()

		filePath := filepath.Join(dir, "config.ini")

		err := parser.SaveToFile(filePath)

		assert.NoError(t, err)

	})

}
