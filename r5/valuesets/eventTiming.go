// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/event-timing
- http://terminology.hl7.org/CodeSystem/v3-TimingEvent
  - HS
  - WAKE
  - C
  - CM
  - CD
  - CV
  - AC
  - ACM
  - ACD
  - ACV
  - PC
  - PCM
  - PCD
  - PCV
*/type EventTiming string

const (
	// Event occurs during the morning. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingMorning EventTiming = "MORN"
	// Event occurs during the early morning. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingEarlyMorning EventTiming = "MORN.early"
	// Event occurs during the late morning. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingLateMorning EventTiming = "MORN.late"
	// Event occurs around 12:00pm. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingNoon EventTiming = "NOON"
	// Event occurs during the afternoon. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingAfternoon EventTiming = "AFT"
	// Event occurs during the early afternoon. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingEarlyAfternoon EventTiming = "AFT.early"
	// Event occurs during the late afternoon. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingLateAfternoon EventTiming = "AFT.late"
	// Event occurs during the evening. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingEvening EventTiming = "EVE"
	// Event occurs during the early evening. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingEarlyEvening EventTiming = "EVE.early"
	// Event occurs during the late evening. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingLateEvening EventTiming = "EVE.late"
	// Event occurs during the night. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingNight EventTiming = "NIGHT"
	// Event occurs [offset] after subject goes to sleep. The exact time is unspecified and established by institution convention or patient interpretation.
	EventTimingAfterSleep EventTiming = "PHS"
	// Event occurs a single time (with no repetitions) as soon as possible after the scheduled or actual start of the overall event.
	EventTimingImmediate EventTiming = "IMD"
	// before meal (from lat. ante cibus)
	EventTimingAc EventTiming = "AC"
	// before lunch (from lat. ante cibus diurnus)
	EventTimingAcd EventTiming = "ACD"
	// before breakfast (from lat. ante cibus matutinus)
	EventTimingAcm EventTiming = "ACM"
	// before dinner (from lat. ante cibus vespertinus)
	EventTimingAcv EventTiming = "ACV"
	// meal (from lat. ante cibus)
	EventTimingC EventTiming = "C"
	// Prior to beginning a regular period of extended sleep (this would exclude naps). Note that this might occur at different times of day depending on a person's regular sleep schedule.
	EventTimingHs EventTiming = "HS"
	// after meal (from lat. post cibus)
	EventTimingPc EventTiming = "PC"
	// after lunch (from lat. post cibus diurnus)
	EventTimingPcd EventTiming = "PCD"
	// after breakfast (from lat. post cibus matutinus)
	EventTimingPcm EventTiming = "PCM"
	// after dinner (from lat. post cibus vespertinus)
	EventTimingPcv EventTiming = "PCV"
	// Upon waking up from a regular period of sleep, in order to start regular activities (this would exclude waking up from a nap or temporarily waking up during a period of sleep)
	EventTimingWake EventTiming = "WAKE"
)
