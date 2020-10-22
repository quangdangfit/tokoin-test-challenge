#### TOKOIN CODE CHALLENGE
Using the provided data (**tickets.json** and **users.json** and **organization.json**)

Write a simple command line application to search the data and return the results
in a human readable format.

* Feel free to use libraries or roll your own code as you see fit.
* Where the data exists, values from any related entities **MUST** be included in
the results, i.e.
    * Searching **organization MUST** return its **ticket subject** and **users name**.
    * Searching **users MUST** return his/her **assignee ticket subject** and **submitted ticket subject** and his/her **organization name**.
    * Searching **tickets MUST** return its **assignee name**, **submitter name**, and **organization name**.
* The user should be able to search on any field, full value matching is fine
(e.g. "mar" won't return "mary").
* The user should also be able to search for empty values, e.g. where
description is empty.
* Search can get pretty complicated pretty easily, we just want to see that you
can code a basic search application.

#### Setup
* Copy file config `cp config/config.sample.yaml config/config.yaml`
* Install require packages: `go mod vendor`

#### Build and Run
##### Run
```shell script
go run main.go
```
##### Test
```shell script
go test ./tests -v
```

#### Structure
```shell script
├── config          # Contain config file
├── data            # Contain data files
├── models          # Models Layer: contain models
├── repositories    # Repositories Layer: contain interface and implement
│   └── files       # Repositories implement: handle get data from files
├── schema          # Contain schemas
├── services        # Business Logic Layer  
├── tests           # Tests package: defaine testcases and implement unittest
│   └── data        # Contain data test files
├── utils           # Utilities package
```

#### Result
* Testcases passed: **234/234**.
* Coverage: **88.2%** files, **89.3%** statements.

#### Contributing
If you want to contribute to this boilerplate, clone the repository and just start making pull requests.