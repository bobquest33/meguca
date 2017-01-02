// This file is automatically generated by qtc from "forms.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line forms.qtpl:1
package templates

//line forms.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line forms.qtpl:1
import "github.com/bakape/meguca/config"

//line forms.qtpl:2
import "github.com/bakape/meguca/lang"

// OwnedBoard renders a form for selecting one of several boards owned by the
// user

//line forms.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line forms.qtpl:6
func StreamOwnedBoard(qw422016 *qt422016.Writer, boards config.BoardTitles, lang map[string]string) {
	//line forms.qtpl:7
	if len(boards) != 0 {
		//line forms.qtpl:7
		qw422016.N().S(`<select name="boards" required>`)
		//line forms.qtpl:9
		for _, b := range boards {
			//line forms.qtpl:9
			qw422016.N().S(`<option value="`)
			//line forms.qtpl:10
			qw422016.N().S(b.ID)
			//line forms.qtpl:10
			qw422016.N().S(`">`)
			//line forms.qtpl:11
			streamformatTitle(qw422016, b.ID, b.Title)
			//line forms.qtpl:11
			qw422016.N().S(`</option>`)
			//line forms.qtpl:13
		}
		//line forms.qtpl:13
		qw422016.N().S(`</select><br><input type="submit" value="`)
		//line forms.qtpl:16
		qw422016.N().S(lang["submit"])
		//line forms.qtpl:16
		qw422016.N().S(`">`)
		//line forms.qtpl:17
	} else {
		//line forms.qtpl:18
		qw422016.N().S(lang["ownNoBoards"])
		//line forms.qtpl:18
		qw422016.N().S(`<br><br>`)
		//line forms.qtpl:21
	}
	//line forms.qtpl:21
	qw422016.N().S(`<input type="button" name="cancel" value="`)
	//line forms.qtpl:22
	qw422016.N().S(lang["cancel"])
	//line forms.qtpl:22
	qw422016.N().S(`"><div class="form-response admin"></div>`)
//line forms.qtpl:24
}

//line forms.qtpl:24
func WriteOwnedBoard(qq422016 qtio422016.Writer, boards config.BoardTitles, lang map[string]string) {
	//line forms.qtpl:24
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:24
	StreamOwnedBoard(qw422016, boards, lang)
	//line forms.qtpl:24
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:24
}

//line forms.qtpl:24
func OwnedBoard(boards config.BoardTitles, lang map[string]string) string {
	//line forms.qtpl:24
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:24
	WriteOwnedBoard(qb422016, boards, lang)
	//line forms.qtpl:24
	qs422016 := string(qb422016.B)
	//line forms.qtpl:24
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:24
	return qs422016
//line forms.qtpl:24
}

//line forms.qtpl:26
func streamformatTitle(qw422016 *qt422016.Writer, id, title string) {
	//line forms.qtpl:26
	qw422016.N().S(`/`)
	//line forms.qtpl:27
	qw422016.N().S(id)
	//line forms.qtpl:27
	qw422016.N().S(`/ -`)
	//line forms.qtpl:27
	qw422016.E().S(title)
//line forms.qtpl:28
}

//line forms.qtpl:28
func writeformatTitle(qq422016 qtio422016.Writer, id, title string) {
	//line forms.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:28
	streamformatTitle(qw422016, id, title)
	//line forms.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:28
}

//line forms.qtpl:28
func formatTitle(id, title string) string {
	//line forms.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:28
	writeformatTitle(qb422016, id, title)
	//line forms.qtpl:28
	qs422016 := string(qb422016.B)
	//line forms.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:28
	return qs422016
//line forms.qtpl:28
}

// BoardNavigation renders a board selection and search form

//line forms.qtpl:31
func StreamBoardNavigation(qw422016 *qt422016.Writer, lang map[string]string) {
	//line forms.qtpl:31
	qw422016.N().S(`<input type="text" class="full-width" name="search" placeholder="`)
	//line forms.qtpl:32
	qw422016.N().S(lang["search"])
	//line forms.qtpl:32
	qw422016.N().S(`"><br><form>`)
	//line forms.qtpl:35
	streamsubmit(qw422016, true, lang)
	//line forms.qtpl:35
	qw422016.N().S(`<br>`)
	//line forms.qtpl:37
	for _, b := range config.GetBoardTitles() {
		//line forms.qtpl:37
		qw422016.N().S(`<label><input type="checkbox" name="`)
		//line forms.qtpl:39
		qw422016.N().S(b.ID)
		//line forms.qtpl:39
		qw422016.N().S(`"><a href="/`)
		//line forms.qtpl:40
		qw422016.N().S(b.ID)
		//line forms.qtpl:40
		qw422016.N().S(`/" class="history">`)
		//line forms.qtpl:41
		streamformatTitle(qw422016, b.ID, b.Title)
		//line forms.qtpl:41
		qw422016.N().S(`</a><br></label>`)
		//line forms.qtpl:45
	}
	//line forms.qtpl:45
	qw422016.N().S(`</form>`)
//line forms.qtpl:47
}

