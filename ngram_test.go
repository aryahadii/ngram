package ngram

import (
	"testing"
)

type getTestStruct struct {
	text   string
	n      int
	result []string
	err    bool
}

type getFrequencyTestStruct struct {
	text   string
	n      int
	result map[string]int
	err    bool
}

func TestGet(t *testing.T) {
	tests := [...]getTestStruct{
		{"ABCDE", 1, []string{"A", "B", "C", "D", "E"}, false},
		{"ABCDE", 2, []string{"AB", "BC", "CD", "DE"}, false},
		{"ABCDE", 3, []string{"ABC", "BCD", "CDE"}, false},
		{"AB", 2, []string{"AB"}, false},
		{"ABCDE", 0, nil, true},
		{"AB", 3, nil, true},
		{"", 1, nil, true},
	}

	for index, test := range tests {
		res, err := Get(test.text, test.n)

		if (err != nil && !test.err) || (err == nil && test.err) {
			t.Errorf("Wrong error: %s\n", err)
			t.FailNow()
		}
		if err != nil {
			continue
		}

		if res == nil && test.result != nil {
			t.Errorf("Wrong result %s %s \n", res, test.result)
			t.FailNow()
		}
		if res != nil && test.result == nil {
			t.Errorf("Wrong result %s %s \n", res, test.result)
			t.FailNow()
		}

		if len(res) != len(test.result) {
			t.Errorf("Different lengths %s %s \n", res, test.result)
			t.FailNow()
		}

		for index := range res {
			if res[index] != test.result[index] {
				t.Errorf("Wrong result %s %s\n", res, test.result)
				t.FailNow()
			}
		}

		t.Logf("Test %d successed!\n", index)
	}
}

func TestGetFrequency(t *testing.T) {
	tests := [...]getFrequencyTestStruct{
		{"AB", 1, map[string]int{"A": 1, "B": 1}, false},
		{"ABC", 2, map[string]int{"AB": 1, "BC": 1}, false},
		{"ABBB", 2, map[string]int{"AB": 1, "BB": 2}, false},
		{"ABCDE", 0, nil, true},
		{"AB", 3, nil, true},
		{"", 1, nil, true},
	}

	for index, test := range tests {
		res, err := GetFrequency(test.text, test.n)

		if (err != nil && !test.err) || (err == nil && test.err) {
			t.Errorf("Wrong error: %v\n", err)
			t.FailNow()
		}
		if err != nil {
			continue
		}

		if res == nil && test.result != nil {
			t.Errorf("Wrong result %v %v \n", res, test.result)
			t.FailNow()
		}
		if res != nil && test.result == nil {
			t.Errorf("Wrong result %v %v \n", res, test.result)
			t.FailNow()
		}

		if len(res) != len(test.result) {
			t.Errorf("Different lengths %v %v \n", res, test.result)
			t.FailNow()
		}

		for key, val1 := range res {
			if val2, ok := test.result[key]; ok {
				if val1 != val2 {
					t.Errorf("Wrong result %v %v\n", res, test.result)
					t.FailNow()
				}
			} else {
				t.Errorf("Wrong result %v %v\n", res, test.result)
				t.FailNow()
			}
		}

		t.Logf("Test %d successed!\n", index)
	}
}
