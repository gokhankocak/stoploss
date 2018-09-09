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
    Implements MACD (Moving Average Convergence Divergence)
*/

package stoploss

import(
  "math"
)

type Macd struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewMacd() *Macd {

  m := new(Macd)
  m.DataSet = DataSet
  return m
}

type MacdResult struct {
  MacdLine float64
  SignalLine float64
  Histogram float64
}

// https://www.tradingview.com/wiki/MACD_(Moving_Average_Convergence/Divergence)

func (m *Macd) MACD( Period int ) (Result *MacdResult) {

  var tmpMacd Series

  LastIndex := len(m.DataSet) - 1
  if LastIndex - Period - 1 < 0 { return nil }

  for k := LastIndex - Period + 1; k < LastIndex; k++ {

    maTmp := m.DataSet[:k].NewMovingAverage()
    tmp12 := maTmp.EMA(12)
    if tmp12 == math.NaN() { continue }
    tmp26 := maTmp.EMA(26)
    if tmp26 == math.NaN() { continue }
    d := new(Ohlc)
    d.Close = tmp12 - tmp26
    tmpMacd = append(tmpMacd, *d)
  }

  maTmp := m.DataSet.NewMovingAverage()
  Result.MacdLine = maTmp.EMA(12) - maTmp.EMA(26)
  Result.SignalLine = tmpMacd.NewMovingAverage().EMA(9)
  Result.Histogram = Result.MacdLine - Result.SignalLine

  return Result
}
