// Code generated by qtc from "beranda.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line beranda.qtpl:1
package view

//line beranda.qtpl:1
import "github.com/valyala/fasthttp"

//line beranda.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line beranda.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line beranda.qtpl:4
type MainPage struct {
	CTX *fasthttp.RequestCtx
}

//line beranda.qtpl:10
func (p *MainPage) StreamTitle(qw422016 *qt422016.Writer) {
//line beranda.qtpl:10
	qw422016.N().S(`
	Beranda
`)
//line beranda.qtpl:12
}

//line beranda.qtpl:12
func (p *MainPage) WriteTitle(qq422016 qtio422016.Writer) {
//line beranda.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line beranda.qtpl:12
	p.StreamTitle(qw422016)
//line beranda.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line beranda.qtpl:12
}

//line beranda.qtpl:12
func (p *MainPage) Title() string {
//line beranda.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
//line beranda.qtpl:12
	p.WriteTitle(qb422016)
//line beranda.qtpl:12
	qs422016 := string(qb422016.B)
//line beranda.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
//line beranda.qtpl:12
	return qs422016
//line beranda.qtpl:12
}

//line beranda.qtpl:15
func (p *MainPage) StreamBody(qw422016 *qt422016.Writer) {
//line beranda.qtpl:15
	qw422016.N().S(`
	<div>
		Some info about you:<br/>
		IP: <b>`)
//line beranda.qtpl:18
	qw422016.E().S(p.CTX.RemoteIP().String())
//line beranda.qtpl:18
	qw422016.N().S(`</b><br/>
		User-Agent: <b>`)
//line beranda.qtpl:19
	qw422016.E().Z(p.CTX.UserAgent())
//line beranda.qtpl:19
	qw422016.N().S(`</b><br/>
	</div>
`)
//line beranda.qtpl:21
}

//line beranda.qtpl:21
func (p *MainPage) WriteBody(qq422016 qtio422016.Writer) {
//line beranda.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line beranda.qtpl:21
	p.StreamBody(qw422016)
//line beranda.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line beranda.qtpl:21
}

//line beranda.qtpl:21
func (p *MainPage) Body() string {
//line beranda.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
//line beranda.qtpl:21
	p.WriteBody(qb422016)
//line beranda.qtpl:21
	qs422016 := string(qb422016.B)
//line beranda.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
//line beranda.qtpl:21
	return qs422016
//line beranda.qtpl:21
}

//line beranda.qtpl:23
func (p *MainPage) StreamModal(qw422016 *qt422016.Writer) {
//line beranda.qtpl:23
	qw422016.N().S(`
	
`)
//line beranda.qtpl:25
}

//line beranda.qtpl:25
func (p *MainPage) WriteModal(qq422016 qtio422016.Writer) {
//line beranda.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line beranda.qtpl:25
	p.StreamModal(qw422016)
//line beranda.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line beranda.qtpl:25
}

//line beranda.qtpl:25
func (p *MainPage) Modal() string {
//line beranda.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line beranda.qtpl:25
	p.WriteModal(qb422016)
//line beranda.qtpl:25
	qs422016 := string(qb422016.B)
//line beranda.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line beranda.qtpl:25
	return qs422016
//line beranda.qtpl:25
}

//line beranda.qtpl:27
func (p *MainPage) StreamScript(qw422016 *qt422016.Writer) {
//line beranda.qtpl:27
	qw422016.N().S(`
	
`)
//line beranda.qtpl:29
}

//line beranda.qtpl:29
func (p *MainPage) WriteScript(qq422016 qtio422016.Writer) {
//line beranda.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line beranda.qtpl:29
	p.StreamScript(qw422016)
//line beranda.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line beranda.qtpl:29
}

//line beranda.qtpl:29
func (p *MainPage) Script() string {
//line beranda.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
//line beranda.qtpl:29
	p.WriteScript(qb422016)
//line beranda.qtpl:29
	qs422016 := string(qb422016.B)
//line beranda.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
//line beranda.qtpl:29
	return qs422016
//line beranda.qtpl:29
}