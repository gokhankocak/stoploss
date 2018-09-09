/*
Copyright 2018 www.gokhankocak.com gokhan.kocak@mail.ru

The MIT License
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
		Implements OHLC (Open, High, Low, Close and Volume)
		Ohlc is the basic data structure used by the package
*/

package stoploss

import (
	"time"
	"log"
	"encoding/json"
)

type Ohlc struct {
	TimeStamp int64 `json:"TimeStamp"`
	High float64 `json:"High"`
	Low float64 `json:"Low"`
	Open float64 `json:"Open"`
	Close float64 `json:"Close"`
	Volume float64 `json:"Volume"`
}

type Series []Ohlc

func (d Ohlc) TimeStampToString() string {

	return time.Unix(d.TimeStamp, 0).UTC().String()
}

func (d Ohlc) ToJsonString() (string, error) {

	ByteResult, err := json.Marshal(d)
	if err != nil { return "", err }

	return string(ByteResult), nil
}

func (d Ohlc) Dump( Prefix string ) {

	s, err := d.ToJsonString()
	if err != nil {
		log.Println(Prefix, err.Error())
	} else {
		log.Println(Prefix, s)
	}
}

func (ds Series) Dump( Prefix string ) {

	for _, d := range ds { d.Dump(Prefix) }
}
