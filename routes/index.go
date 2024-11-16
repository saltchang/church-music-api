package routes

import (
	"net/http"
)

// GetIndex route (todo)
func GetIndex(response http.ResponseWriter, request *http.Request) {

	response.Write([]byte("<!DOCTYPE html><html lang='en'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><meta http-equiv='X-UA-Compatible' content='ie=edge'><title>Church Music API</title><link href='https://fonts.googleapis.com/css?family=Overpass+Mono&display=swap' rel='stylesheet'><style>.code {background: rgb(199, 199, 199);border-radius: .25rem;font-family: 'Overpass Mono', monospace;}</style></head><body style='padding: 1rem;'><h1>Church Music API</h1><p>A music data API build with Go, MongoDB, deployed on Heroku.<br>This API is used by <a href='https://music.caten-church.org'>Caten-Music</a> .<br>For security, only GET method open to normal users.</p><hr><h2>Usage</h2><h3>GET</h3><ul><li><span class='code'>'/api/songs'</span> : get all songs from the database.</li><li><span class='code'>'/api/songs/sid/{sid0+sid1+sid2+...}'</span> : get songs by multiple sid, splited all sid by '+'.</li><li><span class='code'>'/api/songs/search?lang={lang}&c={c}&to={to}&title={title}&lyrics={lyrics}&test=0'</span> : search songs by multiple arguments.<br><ul><li><span class='code'>lang</span> : language - 'Chinese' and 'Taiwanese'.</li><li><span class='code'>c</span> : collection - from 1 to 11.</li><li><span class='code'>title</span> : title - the route keywords are splited by '+'.</li><li><span class='code'>to</span> : tonality - ex. 'C'</li></ul></li><li><span class='code'>'/api/songs/random/{r}'</span> : get random songs by given a amount r.</li></ul><h3>PUT</h3><h4>*Need token for authority</h4><p>If you need the token, please mail to <a href='mailto:saltchang.tw@gmail.com'>saltchang.tw@gmail.com</a></p><ul><li><span class='code'>'/api/songs/sid/{sid}'</span> : update a song by its {sid} from the input body raw.</li></ul><hr><p>To see more information, please visit <a href='https://github.com/saltchang/church-music-api'>GitHub</a></p></body></html>"))
}
