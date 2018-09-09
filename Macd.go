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
