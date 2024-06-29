# Twitter clone

Project to create an as yet unnamed twitter clone with the goal of learning Go and Next Js.

## Back End

A CRUD API backend written in go using gin for api requests and GORM for a sqlite database ORM.

### TODO

#### Users

- [ ] Add create session token w/ lifetime
- [x] Add check session token function
- [x] Login endpoint needs to return id and sessions token
- [x] Logout endpoint needs to destroy session token

#### Posts

- [x] Add post "/posts"
- [x] Get Post "/posts/:id"
- [x] Patch posts "/posts/:id"
- [x] delete posts "/posts/:id"

## Front End

Simple Website for posting messages written using the NextJS framework with Tailwind for UI.

### TODO

#### INIT

- [x] Init NJS
- [x] Install Tailwind
- [x] init shadcnui
- [x] build navbar
- [x] Init Pages
  - [x] Login
  - [x] Create Account
  - [x] Home
  - [x] Messages
- [x] Messages Form
- [x] Message Component
- [ ] API Req
  - [ ] Login
  - [ ] Create Account
  - [x] Get Messages
  - [ ] Send new Message
  - [ ] logout
  - [ ] update message
  - [ ] delete message
