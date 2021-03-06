package godate

import "testing"

func TestParse(t *testing.T) {
	d, err := Parse(RFC3339, "2017-10-20")
	if err != nil {
		t.Errorf("expected error not raised, but got=%v", err)
	}
	if d.String() != "2017-10-20" {
		t.Errorf("expected String() returns '2017-10-20', but got %s", d.String())
	}

	d, err = Parse("20060102", "20171020")
	if err != nil {
		t.Errorf("expected error not raised, but got=%v", err)
	}
	if d.String() != "2017-10-20" {
		t.Errorf("expected String() returns '2017-10-20', but got %s", d.String())
	}

	_, err = Parse("20161020", "2017-10-20")
	if err == nil {
		t.Errorf("expected error raised, bug got %v", err)
	}
}

func TestDate_Format(t *testing.T) {
	d := New(2017, 10, 27)
	if d.Format(RFC3339) != "2017-10-27" {
		t.Errorf("expected formatted godate.RFC3339, but got %s", d.Format(RFC3339))
	}
	if d.Format("2006/01/02") != "2017/10/27" {
		t.Errorf("expected formatted to '2006/01/02', but got %s", d.Format("2006/01/02"))
	}

	d = Date{}
	if d.Format(RFC3339) != "0001-01-01" {
		t.Errorf("expected zero date is formatted godate.RFC3339, but got %s", d.Format(RFC3339))
	}
}

func TestDate_String(t *testing.T) {
	d := New(2017, 10, 27)
	if d.String() != "2017-10-27" {
		t.Errorf("expected formatted godate.RFC3339, but got %s", d.String())
	}

	d = Date{}
	if d.Format(RFC3339) != "0001-01-01" {
		t.Errorf("expected zero date is formatted godate.RFC3339, but got %s", d.Format(RFC3339))
	}
}
