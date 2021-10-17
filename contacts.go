//**********************************************************
//
// This file is part of lexoffice.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package golexoffice

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ContactsReturn is to decode json data
type ContactsReturn struct {
	Content          []ContactsReturnContent `json:"content"`
	First            bool                    `json:"first"`
	Last             bool                    `json:"last"`
	TotalPages       int                     `json:"totalPages"`
	TotalElements    int                     `json:"totalElements"`
	NumberOfElements int                     `json:"numberOfElements"`
	Size             int                     `json:"size"`
	Number           int                     `json:"number"`
	Sort             []ContactsReturnSort    `json:"sort"`
}

type ContactsReturnContent struct {
	Id             string                    `json:"id,omitempty"`
	Version        int                       `json:"version,omitempty"`
	Roles          ContactBodyRoles          `json:"roles"`
	Company        ContactBodyCompany        `json:"company"`
	Addresses      ContactBodyAddresses      `json:"addresses"`
	EmailAddresses ContactBodyEmailAddresses `json:"emailAddresses"`
	PhoneNumbers   ContactBodyPhoneNumbers   `json:"phoneNumbers"`
	Note           string                    `json:"note"`
	Archived       bool                      `json:"archived,omitempty"`
}

type ContactsReturnRoles struct {
	Customer ContactsReturnCustomer `json:"customer"`
	Vendor   ContactsReturnVendor   `json:"vendor"`
}

type ContactsReturnCustomer struct {
	Number int `json:"number,omitempty"`
}

type ContactsReturnVendor struct {
	Number int `json:"number,omitempty"`
}

type ContactsReturnCompany struct {
	Name                 string                         `json:"name"`
	TaxNumber            string                         `json:"taxNumber"`
	VatRegistrationId    string                         `json:"vatRegistrationId"`
	AllowTaxFreeInvoices bool                           `json:"allowTaxFreeInvoices"`
	ContactPersons       []ContactsReturnContactPersons `json:"contactPersons"`
}

type ContactsReturnContactPersons struct {
	Salutation   string `json:"salutation"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"phoneNumber"`
}

type ContactsReturnAddresses struct {
	Billing  []ContactsReturnBilling  `json:"billing"`
	Shipping []ContactsReturnShipping `json:"shipping"`
}

type ContactsReturnBilling struct {
	Supplement  string `json:"supplement"`
	Street      string `json:"street"`
	Zip         string `json:"zip"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
}

type ContactsReturnShipping struct {
	Supplement  string `json:"supplement"`
	Street      string `json:"street"`
	Zip         string `json:"zip"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
}

type ContactsReturnEmailAddresses struct {
	Business []string `json:"business"`
	Office   []string `json:"office"`
	Private  []string `json:"private"`
	Other    []string `json:"other"`
}

type ContactsReturnPhoneNumbers struct {
	Business []string `json:"business"`
	Office   []string `json:"office"`
	Mobile   []string `json:"mobile"`
	Private  []string `json:"private"`
	Fax      []string `json:"fax"`
	Other    []string `json:"other"`
}

type ContactsReturnSort struct {
	Property     string `json:"property"`
	Direction    string `json:"direction"`
	IgnoreCase   bool   `json:"ignoreCase"`
	NullHandling string `json:"nullHandling"`
	Ascending    bool   `json:"ascending"`
}

// ContactBody is to create a new contact
type ContactBody struct {
	Id             string                    `json:"id,omitempty"`
	Version        int                       `json:"version,omitempty"`
	Roles          ContactBodyRoles          `json:"roles"`
	Company        ContactBodyCompany        `json:"company"`
	Addresses      ContactBodyAddresses      `json:"addresses"`
	EmailAddresses ContactBodyEmailAddresses `json:"emailAddresses"`
	PhoneNumbers   ContactBodyPhoneNumbers   `json:"phoneNumbers"`
	Note           string                    `json:"note"`
	Archived       bool                      `json:"archived,omitempty"`
}

