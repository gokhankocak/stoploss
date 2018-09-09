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
    Implements Average True Range
*/

package stoploss

import(
  "math"
)

type TrueRange struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewTrueRange() *TrueRange {
  tr := new(TrueRange)
  tr.DataSet = DataSet
  return tr
}

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/atr

func (tr *TrueRange) TrueRange() float64 {

  LastIndex := len(tr.DataSet) - 1
	if LastIndex - 1 < 0 { return math.NaN() }

	TheMax := math.Max(tr.DataSet[LastIndex].High - tr.DataSet[LastIndex].Low, math.Abs(tr.DataSet[LastIndex].High - tr.DataSet[LastIndex - 1].Close))
	TheMax = math.Max(TheMax, math.Abs(tr.DataSet[LastIndex].Low - tr.DataSet[LastIndex - 1].Close))

	return TheMax
}

func (tr *TrueRange) AverageTrueRange( Period int ) float64 {

	var TrData Series

  LastIndex := len(tr.DataSet) - 1
	if LastIndex - Period - 3 < 0 { return math.NaN() }

	for k, _ := range tr.DataSet[LastIndex - Period - 2: LastIndex] {

    trTmp := tr.DataSet[0:LastIndex - k].NewTrueRange()
		r := new(Ohlc)
		r.Close = trTmp.TrueRange()
		if math.IsNaN(r.Close) { break }
		TrData = append(TrData, *r)
	}

	return TrData.NewMovingAverage().SMA(Period)
}
