package transcoder_test

import (
	"github.com/gophergala/abbita/transcoder"
	"github.com/modocache/gory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTranscoder(t *testing.T) {
	RegisterFailHandler(Fail)
	defineFactories()
	RunSpecs(t, "Transcoder Suite")
}

func defineFactories() {
	gory.Define("mp3", transcoder.File{},
		func(factory gory.Factory) {
			factory["Name"] = "Elly-Wamala-Ebinyumu-Ebyaffe.mp3"
			factory["Type"] = "audio/mp3"
		})
}
