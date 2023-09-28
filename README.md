
## Go Nordigen API Client

A Go client for the Nordigen (now GoCardless) API, providing an easy-to-use wrapper around the API's endpoints. This client helps developers seamlessly integrate and fetch banking data.

### Features:

- **Token Management**: Generate and manage authentication tokens for the API.
- **Account Information**: Fetch details about linked accounts and their transactions.
- **End-User Agreements**: Manage and fetch end-user agreements.
- **Requisitions**: Handle requisitions and their associated data.
- **Supported Institutions**: Get a list of supported banking institutions based on country code.

### Installation:

To install this package, simply run:
```bash
go get github.com/ricardograndecros/go_nordigen
```

### Usage:

**Initialization**:
   ```go
   client, err := go_nordigen.NewClient(secretID, secretKey)
   if err != nil {
       log.Fatal(err)
   }
   ```
I recommend referring to [Nordigen's official documentation](https://developer.gocardless.com/bank-account-data/overview) 
for process flows and more information on terminology.

### Files Overview:

1. **client.go**: 
   - Contains the main `Client` struct for the Nordigen API.
   - Provides the base URL and initialization functionality.

2. **token.go**: 
   - Defines structs for tokens and secrets.
   - Contains methods to fetch new tokens.

3. **accounts.go**: 
   - Defines the `Account` struct.
   - Contains methods to fetch account details and transactions.

4. **agreement.go**: 
   - Defines the `EndUserAgreement` struct.
   - Contains methods to fetch and manage end-user agreements.

5. **requisition.go**: 
   - Defines the `Requisition` struct.
   - Contains methods related to requisitions.

6. **institutions.go**: 
   - Defines the `Institution` struct.
   - Contains methods to fetch supported banking institutions based on a country code.

### Contributing:

Feel free to fork this repository, make changes, and submit pull requests. Feedback and contributions are always welcome!

**ROADMAP**:
1. Implement all Nordigen endpoints
2. Add examples and more usage documentation
3. Grow the project to include more logic and user flows (e.g. a flow to retrieve transactions from an account)

I understand this code could have been autogenerated from the official Swagger. I built this repo just for educational 
purposes. I use it for my personal projects and is not intended to be used in any other scope. 

### License:

This project is licensed under the MIT License. See the `LICENSE` file for more details.



_Note: The README documentation was assisted by OpenAI's ChatGPT._