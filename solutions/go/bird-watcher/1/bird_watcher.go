package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    count := 0
    for i := 0; i < len(birdsPerDay); i++ {
        count += birdsPerDay[i]
    }
    return count
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    count := 0
    start, end := 7 * (week - 1), 7 * (week)
    Weekly := birdsPerDay[start:end]
    for i := 0; i < len(Weekly); i++ {
        count += Weekly[i]
    }
    return count
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i++ {
        if i % 2 == 0 {
            birdsPerDay[i] = birdsPerDay[i] + 1
        }
    }
    return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}
