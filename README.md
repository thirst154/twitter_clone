# Twitter clone

Project to create an as yet unnamed twitter clone with the goal of learning Go and Next Js.

## Back End

A CRUD API backend written in go using gin for api requests and GORM for a sqlite database ORM.

### TODO

#### Users
- [ ] Add create session token w/ lifetime
- [X] Add check session token function
- [X] Login endpoint needs to return id and sessions token 
- [X] Logout endpoint needs to destroy session token

#### Posts
- [X] Add post "/posts"
- [X] Get Post "/posts/:id"
- [X] Patch posts "/posts/:id"
- [X] delete posts "/posts/:id"

## Front End

Simple Website for posting messages written using the NextJS framework with Tailwind for UI.

### TODO

#### INIT
- [X] Init NJS
- [X] Install Tailwind
- [X] init shadcnui
- [X] build navbar
- [X] Init Pages
  - [X] Login
  - [X] Create Account
  - [X] Home
  - [X] Messages
- [ ] Messages Form
- [ ] Message Component
- [ ] API Req
  - [ ] Login
  - [ ] Create Account
  - [ ] Get Messages
  - [ ] Send new Message
  - [ ] logout
  - [ ] update message
  - [ ] delete message