
# GoShort
License             |  Go Version | Actions | GoReport
:-------------------------:|:-------------------------:|:--:|:--:
[![MIT License](https://img.shields.io/github/license/batt0s/goshort?style=flat-square)](https://github.com/batt0s/goshort/blob/master/LICENSE) | ![Go Version](https://img.shields.io/github/go-mod/go-version/batt0s/goshort/master?label=Go%20Version&logo=go&style=flat-square) | ![Build](https://img.shields.io/github/workflow/status/batt0s/goshort/Go/master?style=flat-square) | [![Go Report Card](https://goreportcard.com/badge/github.com/batt0s/goshort)](https://goreportcard.com/report/github.com/batt0s/goshort)

Yet another URL shortener made in Go. \
A hobby project of me that i made for improving my coding skills.\
[GoShort on Heroku](https://goshrt.herokuapp.com).




## Tech Stack

**Server:** Go (used net/http and [chi](https://github.com/go-chi/chi)), PostgreSQL

**Client:** Pure JS, Html, Css written by me


## Features

- URL Shortening
- URL Shortening with a custom short URL
- Getting original URL from a short URL



## To-Do
- Dark Mode
- Expireable Links
- A good icon/logo
- Extensions for browsers
## API Reference

#### Shorten URL

```http
  POST /api/latest/shorten
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `url` | `string` | **Required**. URL to shorten. |

#### Shorten URL with a custom short URL

```http
  POST /api/latest/customShorten
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `url`      | `string` | **Required**. URL to shorten. |
| `custom` | `string` | **Required**. Custom Short URL. |

#### Get Original URL

```http
  POST /api/latest/getOrigin
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `url` | `string` | **Required**. Short URL. |


*[API page on website](http://www.goshort.xyz/api/latest/docs)*
## Screenshots

![App Screenshot](https://camo.githubusercontent.com/37cb45eaca67f5a48036f501d4aa56d29982d3c80ab772da33f823f9c1bde2e8/68747470733a2f2f692e696d6775722e636f6d2f54415a6c6339352e706e67)

