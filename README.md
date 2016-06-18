# GoAngularBrowserifyBoilerplate
Boilerplate for web applications using Angular, Browserify, Gulp, MongoDB & Go.
The boilerplate provides these functions:
1. A static web host that serve static files: html/javascript/css using [GO's echo](https://labstack.com/echo).
2. REST Services that provide the data from the database using [GO's echo](https://labstack.com/echo).
3. OAuth2 functionality to login using social networks (facebook, google+...)
4. Exchange tokens for security based on [JWT](https://jwt.io) standard.
5. Store OAuth tokens in MongoDB database.
6. Build system using Gulp.

## Dependencies

### Development

The following tools are required to build and deploy the app:

* [Node.js](http://nodejs.org/)
* [Bower](http://bower.io/)
* [Gulp](http://gulp.com/)
* [Git](http://msysgit.github.io/)
* [Go](http://golang.org/)
* [MongoDB](https://www.mongodb.com/)
* [echo](https://labstack.com/echo)


### Runtime

The system has no runtime dependencies other than the database storage (MongoDB). It is built on the [echo](http://echo.labstack.com/)

### Build
The build is using [Gulp](http://gulp.com/)

