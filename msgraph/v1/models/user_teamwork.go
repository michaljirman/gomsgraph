package models

import (
	"encoding/json"
	"fmt"
)

// UserTeamwork struct for UserTeamwork
type UserTeamwork struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The apps installed in the personal scope of this user.
	InstalledApps *[]UserScopeTeamsAppInstallation `json:"installedApps,omitempty"`
}

// UserScopeTeamsAppInstallation struct for UserScopeTeamsAppInstallation
type UserScopeTeamsAppInstallation struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The app that is installed.
	TeamsApp *TeamsApp `json:"teamsApp,omitempty"`
	// The details of this version of the app.
	TeamsAppDefinition *TeamsAppDefinition `json:"teamsAppDefinition,omitempty"`
	// The chat between the user and Teams app.
	Chat *Chat `json:"chat,omitempty"`
}

// TeamsApp struct for TeamsApp
type TeamsApp struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.
	DisplayName *string `json:"displayName,omitempty"`
	// The method of distribution for the app. Read-only.
	DistributionMethod *TeamsAppDistributionMethod `json:"distributionMethod,omitempty"`
	// The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.
	ExternalId *string `json:"externalId,omitempty"`
	// The details for each version of the app.
	AppDefinitions *[]TeamsAppDefinition `json:"appDefinitions,omitempty"`
}

// TeamsAppDefinition struct for TeamsAppDefinition
type TeamsAppDefinition struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The name of the app provided by the app developer.
	DisplayName *string `json:"displayName,omitempty"`
	// The ID from the Teams app manifest.
	TeamsAppId *string `json:"teamsAppId,omitempty"`
	// The version number of the application.
	Version *string `json:"version,omitempty"`
}

// Chat struct for Chat
type Chat struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
}

// TeamsAppDistributionMethod the model 'TeamsAppDistributionMethod'
type TeamsAppDistributionMethod string

// List of microsoft.graph.teamsAppDistributionMethod
const (
	AppDistributionMethod_STORE                TeamsAppDistributionMethod = "store"
	AppDistributionMethod_ORGANIZATION         TeamsAppDistributionMethod = "organization"
	AppDistributionMethod_SIDELOADED           TeamsAppDistributionMethod = "sideloaded"
	AppDistributionMethod_UNKNOWN_FUTURE_VALUE TeamsAppDistributionMethod = "unknownFutureValue"
)

func (v *TeamsAppDistributionMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TeamsAppDistributionMethod(value)
	for _, existing := range []TeamsAppDistributionMethod{"store", "organization", "sideloaded", "unknownFutureValue"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TeamsAppDistributionMethod", value)
}
