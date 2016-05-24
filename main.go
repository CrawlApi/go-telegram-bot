package main

import (
	"fmt"
	"github.com/llitfkitfk/go-tele/pkg/api"
	"github.com/tucnak/telebot"
	"log"
	"time"
	"regexp"
	"github.com/parnurzeal/gorequest"
)

var bot *telebot.Bot

func main() {
	if newBot, err := telebot.NewBot(api.BOT_TOKEN); err != nil {
		return
	} else {
		// shadowing, remember?
		bot = newBot
	}

	bot.Messages = make(chan telebot.Message, 1000)
	bot.Queries = make(chan telebot.Query, 1000)

	go messages()
	go queries()

	bot.Start(1 * time.Second)
}

func GetApi(url string) string {
	_, body, errs := gorequest.New().Timeout(5 * time.Second).Get(url).End()
	if errs != nil {
		return ""
	}
	log.Println(body)
	return body
}

func messages() {
	r, _ := regexp.Compile(api.REGEXP_MESSAGE)

	for message := range bot.Messages {
		matcher := r.FindStringSubmatch(message.Text)

		if len(matcher) > 0 {

			responseCh := make(chan string)

			if matcher[2] == "uid" {
				go func() {
					body := GetApi("http://192.168.30.64:10086/api/" + matcher[1] + "/" + matcher[2] + "?url=" + matcher[3])
					responseCh <- body
				}()
			} else {
				go func() {
					body := GetApi("http://192.168.30.64:10086/api/" + matcher[1] + "/" + matcher[2] + "/" + matcher[3])
					responseCh <- body
				}()
			}
			bot.SendMessage(message.Chat, <-responseCh, nil)
		}
		if message.Text == "/upload" {
			file, err := telebot.NewFile("index.html")
			if err != nil {
				bot.SendMessage(message.Chat, "Fail to upload", nil)
			}
			htmlFile := telebot.Document{
				File: file,
			}
			err = bot.SendDocument(message.Chat, &htmlFile, nil)

		}

		if message.Text == "/stringlen" {
			// length 4096
			strTest := `{"{"{"{"{"{"{"114111114411114141114111141111114111141111411114115111678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890{"data":[{"actions":[{"name":"Share","link":"https:\/\/www.facebook.com\/5718732097\/posts\/10153725820857098"}],"caption":"youtube.com","created_time":"2016-05-23T19:28:12+0000","description":"Congratulations to the more than 3,000 graduating students from the University of Wisconsin-Milwaukee. We \u201cCAN\u2019T STOP THE FEELING,\u201d we\u2019re so Panther Proud of...","from":{"name":"Justin Timberlake","id":"5718732097"},"full_picture":"https:\/\/fbexternal-a.akamaihd.net\/safe_image.php?d=AQAu5XNvzXkx7kJR&w=720&h=720&url=https\u00253A\u00252F\u00252Fi.ytimg.com\u00252Fvi\u00252FtFNU8ai1HNI\u00252Fmaxresdefault.jpg&cfs=1&sx=54&sy=0&sw=720&sh=720","icon":"https:\/\/www.facebook.com\/images\/icons\/post.gif","id":"5718732097_10153725820857098","is_expired":false,"instagram_eligibility":"eligible","is_hidden":false,"is_instagram_eligible":true,"is_published":true,"is_spherical":false,"link":"https:\/\/www.youtube.com\/watch?v=tFNU8ai1HNI&feature=youtu.be","message":"Congrats University of Wisconsin Milwaukee grads! \ud83c\udf89\ud83d\udc4f\ud83c\udffc\ud83d\udc4f\ud83c\udffc\ud83d\udc4f\ud83c\udffcImpressive!! #uwmgrad #CantStopTheFeeling #Graduation edition:","message_tags":[{"id":"14716860095","name":"University of Wisconsin Milwaukee","type":"page","offset":9,"length":33}],"name":"\u201cCAN\u2019T STOP THE FEELING!\u201d - UW-Milwaukee Graduation Edition","permalink_url":"https:\/\/www.facebook.com\/justintimberlake\/posts\/10153725820857098","picture":"https:\/\/fbexternal-a.akamaihd.net\/safe_image.php?d=AQAKSXF6w1HYXfNY&w=130&h=130&url=https\u00253A\u00252F\u00252Fi.ytimg.com\u00252Fvi\u00252FtFNU8ai1HNI\u00252Fmaxresdefault.jpg&cfs=1&sx=54&sy=0&sw=720&sh=720","privacy":{"value":"","description":"","friends":"","allow":"","deny":""},"promotion_status":"ineligible","shares":{"count":1684},"source":"https:\/\/www.youtube.com\/embed\/tFNU8ai1HNI?autoplay=1","status_type":"shared_story","timeline_visibility":"normal","type":"video","updated_time":"2016-05-24T10:41:48+0000","likes":{"data":[{"id":"280212842313003","name":"Len Villaruz"}],"paging":{"cursors":{"before":"MjgwMjEyODQyMzEzMDAz","after":"MjgwMjEyODQyMzEzMDAz"},"next":"https:\/\/graph.facebook.com\/v2.6\/5718732097_10153725820857098\/likes?access_token=490895874437565\u00257C3ce74d840577a6d598af56cd46fd0450&summary=true&limit=1&after=MjgwMjEyODQyMzEzMDAz"},"summary":{"total_count":19924,"can_like":false,"has_liked":false}},"comments":{"data":[{"created_time":"2016-05-23T19:34:58+0000","from":{"name":"Brandi Mariah Chapman","id":"10154161724961168"},"message":"I'm realizing today that this song IS the new Uptown Funk. Everyone legit loves it, even my Grandma. So catchy, you can't help it!","id":"10153725820857098_10153725836262098"}],"paging":{"cursors":{"before":"NDYy","after":"NDYy"},"next":"https:\/\/graph.facebook.com\/v2.6\/5718732097_10153725820857098\/comments?access_token=490895874437565\u00257C3ce74d840577a6d598af56cd46fd0450&summary=true&limit=1&after=NDYy"},"summary":{"order":"ranked","total_count":462,"can_comment":false}}}],"paging":{"previous":"https:\/\/graph.facebook.com\/v2.6\/5718732097\/feed?fields=actions,admin_creator,application,call_to_action,child_attachments,caption,comments_mirroring_domain,coordinates,created_time,description,event,expanded_height,expanded_width,feed_targeting,from,full_picture,height,icon,id,is_expired,is_crossposting_eligible,instagram_eligibility,is_hidden,is_instagram_eligible,is_popular,is_published,is_spherical,link,message,message_tags,name,object_id,parent_id,permalink_url,picture,place,privacy,promotion_status,properties,scheduled_publish_time,shares,source,status_type,story,story_tags,subscribed,target,targeting,timeline_visibility,type,updated_time,via,width,with_tags,comments.limit\u0025281\u002529.summary\u002528true\u002529,likes.limit\u0025281\u002529.summary\u002528true\u002529&limit=1&since=1464031692&access_token=490895874437565`
			log.Println(len(strTest))
			bot.SendMessage(message.Chat, strTest, nil)


		}

		if message.Text == "/hi" {
			bot.SendMessage(message.Chat, "Hello, " + message.Sender.FirstName + "!", nil)
		}
		if message.Text == "/userinfo" {
			bot.SendMessage(message.Chat, typeof(message.Sender), nil)
		}

		if message.Text == "/chatinfo" {
			bot.SendMessage(message.Chat, typeof(message.Chat), nil)
		}


		if message.Text == "/optionstest" {

			bot.SendMessage(message.Chat, "pong", &telebot.SendOptions{
				ReplyTo: message,

			})
		}

		if message.Text == "/test2bot" {

			bot.SendMessage(message.Chat, "pong", &telebot.SendOptions{
				ReplyMarkup: telebot.ReplyMarkup{
					ForceReply: true,
					Selective: true,

					CustomKeyboard: [][]string{
						[]string{"1", "2", "3"},
						[]string{"4", "5", "6"},
						[]string{"7", "8", "9"},
						[]string{"*", "0", "#"},
					},
				},
			},
			)
		}

		if message.Document.FileName == "index.html" {
			bot.SendMessage(message.Chat, typeof(message.Document), nil)
		}

		if message.Text == "/info" {
			info := "message.Unixtime: " + typeof(message.Unixtime) + " message.Location:" + typeof(message.Location) + " " + message.Text

			bot.SendMessage(message.Chat, info, nil)
		}

	}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}

func queries() {
	for query := range bot.Queries {
		log.Println("--- new query ---")
		log.Println("from:", query.From)
		log.Println("text:", query.Text)

		// There you build a slice of let's say, article results:
		results := []telebot.Result{}

		// And finally respond to the query:
		if err := bot.Respond(query, results); err != nil {
			log.Println("ouch:", err)
		}
	}
}
