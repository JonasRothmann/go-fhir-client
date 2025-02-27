// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/observation-statistics
*/type StatisticsCode string

const (
	// The [mean](https://en.wikipedia.org/wiki/Arithmetic_mean) of N measurements over the stated period.
	StatisticsCodeAverage StatisticsCode = "average"
	// The [maximum](https://en.wikipedia.org/wiki/Maximal_element) value of N measurements over the stated period.
	StatisticsCodeMaximum StatisticsCode = "maximum"
	// The [minimum](https://en.wikipedia.org/wiki/Minimal_element) value of N measurements over the stated period.
	StatisticsCodeMinimum StatisticsCode = "minimum"
	// The [number] of valid measurements over the stated period that contributed to the other statistical outputs.
	StatisticsCodeCount StatisticsCode = "count"
	// The total [number] of valid measurements over the stated period, including observations that were ignored because they did not contain valid result values.
	StatisticsCodeTotalCount StatisticsCode = "total-count"
	// The [median](https://en.wikipedia.org/wiki/Median) of N measurements over the stated period.
	StatisticsCodeMedian StatisticsCode = "median"
	// The [standard deviation](https://en.wikipedia.org/wiki/Standard_deviation) of N measurements over the stated period.
	StatisticsCodeStandardDeviation StatisticsCode = "std-dev"
	// The [sum](https://en.wikipedia.org/wiki/Summation) of N measurements over the stated period.
	StatisticsCodeSum StatisticsCode = "sum"
	// The [variance](https://en.wikipedia.org/wiki/Variance) of N measurements over the stated period.
	StatisticsCodeVariance StatisticsCode = "variance"
	// The 20th [Percentile](https://en.wikipedia.org/wiki/Percentile) of N measurements over the stated period.
	StatisticsCode20thPercentile StatisticsCode = "20-percent"
	// The 80th [Percentile](https://en.wikipedia.org/wiki/Percentile) of N measurements over the stated period.
	StatisticsCode80thPercentile StatisticsCode = "80-percent"
	// The lower [Quartile](https://en.wikipedia.org/wiki/Quartile) Boundary of N measurements over the stated period.
	StatisticsCodeLowerQuartile StatisticsCode = "4-lower"
	// The upper [Quartile](https://en.wikipedia.org/wiki/Quartile) Boundary of N measurements over the stated period.
	StatisticsCodeUpperQuartile StatisticsCode = "4-upper"
	// The difference between the upper and lower [Quartiles](https://en.wikipedia.org/wiki/Quartile) is called the Interquartile range. (IQR = Q3-Q1) Quartile deviation or Semi-interquartile range is one-half the difference between the first and the third quartiles.
	StatisticsCodeQuartileDeviation StatisticsCode = "4-dev"
	// The lowest of four values that divide the N measurements into a frequency distribution of five classes with each containing one fifth of the total population.
	StatisticsCode1stQuintile StatisticsCode = "5-1"
	// The second of four values that divide the N measurements into a frequency distribution of five classes with each containing one fifth of the total population.
	StatisticsCode2ndQuintile StatisticsCode = "5-2"
	// The third of four values that divide the N measurements into a frequency distribution of five classes with each containing one fifth of the total population.
	StatisticsCode3rdQuintile StatisticsCode = "5-3"
	// The fourth of four values that divide the N measurements into a frequency distribution of five classes with each containing one fifth of the total population.
	StatisticsCode4thQuintile StatisticsCode = "5-4"
	// Skewness is a measure of the asymmetry of the probability distribution of a real-valued random variable about its mean. The skewness value can be positive or negative, or even undefined.  Source: [Wikipedia](https://en.wikipedia.org/wiki/Skewness).
	StatisticsCodeSkew StatisticsCode = "skew"
	// Kurtosis  is a measure of the "tailedness" of the probability distribution of a real-valued random variable.   Source: [Wikipedia](https://en.wikipedia.org/wiki/Kurtosis).
	StatisticsCodeKurtosis StatisticsCode = "kurtosis"
	// Linear regression is an approach for modeling two-dimensional sample points with one independent variable and one dependent variable (conventionally, the x and y coordinates in a Cartesian coordinate system) and finds a linear function (a non-vertical straight line) that, as accurately as possible, predicts the dependent variable values as a function of the independent variables. Source: [Wikipedia](https://en.wikipedia.org/wiki/Simple_linear_regression)  This Statistic code will return both a gradient and an intercept value.
	StatisticsCodeRegression StatisticsCode = "regression"
)
