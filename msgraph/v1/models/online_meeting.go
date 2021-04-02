package models

// OnlineMeeting struct for OnlineMeeting
type OnlineMeeting struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//// Specifies who can be a presenter in a meeting. Possible values are listed in the following table.
	//AllowedPresenters *OnlineMeetingPresenters `json:"allowedPresenters,omitempty"`
	//// The phone access (dial-in) information for an online meeting. Read-only.
	//AudioConferencing *AudioConferencing `json:"audioConferencing,omitempty"`
	//// The chat information associated with this online meeting.
	//ChatInfo *ChatInfo `json:"chatInfo,omitempty"`
	//// The meeting creation time in UTC. Read-only.
	//CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
	//// The meeting end time in UTC.
	//EndDateTime *time.Time     `json:"endDateTime,omitempty"`
	//ExternalId  *string `json:"externalId,omitempty"`
	//// Whether or not to announce when callers join or leave.
	//IsEntryExitAnnounced *bool `json:"isEntryExitAnnounced,omitempty"`
	//// The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
	//JoinInformation *ItemBody `json:"joinInformation,omitempty"`
	//// The join URL of the online meeting. Read-only.
	//JoinWebUrl *string `json:"joinWebUrl,omitempty"`
	//// Specifies which participants can bypass the meeting   lobby.
	//LobbyBypassSettings *LobbyBypassSettings `json:"lobbyBypassSettings,omitempty"`
	//// The participants associated with the online meeting.  This includes the organizer and the attendees.
	//Participants *MeetingParticipants `json:"participants,omitempty"`
	//// The meeting start time in UTC.
	//StartDateTime *time.Time `json:"startDateTime,omitempty"`
	//// The subject of the online meeting.
	//Subject *string `json:"subject,omitempty"`
	//// The video teleconferencing ID. Read-only.
	//VideoTeleconferenceId *string `json:"videoTeleconferenceId,omitempty"`
}
