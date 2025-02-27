// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/goal-achievement
*/type GoalAchievementStatus string

const (
	// The goal is being sought but has not yet been reached. (Also applies if the goal was reached in the past but there has been regression and the goal is again being sought).
	GoalAchievementStatusInProgress GoalAchievementStatus = "in-progress"
	// The goal has been met.
	GoalAchievementStatusAchieved GoalAchievementStatus = "achieved"
	// The goal has not been met and there might or might not have been progress towards target.
	GoalAchievementStatusNotAchieved GoalAchievementStatus = "not-achieved"
)
