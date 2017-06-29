package trace

import (
	"testing"
	"bytes"
)

func TestNew(t * testing.T) {

	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newからの戻り値がnilです")
	} else {
		tracer.Trace("こんにちは,Traceパッケージ")
		if buf.String() != "こんにちは,Traceパッケージ\n" {
			t.Error("'%s'という誤った文字列が出力されました", buf.String())
		}
	}
}

func TestOff(t * testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("データ")
}