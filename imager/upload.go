// Package imager handles image, video, etc. upload requests and processing.
package imager

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/bakape/meguca/config"
	"github.com/bakape/meguca/server/websockets"
	"github.com/bakape/meguca/types"
)

// Supported file formats
const (
	jpeg = iota
	png
	gif
	webm
	pdf
	svg
	mp4
	mp3
	ogg
)

// Map of oficial MIME types to the extension representations we deal with
var mimeTypes = map[string]uint8{
	"image/jpeg":      jpeg,
	"image/png":       png,
	"image/gif":       gif,
	"video/webm":      webm,
	"application/pdf": pdf,
}

// NewImageUpload  handles the clients' image (or other file) upload request
func NewImageUpload(res http.ResponseWriter, req *http.Request) {
	// Limit data received to the maximum uploaded file size limit
	conf := config.Get()
	req.Body = http.MaxBytesReader(res, req.Body, conf.Images.Max.Size)

	defer func() {
		if err := req.MultipartForm.RemoveAll(); err != nil {
			log.Printf("Error removing temporary files: %s\n", err)
		}
	}()

	head := res.Header()
	head.Set("Content-Type", "text/html; charset=UTF-8")
	head.Set("Access-Control-Allow-Origin", conf.HTTP.Origin)

	_, _, err := parseUploadForm(req)
	if err != nil {
		passError(res, req, err, 400)
		return
	}

	file, _, err := req.FormFile("image")
	if err != nil {
		passError(res, req, err, 400)
		return
	}
	defer file.Close()

	// img.Image.Imgnm = fileHeader.Filename
	//
	// fileType, err := detectFileType(file)
	// if err != nil {
	// 	passError(res, req, err, 400)
	// 	return
	// }
	// img.fileType = fileType
	//
	// if err := processFile(file, img); err != nil {
	// 	passError(res, req, err, 400)
	// } else {
	//
	// 	// TODO: Call a method on the client to allocate the image.
	//
	// }
}

// Pass error message to client and log server-side
func passError(
	res http.ResponseWriter,
	req *http.Request,
	err error,
	code int,
) {
	text := err.Error()
	http.Error(res, text, code)
	log.Printf("Upload error: %s : %s\n", req.RemoteAddr, text)
}

// Parse and validate the form of the upload request
func parseUploadForm(req *http.Request) (
	clientID string, spoiler uint8, err error,
) {
	length, err := strconv.ParseInt(req.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return
	}
	if length > config.Get().Images.Max.Size {
		err = errors.New("File too large")
		return
	}

	err = req.ParseMultipartForm(0)
	if err != nil {
		return
	}

	clientID = req.FormValue("id")
	if clientID == "" {
		err = errors.New("No client ID specified")
		return
	}
	if !websockets.Clients.Has(clientID) {
		err = fmt.Errorf("Bad client ID: %s", clientID)
		return
	}
	spoiler, err = extractSpoiler(req)

	return
}

// Extracts and validates a spoiler number from the form
func extractSpoiler(req *http.Request) (sp uint8, err error) {
	// Read the spoiler the client had chosen for the image, if any
	if unparsed := req.FormValue("spoiler"); unparsed != "" {
		var unconverted int
		unconverted, err = strconv.Atoi(unparsed)
		sp = uint8(unconverted)
		if !(err == nil && isValidSpoiler(sp)) {
			err = fmt.Errorf("Invalid spoiler ID: %s", unparsed)
		}
	}
	return
}

// Confirms a spoiler exists in configuration
func isValidSpoiler(id uint8) bool {
	for _, valid := range config.Get().Images.Spoilers {
		if id == valid {
			return true
		}
	}
	return false
}

// detectFileType detects if the upload is of a supported file type, by reading
// its first 512 bytes. OGG and MP4 are also cheked to contain HTML5 supported
// video and audio streams.
func detectFileType(file io.Reader) (uint8, error) {
	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		return 0, err
	}
	mimeType := http.DetectContentType(buf)
	mime, ok := mimeTypes[mimeType]
	if !ok {
		switch {
		case detectSVG(buf):
			return svg, nil
		case detectMP3(buf):
			return mp3, nil
		default:
			is, err := detectCompatibleMP4(buf)
			if is {
				return mp4, err
			}
			is, err = detectCompatibleOGG(buf)
			if is {
				return ogg, err
			}
			return 0, fmt.Errorf("Unsupported mime type: %s", mimeType)
		}
	}
	return mime, nil
}

// TODO: Waiting on Soreil

func detectSVG(buf []byte) bool {
	return false
}

func detectMP3(buf []byte) bool {
	return false
}

func detectCompatibleOGG(buf []byte) (bool, error) {
	return false, nil
}

func detectCompatibleMP4(buf []byte) (bool, error) {
	return false, nil
}

// Delegate the processing of the file to an apropriate function by file type
func processFile(file io.ReadSeeker, postID int64, img *types.ProtoImage) (
	io.Reader, io.Reader, error,
) {
	switch img.FileType {
	case webm:
		return processWebm(file, postID, img)
	case jpeg, png, gif:
		return processImage(file, postID, img)
	default:
		return nil, nil, fmt.Errorf("File type slipped in: %d", img.FileType)
	}
}