//line forms.qtpl:47
func WriteBoardNavigation(qq422016 qtio422016.Writer, lang map[string]string) {
	//line forms.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:47
	StreamBoardNavigation(qw422016, lang)
	//line forms.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:47
}

//line forms.qtpl:47
func BoardNavigation(lang map[string]string) string {
	//line forms.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:47
	WriteBoardNavigation(qb422016, lang)
	//line forms.qtpl:47
	qs422016 := string(qb422016.B)
	//line forms.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:47
	return qs422016
//line forms.qtpl:47
}

// CreateBoard renders a the form for creating new boards

//line forms.qtpl:50
func StreamCreateBoard(qw422016 *qt422016.Writer, ln lang.Pack) {
	//line forms.qtpl:51
	streamtable(qw422016, specs["createBoard"], ln)
	//line forms.qtpl:52
	StreamCaptchaConfirmation(qw422016, ln)
//line forms.qtpl:53
}

//line forms.qtpl:53
func WriteCreateBoard(qq422016 qtio422016.Writer, ln lang.Pack) {
	//line forms.qtpl:53
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:53
	StreamCreateBoard(qw422016, ln)
	//line forms.qtpl:53
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:53
}

//line forms.qtpl:53
func CreateBoard(ln lang.Pack) string {
	//line forms.qtpl:53
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:53
	WriteCreateBoard(qb422016, ln)
	//line forms.qtpl:53
	qs422016 := string(qb422016.B)
	//line forms.qtpl:53
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:53
	return qs422016
//line forms.qtpl:53
}

// CaptchaConfirmation renders a confirmation form with an optional captcha

//line forms.qtpl:56
func StreamCaptchaConfirmation(qw422016 *qt422016.Writer, ln lang.Pack) {
	//line forms.qtpl:57
	streamcaptcha(qw422016, "confirmation", ln.UI)
	//line forms.qtpl:58
	streamsubmit(qw422016, true, ln.UI)
//line forms.qtpl:59
}

//line forms.qtpl:59
func WriteCaptchaConfirmation(qq422016 qtio422016.Writer, ln lang.Pack) {
	//line forms.qtpl:59
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:59
	StreamCaptchaConfirmation(qw422016, ln)
	//line forms.qtpl:59
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:59
}

//line forms.qtpl:59
func CaptchaConfirmation(ln lang.Pack) string {
	//line forms.qtpl:59
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:59
	WriteCaptchaConfirmation(qb422016, ln)
	//line forms.qtpl:59
	qs422016 := string(qb422016.B)
	//line forms.qtpl:59
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:59
	return qs422016
//line forms.qtpl:59
}

