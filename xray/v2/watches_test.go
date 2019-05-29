package v2

import (
	"encoding/json"
	"testing"
)

func TestWatchFilterValueWrapperTestUnmarshalJSON_stringValue(t *testing.T) {
	blob := `"application/deb"`

	var value WatchFilterValueWrapper
	if err := json.Unmarshal([]byte(blob), &value); err != nil {
		t.Errorf("Got the following error: %s", err.Error())
	}

	if value.Key != nil {
		t.Errorf("Expected key to be nil but got: %s", *value.Key)
	}

	if *value.Value != "application/deb" {
		t.Errorf("Expected value to be 'application/deb' but got: %s", *value.Value)
	}

	if value.IsPropertyFilter {
		t.Errorf("Expected IsPropertyFilter to be false but got true")
	}

}

func TestWatchFilterValueWrapperTestUnmarshalJSON_objectValue(t *testing.T) {
	blob := `{"key": "prop_key","value": "prop_value"}`

	var value WatchFilterValueWrapper
	if err := json.Unmarshal([]byte(blob), &value); err != nil {
		t.Errorf("Got the following error: %s", err.Error())
	}

	if *value.Key != "prop_key" {
		t.Errorf("Expected key to be 'prop_key' but got: %s", *value.Key)
	}

	if *value.Value != "prop_value" {
		t.Errorf("Expected value to be 'prop_value' but got: %s", *value.Value)
	}

	if !value.IsPropertyFilter {
		t.Errorf("Expected IsPropertyFilter to be true but got false")
	}
}

func TestWatchFilterValueWrapperTestMarshalJSON_stringValue(t *testing.T) {
	v := "application/deb"

	value := WatchFilterValueWrapper{
		WatchFilterValue: WatchFilterValue{
			Value: &v,
		},
		IsPropertyFilter: false,
	}

	result, err := json.Marshal(value)

	if err != nil {
		t.Errorf("Got the following error: %s", err.Error())
	}

	if string(result) != "\"application/deb\"" {
		t.Errorf("Expected result to be 'application/deb' but got: %s", string(result))
	}

}

func TestWatchFilterValueWrapperTestMarshalJSON_objectValue(t *testing.T) {
	k := "prop_key"
	v := "prop_value"

	value := WatchFilterValueWrapper{
		WatchFilterValue: WatchFilterValue{
			Key:   &k,
			Value: &v,
		},
		IsPropertyFilter: true,
	}

	data, err := json.Marshal(value)

	if err != nil {
		t.Errorf("Got the following error: %s", err.Error())
	}

	var result WatchFilterValue

	err = json.Unmarshal(data, &result)

	if err != nil {
		t.Errorf("Got the following error when attempting to unmarshal: %s", err.Error())
	}

	if *result.Key != k {
		t.Errorf("Expected key to equal %s but got: %s", k, *result.Key)
	}

	if *result.Value != v {
		t.Errorf("Expected value to equal %s but got: %s", v, *result.Value)
	}
}
