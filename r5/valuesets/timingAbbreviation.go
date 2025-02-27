// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/timing-abbreviation
- http://terminology.hl7.org/CodeSystem/v3-GTSAbbreviation
  - BID
  - TID
  - QID
  - AM
  - PM
  - QD
  - QOD
  - Q1H
  - Q2H
  - Q3H
  - Q4H
  - Q6H
  - Q8H
  - BED
  - WK
  - MO
*/type TimingAbbreviation string

const (
	// Every morning at institution specified times.
	TimingAbbreviationAm TimingAbbreviation = "AM"
	// At bedtime (institution specified time).
	TimingAbbreviationAtBedtime TimingAbbreviation = "BED"
	// Two times a day at institution specified time
	TimingAbbreviationBid TimingAbbreviation = "BID"
	// Monthly at institution specified time.
	TimingAbbreviationMonthly TimingAbbreviation = "MO"
	// Every afternoon at institution specified times.
	TimingAbbreviationPm TimingAbbreviation = "PM"
	// Every hour at institution specified times.
	TimingAbbreviationQ1h TimingAbbreviation = "Q1H"
	// Every 2 hours at institution specified times.
	TimingAbbreviationQ2h TimingAbbreviation = "Q2H"
	// Every 3 hours at institution specified times.
	TimingAbbreviationQ3h TimingAbbreviation = "Q3H"
	// Every 4 hours at institution specified time
	TimingAbbreviationQ4h TimingAbbreviation = "Q4H"
	// Every 6 hours at institution specified time
	TimingAbbreviationQ6h TimingAbbreviation = "Q6H"
	// Every 8 hours at institution specified times.
	TimingAbbreviationQ8h TimingAbbreviation = "Q8H"
	// Every day at institution specified times.
	TimingAbbreviationQd TimingAbbreviation = "QD"
	// Four times a day at institution specified time
	TimingAbbreviationQid TimingAbbreviation = "QID"
	// Every other day at institution specified times.
	TimingAbbreviationQod TimingAbbreviation = "QOD"
	// Three times a day at institution specified time
	TimingAbbreviationTid TimingAbbreviation = "TID"
	// Weekly at institution specified time.
	TimingAbbreviationWeekly TimingAbbreviation = "WK"
)
