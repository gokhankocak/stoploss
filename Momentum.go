package stoploss

import (
  "math"
)

type Momentum struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewMomentum() *Momentum {
  m := new(Momentum)
  m.DataSet = DataSet
  return m
}

//
// TODO: Implement https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/momentum-oscillator
//

func (m *Momentum) GenericMomentum( Period int ) float64 {

  var OnePlusR float64
  var M float64

  LastIndex := len(m.DataSet) - 1
  if LastIndex - Period - 1 < 0 { return math.NaN() }

  M = 1.0

  for k := LastIndex - Period; k < LastIndex; k++ {

    OnePlusR = 1.0 + ((m.DataSet[k].Close - m.DataSet[k - 1].Close) / m.DataSet[k - 1].Close)
    M *= OnePlusR
  }

  return (100.0 * (M - 1.0))
}

func (m *Momentum) QuantitativeUsingDailyData( Period int ) float64 {

  var Monthly Series

  LastIndex := len(m.DataSet) - 1
  for k := LastIndex; k > 0; k -= 21 {

    d := new(Ohlc)
    d.Close = m.DataSet[k].Close
    Monthly = append(Monthly, *d)
  }

  for k := 0; k < len(Monthly) / 2; k++ {

    tmp := Monthly[len(Monthly) - k - 1]
    Monthly[len(Monthly) - k - 1] = Monthly[k]
    Monthly[k] = tmp
  }

  return Monthly.NewMomentum().GenericMomentum(12)
}

func (m *Momentum) QuantitativeUsingWeeklyData( Period int ) float64 {

  var Monthly Series

  LastIndex := len(m.DataSet) - 1
  for k := LastIndex; k > 0; k -= 4 {

    d := new(Ohlc)
    d.Close = m.DataSet[k].Close
    Monthly = append(Monthly, *d)
  }

  for k := 0; k < len(Monthly) / 2; k++ {

    tmp := Monthly[len(Monthly) - k - 1]
    Monthly[len(Monthly) - k - 1] = Monthly[k]
    Monthly[k] = tmp
  }

  return Monthly.NewMomentum().GenericMomentum(12)
}

func (m *Momentum) RawFrogInPan( Period int ) float64 {

	var PositiveCount int
	var NegativeCount int
	var PositivePercent float64
	var NegativePercent float64
	var Total float64

  LastIndex := len(m.DataSet) - 1
	if LastIndex - Period - 1 < 0 { return math.NaN() }

	PositiveCount = 0
	NegativeCount = 0

	for k := LastIndex - Period; k < LastIndex; k++ {

		if m.DataSet[k].Close - m.DataSet[k - 1].Close > 0.0 {

			PositiveCount++
		} else if m.DataSet[k].Close - m.DataSet[k - 1].Close < 0.0 {

			NegativeCount++
		}
	}

	Total = float64(PositiveCount + NegativeCount)
	PositivePercent = float64(PositiveCount) / Total
	NegativePercent = float64(NegativeCount) / Total

	return (NegativePercent - PositivePercent)
}

//
// Kitapta uygulanan yöntem bundan farklı.
// Aylık verinin momentumu hesaplanıyor
// Günlük veri için de rawFrogInPan hesaplanıyor

func (m *Momentum) FrogInPan( Period int ) float64 {

  var RetVal float64

  LastIndex := len(m.DataSet) - 1
	if LastIndex - Period - 1 < 0 { return math.NaN() }

  if m.GenericMomentum(Period) > 0.0 {
    RetVal = m.RawFrogInPan(Period)
  } else {
    RetVal = -1.0 * m.RawFrogInPan(Period)
  }

  return RetVal
}
