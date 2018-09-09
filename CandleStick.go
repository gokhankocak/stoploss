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
    Implements Japanese Candle Sticks
*/

package stoploss

import (
  "math"
)

type CandleStick struct {
  DataSet Series
  LastIndex int
  LastData Ohlc
  PrevData Ohlc
  DebugMode bool
}

func (DataSet Series) NewCandleStick() *CandleStick {

  cs := new(CandleStick)
  cs.DataSet = DataSet
  cs.LastIndex = len(DataSet) - 1
  cs.LastData = cs.DataSet[cs.LastIndex]
  cs.PrevData = cs.DataSet[cs.LastIndex - 1]
  return cs
}

//
// Technical Analysis - The Complete Resource for Financial Market Technicians
//

func (cs *CandleStick) StatusSummary() (Action string, Summary string) {

  Action  = ""
  Summary = ""

  if cs.IsSpinningTop() {

    Summary += "[SpinningTop]"
    Action  += "[Reversal]"
  }

  if cs.IsHammer() {

    Summary += "[Hammer]"
    Action  += "[ReversalUpwards]"
  }

  if cs.IsHangingMan() {

    Summary += "[HangingMan]"
    Action  += "[ReversalDownwards]"
  }

  if cs.IsInvertedHammer() {

    Summary += "[InvertedHammer]"
  }

  if cs.IsShootingStar() {

    Summary += "[ShootingStar]"
    Action  += "[ReversalDownwards]"
  }

  if cs.IsEveningStar() {

    Summary += "[EveningStar]"
    Action  += "[Top]"
  }

  if cs.IsMorningStar() {

    Summary += "[MorningStar]"
    Action  += "[Bottom]"
  }

  if cs.IsInsideDown() {

    Summary += "[InsideDown]"
  }

  if cs.IsInsideUp() {

    Summary += "[InsideUp]"
  }

  if cs.IsThreeBlackCrows() {

    Summary += "[ThreeBlackCrows]"
  }

  if cs.IsThreeWhiteSoldiers() {

    Summary += "[ThreeWhiteSoldiers]"
  }

  if cs.IsBullishBeltHold() {

    Summary += "[BullishBeltHold]"
  }

  if cs.IsBearishBeltHold() {

    Summary += "[BearishBeltHold]"
  }

  return Action, Summary
}

//
// https://www.nasdaq.com/forex/education/single-candlestick-patterns-overview.aspx
//

func (cs* CandleStick) IsFuzzyEqual( a float64, b float64 ) bool {

  var e float64

  if a >= 100000.0 {

    e = 1000.0
  } else if a >= 10000.0 {

    e = 100.0
  } else if a >= 1000.0 {

    e = 10.0
  } else if a >= 100.0 {

    e = 1.0
  } else if a >= 10.0 {

    e = 0.1
  } else if a >= 1.0 {

    e = 0.01
  } else if a >= 0.1 {

    e = 0.001
  } else if a >= 0.01 {

    e = 0.0001
  } else if a >= 0.001 {

    e = 0.00001
  } else if a >= 0.0001 {

    e = 0.000001
  } else {

    e = 0.0000001
  }

  return (math.Abs(a - b) <= e)
}

func (cs *CandleStick) IsBlack() bool {

  return cs.LastData.Open > cs.LastData.Close
}

func (cs *CandleStick) IsWhite() bool {

  return cs.LastData.Open < cs.LastData.Close
}

func (cs *CandleStick) IsDownTrend() bool {

  var BlackCount int

  if cs.LastIndex - 3 < 0 { return false }

  BlackCount = 0
  if cs.DataSet[:cs.LastIndex - 1].NewCandleStick().IsBlack() { BlackCount++ }
  if cs.DataSet[:cs.LastIndex - 2].NewCandleStick().IsBlack() { BlackCount++ }
  if cs.DataSet[:cs.LastIndex - 3].NewCandleStick().IsBlack() { BlackCount++ }

  if BlackCount > 1 { return true }
  return false
}

