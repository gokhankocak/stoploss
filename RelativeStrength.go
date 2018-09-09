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
    Implements RSI (Relative Strength Index)
*/

package stoploss

import(
  "math"
)

type RelativeStrength struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewRelativeStrength() *RelativeStrength {
  rs := new(RelativeStrength)
  rs.DataSet = DataSet
  return rs
}

func (rs *RelativeStrength) RSI( Period int ) float64 {

	var Gains []float64
	var Losses []float64
	var Diff float64
	var SumGain float64
	var SumLoss float64
	var Index int

  LastIndex := len(rs.DataSet) - 1

	if LastIndex - Period - 1 < 0 { return math.NaN() }

	Gains = make([]float64, Period)
	Losses = make([]float64, Period)

	SumGain = 0.0
	SumLoss = 0.0
	Index = Period - 1
	for k := LastIndex; k > LastIndex - Period; k-- {

		Diff = rs.DataSet[k].Close - rs.DataSet[k - 1].Close
		if Diff > 0.0 {
			Gains[Index] = Diff
			SumGain += Diff
			Losses[Index] = 0.0
		} else if Diff < 0.0 {
			Gains[Index] = 0.0
			Losses[Index] = -1.0 * Diff
			SumLoss += -1.0 * Diff
		} else {
			Gains[Index] = 0.0
			Losses[Index] = 0.0
		}
		Index--
	}

	return (100.0 - 100.0 / (1.0 + SumGain / SumLoss))
}
