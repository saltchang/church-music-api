# Music API with Go for church

A music data API build with [Go](https://golang.org/), [Gorilla-Mux](https://github.com/gorilla/mux), [MogoDB Go Driver](https://github.com/mongodb/mongo-go-driver) from [MongoDB](https://www.mongodb.com), deploy on [AWS EC2](https://aws.amazon.com/tw/ec2) and [Ubuntu Linux OS](https://www.ubuntu.com)

This API is used by [Caten-Worship](https://caten-worship.herokuapp.com).

## Installation

Change directory to `$GOPATH/src` first:

```shell

$ cd $GOPATH/src

```

Clone the files,

use HTTPS:

```shell

$ git clone https://github.com/saltchang/songs-go-api-for-caten.git

```

or use SSH:

```shell

$ git clone git@github.com:saltchang/songs-go-api-for-caten.git

```

Go into the folder:

```shell

$ cd songs-go-api-for-caten

```

Install or update the [Dep](https://github.com/golang/dep)

by using [Homebrew](https://brew.sh/)) on macOS:

```shell

$ brew install dep
$ brew upgrade dep

```

or by using the following command on Linux:
Your will need to create the GOBIN (`$GOPATH/bin`) directory first.

```shell

$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

```

Then install the requirements through Dep:

```shell

$ dep ensure -v

```

## Run locally

```shell

$ go run main.go

```

or

```shell

$ go build && ./main.go

```

and then visit the site at [http://localhost:7700](http://localhost:7700)

## Usage

### GET

- `"/api/songs"` : get all songs from the database.
  獲取所有歌曲

- `"/api/songs/sid/{sid}"` : get the song by its `{sid}`.
  透過指定的 SID 獲取歌曲

- `"/api/songs/search?lang={lang}&c={c}&to={to}&title={title}"` : search songs by multiple arguments.
  透過複數條件搜尋歌曲

  - `lang` : language - `"Chinese"` and `"Taiwanese"`
    加入語言參數

  - `c` : collection - from `1` to `11`
    加入集數參數

  - `title` : title - you can sperate multiple keyword by `"+"`
    標題關鍵字，支援複數關鍵字，使用"+"來區分關鍵字

### PUT

- `"api/songs/sid/{sid}"` : update a song by its `{sid}` from the input document file.
  更新一首歌的資料

## Example

### ex. GET

#### Get songs by SID

```http
http://localhost:7700/api/songs/sid/1010066
```

Response:

```json
[{
    "_id":        "5cceafa94a38b40395f5adc8",
    "sid":        "1010066",
    "num_c":      "10",
    "num_i":      "66",
    "title":      "前來敬拜",
    "album":      "讚美之泉20-新的事將要成就，6",
    "tonality":   "G",
    "year":       "2015",
    "language":   "Chinese",
    "lyrics":     ["..."]
}]

```

#### Search songs

```http
http://localhost:7700/api/songs/search?lang=Chinese&c=7&to=A&title=來+歡
```

Response:

```json
[{
    "_id":        "5cceafa94a38b40395f5acc0",
    "title":      "我真歡喜來讚美你",
    "num_c":      "7",
    "num_i":      "71",
    "sid":        "1007071",
    "language":   "Chinese",
    "tonality":   "A",
    "album":      "約書亞02-祂的國度，祂的榮耀，1",
    "lyrics":     [
                    [
                     "p",
                     "睜開眼睛，感覺好熟悉，在你面前，",
                     "一切都不會在意，拋開憂慮，煩惱傷心，",
                     "現在只想和你一起，哦，我真歡喜來讚美你。"
                    ],
                    "..."
    ]
}]

```

### ex. PUT

#### Update a song

```http
http://localhost:7700/api/songs/sid/1010066
```

body raw:

```json
{
    "tonality":     "GGG",
    "year":         "200015",
    "language":     "Japanese"
}
```

Response:

```json
{
    "MatchedCount": 1,
    "ModifiedCount": 1,
    "UpsertedCount": 0,
    "UpsertedID": null
}

```

the new song data in the db:

```json
[{
    "_id":        "5cceafa94a38b40395f5adc8",
    "sid":        "1010066",
    "num_c":      "10",
    "num_i":      "66",
    "title":      "前來敬拜",
    "album":      "讚美之泉20-新的事將要成就，6",
    "tonality":   "GGG",
    "year":       "200015",
    "language":   "Japanese",
    "lyrics":     ["..."]
}]

```
