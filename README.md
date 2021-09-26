# Lexoffice for GO

## Install

```console
go get github.com/jjideenschmiede/golexoffice
```

## How to use?

### Get all contacts

To get all contacts you can perform the following function.

```go
// Get all contacts
contacts, err := golexoffice.Contacts("token")
if err != nil {
fmt.Println(err)
} else {
fmt.Println(contacts)
}
```

### Get a contact by id

If you want to read out a specific contact, you can do this via the id (UUID).

```go
// Get a contact by id
contact, err := golexoffice.Contact("b324c2be-b745-4128-9ecd-e262a0a761cd", "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(contact)
}
```

### Create a new contact

To create a new contact, lexoffice needs some data. These must be entered in a structure.

```go
// Define body
body := golexoffice.ContactBody{
    "",
    0,
	golexoffice.ContactBodyRoles{
    golexoffice.ContactBodyCustomer{},
    golexoffice.ContactBodyVendor{},
	},
    golexoffice.ContactBodyCompany{
        "J&J Ideenschmiede GmbH",
        "12345/12345",
        "DE123456789",
        true,
        []golexoffice.ContactBodyContactPersons{{
            "Herr",
            "Jonas",
            "Kwiedor",
            "jonas.kwiedor@jj-ideenschmiede.de",
            "04152 8903730",
        }},
    },
    golexoffice.ContactBodyAddresses{
        []golexoffice.ContactBodyBilling{{
            "Rechnungsadressenzusatz",
            "Fährstraße 31",
            "21502",
            "Geesthacht",
            "DE",
        }},
        []golexoffice.ContactBodyShipping{{
            "Lieferadressenzusatz",
            "Fährstraße 31",
            "21502",
            "Geesthacht",
            "DE",
        }},
    },
    golexoffice.ContactBodyEmailAddresses{
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
    },
    golexoffice.ContactBodyPhoneNumbers{
        []string{"04152 8903730"},
        []string{"04152 8903730"},
        []string{"04152 8903730"},
        []string{"04152 8903730"},
        []string{"04152 8903730"},
        []string{"04152 8903730"},
    },
    "Testnotiz",
    false,
}

// Create new contact
contactReturn, err := golexoffice.AddContact(body, "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(contactReturn)
}
```

### Update a contact

If you want to update a contact, then some information is very important. You need the ID of the contact & the version.

```go
// Define body
body := golexoffice.ContactBody{
    "ID",
    1,
	golexoffice.ContactBodyRoles{
        golexoffice.ContactBodyCustomer{},
        golexoffice.ContactBodyVendor{},
    },
    golexoffice.ContactBodyCompany{
        "J&J Ideenschmiede GmbH",
        "12345/12345",
        "DE123456789",
        true,
        []golexoffice.ContactBodyContactPersons{{
            "Herr",
            "Jonas",
            "Kwiedor",
            "jonas.kwiedor@jj-ideenschmiede.de",
            "017684714777",
        }},
    },
    golexoffice.ContactBodyAddresses{
        []golexoffice.ContactBodyBilling{{
            "Rechnungsadressenzusatz",
            "Fährstraße 31",
            "21502",
            "Geesthacht",
            "DE",
        }},
        []golexoffice.ContactBodyShipping{{
            "Lieferadressenzusatz",
            "Fährstraße 31",
            "21502",
            "Geesthacht",
            "DE",
        }},
    },
    golexoffice.ContactBodyEmailAddresses{
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
    },
    golexoffice.ContactBodyPhoneNumbers{
        []string{"04152 8903730"},
        []string{"04152 8903730"},
        []string{"04152 8903730"},
        []string{"04152 8903730"},
        []string{"04152 8903730"},
        []string{"04152 8903730"},
    },
    "Testnotiz",
    false,
}

// Create new contact
contactReturn, err := UpdateContact(body, "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(contactReturn)
}
```
    
### Get a invoice

If you want to read out a specific invoice, you can do this via the id (UUID).

```go
// Invoice is to get a invoice by id
invoice, err := Invoice("0cf8142b-6f54-4c96-9766-6f44a9a4814b", "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(invoice)
}
```

### Create a invoice

In order to create a new invoice, the data must be sent in a certain format. This works as follows.

If you can assign an invoice to an existing customer, then in the struct InvoiceBodyAddress{} please specify only the Id and leave the rest empty.

For more information, please refer to the [documentation](https://developers.lexoffice.io/docs/?shell#invoices-endpoint-create-an-invoice).

```go
// Define body
body := golexoffice.InvoiceBody{
    "",
    "",
    "",
    "",
    1,
    false,
    "",
    "",
    "2021-07-20T00:00:00.000+01:00",
    "",
    golexoffice.InvoiceBodyAddress{
        "",
        "Test Company",
        "",
        "Teststreet 12",
        "Geesthacht",
        "21502",
        "DE",
    },
    []golexoffice.InvoiceBodyLineItems{{
        "",
        "custom",
        "Testarticle",
        "Very nice article!",
        1,
        "Stück",
		golexoffice.InvoiceBodyUnitPrice{
            "EUR",
            13.4,
            15.59,
            19,
        },
        0,
        13.4,
    }},
    golexoffice.InvoiceBodyTotalPrice{
        "EUR",
        13.4,
        15.95,
        nil,
        2.55,
        nil,
        nil,
    },
    []golexoffice.InvoiceBodyTaxAmounts{{
        19,
        2.55,
        15.95,
    }},
    golexoffice.InvoiceBodyTaxConditions{
        "net",
        nil,
    },
    golexoffice.InvoiceBodyPaymentConditions{
        "Please pay within the next 30 days.",
        30,
		golexoffice.InvoiceBodyPaymentDiscountConditions{
            0,
            0,
        },
    },
    golexoffice.InvoiceBodyShippingConditions{
        "2021-07-20T00:00:00.000+01:00",
        nil,
        "none",
    },
    "Invoice",
    "We hereby invoice you for the items you have ordered",
    "Thank you for your purchase",
}

// Create new contact
invoice, err := golexoffice.AddInvoice(body, "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(invoice)
}
```

### Upload a file

Here you will find a function with which you can upload a file to lexoffice. You only need the path to the file and the lexoffice token.

For more information, please refer to the [documentation](https://developers.lexoffice.io/docs/#files-endpoint-upload-a-file).

```go
// Open file
file, err := os.Open("/Users/jonaskwiedor/Downloads/Rechnung 201912101300005.pdf")
if err != nil {
    fmt.Println(err)
}

// Files is to create a new file
files, err := golexoffice.AddFile(file, "Rechnung 201912101300005.pdf", "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(files)
}
```