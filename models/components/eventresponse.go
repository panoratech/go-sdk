// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// EventResponseType - Scope of the event
type EventResponseType string

const (
	EventResponseTypeCrmContactCreated           EventResponseType = "crm.contact.created"
	EventResponseTypeCrmContactPulled            EventResponseType = "crm.contact.pulled"
	EventResponseTypeCrmCompanyCreated           EventResponseType = "crm.company.created"
	EventResponseTypeCrmCompanyPulled            EventResponseType = "crm.company.pulled"
	EventResponseTypeCrmDealCreated              EventResponseType = "crm.deal.created"
	EventResponseTypeCrmDealPulled               EventResponseType = "crm.deal.pulled"
	EventResponseTypeCrmEngagementCreated        EventResponseType = "crm.engagement.created"
	EventResponseTypeCrmEngagementPulled         EventResponseType = "crm.engagement.pulled"
	EventResponseTypeCrmNoteCreated              EventResponseType = "crm.note.created"
	EventResponseTypeCrmNotePulled               EventResponseType = "crm.note.pulled"
	EventResponseTypeCrmStagePulled              EventResponseType = "crm.stage.pulled"
	EventResponseTypeCrmTaskPulled               EventResponseType = "crm.task.pulled"
	EventResponseTypeCrmTaskCreated              EventResponseType = "crm.task.created"
	EventResponseTypeCrmUserPulled               EventResponseType = "crm.user.pulled"
	EventResponseTypeTicketingTicketCreated      EventResponseType = "ticketing.ticket.created"
	EventResponseTypeTicketingTicketPulled       EventResponseType = "ticketing.ticket.pulled"
	EventResponseTypeTicketingCommentCreated     EventResponseType = "ticketing.comment.created"
	EventResponseTypeTicketingCommentPulled      EventResponseType = "ticketing.comment.pulled"
	EventResponseTypeTicketingAttachmentCreated  EventResponseType = "ticketing.attachment.created"
	EventResponseTypeTicketingAttachmentPulled   EventResponseType = "ticketing.attachment.pulled"
	EventResponseTypeTicketingCollectionPulled   EventResponseType = "ticketing.collection.pulled"
	EventResponseTypeTicketingAccountPulled      EventResponseType = "ticketing.account.pulled"
	EventResponseTypeTicketingContactPulled      EventResponseType = "ticketing.contact.pulled"
	EventResponseTypeTicketingTagPulled          EventResponseType = "ticketing.tag.pulled"
	EventResponseTypeTicketingTeamPulled         EventResponseType = "ticketing.team.pulled"
	EventResponseTypeTicketingUserPulled         EventResponseType = "ticketing.user.pulled"
	EventResponseTypeAtsActivityCreated          EventResponseType = "ats.activity.created"
	EventResponseTypeAtsActivityPulled           EventResponseType = "ats.activity.pulled"
	EventResponseTypeAtsApplicationCreated       EventResponseType = "ats.application.created"
	EventResponseTypeAtsApplicationPulled        EventResponseType = "ats.application.pulled"
	EventResponseTypeAtsAttachmentCreated        EventResponseType = "ats.attachment.created"
	EventResponseTypeAtsAttachmentPulled         EventResponseType = "ats.attachment.pulled"
	EventResponseTypeAtsCandidateCreated         EventResponseType = "ats.candidate.created"
	EventResponseTypeAtsCandidatePulled          EventResponseType = "ats.candidate.pulled"
	EventResponseTypeAtsDepartmentPulled         EventResponseType = "ats.department.pulled"
	EventResponseTypeAtsEecosPulled              EventResponseType = "ats.eecos.pulled"
	EventResponseTypeAtsInterviewCreated         EventResponseType = "ats.interview.created"
	EventResponseTypeAtsInterviewPulled          EventResponseType = "ats.interview.pulled"
	EventResponseTypeAtsJobPulled                EventResponseType = "ats.job.pulled"
	EventResponseTypeAtsJobinterviewstagePulled  EventResponseType = "ats.jobinterviewstage.pulled"
	EventResponseTypeAtsOfferCreated             EventResponseType = "ats.offer.created"
	EventResponseTypeAtsOfficePulled             EventResponseType = "ats.office.pulled"
	EventResponseTypeAtsRejectreasonPulled       EventResponseType = "ats.rejectreason.pulled"
	EventResponseTypeAtsScorecardPulled          EventResponseType = "ats.scorecard.pulled"
	EventResponseTypeAtsTagPulled                EventResponseType = "ats.tag.pulled"
	EventResponseTypeAtsUserPulled               EventResponseType = "ats.user.pulled"
	EventResponseTypeFilestorageFileCreated      EventResponseType = "filestorage.file.created"
	EventResponseTypeFilestorageFilePulled       EventResponseType = "filestorage.file.pulled"
	EventResponseTypeFilestorageFolderCreated    EventResponseType = "filestorage.folder.created"
	EventResponseTypeFilestorageFolderPulled     EventResponseType = "filestorage.folder.pulled"
	EventResponseTypeFilestorageGroupPulled      EventResponseType = "filestorage.group.pulled"
	EventResponseTypeFilestorageUserPulled       EventResponseType = "filestorage.user.pulled"
	EventResponseTypeFilestorageDrivePulled      EventResponseType = "filestorage.drive.pulled"
	EventResponseTypeFilestoragePermissionPulled EventResponseType = "filestorage.permission.pulled"
	EventResponseTypeFilestorageSharedlinkPulled EventResponseType = "filestorage.sharedlink.pulled"
	EventResponseTypeConnectionCreated           EventResponseType = "connection.created"
)

