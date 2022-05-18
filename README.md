
# GoShort
License             |  Go Version | Actions | GoReport
:-------------------------:|:-------------------------:|:--:|:--:
[![MIT License](https://img.shields.io/github/license/batt0s/goshort?style=flat-square)](https://github.com/batt0s/goshort/blob/master/LICENSE) | ![Go Version](https://img.shields.io/github/go-mod/go-version/batt0s/goshort/master?label=Go%20Version&logo=go&style=flat-square) | ![Build](https://img.shields.io/github/workflow/status/batt0s/goshort/Go/master?style=flat-square) | [![Go Report Card](https://goreportcard.com/badge/github.com/batt0s/goshort)](https://goreportcard.com/report/github.com/batt0s/goshort)

Yet another URL shortener made in Go. \
A hobby project of me that i made for improving my coding skills.\
[GoShort on Heroku](https://goshrt.herokuapp.com).




## Tech Stack

**Server:** Go (used net/http and [chi](https://github.com/go-chi/chi)), PostgreSQL

**Client:** Pure JS, Html, Css (TailwindCSS)


## Features

- URL Shortening
- URL Shortening with a custom short URL
- Getting original URL from a short URL
- Sing in with Google
- Dashboard (See your urls and how many click did it get, if you created while you logged in with google)


## To-Do
- [x] Sing in with Google
- [x] Dashboard (See your urls and how many click did it get, if you created while you logged in with google)
- [ ] Expireable Links
- [x] Dark Mode
- [ ] A good icon/logo
- [ ] Extensions for browsers

## API Reference

#### Shorten URL

```
  POST /api/v2/shorten
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `url` | `string` | **Required**. URL to shorten. |
| `is_custom` | `bool` | *Optional* If you want a custom short |
| `custom` | `string` | *Optional* Custom short URL (`is_custom` should be true) |
| `author` | `string` | *Optional* Google user ID |



#### Get Original URL

```
  POST /api/v2/getOrigin
```


| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `url` | `string` | **Required**. Short URL. |


## Screenshots

![App Screenshot](https://i.imgur.com/9qoWbQd.png)
