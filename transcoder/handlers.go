package transcoder

import "github.com/martini-contrib/render"

func Home() string {
	return "Hello abbita!"
}

func GetActions(file File, r render.Render) {
	if file.valid() {
		r.JSON(200, file.actions())
	} else {
		r.JSON(405, "File is not valid")
	}

}
