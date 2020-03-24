
# Milpost Backend

This project contains the backend of the post app, below a deployment guide
can be found and the requirements for this project

- [milpost-backend](#milpost-backend)

## Requirements

Below is a list of requirements and an explanation of the terms used to
describe the app.

### Terminology

As this is a post app it will mainly be dealing with parcels and letters
which will be referred to as an `item` in the future. Each item will belong to
a recipient or a person, this will be know as the `reciever`.

Each `Item` has the following properties:

- Receiver
- Receive date
- Type (Package or Letter)

Each `Receiver` has the following properties:

- Name
- Military Rank (the list of Swiss Military Rank can be found [here](https://de.wikipedia.org/wiki/Grade_der_Schweizer_Armee))
- `Item` associated with this person.

**Note**: The a `reciever` cannot exist without an `item` and an `item` cannot exist without a `receiver`.

Definitions:

- **Item**: Letter of Parcel
- **Letter**: Simple letter with size 25cm x17.6cm x 5cm or less and up to 250g.
- **Parcel**: Parcel or Package is anything that is not a letter
- **Receiver**: The person who is to recieve the item.
- **Post list**: The list of items and corresponding receivers.

### Specification

The app will be used to display the list of `items` required for collection
from the post and to add `items` to the be displayed by an admin user. Hence
there must authentication for the admin user which must use SSL. This must be a
web-app as it must be accessible for mobile and desktop.

#### TLDR

- Web-app displays list of `items` to be collected
- There are normal and admin users
- Admin must be authenticated
- Whole site must be SSL
- Must be mobile and desktop compatible

### Functionality

This should be a very simple web-app with easy to use and simple functionality.
The normal users must be able to:

- See a list of all the `items`
- Search for their name (without rank)
- Sort by date, rank, name, company, section and type.
- See interesting visualisations and statistics about `items`

The admin users should be able to:

- See a list of all the `items`
- Search for there name (without rank)
- See interesting visualisations and statistics about `items`
- Add `items` to the list
- remove `items` to the list
- Select multiple `items` for deletion
- Print the `post list`

**Note**: Normal users must not have access to the aditional functionality

Ideally there are two possible print outs for the `post list`, one to display
and one to be used by the admin to collect signatures from the recievers. The
minimum columns for the admin list are as follows:

- Rank
- Name
- Signature (To be filled in)

The `post list` should display the following information for each `item` to normal users:

- Date
- Rank
- Name
- Company
- Section
- Type

**Note**: For improved vilisbility when `items` are more than 36 hours old they should
be displayed and printed in red.

There is no requirement for the normal users to have any kind of authentication.
