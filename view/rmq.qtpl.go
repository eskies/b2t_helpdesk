// Code generated by qtc from "rmq.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line rmq.qtpl:1
package view

//line rmq.qtpl:1
import (
	"b2t_helpdesk/injector"
	"github.com/valyala/fasthttp"
)

//line rmq.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line rmq.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line rmq.qtpl:7
type RmqPage struct {
	CTX       *fasthttp.RequestCtx
	Dinjector *injector.Injector
}

//line rmq.qtpl:14
func (p *RmqPage) StreamTitle(qw422016 *qt422016.Writer) {
//line rmq.qtpl:14
	qw422016.N().S(`
	Message Queue Stats
`)
//line rmq.qtpl:16
}

//line rmq.qtpl:16
func (p *RmqPage) WriteTitle(qq422016 qtio422016.Writer) {
//line rmq.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line rmq.qtpl:16
	p.StreamTitle(qw422016)
//line rmq.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line rmq.qtpl:16
}

//line rmq.qtpl:16
func (p *RmqPage) Title() string {
//line rmq.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line rmq.qtpl:16
	p.WriteTitle(qb422016)
//line rmq.qtpl:16
	qs422016 := string(qb422016.B)
//line rmq.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line rmq.qtpl:16
	return qs422016
//line rmq.qtpl:16
}

//line rmq.qtpl:19
func (p *RmqPage) StreamBody(qw422016 *qt422016.Writer) {
//line rmq.qtpl:19
	qw422016.N().S(`
	<div class="card shadow mb-4">
		<div class="card-body">
			`)
//line rmq.qtpl:23
	stat, _ := p.Dinjector.QConn.CollectStats([]string{"outbox-b2t"})

//line rmq.qtpl:24
	qw422016.N().S(`
			<pre>`)
//line rmq.qtpl:25
	qw422016.E().S(stat.String())
//line rmq.qtpl:25
	qw422016.N().S(`</pre>
		</div>
	</div>

`)
//line rmq.qtpl:29
}

//line rmq.qtpl:29
func (p *RmqPage) WriteBody(qq422016 qtio422016.Writer) {
//line rmq.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line rmq.qtpl:29
	p.StreamBody(qw422016)
//line rmq.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line rmq.qtpl:29
}

//line rmq.qtpl:29
func (p *RmqPage) Body() string {
//line rmq.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
//line rmq.qtpl:29
	p.WriteBody(qb422016)
//line rmq.qtpl:29
	qs422016 := string(qb422016.B)
//line rmq.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
//line rmq.qtpl:29
	return qs422016
//line rmq.qtpl:29
}

//line rmq.qtpl:31
func (p *RmqPage) StreamModal(qw422016 *qt422016.Writer) {
//line rmq.qtpl:31
	qw422016.N().S(`
	
`)
//line rmq.qtpl:33
}

//line rmq.qtpl:33
func (p *RmqPage) WriteModal(qq422016 qtio422016.Writer) {
//line rmq.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line rmq.qtpl:33
	p.StreamModal(qw422016)
//line rmq.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line rmq.qtpl:33
}

//line rmq.qtpl:33
func (p *RmqPage) Modal() string {
//line rmq.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
//line rmq.qtpl:33
	p.WriteModal(qb422016)
//line rmq.qtpl:33
	qs422016 := string(qb422016.B)
//line rmq.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
//line rmq.qtpl:33
	return qs422016
//line rmq.qtpl:33
}

//line rmq.qtpl:35
func (p *RmqPage) StreamScript(qw422016 *qt422016.Writer) {
//line rmq.qtpl:35
	qw422016.N().S(`

`)
//line rmq.qtpl:37
}

//line rmq.qtpl:37
func (p *RmqPage) WriteScript(qq422016 qtio422016.Writer) {
//line rmq.qtpl:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line rmq.qtpl:37
	p.StreamScript(qw422016)
//line rmq.qtpl:37
	qt422016.ReleaseWriter(qw422016)
//line rmq.qtpl:37
}

//line rmq.qtpl:37
func (p *RmqPage) Script() string {
//line rmq.qtpl:37
	qb422016 := qt422016.AcquireByteBuffer()
//line rmq.qtpl:37
	p.WriteScript(qb422016)
//line rmq.qtpl:37
	qs422016 := string(qb422016.B)
//line rmq.qtpl:37
	qt422016.ReleaseByteBuffer(qb422016)
//line rmq.qtpl:37
	return qs422016
//line rmq.qtpl:37
}