package trace

import (
	"io"
	"fmt"
)

//traceはコード内の出来事を記録できるオブジェクトを　表すインターフェース
// 先頭が大文字だからこの方は公開される
type Tracer interface {
	Trace(...interface{} )
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct {}
func (t *nilTracer) Trace(a ...interface{} ) {}

//offはTraceメソッドの呼び出しを無視するTracerを返します
func Off() Tracer {
	return &nilTracer{}
}