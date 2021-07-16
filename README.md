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
contacts, err := golexware.Contacts("token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(contacts)
}
```

### Create a new contact

To create a new contact, lexoffice needs some data. These must be entered in a structure.

```go
// Define body
body := &golexware.ContactBody{
    "",
    0,
	ContactBodyRoles{
        ContactBodyCustomer{},
        ContactBodyVendor{},
	},
    ContactBodyCompany{
        "J&J Ideenschmiede GmbH",
        "12345/12345",
        "DE123456789",
        true,
        []ContactBodyContactPersons{{
            "Herr",
            "Jonas",
            "Kwiedor",
            "jonas.kwiedor@jj-ideenschmiede.de",
            "04152 8903730",
        }},
    },
    ContactBodyAddresses{
        []ContactBodyBilling{{
            "Rechnungsadressenzusatz",
            "Fährstraße 31",
            "21502",
            "Geesthacht",
            "DE",
        }},
        []ContactBodyShipping{{
            "Lieferadressenzusatz",
            "Fährstraße 31",
            "21502",
            "Geesthacht",
            "DE",
        }},
    },
    ContactBodyEmailAddresses{
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
    },
    ContactBodyPhoneNumbers{
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
contactReturn, err := golexware.AddContact(body, "token")
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
body := &ContactBody{
    "ID",
    1,
    ContactBodyRoles{
        ContactBodyCustomer{},
        ContactBodyVendor{},
    },
    ContactBodyCompany{
        "J&J Ideenschmiede GmbH",
        "12345/12345",
        "DE123456789",
        true,
        []ContactBodyContactPersons{{
            "Herr",
            "Jonas",
            "Kwiedor",
            "jonas.kwiedor@jj-ideenschmiede.de",
            "017684714777",
        }},
    },
    ContactBodyAddresses{
        []ContactBodyBilling{{
            "Rechnungsadressenzusatz",
            "Fährstraße 31",
            "21502",
            "Geesthacht",
            "DE",
        }},
        []ContactBodyShipping{{
            "Lieferadressenzusatz",
            "Fährstraße 31",
            "21502",
            "Geesthacht",
            "DE",
        }},
    },
    ContactBodyEmailAddresses{
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
        []string{"info@jj-ideenschmiede.de"},
    },
    ContactBodyPhoneNumbers{
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
contactReturn, err := golexware.UpdateContact(body, "token")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(contactReturn)
}
```
