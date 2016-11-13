package fruit_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/jwfriese/omgfruitapi/fruit"
	"github.com/jwfriese/omgfruitapi/fruit/fruitfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("/fruit", func() {
	var (
		fakeFruitSource *fruitfakes.FakeFruitSource
		testServer      *httptest.Server
		testClient      *http.Client
		baseURL         string
	)

	BeforeEach(func() {
		fakeFruitSource = new(fruitfakes.FakeFruitSource)
		handler := fruit.GetFruitHandler(fakeFruitSource)
		server := mux.NewRouter()
		server.Handle("/fruit", handler).Methods("GET")
		testServer = httptest.NewServer(server)
		baseURL = testServer.URL
		testClient = http.DefaultClient
	})

	AfterEach(func() {
		testServer.Close()
	})

	Describe("GET", func() {
		var (
			response *http.Response
			err      error
			body     []byte
		)

		BeforeEach(func() {
			fakeFruitSource.GetNextFruitStub = func() (string, string, io.Reader) {
				imageData := bytes.NewBuffer([]byte(`image-data`))
				return "turtle fruit", "it's a fruit like a turtle", imageData
			}

			url := fmt.Sprintf("%v/%v", baseURL, "/fruit")
			response, err = http.Get(url)

			if err != nil {
				log.Fatal(err)
			}

			body, err = ioutil.ReadAll(response.Body)
			response.Body.Close()

			if err != nil {
				log.Fatal(err)
			}
		})

		It("returns a 200", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns the fruit information from the FruitSource", func() {
			Expect(body).To(MatchJSON([]byte(`{"name":"turtle fruit","description":"it's a fruit like a turtle","image":"image-data"}`)))
		})
	})
})
