/*
 * iconik_shared
 *
 * iconik client code shared by all API packages
 *
 * API version: 1.0
 */

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type Time time.Time

//	implement the json.Unmarshaler interface for type Time
func (p *Time) UnmarshalJSON(xBytes []byte) (err error) {
	//	first unmarshal into a string
	var val string
	if err = json.Unmarshal(xBytes, &val); nil == err {
		//	then parse the string using our custom format
		//	because of the need to cast, we can't assign directly to *p
		var dt time.Time
		//	In the absence of time zone information, time.Parse() interprets a time as UTC, which is precisely what we want.
		if dt, err = time.Parse("2006-01-02T15:04:05.999999", val); nil == err {
			*p = Time(dt)
		}
	}

	return
}

//	implement the json.Marshaler interface for type Time
//<<<<	TO DO:

//	implement the fmt.Stringer and fmt.GoStringer interfaces for type Time
func (dt Time) String() string {
	return time.Time(dt).Format(time.RFC3339Nano)
}
func (dt Time) GoString() string {
	return fmt.Sprintf("%#v", time.Time(dt))
}
