package main

import (
	"fmt"
	"strconv"
	"strings"
)

type StringMeeting struct {
	Start string
	End   string
}

type timeint struct {
	startTime int
	endTime int
}

func main() {
	// Define Person 1's calendar as a slice of StringMeeting
	calendar1 := []StringMeeting{
		{"7:00", "7:45"},
		{"8:15", "8:30"},
		{"9:00", "10:30"},
		{"12:00", "14:00"},
		{"14:00", "15:00"},
		{"15:15", "15:30"},
		{"16:30", "18:30"},
		{"20:00", "21:00"},
	}

	// Define Person 2's calendar as a slice of StringMeeting
	calendar2 := []StringMeeting{
		{"9:00", "10:00"},
		{"11:15", "11:30"},
		{"11:45", "17:00"},
		{"17:30", "19:00"},
		{"20:00", "22:15"},
	}

	// Define working hours for Person 1
	workingHours1 := StringMeeting{"6:30", "22:00"}

	// Define working hours for Person 2
	workingHours2 := StringMeeting{"8:00", "22:30"}

	// Define the meeting duration in minutes
	meetingDuration := 30

	val := CalendarMatching(calendar1, workingHours1, calendar2, workingHours2, meetingDuration)
	fmt.Println(val)
}

func CalendarMatching(
	calendar1 []StringMeeting, dailyBounds1 StringMeeting,
	calendar2 []StringMeeting, dailyBounds2 StringMeeting,
	meetingDuration int,
) []StringMeeting {
	// Write your code here.
	// converting calender1 to mins in calendar1Time
	calendar1Time := make([]timeint, len(calendar1))
	
	for i, v := range calendar1 {
		calendar1Time[i].startTime = convertToMins(v.Start)
		calendar1Time[i].endTime = convertToMins(v.End)
	}
	fmt.Println("calendar1Time -> ", calendar1Time)
	//converting calendar2 to mins in calendar2Time
	calendar2Time := make([]timeint, len(calendar2))
	for i, v := range calendar2 {
		calendar2Time[i].startTime = convertToMins(v.Start)
		calendar2Time[i].endTime = convertToMins(v.End)
	}
	fmt.Println("calendar2Time -> ", calendar2Time)

	// func to get adujsted working bound times for both the person to be available for meeting
	adjustedBounds := adujstedWorkingBounds(dailyBounds1, dailyBounds2)
	fmt.Println("adjustedBound -> ", adjustedBounds)

	// mergeing both persons working hours 
	mergedTimeslots := mergeOccupiedBounds(calendar1Time, calendar2Time)
	//availableSlots := []timeint{}

	if len(mergedTimeslots) == 0 {
		// if their is no meeting, then available slot is full working hours
		if adjustedBounds.endTime - adjustedBounds.startTime >= meetingDuration{
			return []StringMeeting{
				{
					Start: minsToTime(adjustedBounds.startTime),
					End: minsToTime(adjustedBounds.endTime),
				},
			}
		}
		return []StringMeeting{}
	}
	
	fmt.Println("mergedTimeslots -> ", mergedTimeslots)

	availableSlots := findAvailableSlots(mergedTimeslots, adjustedBounds, meetingDuration)

	return availableSlots
}

func findAvailableSlots(mergedTimeslots []timeint, bounds timeint, meetingDuration int) []StringMeeting {
	availableSlots := []StringMeeting{}
	fmt.Println("bounds.startTime, mergedTimeslots[0].starttime ", bounds.startTime, mergedTimeslots[0], availableSlots)
	// check if there is any free time between the start of the working hours & first meeting
	if bounds.startTime < mergedTimeslots[0].startTime{
		if mergedTimeslots[0].startTime - bounds.startTime >= meetingDuration {
			availableSlots = append(availableSlots, StringMeeting{
				Start: minsToTime(bounds.startTime),
				End: minsToTime(mergedTimeslots[0].startTime),
			})
		}
	}

	// check for gaps between merged meetings
	for i := 1; i < len(mergedTimeslots)-1; i++ {
		start := mergedTimeslots[i].endTime
		end := mergedTimeslots[i+1].startTime
		if end-start >= meetingDuration {
			availableSlots = append(availableSlots, StringMeeting{
				Start: minsToTime(start),
				End: minsToTime(end),
			})
		}
	}

	//check if their is any free time between the last meeting and end of the working day
	lastMeetingEnd := mergedTimeslots[len(mergedTimeslots) - 1].endTime
	if lastMeetingEnd < bounds.endTime {
		if bounds.endTime - lastMeetingEnd >= meetingDuration {
			availableSlots = append(availableSlots, StringMeeting{
				Start: minsToTime(lastMeetingEnd),
				End: minsToTime(bounds.endTime),
			})
		}
	}

	return availableSlots
}


