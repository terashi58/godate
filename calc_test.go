package godate

import (
	"testing"
	"time"
)

type comparingTest struct {
	u              Date
	expectedResult bool
}

func TestDate_Equal(t *testing.T) {
	d := Today()

	tests := []comparingTest{
		{d, true},
		{d.Add(1, 0, 0), false},
		{d.Add(0, 1, 0), false},
		{d.Add(0, 0, 1), false},
		{d.Add(-1, 0, 0), false},
		{d.Add(0, -1, 0), false},
		{d.Add(0, 0, -1), false},
	}

	for i, test := range tests {
		if d.Equal(test.u) != test.expectedResult {
			t.Errorf("expected result is %v, but got %v. d=%v, u=%v, test idx=%d",
				test.expectedResult, d.Equal(test.u), d, test.u, i)
		}
	}
}

func TestDate_After(t *testing.T) {
	d := Today()

	tests := []comparingTest{
		{d.Add(-1, 1, 1), true},
		{d.Add(0, -1, 1), true},
		{d.Add(0, 0, -1), true},
		{d, false},
		{d.Add(1, -1, -1), false},
		{d.Add(0, 1, -1), false},
		{d.Add(0, 0, 1), false},
	}

	for i, test := range tests {
		if d.After(test.u) != test.expectedResult {
			t.Errorf("expected result is %v, but got %v. d=%v, u=%v, test idx=%d",
				test.expectedResult, d.After(test.u), d, test.u, i)
		}
	}
}

func TestDate_Before(t *testing.T) {
	d := Today()

	tests := []comparingTest{
		{d.Add(1, -1, -1), true},
		{d.Add(0, 1, -1), true},
		{d.Add(0, 0, 1), true},
		{d, false},
		{d.Add(-1, 1, 1), false},
		{d.Add(0, -1, 1), false},
		{d.Add(0, 0, -1), false},
	}

	for i, test := range tests {
		if d.Before(test.u) != test.expectedResult {
			t.Errorf("idx=%d: expected result is %v, but got %v. d=%v, u=%v",
				i, test.expectedResult, d.Before(test.u), d, test.u)
		}
	}
}

func TestDate_Add(t *testing.T) {
	d := New(2017, 10, 26)
	u := d.Add(1, 1, 1)
	if u.String() != "2018-11-27" {
		t.Errorf("got unexpected date string=%s", u.String())
	}

	u = d.Add(-1, -1, -1)
	if u.String() != "2016-09-25" {
		t.Errorf("got unexpected date string=%s", u.String())
	}
}

type subTest struct {
	d    Date
	u    Date
	days int64
}

func TestDate_Sub(t *testing.T) {
	tests := []subTest{
		// in same month
		{New(2017, 10, 1), New(2017, 10, 1), 0},
		{New(2017, 10, 2), New(2017, 10, 1), 1},
		{New(2017, 10, 1), New(2017, 10, 2), -1},
		// in same year that is not leap
		{New(2017, 11, 1), New(2017, 10, 1), 31},
		{New(2017, 12, 1), New(2017, 11, 1), 30},
		{New(2017, 12, 1), New(2017, 10, 1), 31 + 30},
		{New(2017, 10, 1), New(2017, 11, 1), -31},
		{New(2017, 11, 1), New(2017, 12, 1), -30},
		{New(2017, 10, 1), New(2017, 12, 1), -31 - 30},
		// over some years that are not leap
		{New(2018, 1, 1), New(2017, 1, 1), 365},
		{New(2019, 1, 1), New(2017, 1, 1), 365 * 2},
		{New(2017, 1, 1), New(2018, 1, 1), -365},
		{New(2017, 1, 1), New(2019, 1, 1), -365 * 2},
		// over a leap year
		{New(2021, 1, 1), New(2020, 1, 1), 366},
		{New(2021, 1, 1), New(2019, 1, 1), 366 + 365},
		{New(2020, 1, 1), New(2021, 1, 1), -366},
		{New(2019, 1, 1), New(2021, 1, 1), -366 - 365},
	}

	for i, test := range tests {
		if result := test.d.Sub(test.u); result != time.Duration(test.days)*Day {
			t.Errorf("idx=%d: expected value is %d, but got %d", i, test.days, result)
		}
	}
}