func (cs *CandleStick) IsUpTrend() bool {

  var WhiteCount int

  if cs.LastIndex - 3 < 0 { return false }

  WhiteCount = 0
  if cs.DataSet[:cs.LastIndex - 1].NewCandleStick().IsWhite() { WhiteCount++ }
  if cs.DataSet[:cs.LastIndex - 2].NewCandleStick().IsWhite() { WhiteCount++ }
  if cs.DataSet[:cs.LastIndex - 3].NewCandleStick().IsWhite() { WhiteCount++ }

  if WhiteCount > 1 { return true }
  return false
}

/*
A doji is a candlestick where the opening price is almost the exact same as the opening price, with long shadows in one direction or both.
What this can signal is indecision between buyers and sellers.
If these occur at the top or bottom of a trend it can signal a reversal as it shows a slowing of momentum.
*/
func (cs *CandleStick) IsDoji() bool {

  return cs.IsFuzzyEqual(cs.LastData.Open, cs.LastData.Close)
}

/*
A spinning top has two long equal length shadows with a small body and typically signals a reversal when they occur during a trend.
The reason behind the reversal is that it shows indecision between buyers and sellers, and that neither of them can close much higher or lower than the opening.
*/
func (cs *CandleStick) IsSpinningTop() bool {

  Body   := math.Abs(cs.LastData.Open - cs.LastData.Close)
  Height := math.Abs(cs.LastData.High - cs.LastData.Low)

  if Body <= Height / 3.0 { return true }
  return false
}

/*
A Maruboza is when a candlestick forms with a long body and little to no shadow.
This signals strong movement in one direction, which will likely continue movement in that direction in the near future.
In the bullish Maruboza case, the opening price is equal to the low and the closing price is equal to the high.
With a bearish Maruboza the opening price is the high and the closing is the low.
*/
func (cs *CandleStick) IsMaruboza() bool {

  if cs.IsFuzzyEqual(cs.LastData.Open, cs.LastData.Low) && cs.IsFuzzyEqual(cs.LastData.Close, cs.LastData.High) { return true }

  return false
}

/*
The hammer chart pattern is a Japanese candlestick that has a small body with a short to no shadow on top of the body with a long shadow on the bottom.
When this candlestick occurs at the bottom of a trend, it can signal for a reversal.
*/
func (cs *CandleStick) IsHammer() bool {

  var DiffHighLow float64
  var DiffOpenClose float64

  if cs.LastIndex < 0 { return false }

  if cs.IsFuzzyEqual(cs.LastData.Close, cs.LastData.High) || cs.IsFuzzyEqual(cs.LastData.Open, cs.LastData.High) == false { return false }

  DiffOpenClose = math.Abs(cs.LastData.Open - cs.LastData.Close)
  DiffHighLow   = math.Abs(cs.LastData.High - cs.LastData.Low)
  if DiffHighLow / 2.0 >= DiffOpenClose { return true }

  return false
}

/*
The hanging man candlestick pattern has the exact same candlestick as the hammer but has different price action before it, so it signals for a reversal downwards
*/
func (cs *CandleStick) IsHangingMan() bool {

  return cs.IsHammer()
}

/*
The inverted hammer is a candlestick similar to the hammer and hanging man patterns in that it can signal a reversal.
With an inverted hammer, a small bullish candlestick body forms with a long shadow on top, and occurs during a downtrend.
*/
func (cs *CandleStick) IsInvertedHammer() bool {

  return cs.IsShootingStar()
}

