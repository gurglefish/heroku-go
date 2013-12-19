// Copyright (c) 2013 Blake Gentry. All rights reserved. Use of
// this source code is governed by an MIT license that can be
// found in the LICENSE file.

/*
Package heroku is a client interface to the Heroku API.

Background

This package provides a complete interface to all of the Heroku Platform API v3
actions, and is almost entirely auto-generated based on the API's JSON Schema.
The exceptions are the files heroku.go, heroku_test.go, and app_test.go, as well
as the generator itself. All models are auto-generated by the Ruby script in
gen/gen.rb.

The client leverages Go's powerful net/http Client. That means that,
out-of-the-box, it has keep-alive support and the ability to handle many
concurrent requests from different goroutines.

You should have at least some understanding of Heroku and its Platform API:
https://devcenter.heroku.com/articles/platform-api-reference

Installation

This package is targeted towards Go 1.2 or later, though it may work on
earlier versions as well.

Run `go get github.com/bgentry/heroku-go` to download, build, and install the
package.

Getting Started

To use the client, first add it to your Go file's imports list:

  import (
    "github.com/bgentry/heroku-go"
  )

Then create a Client object and make calls to it:

  client := heroku.Client{Username: "email@me.com", Password: "my-api-key"}

  // pass nil for options if you don't need to set any optional params
  app, err := client.AppCreate(nil)
  if err != nil {
    panic(err)
  }
  fmt.Println("Created", app.Name)

  // Output:
  // Created dodging-samurai-42

That's it! Here is a more advanced example that also sets some options on the
new app:

  name := "myapp"
  region := "region"

  // Optional values need to be provided as pointers. If a field in an option
  // struct is nil (not provided), the option is omitted from the API request.
  opts := heroku.AppCreateOpts{Name: &name, Region: &region}

  // Create an app with options set:
  app2, err := client.AppCreate(&opts)
  if err != nil {
    // if this is a heroku.Error, it will contain details about the error
    if hkerr, ok := err.(heroku.Error); ok {
      panic(fmt.Sprintf("Error id=%s message=%q", hkerr.Id, hkerr))
    }
  }
  fmt.Printf("created app2: name=%s region=%s", app2.Name, app2.Region.Name)

  // Output:
  // created app2: name=myapp region=eu

Optional Parameters

Many of the Heroku Platform API actions have optional parameters. For example,
when creating an app, you can either specify a custom name, or allow the API to
choose a random haiku name for your app.

Optional parameters in heroku-go are always provided to functions as a pointer
to a struct, such as AppCreateOpts for the function AppCreate(). If you do not
wish to set any optional parameters, simply provide a nil in place of the
options struct, and the options will be omitted from the API request entirely.
For any individual options that you don't want to set, simply leave them as nil,
and they will be omitted from the API request.
*/
package heroku
