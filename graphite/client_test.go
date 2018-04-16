/*
Copyright (C) 2018 Expedia Group.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package graphite

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGraphiteClient_AddEvent(t *testing.T) {

	want := `{"what":"what","tags":"tags","data":"data"}`
	ts := httptest.NewServer(testPostHandler(t, "/events/", want))
	defer ts.Close()

	client, err := DefaultGraphiteClient(ts.URL)

	if err != nil {
		t.Error(err)
	}

	err = client.AddEvent(GraphiteEvent{
		What:"what",
		Data:"data",
		Tags:"tags",
	})

	if err != nil {
		t.Fatalf("%v", err)
	}

}

func testPostHandler(t *testing.T, url, want string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != url {
			t.Errorf("Page Not Found\ngot  = %s\nwant = %s", r.RequestURI, url)
			w.WriteHeader(http.StatusNotFound)
		} else if r.Method != http.MethodPost {

			w.WriteHeader(http.StatusBadRequest)
		} else {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Error(err)
			}
			got := string(b)

			if got != want {
				w.WriteHeader(http.StatusBadRequest)
				t.Errorf("Bad request\ngot  = %s\nwant = %s", got, want)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}
