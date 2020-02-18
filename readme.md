# Store Employee and Inventory Manager
## A full stack application using React, Go, graphql-go, Postgresql, and Redis

This application demonstrates a full-stack application with a React client and a GraphQL backend API.

## Demo Video
[![Demo Video](http://img.youtube.com/vi/lO0OoP4jmBw/0.jpg)](https://youtu.be/lO0OoP4jmBw)

## WIP
This project is currently in progress. See issues for details of features yet to be implemented, or to suggest a feature.


## Getting started

### Start development server
You will need to create Google and Facebook Developer accounts as this project makes use of Google and Facebook for 3rd party user identities (Google uses OpenID, and Facebook has it's own auth mechanism via its graphAPI).

Rename the file *.env.example* to *.env* . Then, complete the details for the following:

```
GOOGLE_CLIENT_ID=####
GOOGLE_CLIENT_SECTRET=####
FACEBOOK_CLIENT_ID=####
FACEBOOK_CLIENT_SECRET=####
JWT_SECRET=####
```

Info for using OpenID with Google can be found [here](https://developers.google.com/identity/protocols/OpenIDConnect)

Info for Facebook-login can be found [here](https://developers.facebook.com/docs/facebook-login/)

This application uses Google and Facebook on the server to verify a users identity. However, we issue our own tokens which are used to manage logged-in users inside of a redis cache. Therefore, you may use a secret of your choice. You will also need Google and Facebook client IDs for the client app as logging-in takes place on the front-end, and user verification of credentials provided to the front-end client libraries takes place on the back-end. In other words, the front end gets a token after the user logs in with Google or Facebook. These tokens are passed to the backend and verified with secrets provided by Google and Facebook developer consoles. After successful verification, our app issues a JWT and stores it on the client as two cookies as seen this [fine article](https://medium.com/lightrail/getting-token-authentication-right-in-a-stateless-single-page-application-57d0c6474e3). 

```
docker-compose up
go run main.go
```

You can now access GraphQL playground on http://localhost:8080/graphql. You can test the products and categories queries without authorization. 

Also, make sure to use the following in HTTP headers (used as part of the auth flow)
```
{
  "X-Requested-With": "XMLHttpRequest"
}
```

Docker is currently used only for starting a development redis cache and postgres database. In the future it may be used to bundle the entire app and for CI.

Running main.go will start the graphql server.

### Run the React App Client
This is a lot easier. 
Rename the file *.env.example* to *.env.local* . Then, complete the details for the following:

```
REACT_APP_GOOGLE_CLIENT_ID=#####
REACT_APP_FACEBOOK_CLIENT_ID=#####
REACT_APP_URI_GQL="http://localhost:8080/graphql"
```

You may change the URI of the GraphQL backend in the environment variable if you have changed the port in the server application. 

Then install packages and run the development server with:
```
yarn install
yarn start
```

## Libraries and tools used for the application

### [graphql-go](https://github.com/graphql-go/graphql)
An implementation of the reference implementation of GraphQL written in Go. This implementation was used to better understand GraphQL as opposed to using a GraphQL generator from schema such as [gqlgen](https://github.com/99designs/gqlgen)

### [Redis](https://redis.io/)
For storing logged in users. This acts as a sort of light-weight session storage. It is accessed in all requests to retrieve user's role

### [Postgresql](https://www.postgresql.org/)
Used as the principle database for storing users, roles, products, and product categories.

### [Dataloaden](https://github.com/vektah/dataloaden)
Used to generate dataloaders (in models folder) to prevent the n+1 problem with GraphQL APIs. 

### [create-react-app](https://reactjs.org/docs/create-a-new-react-app.html)

### [Apollo Client](https://www.apollographql.com/docs/react/)
Used for making graphQL requests to server and managing and caching state from fetched data. 

### [React Context](https://reactjs.org/docs/context.html)
Since authorization state of the user changes infrequently, the application is wrapped in a Context Provider which provides the user to the application. 
