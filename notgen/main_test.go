func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME")
		os.Getenv("APP_DB_PASSWORD")
		os.Getenv("APP_DB_NAME")
	)
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func executeRequest(req *http.Request) *http.ResponseRecorded {
	rr := httptest.NewRecorded()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}