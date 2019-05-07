# Songs API in Go

歡迎前往 [Caten-Worship](https://caten-worship.herokuapp.com) 參觀

Welcome to visit [Caten-Worship](https://caten-worship.herokuapp.com) .

## Installation

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

$ cd go-api-songs

```

Install or update the Dep:

```shell

$ brew install dep
$ brew upgrade dep

```

Then install the requirements through dep:

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

- `"/api/songs"` : get all songs from the database.

- `"/api/songs/sid/{sid}"` : get the song by its `{sid}`.

- `"/api/songs/search/?lang={lang}&c={c}&title={title}"` : search songs by multiple arguments.
  - `lang` : language - `"Chinese"` and `"Taiwanese"`
  - `c` : collection - from `1` to `11`
  - `title` : title - you can sperate multiple keyword by `"+"`

## Example

### Get songs by SID

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
    "lyrics":     []
}]

```

### Search songs

```http
http://localhost:7700/api/songs/search/?lang=Chinese&c=7&title=來+歡
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
    ]
}]

```
