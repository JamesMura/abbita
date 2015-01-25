package transcoder_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/gophergala/abbita/transcoder"
	"github.com/modocache/gory"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	var dbName string
	var session *DatabaseSession
	var server Server
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	BeforeEach(func() {
		dbName = "transcoder_test"
		session = NewSession(dbName)
		server = NewServer(session)

		recorder = httptest.NewRecorder()
	})

	AfterEach(func() {
		session.DB(dbName).DropDatabase()
	})

	Describe("POST /api/actions", func() {

		Context("when a file is present", func() {
			BeforeEach(func() {
				body, _ := json.Marshal(
					gory.Build("mp3"))
				request, _ = http.NewRequest("POST", "/api/actions", bytes.NewReader(body))
			})

			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns a list of available actions", func() {
				server.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).To(Equal("[{\"name\":\"Convert MP3 to WAV\",\"id\":\"mp3_wav\"}]"))
			})
		})

		Context("when no file is present", func() {
			BeforeEach(func() {
				request, _ = http.NewRequest("POST", "/api/actions", nil)
			})

			It("returns a status code of 405", func() {
				server.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(405))
			})
		})

	})

})
