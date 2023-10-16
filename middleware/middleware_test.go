package middleware

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func TestEditorRegex(t *testing.T) {
	if matches("GET", "^(POST|PUT|PATCH)$") != false {
		t.Error("not match")
	}
	if matches("POST", "^(POST|PUT|PATCH)$") != true {
		t.Error("not match")
	}
	if matches("PUT", "^(POST|PUT|PATCH)$") != true {
		t.Error("not match")
	}
	if matches("PATCH", "^(POST|PUT|PATCH)$") != true {
		t.Error("not match")
	}
}
