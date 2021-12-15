package pkg

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"backend-record/pkg/model/dao"
)

func TestMain(m *testing.M){
	err:=dao.Init()
	if err != nil {
		os.Exit(500)
	}
	status := m.Run()
	os.Exit(status)

}
func TestServer(t *testing.T) {
	server := Server
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/read/tag/articles?tag=1_sampleTag1", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// ...
	assert.Equal(t, w.Body.String(), "{\"articles\":[{\"id\":\"1\",\"title\":\"Title_sample1\",\"imagePath\":\"ImageURL\",\"tags\":[\"1_sampleTag1\",\"1_sampleTag2\",\"1_sampleTag3\",\"1_sampleTag4\"]}]}")
}
