package events

import "testing"

func TestValidateTitle(t *testing.T) {
	title := "My new event"
	if !ValidateTitle(title) {
		t.Errorf("Title %v is not valid", title)
	}
}

func TestValidatePriority(t *testing.T) {
	priority := "low"
	err := Priority(priority).Validate()
	if err != nil {
		t.Errorf("Priority %v is not valid", priority)
	}

	priority = "high"
	err = Priority(priority).Validate()
	if err != nil {
		t.Errorf("Priority %v is not valid", priority)
	}

	priority = "medium"
	err = Priority(priority).Validate()
	if err != nil {
		t.Errorf("Priority %v is not valid", priority)
	}

	priority = "mew"
	err = Priority(priority).Validate()
	if err == nil {
		t.Errorf("Priority %v should be invalid", priority)
	}

	priority = ""
	err = Priority(priority).Validate()
	if err == nil {
		t.Errorf("Priority is empty and not valid")
	}

}
