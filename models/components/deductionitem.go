// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DeductionItem struct {
	// The name of the deduction
	Name *string `json:"name,omitempty"`
	// The amount of employee deduction
	EmployeeDeduction *float64 `json:"employee_deduction,omitempty"`
	// The amount of company deduction
	CompanyDeduction *float64 `json:"company_deduction,omitempty"`
}

func (o *DeductionItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *DeductionItem) GetEmployeeDeduction() *float64 {
	if o == nil {
		return nil
	}
	return o.EmployeeDeduction
}

func (o *DeductionItem) GetCompanyDeduction() *float64 {
	if o == nil {
		return nil
	}
	return o.CompanyDeduction
}
