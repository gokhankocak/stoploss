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
    Implements Stochastic Technical Analysis
    Stochastic Slow
*/

package stoploss

import(
  "math"
)

type Stochastic struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewStochastic() *Stochastic {
  st := new(Stochastic)
  st.DataSet = DataSet
  return st
}

/*
%K = (Current Close - Lowest Low)/(Highest High - Lowest Low) * 100
%D = 3-day SMA of %K

Lowest Low = lowest low for the look-back period
Highest High = highest high for the look-back period
%K is multiplied by 100 to move the decimal point two places
*/

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/slow-stochastic
// 0 1 2 3 4 5 6 7 8 9 10 11 12
// Period = 5
// PeriodK = 3
// PeriodD = 2

func (st *Stochastic) Slow( Period int, PeriodK int, PeriodD int ) (float64, float64) {

	var Highest float64
	var Lowest float64
	var MaxKD int
	var Kpercents Series
	var Dpercent float64

  LastIndex := len(st.DataSet) - 1
	//log.Printf("StochasticSlow: LastIndex %d Period %d PeriodK %d PeriodD %d\n", LastIndex, Period, PeriodK, PeriodD)

	MaxKD = int(math.Max(float64(PeriodK), float64(PeriodD)))

	if LastIndex - Period - MaxKD - 1 < 0 { return math.NaN(), math.NaN() }
	if Period < PeriodK || Period < PeriodD { return math.NaN(), math.NaN() }

	Kpercents = make(Series, MaxKD)

	for c := 0; c < MaxKD; c++ {

		Start := LastIndex - Period - MaxKD + 2 + c
		Highest = math.SmallestNonzeroFloat64
		Lowest  = math.MaxFloat64

		for _, d := range st.DataSet[Start:Start + Period] {

			//log.Printf("StochasticSlow: c %d Start %d k %d\n", c, Start, k)

			Highest = math.Max(Highest, d.High)
			Lowest  = math.Min(Lowest, d.Low)
		}

		Kpercents[c].Close = 100.0 * (st.DataSet[Start + Period - 1].Close - Lowest) / (Highest - Lowest)
	}

	//log.Printf("StochasticSlow: Kpercents %v\n", Kpercents)

	Dpercent = Kpercents.NewMovingAverage().SMA(PeriodD)

	return Kpercents[len(Kpercents) - 1].Close, Dpercent
}
