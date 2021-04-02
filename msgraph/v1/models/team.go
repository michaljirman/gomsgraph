package models

import (
	"encoding/json"
	"fmt"
)

// Team struct for Team
type Team struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.
	Classification *string `json:"classification,omitempty"`
	// An optional description for the team.
	Description *string `json:"description,omitempty"`
	// The name of the team.
	DisplayName *string `json:"displayName,omitempty"`
	// Settings to configure use of Giphy, memes, and stickers in the team.
	FunSettings *TeamFunSettings `json:"funSettings,omitempty"`
	// Settings to configure whether guests can create, update, or delete channels in the team.
	GuestSettings *TeamGuestSettings `json:"guestSettings,omitempty"`
	// A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.
	InternalId *string `json:"internalId,omitempty"`
	// Whether this team is in read-only mode.
	IsArchived *bool `json:"isArchived,omitempty"`
	// Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team.
	MemberSettings *TeamMemberSettings `json:"memberSettings,omitempty"`
	// Settings to configure messaging and mentions in the team.
	MessagingSettings *TeamMessagingSettings `json:"messagingSettings,omitempty"`
	// Optional. Indicates whether the team is intended for a particular use case.  Each team specialization has access to unique behaviors and experiences targeted to its use case.
	Specialization *TeamSpecialization `json:"specialization,omitempty"`
	// The visibility of the group and team. Defaults to Public.
	Visibility *TeamVisibilityType `json:"visibility,omitempty"`
	// A hyperlink that will go to the team in the Microsoft Teams client. This is the URL that you get when you right-click a team in the Microsoft Teams client and select Get link to team. This URL should be treated as an opaque blob, and not parsed.
	WebUrl *string `json:"webUrl,omitempty"`
	// The schedule of shifts for this team.
	Schedule *Schedule `json:"schedule,omitempty"`
	// The collection of channels & messages associated with the team.
	Channels *[]Channel `json:"channels,omitempty"`
	Group    *Group     `json:"group,omitempty"`
	// The apps installed in this team.
	InstalledApps *[]TeamsAppInstallation `json:"installedApps,omitempty"`
	// Members and owners of the team.
	Members *[]ConversationMember `json:"members,omitempty"`
	// The async operations that ran or are running on this team.
	Operations *[]TeamsAsyncOperation `json:"operations,omitempty"`
	// The general channel for the team.
	PrimaryChannel *Channel `json:"primaryChannel,omitempty"`
	// The template this team was created from. See available templates.
	Template *TeamsTemplate `json:"template,omitempty"`
}

// TeamFunSettings struct for TeamFunSettings
type TeamFunSettings struct {
	// If set to true, enables users to include custom memes.
	AllowCustomMemes *bool `json:"allowCustomMemes,omitempty"`
	// If set to true, enables Giphy use.
	AllowGiphy *bool `json:"allowGiphy,omitempty"`
	// If set to true, enables users to include stickers and memes.
	AllowStickersAndMemes *bool `json:"allowStickersAndMemes,omitempty"`
	// Giphy content rating. Possible values are: moderate, strict.
	GiphyContentRating *GiphyRatingType `json:"giphyContentRating,omitempty"`
}

// TeamGuestSettings struct for TeamGuestSettings
type TeamGuestSettings struct {
	// If set to true, guests can add and update channels.
	AllowCreateUpdateChannels *bool `json:"allowCreateUpdateChannels,omitempty"`
	// If set to true, guests can delete channels.
	AllowDeleteChannels *bool `json:"allowDeleteChannels,omitempty"`
}

// TeamMemberSettings struct for TeamMemberSettings
type TeamMemberSettings struct {
	// If set to true, members can add and remove apps.
	AllowAddRemoveApps *bool `json:"allowAddRemoveApps,omitempty"`
	// If set to true, members can add and update private channels.
	AllowCreatePrivateChannels *bool `json:"allowCreatePrivateChannels,omitempty"`
	// If set to true, members can add and update channels.
	AllowCreateUpdateChannels *bool `json:"allowCreateUpdateChannels,omitempty"`
	// If set to true, members can add, update, and remove connectors.
	AllowCreateUpdateRemoveConnectors *bool `json:"allowCreateUpdateRemoveConnectors,omitempty"`
	// If set to true, members can add, update, and remove tabs.
	AllowCreateUpdateRemoveTabs *bool `json:"allowCreateUpdateRemoveTabs,omitempty"`
	// If set to true, members can delete channels.
	AllowDeleteChannels *bool `json:"allowDeleteChannels,omitempty"`
}

