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
    Implements Accumulation Distribution
*/

package stoploss

import(
  "math"
)

type AccDist struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewAccDist() *AccDist {
  ad := new(AccDist)
  ad.DataSet = DataSet
  return ad
}

func (Data Ohlc) AccDist() float64 {

  return (Data.Volume * ((Data.Close - Data.Low) - (Data.High - Data.Close)) / (Data.High - Data.Low))
}

func (ad *AccDist) AccDist( Period int ) float64 {

  var Result float64 = 0.0
  LastIndex := len(ad.DataSet) - 1
  if LastIndex < 0 { return math.NaN() }

  for _, d := range ad.DataSet[LastIndex - Period + 1:] {

    Result += d.AccDist()
  }

  return Result
}

func (ad *AccDist) AccDistSeries( Period int ) (Result Series) {

  LastIndex := len(ad.DataSet) - 1
  if LastIndex - Period - 1 < 0 { return nil }

  for k := LastIndex - Period + 1; k < LastIndex; k++ {

    adTmp := ad.DataSet[:k].NewAccDist()
    v := adTmp.AccDist(Period)
    if v != math.NaN() {

      Data := new(Ohlc)
      Data.Close = v
      Result = append(Result, *Data)
    }
  }

  return Result
}