//line forms.qtpl:61
func streamcaptcha(qw422016 *qt422016.Writer, id string, lang map[string]string) {
	//line forms.qtpl:62
	conf := config.Get()

	//line forms.qtpl:63
	if !conf.Captcha {
		//line forms.qtpl:64
		return
		//line forms.qtpl:65
	}
	//line forms.qtpl:65
	qw422016.N().S(`<div class="captcha-container" id="captcha-`)
	//line forms.qtpl:66
	qw422016.N().S(id)
	//line forms.qtpl:66
	qw422016.N().S(`"><div id="adcopy-outer-`)
	//line forms.qtpl:67
	qw422016.N().S(id)
	//line forms.qtpl:67
	qw422016.N().S(`"><div id="adcopy-puzzle-image-`)
	//line forms.qtpl:68
	qw422016.N().S(id)
	//line forms.qtpl:68
	qw422016.N().S(`" class="captcha-image" title="`)
	//line forms.qtpl:68
	qw422016.N().S(lang["reloadCaptcha"])
	//line forms.qtpl:68
	qw422016.N().S(`"></div><div id="adcopy-puzzle-audio-`)
	//line forms.qtpl:69
	qw422016.N().S(id)
	//line forms.qtpl:69
	qw422016.N().S(`" class="hidden"></div><div id="adcopy-pixel-image-`)
	//line forms.qtpl:70
	qw422016.N().S(id)
	//line forms.qtpl:70
	qw422016.N().S(`" class="hidden"></div><div><span id="adcopy-instr-`)
	//line forms.qtpl:72
	qw422016.N().S(id)
	//line forms.qtpl:72
	qw422016.N().S(`" class="hidden"></span></div><input id="adcopy_response-`)
	//line forms.qtpl:74
	qw422016.N().S(id)
	//line forms.qtpl:74
	qw422016.N().S(`" name="adcopy_response" class="full-width" type="text" placeholder="`)
	//line forms.qtpl:74
	qw422016.N().S(lang["focusForCaptcha"])
	//line forms.qtpl:74
	qw422016.N().S(`" required><input type="hidden" name="adcopy_challenge" id="adcopy_challenge-`)
	//line forms.qtpl:75
	qw422016.N().S(id)
	//line forms.qtpl:75
	qw422016.N().S(`" hidden><a id="adcopy-link-refresh-`)
	//line forms.qtpl:76
	qw422016.N().S(id)
	//line forms.qtpl:76
	qw422016.N().S(`" class="hidden"></a><a id="adcopy-link-audio-`)
	//line forms.qtpl:77
	qw422016.N().S(id)
	//line forms.qtpl:77
	qw422016.N().S(`" class="hidden"></a><a id="adcopy-link-image-`)
	//line forms.qtpl:78
	qw422016.N().S(id)
	//line forms.qtpl:78
	qw422016.N().S(`" class="hidden"></a><a id="adcopy-link-info-`)
	//line forms.qtpl:79
	qw422016.N().S(id)
	//line forms.qtpl:79
	qw422016.N().S(`" class="hidden"></a><!-- Temporary, till better noscript support ><noscript><iframe src="https://api-secure.solvemedia.com/papi/challenge.noscript?k=`)
	//line forms.qtpl:82
	qw422016.N().S(conf.CaptchaPublicKey)
	//line forms.qtpl:82
	qw422016.N().S(`"height="260" width="350" frameborder="0"></iframe><br><input name="adcopy_challenge" type="text" required></noscript><--></div></div>`)
//line forms.qtpl:90
}

//line forms.qtpl:90
func writecaptcha(qq422016 qtio422016.Writer, id string, lang map[string]string) {
	//line forms.qtpl:90
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:90
	streamcaptcha(qw422016, id, lang)
	//line forms.qtpl:90
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:90
}

//line forms.qtpl:90
func captcha(id string, lang map[string]string) string {
	//line forms.qtpl:90
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:90
	writecaptcha(qb422016, id, lang)
	//line forms.qtpl:90
	qs422016 := string(qb422016.B)
	//line forms.qtpl:90
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:90
	return qs422016
//line forms.qtpl:90
}

// Form for inputting key-value map-like data

//line forms.qtpl:93
func streamkeyValueForm(qw422016 *qt422016.Writer, k, v string) {
	//line forms.qtpl:93
	qw422016.N().S(`<span><input type="text" class="map-field" value="`)
	//line forms.qtpl:95
	qw422016.E().S(k)
	//line forms.qtpl:95
	qw422016.N().S(`"><input type="text" class="map-field" value="`)
	//line forms.qtpl:96
	qw422016.E().S(v)
	//line forms.qtpl:96
	qw422016.N().S(`"><a class="map-remove">[X]</a><br></span>`)
//line forms.qtpl:102
}

//line forms.qtpl:102
func writekeyValueForm(qq422016 qtio422016.Writer, k, v string) {
	//line forms.qtpl:102
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:102
	streamkeyValueForm(qw422016, k, v)
	//line forms.qtpl:102
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:102
}

//line forms.qtpl:102
func keyValueForm(k, v string) string {
	//line forms.qtpl:102
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:102
	writekeyValueForm(qb422016, k, v)
	//line forms.qtpl:102
	qs422016 := string(qb422016.B)
	//line forms.qtpl:102
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:102
	return qs422016
//line forms.qtpl:102
}

// Form formatted as a table, with cancel and submit buttons

//line forms.qtpl:105
func streamtableForm(qw422016 *qt422016.Writer, specs []inputSpec, needCaptcha bool, ln lang.Pack) {
	//line forms.qtpl:106
	streamtable(qw422016, specs, ln)
	//line forms.qtpl:107
	if needCaptcha {
		//line forms.qtpl:108
		streamcaptcha(qw422016, "ajax", ln.UI)
		//line forms.qtpl:109
	}
	//line forms.qtpl:110
	streamsubmit(qw422016, true, ln.UI)
//line forms.qtpl:111
}