// TeamMessagingSettings struct for TeamMessagingSettings
type TeamMessagingSettings struct {
	// If set to true, @channel mentions are allowed.
	AllowChannelMentions *bool `json:"allowChannelMentions,omitempty"`
	// If set to true, owners can delete any message.
	AllowOwnerDeleteMessages *bool `json:"allowOwnerDeleteMessages,omitempty"`
	// If set to true, @team mentions are allowed.
	AllowTeamMentions *bool `json:"allowTeamMentions,omitempty"`
	// If set to true, users can delete their messages.
	AllowUserDeleteMessages *bool `json:"allowUserDeleteMessages,omitempty"`
	// If set to true, users can edit their messages.
	AllowUserEditMessages *bool `json:"allowUserEditMessages,omitempty"`
}

// TeamSpecialization the model 'TeamSpecialization'
type TeamSpecialization string

// List of microsoft.graph.teamSpecialization
const (
	TeamSpecialization_NONE                                      TeamSpecialization = "none"
	TeamSpecialization_EDUCATION_STANDARD                        TeamSpecialization = "educationStandard"
	TeamSpecialization_EDUCATION_CLASS                           TeamSpecialization = "educationClass"
	TeamSpecialization_EDUCATION_PROFESSIONAL_LEARNING_COMMUNITY TeamSpecialization = "educationProfessionalLearningCommunity"
	TeamSpecialization_EDUCATION_STAFF                           TeamSpecialization = "educationStaff"
	TeamSpecialization_HEALTHCARE_STANDARD                       TeamSpecialization = "healthcareStandard"
	TeamSpecialization_HEALTHCARE_CARE_COORDINATION              TeamSpecialization = "healthcareCareCoordination"
	TeamSpecialization_UNKNOWN_FUTURE_VALUE                      TeamSpecialization = "unknownFutureValue"
)

func (v *TeamSpecialization) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TeamSpecialization(value)
	for _, existing := range []TeamSpecialization{"none", "educationStandard", "educationClass", "educationProfessionalLearningCommunity", "educationStaff", "healthcareStandard", "healthcareCareCoordination", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TeamSpecialization", value)
}

// TeamVisibilityType the model 'TeamVisibilityType'
type TeamVisibilityType string

// List of microsoft.graph.teamVisibilityType
const (
	TeamVisibilityType_PRIVATE              TeamVisibilityType = "private"
	TeamVisibilityType_PUBLIC               TeamVisibilityType = "public"
	TeamVisibilityType_HIDDEN_MEMBERSHIP    TeamVisibilityType = "hiddenMembership"
	TeamVisibilityType_UNKNOWN_FUTURE_VALUE TeamVisibilityType = "unknownFutureValue"
)

