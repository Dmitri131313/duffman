package fuzz

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/kmarkela/duffman/internal/pcollection"
)

func startWorker(wg *sync.WaitGroup, wq <-chan workUnit, tr *http.Transport) {

	for wu := range wq {

		if !wu.parBody {

			wu.r.Parameters.Get[wu.param] = wu.word

			endpoint := createEndpoint(wu.r.URL, wu.r.Parameters.Get)
			var r io.Reader = strings.NewReader(wu.r.Body)
			doRequest(endpoint, r, wu, tr)

			continue
		}

		body, _ := encodeBody(wu)
		endpoint := createEndpoint(wu.r.URL, wu.r.Parameters.Get)
		doRequest(endpoint, body, wu, tr)
	}

	wg.Done()
}

func createEndpoint(url string, par map[string]string) string {
	endpoint := fmt.Sprintf("%s?", url)
	for gk, gv := range par {
		endpoint = fmt.Sprintf("%s%s=%s&", endpoint, gk, gv)
	}

	return endpoint
}

func doRequest(endpoint string, body io.Reader, wu workUnit, tr *http.Transport) error {

	req, err := http.NewRequest(wu.r.Method, endpoint, body)
	if err != nil {
		return err
	}

	for k, v := range wu.r.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Transport: tr}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer res.Body.Close()

	return nil
}

func encodeBody(wu workUnit) (io.Reader, error) {

	wu.r.Parameters.Post[wu.param] = wu.word

	// encode Form
	if strings.HasPrefix(wu.r.ContentType, "application/x-www-form-urlencoded") {
		form := url.Values{}
		for k, v := range wu.r.Parameters.Post {
			form.Add(k, v)
		}
		return strings.NewReader(form.Encode()), nil
	}

	// encode json
	if wu.r.ContentType == "application/json" {
		b := pcollection.MarshalJSONBody(wu.r.Parameters.Post)
		return bytes.NewBuffer(b), nil
	}

	// TODO: encode multipart

	// unknown content type
	return strings.NewReader(wu.r.Body), fmt.Errorf("no encoder for: %s", wu.r.ContentType)
}
