package main

import (
	"reflect"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	want := "[NETWORK]\n" + "host= example.com\n port = 7878\n\n" + "[database]\n" + "host = localhost\n port = 5432\n username = postgres \n password = password\n\n" + "[Email]\n" + "username= host_email.com \n password=12345\n\n" + ";here is a comment/n" + " [LOCAL]\n user = terry"

	got, error := LoadFromFile("config.ini")

	if error != nil {
		t.Fatalf("Error: %v", error)
		return
	}

	if got != want {

		t.Errorf("Return value is not equal to the expected value!.\n Expected: %+v\nActual: %+v", want, got)

	}

}

func TestGetSections(t *testing.T) {

	want := INI_Parser{
		"Email": map[string]string{
			"username": "host_email.com",
			"password": "12345",
		},
		"LOCAL": map[string]string{
			"user": "terry",
		},
		"database": map[string]string{
			"host":     "localhost",
			"port":     "5432",
			"username": "postgres",
			"password": "password",
		},
		"NETWORK": map[string]string{
			"host": "example.com",
			"port": "7878",
		},
	}
	ini_data, error := LoadFromFile("config.ini")

	got, error := GetSections(ini_data)

	if error != nil {
		t.Fatalf("Error: %v", error)
		return
	}

	if !(reflect.DeepEqual(got, want)) {

		t.Errorf("Return value is not equal to the expected value!.\n Expected: %+v\nActual: %+v", want, got)

	}

}

func TestGetSectionNames(t *testing.T) {

	want := make([]string, 0)
	want = append(want, "database", "Email", "LOCAL", "NETWORK")

	ini_data, error := LoadFromFile("config.ini")

	Parser, error := GetSections(ini_data)

	if error != nil {
		t.Fatalf("Error: %v", error)
		return
	}

	got := GetSectionNames(Parser)

	if !(reflect.DeepEqual(got, want)) {

		t.Errorf("Return value is not equal to the expected value!.\n Expected: %+v\nActual: %+v", want, got)

	}

}

func TestGetValue(t *testing.T) {

	want := "12345"

	ini_data, error := LoadFromFile("config.ini")

	Parser, error := GetSections(ini_data)

	if error != nil {
		t.Fatalf("Error: %v", error)
		return
	}

	got, err := GetValue(Parser, "Email", "password")

	if err != nil {
		t.Fatalf("Error: %v", error)
		return
	}

	if got != want {

		t.Errorf("Return value is not equal to the expected value!.\n Expected: %+v\nActual: %+v", want, got)

	}

}

func TestSetValue(t *testing.T) {

	want := "ex.com"

	ini_data, error := LoadFromFile("config.ini")

	Parser, error := GetSections(ini_data)

	if error != nil {
		t.Fatalf("Error: %v", error)
		return
	}

	err := SetValue(&Parser, "NETWORK", "host", "ex.com")

	if err != nil {
		t.Fatalf("Error: %v", error)
		return
	}

	got := Parser["NETWORK"]["host"]

	if got != want {

		t.Errorf("Return value is not equal to the expected value!.\n Expected: %+v\nActual: %+v", want, got)

	}

}