func (e EventResponseType) ToPointer() *EventResponseType {
	return &e
}
func (e *EventResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "crm.contact.created":
		fallthrough
	case "crm.contact.pulled":
		fallthrough
	case "crm.company.created":
		fallthrough
	case "crm.company.pulled":
		fallthrough
	case "crm.deal.created":
		fallthrough
	case "crm.deal.pulled":
		fallthrough
	case "crm.engagement.created":
		fallthrough
	case "crm.engagement.pulled":
		fallthrough
	case "crm.note.created":
		fallthrough
	case "crm.note.pulled":
		fallthrough
	case "crm.stage.pulled":
		fallthrough
	case "crm.task.pulled":
		fallthrough
	case "crm.task.created":
		fallthrough
	case "crm.user.pulled":
		fallthrough
	case "ticketing.ticket.created":
		fallthrough
	case "ticketing.ticket.pulled":
		fallthrough
	case "ticketing.comment.created":
		fallthrough
	case "ticketing.comment.pulled":
		fallthrough
	case "ticketing.attachment.created":
		fallthrough
	case "ticketing.attachment.pulled":
		fallthrough
	case "ticketing.collection.pulled":
		fallthrough
	case "ticketing.account.pulled":
		fallthrough
	case "ticketing.contact.pulled":
		fallthrough
	case "ticketing.tag.pulled":
		fallthrough
	case "ticketing.team.pulled":
		fallthrough
	case "ticketing.user.pulled":
		fallthrough
	case "ats.activity.created":
		fallthrough
	case "ats.activity.pulled":
		fallthrough
	case "ats.application.created":
		fallthrough
	case "ats.application.pulled":
		fallthrough
	case "ats.attachment.created":
		fallthrough
	case "ats.attachment.pulled":
		fallthrough
	case "ats.candidate.created":
		fallthrough
	case "ats.candidate.pulled":
		fallthrough
	case "ats.department.pulled":
		fallthrough
	case "ats.eecos.pulled":
		fallthrough
	case "ats.interview.created":
		fallthrough
	case "ats.interview.pulled":
		fallthrough
	case "ats.job.pulled":
		fallthrough
	case "ats.jobinterviewstage.pulled":
		fallthrough
	case "ats.offer.created":
		fallthrough
	case "ats.office.pulled":
		fallthrough
	case "ats.rejectreason.pulled":
		fallthrough
	case "ats.scorecard.pulled":
		fallthrough
	case "ats.tag.pulled":
		fallthrough
	case "ats.user.pulled":
		fallthrough
	case "filestorage.file.created":
		fallthrough
	case "filestorage.file.pulled":
		fallthrough
	case "filestorage.folder.created":
		fallthrough
	case "filestorage.folder.pulled":
		fallthrough
	case "filestorage.group.pulled":
		fallthrough
	case "filestorage.user.pulled":
		fallthrough
	case "filestorage.drive.pulled":
		fallthrough
	case "filestorage.permission.pulled":
		fallthrough
	case "filestorage.sharedlink.pulled":
		fallthrough
	case "connection.created":
		*e = EventResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EventResponseType: %v", v)
	}
}

