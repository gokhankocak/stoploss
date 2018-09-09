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
