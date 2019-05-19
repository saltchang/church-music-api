# Go API of Music for Church

A music data API build with [Go](https://golang.org/), [MongoDB](https://www.mongodb.com), deployed on [Heroku](https://devcenter.heroku.com).

This API is used by [Caten-Worship](https://caten-worship.herokuapp.com).

For security, only GET method open to normal users.

## Dependencies

- [Gorilla-Mux](https://github.com/gorilla/mux)

- [MogoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

[Dep](https://github.com/golang/dep) is used for managing the dependencies.

## Usage

### GET

- `"/api/songs"` : get all songs from the database.

- `"/api/songs/sid/{sid}"` : get a song by its `{sid}`.

- `"/api/songs/search?lang={lang}&c={c}&to={to}&title={title}"` : search songs by multiple arguments.

  - `lang` : language - `"Chinese"` and `"Taiwanese"`.

  - `c` : collection - from `1` to `11`.

  - `title` : title - the route keywords will be sperated by `"+"`.

  - `to` : tonality - ex. `"C"`

### PUT

- **Need token for authority**

- `"/api/songs/sid/{sid}"` : update a song by its `{sid}` from the input body raw.

## Example

### ex. GET

#### Get songs by SID

```http
https://church-music-api.herokuapp.com/api/songs/sid/1010066
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
    "lyrics":     ...
}]

```

#### Search songs

```http
https://church-music-api.herokuapp.com/api/songs/search?lang=Chinese&c=7&to=A&title=來+歡
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
                    ...
    ]
}]

```

### ex. PUT

#### Update a song

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
    "_id":        "5cceafa94a38b40395f5adc8",
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
