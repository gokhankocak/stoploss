package stoploss

import (
	"math"
)

type Bollinger struct {
	DataSet Series
	DebugMode bool
}

func (DataSet Series) NewBollinger() *Bollinger {
	b := new(Bollinger)
	b.DataSet = DataSet
	return b
}

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/bollinger-bands

func (b *Bollinger) Bands( Period int, Multiplier float64 ) (LowerBand float64, UpperBand float64) {

	LastIndex := len(b.DataSet) - 1
	if LastIndex - Period - 1 < 0 { return math.NaN(), math.NaN() }

	// Choose appropriate Multiplier if Multiplier == 0.0
	if Multiplier == 0.0 {
		if Period >= 50 {
			Multiplier = 2.1
		} else if Period >= 20 {
			Multiplier = 2.0
		} else {
			Multiplier = 1.9
		}
	}

	StdDev := b.DataSet.StdDev(Period)
	Sma := b.DataSet.NewMovingAverage().SMA(Period)
	LowerBand = Sma - (Multiplier * StdDev)
	UpperBand = Sma + (Multiplier * StdDev)

	return LowerBand, UpperBand
}

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/percent-b

func (b *Bollinger) PercentB( Period int, Multiplier float64 ) float64 {

	LastIndex := len(b.DataSet) - 1
	if LastIndex - Period - 1 < 0 { return math.NaN() }

	LowerBand, UpperBand := b.Bands(Period, Multiplier)
	return (100.0 * (b.DataSet[LastIndex].Close - LowerBand) / (UpperBand - LowerBand))
}

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/bollinger-band-width

func (b *Bollinger) Bandwidth( Period int, Multiplier float64 ) float64 {

	LastIndex := len(b.DataSet) - 1
	if LastIndex - Period - 1 < 0 { return math.NaN() }

	LowerBand, UpperBand := b.Bands(Period, Multiplier)
	Sma := b.DataSet.NewMovingAverage().SMA(Period)
	return (UpperBand - LowerBand) / Sma
}