//line forms.qtpl:111
func writetableForm(qq422016 qtio422016.Writer, specs []inputSpec, needCaptcha bool, ln lang.Pack) {
	//line forms.qtpl:111
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:111
	streamtableForm(qw422016, specs, needCaptcha, ln)
	//line forms.qtpl:111
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:111
}

//line forms.qtpl:111
func tableForm(specs []inputSpec, needCaptcha bool, ln lang.Pack) string {
	//line forms.qtpl:111
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:111
	writetableForm(qb422016, specs, needCaptcha, ln)
	//line forms.qtpl:111
	qs422016 := string(qb422016.B)
	//line forms.qtpl:111
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:111
	return qs422016
//line forms.qtpl:111
}

// Render a map form for inputting map-like data

//line forms.qtpl:114
func streamrenderMap(qw422016 *qt422016.Writer, spec inputSpec, ln lang.Pack) {
	//line forms.qtpl:114
	qw422016.N().S(`<div class="map-form" name="`)
	//line forms.qtpl:115
	qw422016.N().S(spec.ID)
	//line forms.qtpl:115
	qw422016.N().S(`" title="`)
	//line forms.qtpl:115
	qw422016.N().S(ln.Forms[spec.ID][1])
	//line forms.qtpl:115
	qw422016.N().S(`">`)
	//line forms.qtpl:116
	for k, v := range spec.Val.(map[string]string) {
		//line forms.qtpl:117
		streamkeyValueForm(qw422016, k, v)
		//line forms.qtpl:118
	}
	//line forms.qtpl:118
	qw422016.N().S(`<a class="map-add">`)
	//line forms.qtpl:120
	qw422016.N().S(ln.UI["add"])
	//line forms.qtpl:120
	qw422016.N().S(`</a><br></div>`)
//line forms.qtpl:124
}

//line forms.qtpl:124
func writerenderMap(qq422016 qtio422016.Writer, spec inputSpec, ln lang.Pack) {
	//line forms.qtpl:124
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:124
	streamrenderMap(qw422016, spec, ln)
	//line forms.qtpl:124
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:124
}

//line forms.qtpl:124
func renderMap(spec inputSpec, ln lang.Pack) string {
	//line forms.qtpl:124
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:124
	writerenderMap(qb422016, spec, ln)
	//line forms.qtpl:124
	qs422016 := string(qb422016.B)
	//line forms.qtpl:124
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:124
	return qs422016
//line forms.qtpl:124
}

// Render submit and cancel buttons

//line forms.qtpl:127
func streamsubmit(qw422016 *qt422016.Writer, cancel bool, ln map[string]string) {
	//line forms.qtpl:127
	qw422016.N().S(`<input type="submit" value="`)
	//line forms.qtpl:128
	qw422016.N().S(ln["submit"])
	//line forms.qtpl:128
	qw422016.N().S(`">`)
	//line forms.qtpl:129
	if cancel {
		//line forms.qtpl:129
		qw422016.N().S(`<input type="button" name="cancel" value="`)
		//line forms.qtpl:130
		qw422016.N().S(ln["cancel"])
		//line forms.qtpl:130
		qw422016.N().S(`">`)
		//line forms.qtpl:131
	}
	//line forms.qtpl:131
	qw422016.N().S(`<div class="form-response admin"></div>`)
//line forms.qtpl:133
}

//line forms.qtpl:133
func writesubmit(qq422016 qtio422016.Writer, cancel bool, ln map[string]string) {
	//line forms.qtpl:133
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line forms.qtpl:133
	streamsubmit(qw422016, cancel, ln)
	//line forms.qtpl:133
	qt422016.ReleaseWriter(qw422016)
//line forms.qtpl:133
}

//line forms.qtpl:133
func submit(cancel bool, ln map[string]string) string {
	//line forms.qtpl:133
	qb422016 := qt422016.AcquireByteBuffer()
	//line forms.qtpl:133
	writesubmit(qb422016, cancel, ln)
	//line forms.qtpl:133
	qs422016 := string(qb422016.B)
	//line forms.qtpl:133
	qt422016.ReleaseByteBuffer(qb422016)
	//line forms.qtpl:133
	return qs422016
//line forms.qtpl:133
}
