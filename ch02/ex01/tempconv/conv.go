package tempconv

// CToF は摂氏の温度を華氏へ変換します
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToCは華氏の温度を摂氏へ変換します
func FToC(f Fahrenheit) Celsius { return Celsius(f-32) * 5 / 9 }

// KToC　はケルビンの温度を摂氏へ変換します
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

// CToK は摂氏の温度をケルビンへ変換します
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToF はケルビンの温度を華氏へ変換します
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