type ContactBodyRoles struct {
	Customer ContactBodyCustomer `json:"customer"`
	Vendor   ContactBodyVendor   `json:"vendor"`
}

type ContactBodyCustomer struct {
	Number int `json:"number,omitempty"`
}

type ContactBodyVendor struct {
	Number int `json:"number,omitempty"`
}

type ContactBodyCompany struct {
	Name                 string                      `json:"name"`
	TaxNumber            string                      `json:"taxNumber"`
	VatRegistrationId    string                      `json:"vatRegistrationId"`
	AllowTaxFreeInvoices bool                        `json:"allowTaxFreeInvoices"`
	ContactPersons       []ContactBodyContactPersons `json:"contactPersons"`
}

type ContactBodyContactPersons struct {
	Salutation   string `json:"salutation"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"phoneNumber"`
}

type ContactBodyAddresses struct {
	Billing  []ContactBodyBilling  `json:"billing"`
	Shipping []ContactBodyShipping `json:"shipping"`
}

type ContactBodyBilling struct {
	Supplement  string `json:"supplement"`
	Street      string `json:"street"`
	Zip         string `json:"zip"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
}

type ContactBodyShipping struct {
	Supplement  string `json:"supplement"`
	Street      string `json:"street"`
	Zip         string `json:"zip"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
}

type ContactBodyEmailAddresses struct {
	Business []string `json:"business"`
	Office   []string `json:"office"`
	Private  []string `json:"private"`
	Other    []string `json:"other"`
}

type ContactBodyPhoneNumbers struct {
	Business []string `json:"business"`
	Office   []string `json:"office"`
	Mobile   []string `json:"mobile"`
	Private  []string `json:"private"`
	Fax      []string `json:"fax"`
	Other    []string `json:"other"`
}

// ContactReturn is to decode json return
type ContactReturn struct {
	ID          string `json:"id"`
	ResourceUri string `json:"resourceUri"`
	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
	Version     int    `json:"version"`
}

// Contacts is to get a list of all contacts
func Contacts(token string) ([]ContactsReturnContent, error) {

	// To save the contact data
	var contacts []ContactsReturnContent

	// To call the page
	page := 0

	// Loop over all sites
	for {

		// Set config for new request
		c := Config{fmt.Sprintf("/v1/contacts?page=%d", page), "GET", token, "application/json", nil}

		// Send request
		response, err := c.Send()
		if err != nil {
			return nil, err
		}

		// Decode data
		var decode ContactsReturn

		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return nil, err
		}

		// Close request
		response.Body.Close()

		// Add contacts
		for _, value := range decode.Content {
			contacts = append(contacts, value)
		}

		// Check length & break the loop
		if decode.TotalPages == page {
			break
		} else {
			page++
		}

	}

	// Return data
	return contacts, nil

}

// Contact is to get a contact by id
func Contact(id, token string) (ContactsReturnContent, error) {

	// Set config for new request
	c := Config{"/v1/contacts/" + id, "GET", token, "application/json", nil}

	// Send request
	response, err := c.Send()
	if err != nil {
		return ContactsReturnContent{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ContactsReturnContent

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ContactsReturnContent{}, err
	}

	// Return data
	return decode, nil

}

// AddContact is to add a new contact
func AddContact(body ContactBody, token string) (ContactReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ContactReturn{}, err
	}

	// Set config for new request
	c := Config{"/v1/contacts/", "POST", token, "application/json", bytes.NewBuffer(convert)}

	// Send request
	response, err := c.Send()
	if err != nil {
		return ContactReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ContactReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ContactReturn{}, err
	}

	// Return data
	return decode, nil

}

// UpdateContact is to add a new contact
func UpdateContact(body ContactBody, token string) (ContactReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ContactReturn{}, err
	}

	// Set config for new request
	c := Config{"/v1/contacts/" + body.Id, "PUT", token, "application/json", bytes.NewBuffer(convert)}

	// Send request
	response, err := c.Send()
	if err != nil {
		return ContactReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ContactReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ContactReturn{}, err
	}

	// Return data
	return decode, nil

}
