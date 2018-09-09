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
    Implements Several Moving Averages
    SMA (Simple Moving Average)
    EMA (Exponential Moving Average)
    WMA (Weighted Moving Average)
*/

package stoploss

import(
  "math"
)

type MovingAverage struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewMovingAverage() *MovingAverage {
  ma := new(MovingAverage)
  ma.DataSet = DataSet
  return ma
}

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/sma

func (ma *MovingAverage) SMA( Period int ) float64 {

	var Sum float64
  var SMA float64

  LastIndex := len(ma.DataSet) - 1
	if LastIndex - Period - 1 < 0 { return math.NaN() }

	Sum = 0.0
	for _, d := range ma.DataSet[LastIndex - Period + 1:] { Sum += d.Close }

  SMA = (Sum / float64(Period))
	return SMA
}

func (ma *MovingAverage) EMA( Period int ) float64 {

	var k float64

  LastIndex := len(ma.DataSet) - 1
	if LastIndex - Period - 1 < 0 { return math.NaN() }

	if Period == 0 { return ma.DataSet[LastIndex].Close }

	k = 2.0 / (float64(Period) + 1)
  maTmp := ma.DataSet[0:LastIndex - 1].NewMovingAverage()
	return (ma.DataSet[LastIndex].Close * k + (1.0 - k) * maTmp.EMA(Period - 1))
}

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/wma

func (ma *MovingAverage) WMA( Period int ) float64 {

  var Sum float64 = 0.0
  var SumCount int = 0

  LastIndex := len(ma.DataSet) - 1
  if LastIndex - Period - 1 < 0 { return math.NaN() }

  for k, d := range ma.DataSet[LastIndex - Period + 1:] {
    SumCount += (k + 1)
    Sum += d.Close * float64(k + 1)
  }

  return (Sum / float64(SumCount))
}
