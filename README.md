
# GoShort
|                                                                     License                                                                     |                                                            Go Version                                                             |                                                                                                                     Actions                                                                                                                      |                                                                 GoReport                                                                 |
|:-----------------------------------------------------------------------------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------------------------------------------:|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:----------------------------------------------------------------------------------------------------------------------------------------:|
| [![MIT License](https://img.shields.io/github/license/batt0s/goshort?style=flat-square)](https://github.com/batt0s/goshort/blob/master/LICENSE) | ![Go Version](https://img.shields.io/github/go-mod/go-version/batt0s/goshort/master?label=Go%20Version&logo=go&style=flat-square) | ![Tests](https://img.shields.io/github/actions/workflow/status/batt0s/goshort/test.yml?style=flat-square&label=Test)<br/> ![Build](https://img.shields.io/github/actions/workflow/status/batt0s/goshort/build.yml?style=flat-square&label=Build) | [![Go Report Card](https://goreportcard.com/badge/github.com/batt0s/goshort)](https://goreportcard.com/report/github.com/batt0s/goshort) |

Yet another URL shortener made in Go. <br>
A hobby project of me that I made for improving my coding skills. <br>
[GoShort](https://goshort.battos.dev/).


## Tech Stack

**Server:** 
- Go (used net/http and [chi](https://github.com/go-chi/chi))
- PostgreSQL (production) and SQLite3 (dev and tests)

**Client:** 
- Pure JS
- HTML5
- CSS (TailwindCSS)


## Features

- URL Shortening
- URL Shortening with a custom short URL
- Getting original URL from a short URL


## To-Do
- [ ] Expire able Links
- [ ] A good icon/logo
- [ ] Extensions for browsers

## API Reference

#### Shorten URL

```
  POST /api/v3/shorten
```

| Parameter | Type     | Description                   |
|:----------|:---------|:------------------------------|
| `url`     | `string` | **Required**. URL to shorten. |
| `custom`  | `string` | *Optional* Custom short URL   |



#### Get Original URL

```
  POST /api/v3/getOrigin
```


| Parameter | Type     | Description              |
|:----------|:---------|:-------------------------|
| `url`     | `string` | **Required**. Short URL. |


## Screenshots

![App Screenshot](https://i.imgur.com/9LqMBwu.png)
