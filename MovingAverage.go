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
