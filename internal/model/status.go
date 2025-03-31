package model

import "time"

type Line struct {
	Name         string
	Status       string
	LineStatuses []LineStatus
	Disruption   *string
}

type LineStatus struct {
	StatusSeverityDescription string
	Reason                    string
}

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

type TfLStatus struct {
	Time          time.Time
	Underground   Underground
	Overground    Line
	DLR           Line
	ElizabethLine Line
}
