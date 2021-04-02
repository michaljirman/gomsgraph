package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// TodoTask struct for TodoTask
type TodoTask struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The task body that typically contains information about the task.
	Body *ItemBody `json:"body,omitempty"`
	// The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
	BodyLastModifiedDateTime *time.Time `json:"bodyLastModifiedDateTime,omitempty"`
	// The date in the specified time zone that the task was finished.
	CompletedDateTime *DateTimeTimeZone `json:"completedDateTime,omitempty"`
	// The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The date in the specified time zone that the task is to be finished.
	DueDateTime *DateTimeTimeZone `json:"dueDateTime,omitempty"`
	// The importance of the task. Possible values are: low, normal, high.
	Importance **Importance `json:"importance,omitempty"`
	// Set to true if an alert is set to remind the user of the task.
	IsReminderOn *bool `json:"isReminderOn,omitempty"`
	// The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The recurrence pattern for the task.
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// The date and time for a reminder alert of the task to occur.
	ReminderDateTime *DateTimeTimeZone `json:"reminderDateTime,omitempty"`
	// Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed, waitingOnOthers, deferred.
	Status **TaskStatus `json:"status,omitempty"`
	// A brief description of the task.
	Title *string `json:"title,omitempty"`
	// The collection of open extensions defined for the task. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`
	// A collection of resources linked to the task.
	LinkedResources *[]LinkedResource `json:"linkedResources,omitempty"`
}

// TaskStatus the model 'TaskStatus'
type TaskStatus string

// List of microsoft.graph.taskStatus
const (
	TaskStatus_NOT_STARTED       TaskStatus = "notStarted"
	TaskStatus_IN_PROGRESS       TaskStatus = "inProgress"
	TaskStatus_COMPLETED         TaskStatus = "completed"
	TaskStatus_WAITING_ON_OTHERS TaskStatus = "waitingOnOthers"
	TaskStatus_DEFERRED          TaskStatus = "deferred"
)

func (v *TaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskStatus(value)
	for _, existing := range []TaskStatus{"notStarted", "inProgress", "completed", "waitingOnOthers", "deferred"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskStatus", value)
}

// LinkedResource struct for LinkedResource
type LinkedResource struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Field indicating the app name of the source that is sending the linkedResource.
	ApplicationName *string `json:"applicationName,omitempty"`
	// Field indicating the title of the linkedResource.
	DisplayName *string `json:"displayName,omitempty"`
	// Id of the object that is associated with this task on the third-party/partner system.
	ExternalId *string `json:"externalId,omitempty"`
	// Deep link to the linkedResource.
	WebUrl *string `json:"webUrl,omitempty"`
}
