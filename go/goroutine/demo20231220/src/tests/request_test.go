package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"playground/goroutine/demo20231220/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	dir, err := os.Getwd()
	fmt.Printf("os.Getwd(): %v, Error: %v\n ", dir, err)
	fmt.Println("os.Args: ", os.Args)
}

func buildURL(path string) string {
	return fmt.Sprintf("http://localhost:8080%s", path)
}

func TestSubmit(t *testing.T) {
	assert.True(t, true)

	form := url.Values{}
	form.Add("message", util.CurFuncDesc())

	resp, err := http.PostForm(buildURL("/submit"), form)
	t.Logf("Response=%+v, Error=%+v", resp, err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	t.Logf("Body=%s, Error=%+v", body, err)
}
