package stoploss

import(
  "math"
)

func (DataSet Series) LowestClose() (Index int, Close float64) {

  Index = -1
  Close = math.MaxFloat64
  for k, d := range DataSet {
    if d.Close < Close {
      Close = d.Close
      Index = k
    }
  }

  return Index, Close
}

func (DataSet Series) HighestClose() (Index int, Close float64) {

  Index = -1
  Close = math.SmallestNonzeroFloat64
  for k, d := range DataSet {
    if d.Close > Close {
      Close = d.Close
      Index = k
    }
  }

  return Index, Close
}

func (DataSet Series) LowestLow() (Index int, Low float64) {

  Index = -1
  Low = math.MaxFloat64
  for k, d := range DataSet {
    if d.Low < Low {
      Low = d.Low
      Index = k
    }
  }

  return Index, Low
}

func (DataSet Series) HighestHigh() (Index int, High float64) {

  Index = -1
  High = math.SmallestNonzeroFloat64
  for k, d := range DataSet {
    if d.High > High {
      High = d.High
      Index = k
    }
  }

  return Index, High
}

func (DataSet Series) StdDev( Period int ) float64 {

	var Sum float64
	var Mean float64

  LastIndex := len(DataSet) - 1
	if LastIndex - Period - 1 < 0 { return math.NaN() }

	Sum = 0.0
	// for k := LastIndex; k > LastIndex - Period; k-- { Sum += Data[k].Close }
	for _, d := range DataSet[LastIndex - Period + 1:] { Sum += d.Close }
	Mean = Sum / float64(Period)

	Sum = 0.0
	// for k := LastIndex; k > LastIndex - Period; k-- {	Sum += ((Data[k].Close - Mean) * (Data[k].Close - Mean)) }
	for _, d := range DataSet[LastIndex - Period + 1:] { Sum += ((d.Close - Mean) * (d.Close - Mean)) }

	return math.Sqrt(Sum / float64(Period))
}
