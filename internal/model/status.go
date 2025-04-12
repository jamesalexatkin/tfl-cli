package model

import "time"

// Line represents the status for a particular line.
type Line struct {
	Name         string
	Status       string
	LineStatuses []LineStatus
	Disruption   *string
}

// LineStatus represents a status description for a line.
// Lines can have multiple statuses (e.g. it could be 'Part Closed' for works, while also having 'Minor Delays').
type LineStatus struct {
	StatusSeverityDescription string
	Reason                    string
}

// Underground represents the status of all underground tube lines.
type Underground struct {
	Bakerloo           Line
	Central            Line
	Circle             Line
	District           Line
	HammersmithAndCity Line
	Jubilee            Line
	Metropolitan       Line
	Northern           Line
	Piccadilly         Line
	Victoria           Line
	WaterlooAndCity    Line
}

// Overground represents the status of all overground rail lines.
type Overground struct {
	Liberty     Line
	Lioness     Line
	Mildmay     Line
	Suffragette Line
	Weaver      Line
	Windrush    Line
}

// TfLStatus represents the overall status for all types of lines.
type TfLStatus struct {
	Time          time.Time
	Underground   Underground
	Overground    Overground
	DLR           Line
	ElizabethLine Line
}
