package main

import (
	"fmt"
	pq "github.com/gnanderson/pqueue"
	.  "container/heap"
)

const (
	EMERG   = iota  // Emergency: system is unusable
	ALERT			// Alert: action must be taken immediately
	CRIT			// Critical: critical conditions
	ERR				// Error: error conditions
	WARN			// Warning: warning conditions
	NOTICE			// Notice: normal but significant condition
	INFO			// Informational: informational messages
	DEBUG			// Debug: debug messages
)

type Alert struct {
	priority  int
	message   string
}

func (a *Alert) Priority() int {
	return a.priority
}

func (a *Alert) String() string {
	return fmt.Sprintf("PRIORITY: %d - %s", a.priority, a.message)
}

type Log struct {
	priority int
	message  string
}

func (l *Log) Priority() int {
	return l.priority
}

func (l *Log) String() string {
	return fmt.Sprintf("PRIORITY: %d - %s", l.priority, l.message)
}

type ExampleQueue interface {
	pq.Queueable
	String() string
}

func main() {

	// Add directly to Q object
	q := pq.NewQueue()
	q.Add(&Alert{DEBUG, "Debug Foo"})
	q.Add(&Alert{EMERG, "THIS IS FOOBAR!"})
	q.Add(&Alert{ERR,   "Error Baz"})
	q.Add(&Alert{INFO,  "Info Bar"})

	q.Add(&Log{ALERT, "Alert log message."})
	q.Add(&Log{CRIT, "Critial log message."})
	q.Add(&Log{WARN, "Warning log message."})

	for q.Len() > 0 {
		if prio, ok := q.Remove().(ExampleQueue); ok {
			fmt.Println(prio)
		}
	}

	// Using heap interface
	q = pq.NewQueue()
	Init(q)
	Push(q, &Alert{DEBUG, "Debug Foo"})
	Push(q, &Alert{EMERG, "THIS IS FOOBAR!"})
	Push(q, &Alert{ERR,   "Error Baz"})
	Push(q, &Alert{INFO,  "Info Bar"})

	Push(q, &Log{ALERT, "Alert log message."})
	Push(q, &Log{CRIT, "Critial log message."})
	Push(q, &Log{WARN, "Warning log message."})

	for q.Len() > 0 {
		if prio, ok := Pop(q).(ExampleQueue); ok {
			fmt.Println(prio)
		}
	}
}