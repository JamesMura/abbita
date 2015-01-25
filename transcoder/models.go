package transcoder

type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Action struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (file *File) valid() bool {
	return len(file.Name) > 0 &&
		len(file.Type) > 0
}

func (file *File) actions() []Action {
	mp32Wav := Action{Name: "Convert MP3 to WAV", Id: "mp3_wav"}
	mp3Actions := []Action{
		mp32Wav,
	}
	return mp3Actions
}
