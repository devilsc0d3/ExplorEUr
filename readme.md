# Order
***

### Objectives
The Forum project consists of building a forum using the go language. This forum allows us to bring together a community.
It will allow us to achieve several things:
- a database and to use docker in order to secure our database and our site.
- users who can create an account in order to be able to use the forum other than as read-only. These users can also log in and log out whenever they wish.
- the categories in the forum and that in each category there are the posts of the users to which we can reply.
- likes and dislikes reactions to some posts so that even if users don't want to reply they can show what they think.
- an activity page that can gather notifications, a history of user appearances and for moderators and admins additional features.
This project also focuses on the visualization of the forum website with all its protections, cookies and different roles for users.

# Feature
***

-Forum
-Server
-Docker
-Register 
    -nickname
    -email
    -password (cryptage)
    -roles ?
-middleware register
    -check email
    -check password
    -check no duplicate nickname
    -check no duplicate email
-Error handling register / login
    -unique nickname
    -incorrect / unique address
    -password incorrect
-Crypter data
    -password user
    -bdd
-Cookies
-Hide / Show button
-Category templates
    -locations
    -anecdote
    -advice
-Ester eggs
-Home page
-Ban
    -insult
    -ban account
-Like / dislike / comment
    -like
    -dislike
    -comment
    -number
-All mode
    -view
    -user
    -moderator
    -admin
-Activity page
    -publication
    -comment
    -like
    -...
-Filtering
    -date
    -like
    -category global
-Authentication certificate

# Amelioration
***
    -Add image
    -Profile Manager
    -Reaction / Emoji
    -Language / Universalize
    -Address Blocking
    -Maps
    -Connected by google / github
    -Confidentiality policy

# Installation
***

- Clone repository with command :
```bash
git clone 
```
- Launch the forum project :
Run with the command : 
```go
docker compose up --build
```

# Credit
***
*FAURE Leo*

*BRUN Sasha*

*LALDY-MAQUIHA Adan*

*SENAC Lucille*