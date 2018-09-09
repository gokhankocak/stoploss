package stoploss

import(
  "math"
)

type MoneyFlow struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewMoneyFlow() *MoneyFlow {
  mf := new(MoneyFlow)
  mf.DataSet = DataSet
  return mf
}

// Calculate Money Flow Index
func (mf *MoneyFlow) MFI( Period int ) float64 {

	var TypicalPrice []float64
	var RawMoneyFlow []float64
	var PositiveFlow float64
	var NegativeFlow float64
	var Index int
	var Result float64

	LastIndex := len(mf.DataSet) - 1

	if LastIndex - Period - 1 < 0 { return math.NaN() }

	TypicalPrice = make([]float64, Period)
	RawMoneyFlow = make([]float64, Period)

	Index = Period - 1
	for k := LastIndex; k > LastIndex - Period; k-- {

		TypicalPrice[Index] = (mf.DataSet[k].High + mf.DataSet[k].Low + mf.DataSet[k].Close) / 3.0
		RawMoneyFlow[Index] = TypicalPrice[Index] * mf.DataSet[k].Volume
		Index--
	}

	PositiveFlow = 0.0
	NegativeFlow = 0.0
	for k := 1; k < Period - 1; k++ {

		if TypicalPrice[k] > TypicalPrice[k - 1] {

			PositiveFlow += RawMoneyFlow[k]
		} else if TypicalPrice[k] < TypicalPrice[k - 1] {

			NegativeFlow += RawMoneyFlow[k]
		}
	}

	Result = (100.0 - 100.0 / (1.0 + PositiveFlow / NegativeFlow))

	return Result
}
