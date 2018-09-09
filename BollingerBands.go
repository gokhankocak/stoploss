/*
Copyright 2018 www.gokhankocak.com gokhan.kocak@mail.ru

The MIT License
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
		Implements Bollinger Bands
*/

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
