package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Songs var
var Songs []Song // Songs data model

// Song Struct: This is the songs data model.
type Song struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SID           string             `json:"sid" bson:"sid"`
	NumC          string             `json:"num_c" bson:"num_c"`
	NumI          string             `json:"num_i" bson:"num_i"`
	Title         string             `json:"title" bson:"title"`
	Album         string             `json:"album" bson:"album"`
	Tonality      string             `json:"tonality" bson:"tonality"`
	Year          string             `json:"year" bson:"year"`
	Language      string             `json:"language" bson:"language"`
	Lyrics        []string           `json:"lyrics" bson:"lyrics"`
	Tempo         string             `json:"tempo" bson:"tempo"`
	TimeSignature string             `json:"time_signature" bson:"time_signature"`
	Publisher     string             `json:"publisher" bson:"publisher"`
	Lyricist      string             `json:"lyricist" bson:"lyricist"`
	Composer      string             `json:"composer" bson:"composer"`
	Translator    string             `json:"translator" bson:"translator"`
}

// DummySongs func: Make the dummy songs data for development.
func DummySongs() {
	Songs = append(Songs,
		Song{
			SID:      "1011054",
			NumC:     "11",
			NumI:     "54",
			Title:    "我獻上我心",
			Album:    "這是真愛",
			Tonality: "G",
			Year:     "",
			Language: "Chinese",
			Lyrics: []string{
				"p",
				"我心何等渴望，來尊崇你，主，我用全心來敬拜你，",
				"凡在我裡面的，都讚美你，我一切所愛，在於你。",
				"p",
				"主，我獻上我心，我獻上我的靈，",
				"我活著為了你，我的每個氣息，",
				"生命中的每個時刻，主，成全你旨意。",
				"p",
				"獻上我心，獻上我靈。",
			},
			Tempo:         "",
			TimeSignature: "",
			Publisher:     "",
			Lyricist:      "Reuben Morgan",
			Composer:      "Reuben Morgan",
			Translator:    "周巽光",
		})

	Songs = append(Songs,
		Song{
			SID:      "1010066",
			NumC:     "10",
			NumI:     "66",
			Title:    "前來敬拜",
			Album:    "新的事將要成就",
			Tonality: "G",
			Year:     "",
			Language: "Chinese",
			Lyrics: []string{
				"v",
				"哈利路亞，哈利路亞，前來敬拜永遠的君王，",
				"哈利路亞，哈利路亞，大聲宣告主榮耀降臨。",
				"c",
				"榮耀尊貴，能力權柄歸於你，",
				"你是我的救主，我的救贖，",
				"榮耀尊貴，能力權柄歸於你，",
				"你是配得，你是配得，你是配得我的敬拜。",
				"b",
				"榮耀尊貴，美麗無比，神的兒子，耶穌我的主，",
				"榮耀尊貴，美麗無比，神的兒子，耶穌我的主。",
			},
			Tempo:         "",
			TimeSignature: "",
			Publisher:     "",
			Lyricist:      "Reuben Morgan",
			Composer:      "Reuben Morgan",
			Translator:    "周巽光",
		})
}
