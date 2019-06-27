# Go API of Music for Church

A music data API build with [Go](https://golang.org/), [MongoDB](https://www.mongodb.com), deployed on [Heroku](https://devcenter.heroku.com).

This API is used by [Caten-Music](https://caten-music.herokuapp.com).

For security, only GET method open to normal users.

## Dependencies

- [Gorilla-Mux](https://github.com/gorilla/mux)

- [MogoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

[Dep](https://github.com/golang/dep) is used for managing the dependencies.

## Usage

### GET

- `"/api/songs"` : get all songs from the database.

- `"/api/songs/sid/{sid0+sid1+sid2+...}"` : get songs by multiple `sid`, splited all `sid` by `"+"`.

- `"/api/songs/search?lang={lang}&c={c}&to={to}&title={title}"` : search songs by multiple arguments.

  - `lang` : language - `"Chinese"` and `"Taiwanese"`.

  - `c` : collection - from `1` to `11`.

  - `title` : title - the route keywords are splited by `"+"`.

  - `to` : tonality - ex. `"C"`

- `"/api/songs/random/{r}"` : get random songs by given a amount `r`.

### PUT

- **Need token for authority**

- `"/api/songs/sid/{sid}"` : update a song by its `{sid}` from the input body raw.

## Example

### ex. GET

#### Get songs by one or multiple SIDs

```http
https://church-music-api.herokuapp.com/api/songs/sid/1010066+1010050+1003001
```

Response:

```json
[{
    "sid":        "1010066",
    "num_c":      "10",
    "num_i":      "66",
    "title":      "前來敬拜",
    "album":      "讚美之泉20-新的事將要成就，6",
    "tonality":   "G",
    "year":       "2015",
    "language":   "Chinese",
    ..
},
{
    "sid":        "1010050",
    "num_c":      "10",
    "num_i":      "50",
    "title":      "當我謙卑來到主前",
    "tonality":   "G",
    ...
},
{
    "sid":        "1003001",
    "num_c":      "3",
    "num_i":      "1",
    "title":      "主愛有多少",
    "tonality":   "Eb",
    ...
}]

```

#### Search songs by multiple arguments

```http
https://church-music-api.herokuapp.com/api/songs/search?lang=Chinese&c=7&to=A&title=來+歡
```

Response:

```json
[{
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
                    ...
    ]
}]

```

### ex. PUT

#### Update a song by specific SID

```http
https://church-music-api.herokuapp.com/api/songs/sid/1010066
```

Body raw:

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

The new data in the db:

```json
[{
    "sid":        "1010066",
    "num_c":      "10",
    "num_i":      "66",
    "title":      "前來敬拜",
    "album":      "讚美之泉20-新的事將要成就，6",
    "tonality":   "GGG",
    "year":       "200015",
    "language":   "Japanese",
    "lyrics":     ...
}]

```
