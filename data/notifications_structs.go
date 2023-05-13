package data

type Notifications struct {
	Notifications []Notification `json:"notifications"`
}

type Notification struct {
	AlarmLabel                *string `json:"alarmLabel,omitempty"`
	AlarmTime                 int     `json:"alarmTime,omitempty"`
	CreatedDate               int64   `json:"createdDate,omitempty"`
	DeferredAtTime            *string `json:"deferredAtTime,omitempty"`
	DeviceName                *string `json:"deviceName,omitempty"`
	DeviceSerialNumber        string  `json:"deviceSerialNumber,omitempty"`
	Id                        string  `json:"id,omitempty"`
	LastOccurrenceTimeInMilli int64   `json:"lastOccurrenceTimeInMilli,omitempty"`
	LastTriggerTimeInUtc      *string `json:"lastTriggerTimeInUtc,omitempty"`
	LastUpdatedDate           int64   `json:"lastUpdatedDate,omitempty"`
	LoopCount                 *int    `json:"loopCount,omitempty"`
	OriginalDate              string  `json:"originalDate,omitempty"`
	OriginalDurationInMillis  int64   `json:"originalDurationInMillis,omitempty"`
	OriginalTime              string  `json:"originalTime,omitempty"`
	RemainingTime             int64   `json:"remainingTime,omitempty"`
	ReminderLabel             string  `json:"reminderLabel,omitempty"`
	SnoozedToTime             *string `json:"snoozedToTime,omitempty"`
	Status                    string  `json:"status,omitempty"`
	TimerLabel                *string `json:"timerLabel,omitempty"`
	TriggerTime               int64   `json:"triggerTime,omitempty"`
	Type                      string  `json:"type,omitempty"`
}