func (v *TeamVisibilityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TeamVisibilityType(value)
	for _, existing := range []TeamVisibilityType{"private", "public", "hiddenMembership", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TeamVisibilityType", value)
}

// Schedule struct for Schedule
type Schedule struct {
	//// Read-only.
	//Id *string `json:"id,omitempty"`
	//// Indicates whether the schedule is enabled for the team. Required.
	//Enabled *bool `json:"enabled,omitempty"`
	//// Indicates whether offer shift requests are enabled for the schedule.
	//OfferShiftRequestsEnabled *bool `json:"offerShiftRequestsEnabled,omitempty"`
	//// Indicates whether open shifts are enabled for the schedule.
	//OpenShiftsEnabled *bool `json:"openShiftsEnabled,omitempty"`
	//// The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
	//ProvisionStatus NullableAnyOfOperationStatus `json:"provisionStatus,omitempty"`
	//// Additional information about why schedule provisioning failed.
	//ProvisionStatusCode *string `json:"provisionStatusCode,omitempty"`
	//// Indicates whether swap shifts requests are enabled for the schedule.
	//SwapShiftsRequestsEnabled *bool `json:"swapShiftsRequestsEnabled,omitempty"`
	//// Indicates whether time clock is enabled for the schedule.
	//TimeClockEnabled *bool `json:"timeClockEnabled,omitempty"`
	//// Indicates whether time off requests are enabled for the schedule.
	//TimeOffRequestsEnabled *bool `json:"timeOffRequestsEnabled,omitempty"`
	//// Indicates the time zone of the schedule team using tz database format. Required.
	//TimeZone                *string                          `json:"timeZone,omitempty"`
	//WorkforceIntegrationIds *[]string                               `json:"workforceIntegrationIds,omitempty"`
	//OfferShiftRequests      *[]OfferShiftRequest      `json:"offerShiftRequests,omitempty"`
	//OpenShiftChangeRequests *[]OpenShiftChangeRequest `json:"openShiftChangeRequests,omitempty"`
	//OpenShifts              *[]OpenShift              `json:"openShifts,omitempty"`
	//// The logical grouping of users in the schedule (usually by role).
	//SchedulingGroups *[]SchedulingGroup `json:"schedulingGroups,omitempty"`
	//// The shifts in the schedule.
	//Shifts                   *[]Shift                   `json:"shifts,omitempty"`
	//SwapShiftsChangeRequests *[]SwapShiftsChangeRequest `json:"swapShiftsChangeRequests,omitempty"`
	//// The set of reasons for a time off in the schedule.
	//TimeOffReasons  *[]TimeOffReason  `json:"timeOffReasons,omitempty"`
	//TimeOffRequests *[]TimeOffRequest `json:"timeOffRequests,omitempty"`
	//// The instances of times off in the schedule.
	//TimesOff *[]TimeOff `json:"timesOff,omitempty"`
}

// GiphyRatingType the model 'GiphyRatingType'
type GiphyRatingType string

// List of microsoft.graph.giphyRatingType
const (
	STRICT               GiphyRatingType = "strict"
	MODERATE             GiphyRatingType = "moderate"
	UNKNOWN_FUTURE_VALUE GiphyRatingType = "unknownFutureValue"
)

func (v *GiphyRatingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GiphyRatingType(value)
	for _, existing := range []GiphyRatingType{"strict", "moderate", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GiphyRatingType", value)
}

// Channel struct for Channel
type Channel struct {
	//// Read-only.
	//Id *string `json:"id,omitempty"`
	//// Optional textual description for the channel.
	//Description *string `json:"description,omitempty"`
	//// Channel name as it will appear to the user in Microsoft Teams.
	//DisplayName *string `json:"displayName,omitempty"`
	//// The email address for sending messages to the channel. Read-only.
	//Email *string `json:"email,omitempty"`
	//// Indicates whether the channel should automatically be marked 'favorite' for all members of the team. Can only be set programmatically with Create team. Default: false.
	//IsFavoriteByDefault *bool `json:"isFavoriteByDefault,omitempty"`
	//// The type of the channel. Can be set during creation and cannot be changed. Possible values are: standard - Channel inherits the list of members of the parent team; private - Channel can have members that are a subset of all the members on the parent team.
	//MembershipType NullableAnyOfChannelMembershipType `json:"membershipType,omitempty"`
	//// A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not parsed. Read-only.
	//WebUrl *string `json:"webUrl,omitempty"`
	//// Metadata for the location where the channel's files are stored.
	//FilesFolder NullableAnyOfDriveItem `json:"filesFolder,omitempty"`
	//// A collection of membership records associated with the channel.
	//Members *[]ConversationMember `json:"members,omitempty"`
	//// A collection of all the messages in the channel. A navigation property. Nullable.
	//Messages *[]ChatMessage `json:"messages,omitempty"`
	//// A collection of all the tabs in the channel. A navigation property.
	//Tabs *[]TeamsTab `json:"tabs,omitempty"`
}

// TeamsAppInstallation struct for TeamsAppInstallation
type TeamsAppInstallation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//// The app that is installed.
	//TeamsApp NullableAnyOfTeamsApp `json:"teamsApp,omitempty"`
	//// The details of this version of the app.
	//TeamsAppDefinition NullableAnyOfTeamsAppDefinition `json:"teamsAppDefinition,omitempty"`
}

// ConversationMember struct for ConversationMember
type ConversationMember struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The display name of the user.
	DisplayName *string `json:"displayName,omitempty"`
	// The roles for that user.
	Roles *[]string `json:"roles,omitempty"`
}

// TeamsAsyncOperation struct for TeamsAsyncOperation
type TeamsAsyncOperation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	//// Number of times the operation was attempted before being marked successful or failed.
	//AttemptsCount *int32 `json:"attemptsCount,omitempty"`
	//// Time when the operation was created.
	//CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	//// Any error that causes the async operation to fail.
	//Error NullableAnyOfOperationError `json:"error,omitempty"`
	//// Time when the async operation was last updated.
	//LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	//// Denotes which type of operation is being described.
	//OperationType *AnyOfTeamsAsyncOperationType `json:"operationType,omitempty"`
	//// Operation status.
	//Status *AnyOfTeamsAsyncOperationStatus `json:"status,omitempty"`
	//// The ID of the object that's created or modified as result of this async operation, typically a team.
	//TargetResourceId *string `json:"targetResourceId,omitempty"`
	//// The location of the object that's created or modified as result of this async operation. This URL should be treated as an opaque value and not parsed into its component paths.
	//TargetResourceLocation *string `json:"targetResourceLocation,omitempty"`
}

// TeamsTemplate struct for TeamsTemplate
type TeamsTemplate struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
}