/*
The shooting star is similar to the inverted hammer but occurs during an uptrend and can signal a reversal downwarDataSet.
The candlestick for a shooting star is a small bearish body with a long shadow on top.
*/
func (cs *CandleStick) IsShootingStar() bool {

  var DiffHighLow float64
  var DiffOpenClose float64

  if cs.LastIndex < 0 { return false }

  if cs.IsFuzzyEqual(cs.LastData.Close, cs.LastData.Low) || cs.IsFuzzyEqual(cs.LastData.Open, cs.LastData.Low) == false { return false }

  DiffOpenClose = math.Abs(cs.LastData.Open - cs.LastData.Close)
  DiffHighLow   = math.Abs(cs.LastData.High - cs.LastData.Low)
  if DiffHighLow / 2.0 >= DiffOpenClose { return true }

  return false
}

// Is market TOP?
/*
The evening star is similar to the morning star pattern but occurs during an uptrend and signals a reversal downwarDataSet.
The evening stars' first candle is a bullish candle with a long body.
The second candle is a doji, which signals indecision.
The third and final candle in the chart pattern is the bearish candle that closes past at least the halfway point of the first bullish candle.
*/
func (cs *CandleStick) IsEveningStar() bool {

  var RetVal bool
  var Cond [11]bool

  if cs.LastIndex - 2 < 0 { return false }

  // Siyah
  Cond[0] = cs.LastData.Open > cs.LastData.Close
  Cond[1] = cs.LastData.High > cs.LastData.Open
  Cond[2] = cs.LastData.Low < cs.LastData.Close

  // log.Println(Cond)

  RetVal = true
  for c := 0; c < 3; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Ortadaki Siyah
  Cond[3] = cs.PrevData.Open > cs.PrevData.Close
  Cond[4] = cs.PrevData.High > cs.PrevData.Open
  Cond[5] = cs.PrevData.Low < cs.PrevData.Close

  // log.Println(Cond)

  RetVal = true
  for c := 3; c < 6; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // En baştaki beyaz
  Cond[6] = cs.DataSet[cs.LastIndex - 2].Open < cs.DataSet[cs.LastIndex - 2].Close
  Cond[7] = cs.DataSet[cs.LastIndex - 2].High > cs.DataSet[cs.LastIndex - 2].Close
  Cond[8] = cs.DataSet[cs.LastIndex - 2].Low < cs.DataSet[cs.LastIndex - 2].Open

  // log.Println(Cond)

  RetVal = true
  for c := 6; c < 9; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Ortadaki en baştakinin vücut kısmının dışında olmalı
  Cond[9]  = cs.PrevData.Close > cs.DataSet[cs.LastIndex - 2].Close
  Cond[10] = cs.PrevData.Close > cs.LastData.Open

  // log.Println(Cond)

  RetVal = true
  for c := 9; c < 11; c++ {

    RetVal = RetVal && Cond[c]
  }

  return RetVal
}

// Is market Bottom?
/*
The first candlestick for a morning star is a bearish candle with a long body.
It is then followed by a doji (a small body candle with long shadows on bottom and top).
The doji signals indecisions and doesn't matter if it closes up or down.
The third candlestick is a bullish candlestick that should at least pass the halfway point of the first bearish candle.
The morning star is a buy indicator.
*/
func (cs *CandleStick) IsMorningStar() bool {

  var RetVal bool
  var Cond [11]bool

  if cs.LastIndex - 2 < 0 { return false }

  // Beyaz, en sağdaki
  Cond[0] = cs.LastData.Open < cs.LastData.Close
  Cond[1] = cs.LastData.High > cs.LastData.Close
  Cond[2] = cs.LastData.Low < cs.LastData.Open

  // log.Println(Cond)

  RetVal = true
  for c := 0; c < 3; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Ortadaki Beyaz
  Cond[3] = cs.PrevData.Open < cs.PrevData.Close
  Cond[4] = cs.PrevData.High > cs.PrevData.Close
  Cond[5] = cs.PrevData.Low < cs.PrevData.Open

  // log.Println(Cond)

  RetVal = true
  for c := 3; c < 6; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // En baştaki siyah
  Cond[6] = cs.DataSet[cs.LastIndex - 2].Open > cs.DataSet[cs.LastIndex - 2].Close
  Cond[7] = cs.DataSet[cs.LastIndex - 2].High > cs.DataSet[cs.LastIndex - 2].Open
  Cond[8] = cs.DataSet[cs.LastIndex - 2].Low < cs.DataSet[cs.LastIndex - 2].Close

  // log.Println(Cond)

  RetVal = true
  for c := 6; c < 9; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Ortadaki en baştakinin vücut kısmının dışında olmalı
  Cond[9]  = cs.PrevData.Close < cs.DataSet[cs.LastIndex - 2].Close
  Cond[10] = cs.PrevData.Close < cs.LastData.Open

  // log.Println(Cond)

  RetVal = true
  for c := 9; c < 11; c++ {

    RetVal = RetVal && Cond[c]
  }

  return RetVal
}

