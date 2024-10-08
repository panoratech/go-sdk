// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package gosdk

type Crm struct {
	Companies   *Companies
	Contacts    *PanoraContacts
	Deals       *Deals
	Engagements *Engagements
	Notes       *Notes
	Stages      *Stages
	Tasks       *Tasks
	Users       *PanoraUsers

	sdkConfiguration sdkConfiguration
}

func newCrm(sdkConfig sdkConfiguration) *Crm {
	return &Crm{
		sdkConfiguration: sdkConfig,
		Companies:        newCompanies(sdkConfig),
		Contacts:         newPanoraContacts(sdkConfig),
		Deals:            newDeals(sdkConfig),
		Engagements:      newEngagements(sdkConfig),
		Notes:            newNotes(sdkConfig),
		Stages:           newStages(sdkConfig),
		Tasks:            newTasks(sdkConfig),
		Users:            newPanoraUsers(sdkConfig),
	}
}
