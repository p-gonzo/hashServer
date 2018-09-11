//Package state is used to manage the Application state and provide an interface to manage it
package state

var totalConnections = 0
var averageHashTime = 0

//SubmitHashTimeAndUpdateAverage does a thing
func SubmitHashTimeAndUpdateAverage(microSeconds int64) {
	// fmt.Println(strconv.FormatInt(microSeconds, 10))
	microSecondsAsInt := int(microSeconds)
	var weightedHashTime = averageHashTime * totalConnections
	totalConnections++
	averageHashTime = (weightedHashTime + microSecondsAsInt) / totalConnections
}

//GetHashStats does a thing
func GetHashStats() (int, int) {
	return totalConnections, averageHashTime
}
