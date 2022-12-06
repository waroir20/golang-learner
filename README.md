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
   - [Capstone: Recreate - Full Stack App](#n-capstone-full-stack-app)  

## Modules
### 1. Data Structures 
   Choose one of the common data structures and first implement the Non-generic version, then once all tests pass implement the generic version:  
   *__Note__: your solution should be a self-contained golang struct*
   - [Queue](1/queue/queue.go) and [Generic Queue](1/queue/generic-version/generic-queue.go)
   - [Stack](1/stack/stack.go) and [Generic Stack](1/stack/generic-version/generic-stack.go)
   
   #### Learning Objectives:
   - Interfaces
   - Multiple return from function
   - Generic typing of: function parameters and returns, structs, and interface
### 2. Testing
   Expand on your golang abilities and dev into the beautiful world of testing.  
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
   Implement the APIs [here](3/api.go)

  #### Learning Objectives:
  - Golang routing via the [Gin](https://github.com/gin-gonic/gin) framework
  - Struct tags
  - JSON serialization and deserialization from HTTP request
  - Simple struct validation
  - Http Status codes
### N. Capstone Full-Stack-App
   Previously you have created a working full stack application with a simple UI, a backend logic layer, and a persistence tier.  
   Your goal will be to recreate that logic layer which has been written in Java/Python/C++ in Go 
   (the goal is to understand that go is a fungible substitute for another backend language).
   __It is important to make as few changes to the UI and DB layers as possible__ (*ideally no changes are needed, 
   you are exchanging one set of APIs for another*).

  #### Learning Objectives:
  - Serving a UI from a golang application
  - End-to-end http request to database persistence of data
  - Sticking with a predefined contract
  - Building a full-fledge Go app
   