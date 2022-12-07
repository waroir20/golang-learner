# Golang Learning
Set of exercises to get going on golang
### How to use this repo
1. Fork this git repository into your git account (see [github documentation](https://docs.github.com/en/get-started/quickstart/fork-a-repo))
2. Clone the new repo to your local computer using the command:  
`git clone https://github.com/<your gihub username here>/golang-learner.git` 
3. Initialize your local file system with the needed dependencies using the command:  
`go install`
4. Make sure you can run the hello world app using the command:  
`go run main.go`
5. Go through each of the modules. __Remember to commit and push up often.__  
As you progress through the modules you will be expected to do more and more "fill-in-the-blanks" meaning 100% of the steps you will need to do will not be explicitly stated, this is intentional.
   - [Data Structures](#1-data-structures)
   - [Testing](#2-testing)
   - [Simple APIs](#3-simple-apis)  
   - [Databases](#4-databases)
   - [Capstone: Recreate - Full Stack App](#n-capstone-full-stack-app)

## Modules
### 1. Data Structures 
   Choose one of the common data structures and first implement the Non-generic version, then once all tests pass implement the generic version  
   *__Note__: your solution should be a self-contained golang struct*
   
   #### TODO (choose at least one)
   - [Queue](1/queue/queue.go) and [Generic Queue](1/queue/generic-version/generic-queue.go)
   - [Stack](1/stack/stack.go) and [Generic Stack](1/stack/generic-version/generic-stack.go)
   
   #### Learning Objectives:
   - Interfaces
   - Multiple return from function
   - Generic typing of: function parameters and returns, structs, and interface
### 2. Testing
   Expand on your golang abilities and dev into the beautiful world of testing.  
   
   #### TODO
   Implement the tests [here](2/learn-to-test_test.go)

   #### Learning Objectives:
  - Writing tests
  - Pass by Reference vs Pass by Value
  - Struct tags
  - JSON serialization and deserialization
  - Simple struct validation
### 3. Simple APIs
   Work on building a set of simple REST APIs based on a pre-defined contract.
   Sorry had to get a bit silly on this one (for reference I have started playing Final Fantasy 14 so lots of references). 

   #### TODO
   Implement the APIs [here](3/api.go)

  #### Learning Objectives:
  - Golang routing via the [Gin](https://github.com/gin-gonic/gin) framework
  - Struct tags
  - JSON serialization and deserialization from HTTP request
  - Simple struct validation
  - Http Status codes
### 4. Databases
   Work on implementing the operations All, Create, DeleteById, GetById, and Update methods for this simple User repository.  
   *__Note__: a bare-bones implementation of TestContainers, establishing a Postgres connection, and wiring up the integration test have been provided for you.*  
   
   #### TODO
   Implement the [Database operations](4/repository.go) and their respective [Integration test cases](4/repository_test.go)

  #### Learning Objectives:
  - Relational Database operations using the [GORM](https://github.com/go-gorm/gorm) (Golang Object Relational Model) Library, [documentation](https://gorm.io).
  - Use of TestContainers for integration testing
  - Use of lifecycle hooks for structs (based on them implementing a framework's interfaces)
  - Writing and testing your own code
### N. Capstone Full-Stack-App
   Previously you have created a working full stack application with a simple UI, a backend logic layer, and a persistence tier.  
   Your goal will be to recreate that logic layer which has been written in Java/Python/C++ in Go 
   (the goal is to understand that go is a fungible substitute for another backend language).
   __It is important to make as few changes to the UI and DB layers as possible__ (*ideally no changes are needed, 
   you are exchanging one set of APIs for another*).

   #### TODO
   - Implement your previous Java backend application [TODO]() in golang
   - Utilize your existing UI (HTML/Javascript) app (Minimal changes please)
   - Utilize your existing MySQL database
   - Prepare to give a brief demo of your working project and beable to answer simple questions

  #### Learning Objectives:
  - Serving a UI from a golang application
  - End-to-end http request to database persistence of data
  - Sticking with a predefined contract
  - Building a full-fledge Go app
   