/*
The three inside down pattern is the opposite of the three inside up pattern.
In this case, the pattern is an indicator for a reversal downwarDataSet and must follow a recent uptrend.
The first candlestick in the pattern is a bullish candle with a long body.
The second is a bearish candle that passes at least the halfway point of the first bullish candle.
The last candlestick is another bearish candle that passes at least the low of the first bullish candle.
*/
func (cs *CandleStick) IsInsideDown() bool {

  var RetVal bool
  var Cond [12]bool

  if cs.LastIndex - 2 < 0 { return false }

  // Siyah en sağdaki
  Cond[0] = cs.LastData.Open > cs.LastData.Close
  Cond[1] = cs.LastData.High > cs.LastData.Open
  Cond[2] = cs.LastData.Low < cs.LastData.Close

  RetVal = true
  for c := 0; c < 3; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Siyah ortadaki
  Cond[3] = cs.PrevData.Open > cs.PrevData.Close
  Cond[4] = cs.PrevData.High > cs.PrevData.Open
  Cond[5] = cs.PrevData.Low < cs.PrevData.Close

  RetVal = true
  for c := 3; c < 6; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Soldaki Beyaz
  Cond[6] = cs.DataSet[cs.LastIndex - 2].Open < cs.DataSet[cs.LastIndex - 2].Close
  Cond[7] = cs.DataSet[cs.LastIndex - 2].High > cs.DataSet[cs.LastIndex - 2].Close
  Cond[8] = cs.DataSet[cs.LastIndex - 2].Low < cs.DataSet[cs.LastIndex - 2].Open

  RetVal = true
  for c := 6; c < 9; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Ortadaki soldakinin içinde kalmalı
  Cond[9]  = cs.LastData.Open < cs.PrevData.Open
  Cond[10] = cs.PrevData.Open < cs.DataSet[cs.LastIndex - 2].Close
  Cond[11] = cs.PrevData.Close > cs.DataSet[cs.LastIndex - 2].Open

  RetVal = true
  for c := 9; c < 12; c++ {

    RetVal = RetVal && Cond[c]
  }

  return RetVal
}

