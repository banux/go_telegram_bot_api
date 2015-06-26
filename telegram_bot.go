package telegram_bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Result struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId           int         `json:"message_id"`
	From                User        `json:"from"`
	Date                int         `json:"date"`
	Chat                Chat        `json:"chat"`
	ForwardFrom         User        `json:"forward_from"`
	ForwardDate         int         `json:"forward_date"`
	ReplyToMessage      *Message    `json:"reply_to_message"`
	Text                string      `json:"text"`
	Audio               Audio       `json:"audio"`
	Document            Document    `json:"document"`
	Photo               []PhotoSize `json:"photo"`
	Sticker             Sticker     `json:"sticker"`
	Video               Video       `json:"video"`
	Contact             Contact     `json:"contact"`
	Location            Location    `json:"location"`
	NewChatParticipant  User        `json:"new_chat_participant"`
	LeftChatParticipant User        `json:"left_chat_participant"`
	NewChatTitle        string      `json:"new_chat_title"`
	NewChatPhoto        string      `json:"new_chat_photo"`
	DeleteChatPhoto     bool        `json:"delete_chat_photo"`
	GroupChatCreated    bool        `json:"group_chat_created"`
}

type Document struct {
	FileId   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	Filesize int       `json:"file_size"`
}

type Audio struct {
	FileId   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

type Sticker struct {
	FileId   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	FileSize int       `json:"file_size"`
}

type PhotoSize struct {
	FileId   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type GroupChat struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Chat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Title     string `json:"title"`
}

type Video struct {
	FileId   string    `json:"file_id"`
	Duration int       `json:"duration"`
	MimeType string    `json:"mime_type"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	FileSize int       `json:"file_size"`
	Caption  string    `json:"caption"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      string `json:"user_id"`
}

type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type UserProfilePhotos struct {
	TotalCount int         `json:"total_count"`
	Photos     []PhotoSize `json:"photos"`
}

type ReplyKeyboardMarkup struct {
}

type ReplyKeyboardhide struct {
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

type TelegramBot struct {
	Token   string
	BaseUrl string
}

func NewBot(token string) *TelegramBot {
	return &TelegramBot{Token: token, BaseUrl: "https://api.telegram.org/bot" + token}
}

func (t *TelegramBot) SetToken(token string) {
	t.Token = token
	t.BaseUrl = "https://api.telegram.org/bot" + token
}

func (t *TelegramBot) GetUpdates() *Result {
	var result Result

	client := &http.Client{}

	updatesUrl := t.BaseUrl + "/getUpdates"

	req, _ := http.NewRequest("GET", updatesUrl, nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return &Result{}
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	err = json.Unmarshal(body, &result)
	if err != nil {
		return &Result{}
	} else {
		return &result
	}
}

func (t *TelegramBot) SendMessage(chatId int, text string) {
	// client := &http.Client{}

	query := url.Values{}
	StrChatId := strconv.FormatInt(int64(chatId), 10)
	query.Set("chat_id", StrChatId)
	query.Set("text", text)
	fmt.Println(text)
	fmt.Println(StrChatId)

	sendUrl := t.BaseUrl + "/sendMessage?" + query.Encode()
	fmt.Println(sendUrl)

	resp, err := http.Get(sendUrl)
	if err != nil {
		fmt.Printf("%+v\n", resp)
	}
	fmt.Printf("%+v\n", resp)

}