// EventResponseStatus - Status of the event
type EventResponseStatus string

const (
	EventResponseStatusSuccess EventResponseStatus = "success"
	EventResponseStatusFail    EventResponseStatus = "fail"
)

func (e EventResponseStatus) ToPointer() *EventResponseStatus {
	return &e
}
func (e *EventResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "fail":
		*e = EventResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EventResponseStatus: %v", v)
	}
}

// Method - HTTP method used for the event
type Method string

const (
	MethodGet    Method = "GET"
	MethodPost   Method = "POST"
	MethodPut    Method = "PUT"
	MethodDelete Method = "DELETE"
)

func (e Method) ToPointer() *Method {
	return &e
}
func (e *Method) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GET":
		fallthrough
	case "POST":
		fallthrough
	case "PUT":
		fallthrough
	case "DELETE":
		*e = Method(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Method: %v", v)
	}
}

type EventResponse struct {
	// Unique identifier for the event
	IDEvent string `json:"id_event"`
	// Connection ID associated with the event
	IDConnection string `json:"id_connection"`
	// Project ID associated with the event
	IDProject string `json:"id_project"`
	// Scope of the event
	Type EventResponseType `json:"type"`
	// Status of the event
	Status EventResponseStatus `json:"status"`
	// Direction of the event
	Direction string `json:"direction"`
	// HTTP method used for the event
	Method Method `json:"method"`
	// URL associated with the event
	URL string `json:"url"`
	// Provider associated with the event
	Provider string `json:"provider"`
	// Timestamp of the event
	Timestamp time.Time `json:"timestamp"`
	// Linked user ID associated with the event
	IDLinkedUser string `json:"id_linked_user"`
}

func (e EventResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EventResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EventResponse) GetIDEvent() string {
	if o == nil {
		return ""
	}
	return o.IDEvent
}

func (o *EventResponse) GetIDConnection() string {
	if o == nil {
		return ""
	}
	return o.IDConnection
}

func (o *EventResponse) GetIDProject() string {
	if o == nil {
		return ""
	}
	return o.IDProject
}

func (o *EventResponse) GetType() EventResponseType {
	if o == nil {
		return EventResponseType("")
	}
	return o.Type
}

func (o *EventResponse) GetStatus() EventResponseStatus {
	if o == nil {
		return EventResponseStatus("")
	}
	return o.Status
}

func (o *EventResponse) GetDirection() string {
	if o == nil {
		return ""
	}
	return o.Direction
}

func (o *EventResponse) GetMethod() Method {
	if o == nil {
		return Method("")
	}
	return o.Method
}

func (o *EventResponse) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *EventResponse) GetProvider() string {
	if o == nil {
		return ""
	}
	return o.Provider
}

func (o *EventResponse) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

func (o *EventResponse) GetIDLinkedUser() string {
	if o == nil {
		return ""
	}
	return o.IDLinkedUser
}
