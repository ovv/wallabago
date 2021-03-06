package wallabago

import "testing"

func TestVersion(t *testing.T) {
	// expect string without quotes
	expected := "2.2.3"
	got, _ := Version(mockGetBodyOfURL)
	if expected != got {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func mockGetBodyOfURL(url string, httpMethod string, postData []byte) (string, error) {
	// return string with quotes
	return `"2.2.3"`, nil
}