func mergeOccupiedBounds(calendar1Time []timeint, calendar2Time []timeint) []timeint {
	mergedBookedBounds := make([]timeint, 0, len(calendar1Time) + len(calendar2Time))
	i, j := 0, 0

	for i < len(calendar1Time) && j < len(calendar2Time) {
		if calendar1Time[i].startTime < calendar2Time[j].startTime {
			mergedBookedBounds = append(mergedBookedBounds, calendar1Time[i])
			i++
		} else {
			mergedBookedBounds = append(mergedBookedBounds, calendar2Time[j])
			j++
		} 
	}

	// check for remaining elements in calendersTime & also merge them
	for i < len(calendar1Time) {
		mergedBookedBounds = append(mergedBookedBounds, calendar1Time[i])
		i++
	}

	for j < len(calendar2Time) {
		mergedBookedBounds = append(mergedBookedBounds, calendar2Time[j])
		j++
	}

	if len(mergedBookedBounds) == 0 {
		return mergedBookedBounds
	}

	finalMergedBookedBounds := make([]timeint, 0)
	prev := mergedBookedBounds[0]
	for curr := 1; curr < len(mergedBookedBounds); curr++ {
		if prev.endTime >= mergedBookedBounds[curr].startTime {
			// update endime to be maximum of prev & curr end time since its overlapping
			if mergedBookedBounds[curr].endTime > prev.endTime {
				prev.endTime = mergedBookedBounds[curr].endTime
			}
		} else {
			finalMergedBookedBounds = append(finalMergedBookedBounds, prev)
			prev = mergedBookedBounds[curr]
		}
	}

	finalMergedBookedBounds = append(finalMergedBookedBounds, prev)

	fmt.Println("mergedBookedBounds ", mergedBookedBounds, finalMergedBookedBounds)
	return finalMergedBookedBounds
}

func adujstedWorkingBounds(dailyBounds1 StringMeeting, dailyBounds2 StringMeeting) timeint {
	// convert both persons bounds to ints

	dailyBounds1Time := timeint{}
	dailyBounds1Time.startTime = convertToMins(dailyBounds1.Start)
	dailyBounds1Time.endTime = convertToMins(dailyBounds1.End)

	dailyBounds2Time := timeint{}
	dailyBounds2Time.startTime = convertToMins(dailyBounds2.Start)
	dailyBounds2Time.endTime = convertToMins(dailyBounds2.End)

	fmt.Println(dailyBounds1Time, dailyBounds2Time)
	adjustedBound := timeint{
		startTime: max(dailyBounds1Time.startTime, dailyBounds2Time.startTime),
		endTime: min(dailyBounds1Time.endTime, dailyBounds2Time.endTime),
	}
	
	return adjustedBound
}

func convertToMins(timeStr string) int {
	timesArr := strings.Split(timeStr, ":")
	if len(timesArr) != 2 {
		fmt.Println("Invalid time format. Expected HH:MM")
		return 0 // or handle it differently based on your requirements
	}

	// Convert hour part to integer
	hourPart, err := strconv.Atoi(timesArr[0])
	if err != nil {
		fmt.Println("Hour part conversion error:", err)
		return 0 // return an error code or handle it differently
	}

	// Convert minute part to integer
	minPart, err := strconv.Atoi(timesArr[1])
	if err != nil {
		fmt.Println("Minute part conversion error:", err)
		return 0 // return an error code or handle it differently
	}

	//fmt.Println(hourPart)
	return hourPart * 60 + minPart
}

func minsToTime(mins int) string {
	hours := mins / 60
	minutes := mins % 60
	return fmt.Sprintf("%d:%02d", hours, minutes)
}

func max (a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min (a, b int) int {
	if a < b {
		return a
	}
	return b
}