/*
The three inside up pattern occurs after a recent downtrend and signals for a reversal to an uptrend.
The first candle in the pattern is a bearish candle with a long body.
The next is a bullish candle that passes at least the halfway point of the first bearish candle.
The third and final candle is another bullish candle that passes at least the high of the first bearish candle.
*/
func (cs *CandleStick) IsInsideUp() bool {

  var RetVal bool
  var Cond [12]bool

  if cs.LastIndex - 2 < 0 { return false }

  // Beyaz en sağdaki
  Cond[0] = cs.LastData.Open < cs.LastData.Close
  Cond[1] = cs.LastData.High > cs.LastData.Close
  Cond[2] = cs.LastData.Low < cs.LastData.Open

  RetVal = true
  for c := 0; c < 3; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Beyaz ortadaki
  Cond[3] = cs.PrevData.Open < cs.PrevData.Close
  Cond[4] = cs.PrevData.High > cs.PrevData.Close
  Cond[5] = cs.PrevData.Low < cs.PrevData.Open

  RetVal = true
  for c := 3; c < 6; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Soldaki Siyah
  Cond[6] = cs.DataSet[cs.LastIndex - 2].Open > cs.DataSet[cs.LastIndex - 2].Close
  Cond[7] = cs.DataSet[cs.LastIndex - 2].High > cs.DataSet[cs.LastIndex - 2].Open
  Cond[8] = cs.DataSet[cs.LastIndex - 2].Low < cs.DataSet[cs.LastIndex - 2].Close

  RetVal = true
  for c := 6; c < 9; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Ortadaki soldakinin içinde kalmalı
  Cond[9]  = cs.LastData.Open > cs.PrevData.Open
  Cond[10] = cs.PrevData.Open > cs.DataSet[cs.LastIndex - 2].Close
  Cond[11] = cs.PrevData.Close < cs.DataSet[cs.LastIndex - 2].Open

  RetVal = true
  for c := 9; c < 12; c++ {

    RetVal = RetVal && Cond[c]
  }

  return RetVal
}

/*
The three black crows chart pattern is the opposite of the three white soldiers chart pattern.
Instead of three bullish candles with the three white soldiers, you have three bearish candles instead.
Also, the three black crows pattern neeDataSet to come after an extended uptrend and consolidation for it to confirm a new downtrend.
*/
func (cs *CandleStick) IsThreeBlackCrows() bool {

  var RetVal bool
  var Cond [13]bool

  if cs.LastIndex - 2 < 0 { return false }

  // Siyah sağdaki
  Cond[0] = cs.LastData.Open > cs.LastData.Close
  Cond[1] = cs.LastData.High > cs.LastData.Open
  Cond[2] = cs.LastData.Low < cs.LastData.Close

  RetVal = true
  for c := 0; c < 3; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Siyah ortadaki
  Cond[3] = cs.PrevData.Open > cs.PrevData.Close
  Cond[4] = cs.PrevData.High > cs.PrevData.Open
  Cond[5] = cs.PrevData.Low < cs.PrevData.Close

  RetVal = true
  for c := 3; c < 6; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Siyah soldaki
  Cond[6] = cs.DataSet[cs.LastIndex - 2].Open > cs.DataSet[cs.LastIndex - 2].Close
  Cond[7] = cs.DataSet[cs.LastIndex - 2].High > cs.DataSet[cs.LastIndex - 2].Open
  Cond[8] = cs.DataSet[cs.LastIndex - 2].Low < cs.DataSet[cs.LastIndex - 2].Close

  RetVal = true
  for c := 6; c < 9; c++ {

    RetVal = RetVal && Cond[c]
  }

  if RetVal == false { return false }

  // Soldan sağa azalan
  Cond[9]  = cs.LastData.Open < cs.PrevData.Open
  Cond[10] = cs.PrevData.Open < cs.DataSet[cs.LastIndex - 2].Open
  Cond[11] = cs.LastData.Close < cs.PrevData.Close
  Cond[12] = cs.PrevData.Close < cs.DataSet[cs.LastIndex - 2].Close

  RetVal = true
  for c := 9; c < 13; c++ {

    RetVal = RetVal && Cond[c]
  }

  return RetVal
}

/*
The three white soldiers pattern can appear after an extended downtrend and a period of consolidation.
The first candlestick of the chart pattern that neeDataSet to appear is a bullish candlestick with a long body.
The next candlestick in the pattern is another bullish candlestick, but this candlestick neeDataSet to have a body of greater size than the first candlestick.
This second candlestick also neeDataSet to have little to no shadow. The last candlestick is another bullish candlestick that neeDataSet to be equal or greater length of a body than the second candlestick.
When all three candlesticks appear, this chart pattern can be used to confirm the start of a new uptrend.
*/

