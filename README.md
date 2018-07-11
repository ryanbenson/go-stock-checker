[![Codacy Badge](https://api.codacy.com/project/badge/Grade/869ce453b78444b79ac627ade97d5eb0)](https://www.codacy.com/app/ryanbenson/go-stock-checker?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ryanbenson/go-stock-checker&amp;utm_campaign=Badge_Grade) [![codecov](https://codecov.io/gh/ryanbenson/go-stock-checker/branch/master/graph/badge.svg)](https://codecov.io/gh/ryanbenson/go-stock-checker) [![Build Status](https://travis-ci.org/ryanbenson/go-stock-checker.svg?branch=master)](https://travis-ci.org/ryanbenson/go-stock-checker)

# go-stock-checker
Application to check on out of stock items from stores

## Requirements
* [Golang v1.10.x](https://golang.org/)

Install dependencies:
* `go get github.com/PuerkitoBio/goquery`
* `go get github.com/gen2brain/beeep`
* `go get github.com/joho/godotenv`

Fill out the `.env` file. You'll need a [Twilio](https://www.twilio.com/) account
which is free and you get free credits. There will you need to:
* Make a phone number
* Get your Account SID
* Get your Auth Token

Then fill out the `.env` file and add in who gets the text message.

## List of pages to check
In order to configure which pages to check and how, go to the `buildInventory`
function and update:
* `url` - the page to check`
* `name` - the product name
* `dom` - using a CSS selector, which DOM element has the "sold out" text
* `soldOut` - the text flavor used to convey "sold out"

## Running
It's as simple as either:
* `go run app.go`

Or if you want to build it: `go build` and `./go-stock-checker`
