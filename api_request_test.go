package NetEaseAPI

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestFakeSongInfo(t *testing.T) {
	jsonPath := filepath.Join("testdata", t.Name()+".json")
	json, err := ioutil.ReadFile(jsonPath)
	assert.Nil(t, err, "could not read test json: "+jsonPath)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/linux/forward", r.RequestURI)

		body, err := ioutil.ReadAll(r.Body)
		assert.Nil(t, err, "should not fail read body")
		const expectedBody = "eparams=A0D9583F4C5FF68DE851D2893A49DE98C1A25152CD34D35EBE89AC" +
			"391CA899FA6D3428B98C3E130983C70BE4753892E92DEF1482730790E365C5F7EE3DD472FCCFA97" +
			"2FA5EA3D6247CD6247C8198CB87EC48C52715001F5101580F2A27B5DF824146577B207908A96F22" +
			"361653C12A70"
		assert.Equal(t, expectedBody, string(body))

		_, _ = w.Write(json)
	}))
	defer ts.Close()

	api := New()
	api.BaseURL = ts.URL
	song, _ := api.Song(34341360)
	assert.NotNil(t, song.Songs, "response json be parsed correctly from test server")
}
