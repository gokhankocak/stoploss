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
    Implements Several Statistics Functions
    Lowest Close
    Highest Close
    Lowest Low
    Highest High
    StdDev (Standard Deviation)
*/

package stoploss

import(
  "math"
)

func (DataSet Series) LowestClose() (Index int, Close float64) {

  Index = -1
  Close = math.MaxFloat64
  for k, d := range DataSet {
    if d.Close < Close {
      Close = d.Close
      Index = k
    }
  }

  return Index, Close
}

func (DataSet Series) HighestClose() (Index int, Close float64) {

  Index = -1
  Close = math.SmallestNonzeroFloat64
  for k, d := range DataSet {
    if d.Close > Close {
      Close = d.Close
      Index = k
    }
  }

  return Index, Close
}

func (DataSet Series) LowestLow() (Index int, Low float64) {

  Index = -1
  Low = math.MaxFloat64
  for k, d := range DataSet {
    if d.Low < Low {
      Low = d.Low
      Index = k
    }
  }

  return Index, Low
}

func (DataSet Series) HighestHigh() (Index int, High float64) {

  Index = -1
  High = math.SmallestNonzeroFloat64
  for k, d := range DataSet {
    if d.High > High {
      High = d.High
      Index = k
    }
  }

  return Index, High
}

func (DataSet Series) StdDev( Period int ) float64 {

	var Sum float64
	var Mean float64

  LastIndex := len(DataSet) - 1
	if LastIndex - Period - 1 < 0 { return math.NaN() }

	Sum = 0.0
	// for k := LastIndex; k > LastIndex - Period; k-- { Sum += Data[k].Close }
	for _, d := range DataSet[LastIndex - Period + 1:] { Sum += d.Close }
	Mean = Sum / float64(Period)

	Sum = 0.0
	// for k := LastIndex; k > LastIndex - Period; k-- {	Sum += ((Data[k].Close - Mean) * (Data[k].Close - Mean)) }
	for _, d := range DataSet[LastIndex - Period + 1:] { Sum += ((d.Close - Mean) * (d.Close - Mean)) }

	return math.Sqrt(Sum / float64(Period))
}
