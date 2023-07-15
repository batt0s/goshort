
# GoShort
License             |  Go Version | Actions | GoReport
:-------------------------:|:-------------------------:|:--:|:--:
[![MIT License](https://img.shields.io/github/license/batt0s/goshort?style=flat-square)](https://github.com/batt0s/goshort/blob/master/LICENSE) | ![Go Version](https://img.shields.io/github/go-mod/go-version/batt0s/goshort/master?label=Go%20Version&logo=go&style=flat-square) | [![Go Report Card](https://goreportcard.com/badge/github.com/batt0s/goshort)](https://goreportcard.com/report/github.com/batt0s/goshort)

Yet another URL shortener made in Go. <br>
A hobby project of me that i made for improving my coding skills. <br>
[GoShort on Render](https://goshort.onrender.com/).

**I am going to rewrite almost all of the project. You can find this verison in branch v1**

**Update**: Going to rewrite almost all of the project. My goal is to remove the sign in and make a simplier app. With a simple API. I plan to make extensions for chrome, firefox and opera. Nobody gonna use the front-end propably (because there is bunch of url shorteners out there way more popular) maybe i could delete the front-end completly. 

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
