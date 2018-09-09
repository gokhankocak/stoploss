package stoploss

import(
)

type ParabolicSAR struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewParabolicSAR() *ParabolicSAR {

  psar := new(ParabolicSAR)
  psar.DataSet = DataSet
  return psar
}

// https://www.tradingview.com/wiki/Parabolic_SAR_(SAR)
