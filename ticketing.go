// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package gosdk

type Ticketing struct {
	Tickets     *Tickets
	Users       *Users
	Accounts    *Accounts
	Contacts    *Contacts
	Collections *Collections
	Comments    *Comments
	Tags        *Tags
	Teams       *Teams
	Attachments *PanoraTicketingAttachments

	sdkConfiguration sdkConfiguration
}

func newTicketing(sdkConfig sdkConfiguration) *Ticketing {
	return &Ticketing{
		sdkConfiguration: sdkConfig,
		Tickets:          newTickets(sdkConfig),
		Users:            newUsers(sdkConfig),
		Accounts:         newAccounts(sdkConfig),
		Contacts:         newContacts(sdkConfig),
		Collections:      newCollections(sdkConfig),
		Comments:         newComments(sdkConfig),
		Tags:             newTags(sdkConfig),
		Teams:            newTeams(sdkConfig),
		Attachments:      newPanoraTicketingAttachments(sdkConfig),
	}
}