func (cs *CandleStick) IsThreeWhiteSoldiers() bool {

  if cs.LastIndex - 2 < 0 { return false }

  // Üç tane beyaz olmalı
  if cs.IsWhite() && cs.DataSet[:cs.LastIndex - 1].NewCandleStick().IsWhite() && cs.DataSet[:cs.LastIndex - 2].NewCandleStick().IsWhite() == false { return false }

  // Her üçü de vücut ağırlıklı olmalı
  if cs.IsFuzzyEqual(cs.LastData.Open, cs.LastData.Low) && cs.IsFuzzyEqual(cs.LastData.Close, cs.LastData.High) == false { return false }
  if cs.IsFuzzyEqual(cs.PrevData.Open, cs.PrevData.Low) && cs.IsFuzzyEqual(cs.PrevData.Close, cs.PrevData.High) == false { return false }
  if cs.IsFuzzyEqual(cs.DataSet[cs.LastIndex - 2].Open, cs.DataSet[cs.LastIndex - 2].Low) && cs.IsFuzzyEqual(cs.DataSet[cs.LastIndex - 2].Close, cs.DataSet[cs.LastIndex - 2].High) == false { return false }

  // Sağdan sola boyutları küçülerek gitmeli
  if math.Abs(cs.LastData.Open - cs.LastData.Close) > math.Abs(cs.PrevData.Open - cs.PrevData.Close) == false { return false }
  if math.Abs(cs.PrevData.Open - cs.PrevData.Close) > math.Abs(cs.DataSet[cs.LastIndex - 2].Open - cs.DataSet[cs.LastIndex - 2].Close) == false { return false }

  // Sağdan sola kapanışlar azalarak gitmeli
  if cs.LastData.Close > cs.PrevData.Close && cs.PrevData.Close > cs.DataSet[cs.LastIndex - 2].Close == false { return false }

  // Sağdan sola açılışlar azalarak gitmeli
  if cs.LastData.Open > cs.PrevData.Open && cs.PrevData.Open > cs.DataSet[cs.LastIndex - 2].Open == false { return false }

  return true
}

// A bullish belt-hold line is a tall white candle that has very little or no lower shadow and little or no upper shadow
// In a downtrend, the Low becomes the support. If you are short, then time to take profit
// Steve Nison. The Candlestick Course (A Marketplace Book) (Kindle Locations 623-624). Kindle Edition.

func (cs *CandleStick) IsBullishBeltHold() bool {

  if cs.IsWhite() == false { return false }

  if cs.IsFuzzyEqual(cs.LastData.Open, cs.LastData.Low) {

    if cs.IsFuzzyEqual(cs.LastData.Close, cs.LastData.High) {

      return true
    } else {

      RealBody := math.Abs(cs.LastData.Open - cs.LastData.Close)
      UpperShadow := math.Abs(cs.LastData.Close - cs.LastData.High)

      if 4.0 * UpperShadow <= RealBody { return true }
    }
  }

  return false
}

// A bearish belt-hold line is a long, black real body that opens at the high of the session and closes at or near the low of the session.
// It has small or nonexistent upper or lower shadows.
// In an uptrend, the High becomes the resistance. If you are Long, then time to take profit
// Steve Nison. The Candlestick Course (A Marketplace Book) (Kindle Locations 624-625). Kindle Edition.

func (cs *CandleStick) IsBearishBeltHold() bool {

  if cs.IsBlack() == false { return false }

  if cs.IsFuzzyEqual(cs.LastData.Open, cs.LastData.High) {

    if cs.IsFuzzyEqual(cs.LastData.Close, cs.LastData.Low) {

      return true
    } else {

      RealBody := math.Abs(cs.LastData.Open - cs.LastData.Close)
      LowerShadow := math.Abs(cs.LastData.Close - cs.LastData.Low)

      if 4.0 * LowerShadow <= RealBody { return true }
    }
  }

  return false
}
