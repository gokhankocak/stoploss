package stoploss

import(
  "math"
)

type RelativeStrength struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewRelativeStrength() *RelativeStrength {
  rs := new(RelativeStrength)
  rs.DataSet = DataSet
  return rs
}

func (rs *RelativeStrength) RSI( Period int ) float64 {

	var Gains []float64
	var Losses []float64
	var Diff float64
	var SumGain float64
	var SumLoss float64
	var Index int

  LastIndex := len(rs.DataSet) - 1

	if LastIndex - Period - 1 < 0 { return math.NaN() }

	Gains = make([]float64, Period)
	Losses = make([]float64, Period)

	SumGain = 0.0
	SumLoss = 0.0
	Index = Period - 1
	for k := LastIndex; k > LastIndex - Period; k-- {

		Diff = rs.DataSet[k].Close - rs.DataSet[k - 1].Close
		if Diff > 0.0 {
			Gains[Index] = Diff
			SumGain += Diff
			Losses[Index] = 0.0
		} else if Diff < 0.0 {
			Gains[Index] = 0.0
			Losses[Index] = -1.0 * Diff
			SumLoss += -1.0 * Diff
		} else {
			Gains[Index] = 0.0
			Losses[Index] = 0.0
		}
		Index--
	}

	return (100.0 - 100.0 / (1.0 + SumGain / SumLoss))
}
