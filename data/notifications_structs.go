package data

type Notifications struct {
	Notifications []Notification `json:"notifications"`
}

type Notification struct {
	AlarmLabel                string `json:"alarmLabel"`
	AlarmTime                 int    `json:"alarmTime"`
	CreatedDate               int64  `json:"createdDate"`
	DeferredAtTime            string `json:"deferredAtTime"`
	DeviceName                string `json:"deviceName"`
	DeviceSerialNumber        string `json:"deviceSerialNumber"`
	Id                        string `json:"id"`
	LastOccurrenceTimeInMilli int64  `json:"lastOccurrenceTimeInMilli"`
	LastTriggerTimeInUtc      string `json:"lastTriggerTimeInUtc"`
	LastUpdatedDate           int64  `json:"lastUpdatedDate"`
	LoopCount                 int    `json:"loopCount"`
	OriginalDate              string `json:"originalDate"`
	OriginalDurationInMillis  int64  `json:"originalDurationInMillis"`
	OriginalTime              string `json:"originalTime"`
	RemainingTime             int64  `json:"remainingTime"`
	ReminderLabel             string `json:"reminderLabel"`
	SnoozedToTime             string `json:"snoozedToTime"`
	Status                    string `json:"status"`
	TimerLabel                string `json:"timerLabel"`
	TriggerTime               int64  `json:"triggerTime"`
	Type                      string `json:"type"`
}
