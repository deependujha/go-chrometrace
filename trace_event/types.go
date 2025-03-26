package trace_event

// TraceEventType represents the type of a trace event.
type TraceEventType string

const (
	// Duration trace events
	Begin TraceEventType = "B"
	End   TraceEventType = "E"

	// Complete trace event
	Complete TraceEventType = "X"

	// Instant trace event
	Instant TraceEventType = "I"

	// Counter trace event
	Counter TraceEventType = "C"

	// Async trace events
	NestableAsyncBegin   TraceEventType = "b"
	NestableAsyncEnd     TraceEventType = "e"
	NestableAsyncInstant TraceEventType = "n"

	// Flow trace events
	FlowBegin TraceEventType = "s"
	FlowStep  TraceEventType = "t"
	FlowEnd   TraceEventType = "f"

	// Metadata trace events
	Metadata TraceEventType = "M"

	// Sample trace event
	Sample TraceEventType = "P"

	// Object trace events
	CreateObject   TraceEventType = "N"
	SnapshotObject TraceEventType = "O"
	DeleteObject   TraceEventType = "D"

	// Memory dump trace events
	MemoryDumpGlobal TraceEventType = "V"
	MemoryDump       TraceEventType = "v"

	// Mark trace event
	Mark TraceEventType = "R"

	// Clock sync event
	CLOCK_SYNC TraceEventType = "c"
)
