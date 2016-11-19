package dto

type InstagramUser struct {
	Username       string `json:"username"`
	Bio            string `json:"bio"`
	Website        string `json:"website"`
	ProfilePicture string `json:"profile_picture"`
	FullName       string `json:"full_name"`
	Id             string `json:"id"`
}

type InstagramUsers []InstagramUser

type InstagramMedia struct {
}

type InstagramMedias []InstagramMedia

type InstagramHashtag struct {
	Name  string `json:"hashtag"`
	Count int    `json:"count"`
}

type InstagramHashtags []InstagramHashtag

/*
{
		"attribution": null,
		"tags": [
			"gyeongridan",
			"coffee",
			"카페",
			"카페사루",
			"사루",
			"morning",
			"itaewon",
			"커피",
			"이태원",
			"경리단길",
			"interior",
			"망고플레이트",
			"cafe",
			"saru"
		],
		"type": "image",
		"location": {
			"latitude": 37.539191742988,
			"name": "카페 사루 - saru",
			"longitude": 126.99436941793,
			"id": 1027903396
		},
		"comments": {
			"count": 0
		},
		"filter": "Normal",
		"created_time": "1478743870",
		"link": "https://www.instagram.com/p/BMnMVlRBUKV/",
		"likes": {
			"count": 178
		},
		"images": {
			"low_resolution": {
				"url": "https://scontent.cdninstagram.com/t51.2885-15/s320x320/e35/14705134_151201735350834_5683230614468165632_n.jpg?ig_cache_key=MTM4MDEyNjA3MDQ5MzQ5NTk1Nw%3D%3D.2",
				"width": 320,
				"height": 213
			},
			"thumbnail": {
				"url": "https://scontent.cdninstagram.com/t51.2885-15/s150x150/e35/c180.0.720.720/14705134_151201735350834_5683230614468165632_n.jpg?ig_cache_key=MTM4MDEyNjA3MDQ5MzQ5NTk1Nw%3D%3D.2.c",
				"width": 150,
				"height": 150
			},
			"standard_resolution": {
				"url": "https://scontent.cdninstagram.com/t51.2885-15/s640x640/sh0.08/e35/14705134_151201735350834_5683230614468165632_n.jpg?ig_cache_key=MTM4MDEyNjA3MDQ5MzQ5NTk1Nw%3D%3D.2",
				"width": 640,
				"height": 426
			}
		},
		"users_in_photo": [],
		"caption": {
			"created_time": "1478743870",
			"text": "☕️풀렸다고는 하는데, 그래도 너무 쌀쌀하죠. 창가에 앉아 따스한 커피 한잔 하고싶은 날!\n.\n.\n.\n.\n.\n#망고플레이트 띵시님 멋진 사진 감사드려요.\n#이태원 #경리단길 #사루 #카페사루 #카페 #커피 #itaewon #gyeongridan #saru #cafe #coffee #interior #morning",
			"from": {
				"username": "mangoplate",
				"profile_picture": "https://scontent.cdninstagram.com/t51.2885-19/s150x150/11849045_483082581852341_2060591985_a.jpg",
				"id": "646748795",
				"full_name": "망고플레이트 | MangoPlate"
			},
			"id": "17854915681126961"
		},
		"user_has_liked": false,
		"id": "1380126070493495957_646748795",
		"user": {
			"username": "mangoplate",
			"profile_picture": "https://scontent.cdninstagram.com/t51.2885-19/s150x150/11849045_483082581852341_2060591985_a.jpg",
			"id": "646748795",
			"full_name": "망고플레이트 | MangoPlate"
		}
	}
*/
