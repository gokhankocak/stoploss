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
