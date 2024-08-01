# github.com/panoratech/go-sdk

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=<no value>&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasy.com/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasy.com/docs/advanced-setup/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/panoratech/go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	gosdk "github.com/panoratech/go-sdk"
	"log"
	"os"
)

func main() {
	s := gosdk.New(
		gosdk.WithSecurity(os.Getenv("BEARER")),
	)

	ctx := context.Background()
	res, err := s.Hello(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Panora SDK](docs/sdks/panora/README.md)

* [Hello](docs/sdks/panora/README.md#hello)
* [Health](docs/sdks/panora/README.md#health)

### [Webhooks](docs/sdks/webhooks/README.md)

* [List](docs/sdks/webhooks/README.md#list) - List webhooks
* [Create](docs/sdks/webhooks/README.md#create) - Add webhook metadata

### [Webhooks.{id}](docs/sdks/id/README.md)

* [Delete](docs/sdks/id/README.md#delete) - Delete Webhook
* [UpdateStatus](docs/sdks/id/README.md#updatestatus) - Update webhook status

### [Webhooks.Verifyevent](docs/sdks/verifyevent/README.md)

* [VerifyEvent](docs/sdks/verifyevent/README.md#verifyevent) - Verify payload signature of the webhook


### [Ticketing.Tickets](docs/sdks/tickets/README.md)

* [List](docs/sdks/tickets/README.md#list) - List  Tickets
* [Create](docs/sdks/tickets/README.md#create) - Create Tickets
* [Retrieve](docs/sdks/tickets/README.md#retrieve) - Retrieve Tickets

### [Ticketing.Users](docs/sdks/users/README.md)

* [List](docs/sdks/users/README.md#list) - List  Users
* [Retrieve](docs/sdks/users/README.md#retrieve) - Retrieve Users

### [Ticketing.Accounts](docs/sdks/accounts/README.md)

* [List](docs/sdks/accounts/README.md#list) - List  Accounts
* [Retrieve](docs/sdks/accounts/README.md#retrieve) - Retrieve Accounts

### [Ticketing.Contacts](docs/sdks/contacts/README.md)

* [List](docs/sdks/contacts/README.md#list) - List all Contacts
* [Retrieve](docs/sdks/contacts/README.md#retrieve) - Retrieve Contacts

### [Ticketing.Collections](docs/sdks/collections/README.md)

* [List](docs/sdks/collections/README.md#list) - List  Collections
* [Retrieve](docs/sdks/collections/README.md#retrieve) - Retrieve Collections

### [Ticketing.Comments](docs/sdks/comments/README.md)

* [List](docs/sdks/comments/README.md#list) - List  Comments
* [Create](docs/sdks/comments/README.md#create) - Create Comments
* [Retrieve](docs/sdks/comments/README.md#retrieve) - Retrieve Comments

### [Ticketing.Tags](docs/sdks/tags/README.md)

* [List](docs/sdks/tags/README.md#list) - List  Tags
* [Retrieve](docs/sdks/tags/README.md#retrieve) - Retrieve Tags

### [Ticketing.Teams](docs/sdks/teams/README.md)

* [List](docs/sdks/teams/README.md#list) - List  Teams
* [Retrieve](docs/sdks/teams/README.md#retrieve) - Retrieve Teams

### [Ticketing.Attachments](docs/sdks/panoraticketingattachments/README.md)

* [List](docs/sdks/panoraticketingattachments/README.md#list) - List  Attachments
* [Create](docs/sdks/panoraticketingattachments/README.md#create) - Create Attachments
* [Retrieve](docs/sdks/panoraticketingattachments/README.md#retrieve) - Retrieve Attachments

### [Sync](docs/sdks/sync/README.md)

* [Status](docs/sdks/sync/README.md#status) - Retrieve sync status of a certain vertical
* [Resync](docs/sdks/sync/README.md#resync) - Resync common objects across a vertical


### [Crm.Companies](docs/sdks/companies/README.md)

* [List](docs/sdks/companies/README.md#list) - List  Companies
* [Create](docs/sdks/companies/README.md#create) - Create Companies
* [Retrieve](docs/sdks/companies/README.md#retrieve) - Retrieve Companies

### [Crm.Contacts](docs/sdks/panoracontacts/README.md)

* [List](docs/sdks/panoracontacts/README.md#list) - List CRM Contacts
* [Create](docs/sdks/panoracontacts/README.md#create) - Create Contacts
* [Retrieve](docs/sdks/panoracontacts/README.md#retrieve) - Retrieve Contacts

### [Crm.Deals](docs/sdks/deals/README.md)

* [List](docs/sdks/deals/README.md#list) - List  Deals
* [Create](docs/sdks/deals/README.md#create) - Create Deals
* [Retrieve](docs/sdks/deals/README.md#retrieve) - Retrieve Deals

### [Crm.Engagements](docs/sdks/engagements/README.md)

* [List](docs/sdks/engagements/README.md#list) - List  Engagements
* [Create](docs/sdks/engagements/README.md#create) - Create Engagements
* [Retrieve](docs/sdks/engagements/README.md#retrieve) - Retrieve Engagements

### [Crm.Notes](docs/sdks/notes/README.md)

* [List](docs/sdks/notes/README.md#list) - List  Notes
* [Create](docs/sdks/notes/README.md#create) - Create Notes
* [Retrieve](docs/sdks/notes/README.md#retrieve) - Retrieve Notes

### [Crm.Stages](docs/sdks/stages/README.md)

* [List](docs/sdks/stages/README.md#list) - List  Stages
* [Retrieve](docs/sdks/stages/README.md#retrieve) - Retrieve Stages

### [Crm.Tasks](docs/sdks/tasks/README.md)

* [List](docs/sdks/tasks/README.md#list) - List  Tasks
* [Create](docs/sdks/tasks/README.md#create) - Create Tasks
* [Retrieve](docs/sdks/tasks/README.md#retrieve) - Retrieve Tasks

### [Crm.Users](docs/sdks/panorausers/README.md)

* [List](docs/sdks/panorausers/README.md#list) - List  Users
* [Retrieve](docs/sdks/panorausers/README.md#retrieve) - Retrieve Users

### [LinkedUsers](docs/sdks/linkedusers/README.md)

* [Create](docs/sdks/linkedusers/README.md#create) - Create Linked Users
* [List](docs/sdks/linkedusers/README.md#list) - List Linked Users

### [LinkedUsers.Batch](docs/sdks/batch/README.md)

* [ImportBatch](docs/sdks/batch/README.md#importbatch) - Add Batch Linked Users

### [LinkedUsers.Single](docs/sdks/single/README.md)

* [Retrieve](docs/sdks/single/README.md#retrieve) - Retrieve Linked Users

### [LinkedUsers.Fromremoteid](docs/sdks/fromremoteid/README.md)

* [RemoteID](docs/sdks/fromremoteid/README.md#remoteid) - Retrieve a Linked User From A Remote Id

### [FieldMappings](docs/sdks/fieldmappings/README.md)

* [DefineCustomField](docs/sdks/fieldmappings/README.md#definecustomfield) - Create Custom Field

### [FieldMappings.Define](docs/sdks/define/README.md)

* [Definitions](docs/sdks/define/README.md#definitions) - Define target Field

### [FieldMappings.Map](docs/sdks/map/README.md)

* [Map](docs/sdks/map/README.md#map) - Map Custom Field

### [Passthrough](docs/sdks/passthrough/README.md)

* [Request](docs/sdks/passthrough/README.md#request) - Make a passthrough request


### [Hris.Bankinfos](docs/sdks/bankinfos/README.md)

* [List](docs/sdks/bankinfos/README.md#list) - List  Bankinfos
* [Retrieve](docs/sdks/bankinfos/README.md#retrieve) - Retrieve Bank Infos

### [Hris.Benefits](docs/sdks/benefits/README.md)

* [List](docs/sdks/benefits/README.md#list) - List  Benefits
* [Retrieve](docs/sdks/benefits/README.md#retrieve) - Retrieve Benefits

### [Hris.Companies](docs/sdks/panoracompanies/README.md)

* [List](docs/sdks/panoracompanies/README.md#list) - List  Companys

### [Hris.Dependents](docs/sdks/dependents/README.md)

* [List](docs/sdks/dependents/README.md#list) - List  Dependents
* [Retrieve](docs/sdks/dependents/README.md#retrieve) - Retrieve Dependents

### [Hris.Employeepayrollruns](docs/sdks/employeepayrollruns/README.md)

* [List](docs/sdks/employeepayrollruns/README.md#list) - List  EmployeePayrollRuns
* [Retrieve](docs/sdks/employeepayrollruns/README.md#retrieve) - Retrieve Employee Payroll Runs

### [Hris.Employees](docs/sdks/employees/README.md)

* [List](docs/sdks/employees/README.md#list) - List  Employees
* [Create](docs/sdks/employees/README.md#create) - Create Employees
* [Retrieve](docs/sdks/employees/README.md#retrieve) - Retrieven Employees

### [Hris.Employerbenefits](docs/sdks/employerbenefits/README.md)

* [List](docs/sdks/employerbenefits/README.md#list) - List  EmployerBenefits
* [Retrieve](docs/sdks/employerbenefits/README.md#retrieve) - Retrieve Employer Benefits

### [Hris.Employments](docs/sdks/employments/README.md)

* [List](docs/sdks/employments/README.md#list) - List  Employments
* [Retrieve](docs/sdks/employments/README.md#retrieve) - Retrieve Employments

### [Hris.Groups](docs/sdks/groups/README.md)

* [List](docs/sdks/groups/README.md#list) - List  Groups
* [Retrieve](docs/sdks/groups/README.md#retrieve) - Retrieve Groups

### [Hris.Locations](docs/sdks/locations/README.md)

* [List](docs/sdks/locations/README.md#list) - List  Locations
* [Retrieve](docs/sdks/locations/README.md#retrieve) - Retrieve Locations

### [Hris.Paygroups](docs/sdks/paygroups/README.md)

* [List](docs/sdks/paygroups/README.md#list) - List  PayGroups
* [Retrieve](docs/sdks/paygroups/README.md#retrieve) - Retrieve Pay Groups

### [Hris.Payrollruns](docs/sdks/payrollruns/README.md)

* [List](docs/sdks/payrollruns/README.md#list) - List  PayrollRuns

### [Hris.Timeoffs](docs/sdks/timeoffs/README.md)

* [List](docs/sdks/timeoffs/README.md#list) - List  Timeoffs
* [Create](docs/sdks/timeoffs/README.md#create) - Create Timeoffs
* [Retrieve](docs/sdks/timeoffs/README.md#retrieve) - Retrieve Timeoffs

### [Hris.Timeoffbalances](docs/sdks/timeoffbalances/README.md)

* [List](docs/sdks/timeoffbalances/README.md#list) - List  TimeoffBalances
* [Retrieve](docs/sdks/timeoffbalances/README.md#retrieve) - Retrieve Time off Balances


### [Marketingautomation.Actions](docs/sdks/actions/README.md)

* [List](docs/sdks/actions/README.md#list) - List  Actions
* [Create](docs/sdks/actions/README.md#create) - Create Action
* [Retrieve](docs/sdks/actions/README.md#retrieve) - Retrieve Actions

### [Marketingautomation.Automations](docs/sdks/automations/README.md)

* [List](docs/sdks/automations/README.md#list) - List  Automations
* [Create](docs/sdks/automations/README.md#create) - Create Automation
* [Retrieve](docs/sdks/automations/README.md#retrieve) - Retrieve Automations

### [Marketingautomation.Campaigns](docs/sdks/campaigns/README.md)

* [List](docs/sdks/campaigns/README.md#list) - List  Campaigns
* [Create](docs/sdks/campaigns/README.md#create) - Create Campaign
* [Retrieve](docs/sdks/campaigns/README.md#retrieve) - Retrieve Campaigns

### [Marketingautomation.Contacts](docs/sdks/panoramarketingautomationcontacts/README.md)

* [List](docs/sdks/panoramarketingautomationcontacts/README.md#list) - List  Contacts
* [Create](docs/sdks/panoramarketingautomationcontacts/README.md#create) - Create Contact
* [Retrieve](docs/sdks/panoramarketingautomationcontacts/README.md#retrieve) - Retrieve Contacts

### [Marketingautomation.Emails](docs/sdks/emails/README.md)

* [List](docs/sdks/emails/README.md#list) - List  Emails
* [Retrieve](docs/sdks/emails/README.md#retrieve) - Retrieve Emails

### [Marketingautomation.Events](docs/sdks/events/README.md)

* [List](docs/sdks/events/README.md#list) - List  Events
* [Retrieve](docs/sdks/events/README.md#retrieve) - Retrieve Events

### [Marketingautomation.Lists](docs/sdks/lists/README.md)

* [List](docs/sdks/lists/README.md#list) - List  Lists
* [Create](docs/sdks/lists/README.md#create) - Create Lists
* [Retrieve](docs/sdks/lists/README.md#retrieve) - Retrieve Lists

### [Marketingautomation.Messages](docs/sdks/messages/README.md)

* [List](docs/sdks/messages/README.md#list) - List  Messages
* [Retrieve](docs/sdks/messages/README.md#retrieve) - Retrieve Messages

### [Marketingautomation.Templates](docs/sdks/templates/README.md)

* [List](docs/sdks/templates/README.md#list) - List  Templates
* [Create](docs/sdks/templates/README.md#create) - Create Template
* [Retrieve](docs/sdks/templates/README.md#retrieve) - Retrieve Templates

### [Marketingautomation.Users](docs/sdks/panoramarketingautomationusers/README.md)

* [List](docs/sdks/panoramarketingautomationusers/README.md#list) - List  Users
* [Retrieve](docs/sdks/panoramarketingautomationusers/README.md#retrieve) - Retrieve Users


### [Ats.Activities](docs/sdks/activities/README.md)

* [List](docs/sdks/activities/README.md#list) - List  Activities
* [Create](docs/sdks/activities/README.md#create) - Create Activities
* [Retrieve](docs/sdks/activities/README.md#retrieve) - Retrieve Activities

### [Ats.Applications](docs/sdks/applications/README.md)

* [List](docs/sdks/applications/README.md#list) - List  Applications
* [Create](docs/sdks/applications/README.md#create) - Create Applications
* [Retrieve](docs/sdks/applications/README.md#retrieve) - Retrieve Applications

### [Ats.Attachments](docs/sdks/attachments/README.md)

* [List](docs/sdks/attachments/README.md#list) - List  Attachments
* [Create](docs/sdks/attachments/README.md#create) - Create Attachments
* [Retrieve](docs/sdks/attachments/README.md#retrieve) - Retrieve Attachments

### [Ats.Candidates](docs/sdks/candidates/README.md)

* [List](docs/sdks/candidates/README.md#list) - List  Candidates
* [Create](docs/sdks/candidates/README.md#create) - Create Candidates
* [Retrieve](docs/sdks/candidates/README.md#retrieve) - Retrieve Candidates

### [Ats.Departments](docs/sdks/departments/README.md)

* [List](docs/sdks/departments/README.md#list) - List  Departments
* [Retrieve](docs/sdks/departments/README.md#retrieve) - Retrieve Departments

### [Ats.Interviews](docs/sdks/interviews/README.md)

* [List](docs/sdks/interviews/README.md#list) - List  Interviews
* [Create](docs/sdks/interviews/README.md#create) - Create Interviews
* [Retrieve](docs/sdks/interviews/README.md#retrieve) - Retrieve Interviews

### [Ats.Jobinterviewstages](docs/sdks/jobinterviewstages/README.md)

* [List](docs/sdks/jobinterviewstages/README.md#list) - List  JobInterviewStages
* [Retrieve](docs/sdks/jobinterviewstages/README.md#retrieve) - Retrieve Job Interview Stages

### [Ats.Jobs](docs/sdks/jobs/README.md)

* [List](docs/sdks/jobs/README.md#list) - List  Jobs
* [Retrieve](docs/sdks/jobs/README.md#retrieve) - Retrieve Jobs

### [Ats.Offers](docs/sdks/offers/README.md)

* [List](docs/sdks/offers/README.md#list) - List  Offers
* [Retrieve](docs/sdks/offers/README.md#retrieve) - Retrieve Offers

### [Ats.Offices](docs/sdks/offices/README.md)

* [List](docs/sdks/offices/README.md#list) - List Offices
* [Retrieve](docs/sdks/offices/README.md#retrieve) - Retrieve Offices

### [Ats.Rejectreasons](docs/sdks/rejectreasons/README.md)

* [List](docs/sdks/rejectreasons/README.md#list) - List  RejectReasons
* [Retrieve](docs/sdks/rejectreasons/README.md#retrieve) - Retrieve Reject Reasons

### [Ats.Scorecards](docs/sdks/scorecards/README.md)

* [List](docs/sdks/scorecards/README.md#list) - List  ScoreCards
* [Retrieve](docs/sdks/scorecards/README.md#retrieve) - Retrieve Score Cards

### [Ats.Tags](docs/sdks/panoratags/README.md)

* [List](docs/sdks/panoratags/README.md#list) - List  Tags
* [Retrieve](docs/sdks/panoratags/README.md#retrieve) - Retrieve Tags

### [Ats.Users](docs/sdks/panoraatsusers/README.md)

* [List](docs/sdks/panoraatsusers/README.md#list) - List  Users
* [Retrieve](docs/sdks/panoraatsusers/README.md#retrieve) - Retrieve Users

### [Ats.Eeocs](docs/sdks/eeocs/README.md)

* [List](docs/sdks/eeocs/README.md#list) - List  Eeocss
* [Retrieve](docs/sdks/eeocs/README.md#retrieve) - Retrieve Eeocs


### [Accounting.Accounts](docs/sdks/panoraaccounts/README.md)

* [List](docs/sdks/panoraaccounts/README.md#list) - List  Accounts
* [Create](docs/sdks/panoraaccounts/README.md#create) - Create Accounts
* [Retrieve](docs/sdks/panoraaccounts/README.md#retrieve) - Retrieve Accounts

### [Accounting.Addresses](docs/sdks/addresses/README.md)

* [List](docs/sdks/addresses/README.md#list) - List  Addresss
* [Retrieve](docs/sdks/addresses/README.md#retrieve) - Retrieve Addresses

### [Accounting.Attachments](docs/sdks/panoraattachments/README.md)

* [List](docs/sdks/panoraattachments/README.md#list) - List  Attachments
* [Create](docs/sdks/panoraattachments/README.md#create) - Create Attachments
* [Retrieve](docs/sdks/panoraattachments/README.md#retrieve) - Retrieve Attachments

### [Accounting.Balancesheets](docs/sdks/balancesheets/README.md)

* [List](docs/sdks/balancesheets/README.md#list) - List  BalanceSheets
* [Retrieve](docs/sdks/balancesheets/README.md#retrieve) - Retrieve BalanceSheets

### [Accounting.Cashflowstatements](docs/sdks/cashflowstatements/README.md)

* [List](docs/sdks/cashflowstatements/README.md#list) - List  CashflowStatements
* [Retrieve](docs/sdks/cashflowstatements/README.md#retrieve) - Retrieve Cashflow Statements

### [Accounting.Companyinfos](docs/sdks/companyinfos/README.md)

* [List](docs/sdks/companyinfos/README.md#list) - List  CompanyInfos
* [Retrieve](docs/sdks/companyinfos/README.md#retrieve) - Retrieve Company Infos

### [Accounting.Contacts](docs/sdks/panoraaccountingcontacts/README.md)

* [List](docs/sdks/panoraaccountingcontacts/README.md#list) - List  Contacts
* [Create](docs/sdks/panoraaccountingcontacts/README.md#create) - Create Contacts
* [Retrieve](docs/sdks/panoraaccountingcontacts/README.md#retrieve) - Retrieve Contacts

### [Accounting.Creditnotes](docs/sdks/creditnotes/README.md)

* [List](docs/sdks/creditnotes/README.md#list) - List  CreditNotes
* [Retrieve](docs/sdks/creditnotes/README.md#retrieve) - Retrieve Credit Notes

### [Accounting.Expenses](docs/sdks/expenses/README.md)

* [List](docs/sdks/expenses/README.md#list) - List  Expenses
* [Create](docs/sdks/expenses/README.md#create) - Create Expenses
* [Retrieve](docs/sdks/expenses/README.md#retrieve) - Retrieve Expenses

### [Accounting.Incomestatements](docs/sdks/incomestatements/README.md)

* [List](docs/sdks/incomestatements/README.md#list) - List  IncomeStatements
* [Retrieve](docs/sdks/incomestatements/README.md#retrieve) - Retrieve Income Statements

### [Accounting.Invoices](docs/sdks/invoices/README.md)

* [List](docs/sdks/invoices/README.md#list) - List  Invoices
* [Create](docs/sdks/invoices/README.md#create) - Create Invoices
* [Retrieve](docs/sdks/invoices/README.md#retrieve) - Retrieve Invoices

### [Accounting.Items](docs/sdks/items/README.md)

* [List](docs/sdks/items/README.md#list) - List  Items
* [Retrieve](docs/sdks/items/README.md#retrieve) - Retrieve Items

### [Accounting.Journalentries](docs/sdks/journalentries/README.md)

* [List](docs/sdks/journalentries/README.md#list) - List  JournalEntrys
* [Create](docs/sdks/journalentries/README.md#create) - Create Journal Entries
* [Retrieve](docs/sdks/journalentries/README.md#retrieve) - Retrieve Journal Entries

### [Accounting.Payments](docs/sdks/payments/README.md)

* [List](docs/sdks/payments/README.md#list) - List  Payments
* [Create](docs/sdks/payments/README.md#create) - Create Payments
* [Retrieve](docs/sdks/payments/README.md#retrieve) - Retrieve Payments

### [Accounting.Phonenumbers](docs/sdks/phonenumbers/README.md)

* [List](docs/sdks/phonenumbers/README.md#list) - List  PhoneNumbers
* [Retrieve](docs/sdks/phonenumbers/README.md#retrieve) - Retrieve Phone Numbers

### [Accounting.Purchaseorders](docs/sdks/purchaseorders/README.md)

* [List](docs/sdks/purchaseorders/README.md#list) - List  PurchaseOrders
* [Create](docs/sdks/purchaseorders/README.md#create) - Create Purchase Orders
* [Retrieve](docs/sdks/purchaseorders/README.md#retrieve) - Retrieve Purchase Orders

### [Accounting.Taxrates](docs/sdks/taxrates/README.md)

* [List](docs/sdks/taxrates/README.md#list) - List  TaxRates
* [Retrieve](docs/sdks/taxrates/README.md#retrieve) - Retrieve Tax Rates

### [Accounting.Trackingcategories](docs/sdks/trackingcategories/README.md)

* [List](docs/sdks/trackingcategories/README.md#list) - List  TrackingCategorys
* [Retrieve](docs/sdks/trackingcategories/README.md#retrieve) - Retrieve Tracking Categories

### [Accounting.Transactions](docs/sdks/transactions/README.md)

* [List](docs/sdks/transactions/README.md#list) - List  Transactions
* [Retrieve](docs/sdks/transactions/README.md#retrieve) - Retrieve Transactions

### [Accounting.Vendorcredits](docs/sdks/vendorcredits/README.md)

* [List](docs/sdks/vendorcredits/README.md#list) - List  VendorCredits
* [Retrieve](docs/sdks/vendorcredits/README.md#retrieve) - Retrieve Vendor Credits


### [Filestorage.Drives](docs/sdks/drives/README.md)

* [List](docs/sdks/drives/README.md#list) - List  Drives
* [Retrieve](docs/sdks/drives/README.md#retrieve) - Retrieve Drives

### [Filestorage.Files](docs/sdks/files/README.md)

* [List](docs/sdks/files/README.md#list) - List  Files
* [Create](docs/sdks/files/README.md#create) - Create Files
* [Retrieve](docs/sdks/files/README.md#retrieve) - Retrieve Files

### [Filestorage.Folders](docs/sdks/folders/README.md)

* [List](docs/sdks/folders/README.md#list) - List  Folders
* [Create](docs/sdks/folders/README.md#create) - Create Folders
* [Retrieve](docs/sdks/folders/README.md#retrieve) - Retrieve Folders

### [Filestorage.Groups](docs/sdks/panoragroups/README.md)

* [List](docs/sdks/panoragroups/README.md#list) - List  Groups
* [Retrieve](docs/sdks/panoragroups/README.md#retrieve) - Retrieve Groups

### [Filestorage.Users](docs/sdks/panorafilestorageusers/README.md)

* [List](docs/sdks/panorafilestorageusers/README.md#list) - List  Users
* [Retrieve](docs/sdks/panorafilestorageusers/README.md#retrieve) - Retrieve Users
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	gosdk "github.com/panoratech/go-sdk"
	"github.com/panoratech/go-sdk/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	s := gosdk.New(
		gosdk.WithSecurity(os.Getenv("BEARER")),
	)

	ctx := context.Background()
	res, err := s.Hello(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	gosdk "github.com/panoratech/go-sdk"
	"github.com/panoratech/go-sdk/retry"
	"log"
	"os"
)

func main() {
	s := gosdk.New(
		gosdk.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		gosdk.WithSecurity(os.Getenv("BEARER")),
	)

	ctx := context.Background()
	res, err := s.Hello(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	gosdk "github.com/panoratech/go-sdk"
	"github.com/panoratech/go-sdk/models/sdkerrors"
	"log"
	"os"
)

func main() {
	s := gosdk.New(
		gosdk.WithSecurity(os.Getenv("BEARER")),
	)

	ctx := context.Background()
	res, err := s.Hello(ctx)
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.panora.dev` | None |
| 1 | `https://api-sandbox.panora.dev` | None |
| 2 | `https://api-dev.panora.dev` | None |

#### Example

```go
package main

import (
	"context"
	gosdk "github.com/panoratech/go-sdk"
	"log"
	"os"
)

func main() {
	s := gosdk.New(
		gosdk.WithServerIndex(2),
		gosdk.WithSecurity(os.Getenv("BEARER")),
	)

	ctx := context.Background()
	res, err := s.Hello(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	gosdk "github.com/panoratech/go-sdk"
	"log"
	"os"
)

func main() {
	s := gosdk.New(
		gosdk.WithServerURL("https://api.panora.dev"),
		gosdk.WithSecurity(os.Getenv("BEARER")),
	)

	ctx := context.Background()
	res, err := s.Hello(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name        | Type        | Scheme      |
| ----------- | ----------- | ----------- |
| `Bearer`    | http        | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	gosdk "github.com/panoratech/go-sdk"
	"log"
	"os"
)

func main() {
	s := gosdk.New(
		gosdk.WithSecurity(os.Getenv("BEARER")),
	)

	ctx := context.Background()
	res, err := s.Hello(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=<no value>&utm_campaign=go)
