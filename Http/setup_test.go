package setup

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoints(t *testing.T) {

	app, client := Setup() // Create a new fiber app
	defer client.Close()
	defer app.Shutdown()

	///////////////////////////////////////////////////     USER      //////////////////////////////////////////////////////////

	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer([]byte(`{
        "name":"Dilara2",
		"email":"dilara@gmail.com",
		"password":"busifreyiaslabulamazsınız"
    }`)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1) // Perform the request
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusCreated { // Check the response status code
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.StatusCode)
	}

	expected := `{"message":"User created successfully"}` // Check the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("Expected response body '%s' but got '%s'", expected, string(body))
	}

	//////////////////////////////////////////////   TOPİC    //////////////////////////////////////////////

	req = httptest.NewRequest("POST", "/topic", bytes.NewBuffer([]byte(`{
        "topic":"topic1"
    }`)))
	req.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(req, -1) // Perform the request
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusCreated { // Check the response status code
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.StatusCode)
	}

	expected = `{"message":"Topic created successfully"}` // Check the response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("Expected response body '%s' but got '%s'", expected, string(body))
	}

	/////////////////////////////////////////////////     RESEARCH         ///////////////////////////////////////////////////////

	req = httptest.NewRequest("POST", "/topic/research", bytes.NewBuffer([]byte(`{
			"research_header":"RestApi2"
    }`)))
	req.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(req, -1) // Perform the request
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusCreated { // Check the response status code
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.StatusCode)
	}

	expected = `{"message":"Research created successfully"}` // Check the response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("Expected response body '%s' but got '%s'", expected, string(body))
	}
}
