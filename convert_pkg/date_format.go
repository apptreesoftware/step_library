package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"time"
)

type DateFormat struct {
}

func (DateFormat) Name() string {
	return "date_format"
}

func (DateFormat) Version() string {
	return "1.0"
}

func (DateFormat) Execute(in step.Context) (interface{}, error) {
	input := formatDateInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	d, err := time.Parse(getDateFormat(input.InputFormat), input.Date)
	if err != nil {
		return nil, fmt.Errorf("unable to parse date %s with format %s. Make sure you are using an appropriate format\n%s", input.Date, input.InputFormat, formatExample)
	}
	outFormat := getDateFormat(input.OutputFormat)
	out := d.Format(outFormat)
	return formatDateOutput{
		Date: out,
	}, nil
}

type formatDateInput struct {
	Date         string
	InputFormat  string
	OutputFormat string
}

type formatDateOutput struct {
	Date string
}

func getDateFormat(format string) string {
	switch format {
	case "ANSIC":
		return time.ANSIC
	case "UnixDate":
		return time.UnixDate
	case "RubyDate":
		return time.RubyDate
	case "RFC822":
		return time.RFC822
	case "RFC822Z":
		return time.RFC822Z
	case "RFC850":
		return time.RFC850
	case "RFC1123":
		return time.RFC1123
	case "RFC1123Z":
		return time.RFC1123Z
	case "RFC3339":
		return time.RFC3339
	case "RFC3339Nano":
		return time.RFC3339Nano
	case "Kitchen":
		return time.Kitchen
	case "Stamp":
		return time.Stamp
	case "StampMilli":
		return time.StampMilli
	case "StampMicro":
		return time.StampMicro
	case "StampNano":
		return time.StampNano
	}

	return format
}

var formatExample = `
Year	06   2006
Month	01   1   Jan   January
Day	02   2   _2   (width two, right justified)
Weekday	Mon   Monday
Hours	03   3   15
Minutes	04   4
Seconds	05   5
ms μs ns	.000   .000000   .000000000
ms μs ns	.999   .999999   .999999999   (trailing zeros removed)
am/pm	PM   pm
Timezone	MST
Offset	-0700   -07   -07:00   Z0700   Z07:00

For example, to parse a date in the format of 2019-09-10 you would use the input format of

2006-01-02

You can also specify one of the common formats:

ANSIC       = "Mon Jan _2 15:04:05 2006"
UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
RFC822      = "02 Jan 06 15:04 MST"
RFC822Z     = "02 Jan 06 15:04 -0700"
RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
RFC3339     = "2006-01-02T15:04:05Z07:00"
RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
Kitchen     = "3:04PM"
// Handy time stamps.
Stamp      = "Jan _2 15:04:05"
StampMilli = "Jan _2 15:04:05.000"
StampMicro = "Jan _2 15:04:05.000000"
StampNano  = "Jan _2 15:04:05.000000000"

`
