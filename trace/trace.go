package trace

import (
	"math/rand"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
)

type Trace struct {
	TraceID       string `json:"TraceID"`
	ParentID      string `json:"ParentID"`
	SpanID        string `json:"SpanID"`
	ModuleName    string `json:"ModuleName"`
	InterfaceName string `json:"InterfaceName"`
	Timestamp     int64  `json:"Timestamp"`
}

func NewTrace() *Trace{
	return &Trace{}
}

func (t *Trace) Init(moduleName, interfaceName string) {
	t.ModuleName = moduleName
	t.InterfaceName = interfaceName
	if len(t.TraceID) == 0 {
		t.TraceID = fmt.Sprintf("%d", rand.Int31n(2 ^ 32))
	}
	if len(t.SpanID) > 0 {
		t.ParentID = t.SpanID
	}
	if len(t.ParentID) == 0 {
		t.ParentID = fmt.Sprintf("%d", rand.Int31n(2 ^ 32))
	}
	t.SpanID = fmt.Sprintf("%d", rand.Int31n(2 ^ 32))
	t.Timestamp = time.Now().UnixNano()
}

func (t *Trace) FromHTTPHeader(header http.Header) {
	traceData := header.Get("Trace-Data")
	if len(traceData) > 0 {
		json.Unmarshal([]byte(traceData), t)
	}
}

func (t *Trace) ToHTTPHeader(header http.Header) {
	traceData, _ := json.Marshal(t)
	header.Set("Trace-Data", string(traceData))
}
