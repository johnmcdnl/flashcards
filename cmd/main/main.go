package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/johnmcdnl/flashcards"
)

func main() {

	var deck = NewDeck("deck.db")

	deck.Know = English
	deck.Learning = Russian

	for i := 1; i <= 10; i++ {
		deck.Next()
		deck.Current.PrintQuestion(deck)
		deck.Current.AttemptAnswer(deck.Know, deck.Learning, "")
		deck.SaveState()
		fmt.Println()
		fmt.Println()
	}

	toJSON(deck)

}

func toJSON(i interface{}) {
	j, _ := json.Marshal(i)
	// fmt.Println(string(j))
	ioutil.WriteFile("data.json", j, os.ModePerm)
}

func seed(deck *Deck) {
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "apple")).WithTranslation(NewTranslation(Russian, "яблоко").WithPhonetic(NewPhonetic(English, "yabloko")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "baby")).WithTranslation(NewTranslation(Russian, "детка").WithPhonetic(NewPhonetic(English, "detka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "back")).WithTranslation(NewTranslation(Russian, "назад").WithPhonetic(NewPhonetic(English, "nazad")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "ball")).WithTranslation(NewTranslation(Russian, "мяч").WithPhonetic(NewPhonetic(English, "myach")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "bear")).WithTranslation(NewTranslation(Russian, "медведь").WithPhonetic(NewPhonetic(English, "medved'")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "bed")).WithTranslation(NewTranslation(Russian, "постель").WithPhonetic(NewPhonetic(English, "postel'")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "bell")).WithTranslation(NewTranslation(Russian, "колокол").WithPhonetic(NewPhonetic(English, "kolokol")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "bird")).WithTranslation(NewTranslation(Russian, "птица").WithPhonetic(NewPhonetic(English, "ptitsa")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "birthday")).WithTranslation(NewTranslation(Russian, "день рождения").WithPhonetic(NewPhonetic(English, "den' rozhdeniya")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "boat")).WithTranslation(NewTranslation(Russian, "лодка").WithPhonetic(NewPhonetic(English, "lodka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "box")).WithTranslation(NewTranslation(Russian, "коробка").WithPhonetic(NewPhonetic(English, "korobka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "boy")).WithTranslation(NewTranslation(Russian, "мальчик").WithPhonetic(NewPhonetic(English, "mal'chik")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "bread")).WithTranslation(NewTranslation(Russian, "хлеб").WithPhonetic(NewPhonetic(English, "khleb")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "brother")).WithTranslation(NewTranslation(Russian, "брат").WithPhonetic(NewPhonetic(English, "brat")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "cake")).WithTranslation(NewTranslation(Russian, "кекс").WithPhonetic(NewPhonetic(English, "keks")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "car")).WithTranslation(NewTranslation(Russian, "автомобиль").WithPhonetic(NewPhonetic(English, "avtomobil'")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "cat")).WithTranslation(NewTranslation(Russian, "Кот").WithPhonetic(NewPhonetic(English, "Kot")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "chair")).WithTranslation(NewTranslation(Russian, "стул").WithPhonetic(NewPhonetic(English, "stul")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "chicken")).WithTranslation(NewTranslation(Russian, "курица").WithPhonetic(NewPhonetic(English, "kuritsa")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "children")).WithTranslation(NewTranslation(Russian, "дети").WithPhonetic(NewPhonetic(English, "deti")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "Christmas")).WithTranslation(NewTranslation(Russian, "рождество").WithPhonetic(NewPhonetic(English, "rozhdestvo")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "coat")).WithTranslation(NewTranslation(Russian, "Пальто").WithPhonetic(NewPhonetic(English, "Pal'to")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "corn")).WithTranslation(NewTranslation(Russian, "кукуруза").WithPhonetic(NewPhonetic(English, "kukuruza")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "cow")).WithTranslation(NewTranslation(Russian, "корова").WithPhonetic(NewPhonetic(English, "korova")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "day")).WithTranslation(NewTranslation(Russian, "день").WithPhonetic(NewPhonetic(English, "den'")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "dog")).WithTranslation(NewTranslation(Russian, "собака").WithPhonetic(NewPhonetic(English, "sobaka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "doll")).WithTranslation(NewTranslation(Russian, "кукла").WithPhonetic(NewPhonetic(English, "kukla")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "door")).WithTranslation(NewTranslation(Russian, "дверь").WithPhonetic(NewPhonetic(English, "dver'")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "duck")).WithTranslation(NewTranslation(Russian, "утка").WithPhonetic(NewPhonetic(English, "utka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "egg")).WithTranslation(NewTranslation(Russian, "яйцо").WithPhonetic(NewPhonetic(English, "yaytso")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "eye")).WithTranslation(NewTranslation(Russian, "глаз").WithPhonetic(NewPhonetic(English, "glaz")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "farm")).WithTranslation(NewTranslation(Russian, "ферма").WithPhonetic(NewPhonetic(English, "ferma")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "farmer")).WithTranslation(NewTranslation(Russian, "фермер").WithPhonetic(NewPhonetic(English, "fermer")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "father")).WithTranslation(NewTranslation(Russian, "отец").WithPhonetic(NewPhonetic(English, "otets")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "feet")).WithTranslation(NewTranslation(Russian, "ноги").WithPhonetic(NewPhonetic(English, "nogi")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "fire")).WithTranslation(NewTranslation(Russian, "Огонь").WithPhonetic(NewPhonetic(English, "Ogon'")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "fish")).WithTranslation(NewTranslation(Russian, "рыба").WithPhonetic(NewPhonetic(English, "ryba")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "floor")).WithTranslation(NewTranslation(Russian, "пол").WithPhonetic(NewPhonetic(English, "pol")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "flower")).WithTranslation(NewTranslation(Russian, "цветок").WithPhonetic(NewPhonetic(English, "tsvetok")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "game")).WithTranslation(NewTranslation(Russian, "игра").WithPhonetic(NewPhonetic(English, "igra")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "garden")).WithTranslation(NewTranslation(Russian, "сад").WithPhonetic(NewPhonetic(English, "sad")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "girl")).WithTranslation(NewTranslation(Russian, "девушка").WithPhonetic(NewPhonetic(English, "devushka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "good-bye")).WithTranslation(NewTranslation(Russian, "Прощай").WithPhonetic(NewPhonetic(English, "Proshchay")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "grass")).WithTranslation(NewTranslation(Russian, "трава").WithPhonetic(NewPhonetic(English, "trava")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "ground")).WithTranslation(NewTranslation(Russian, "земля").WithPhonetic(NewPhonetic(English, "zemlya")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "hand")).WithTranslation(NewTranslation(Russian, "рука").WithPhonetic(NewPhonetic(English, "ruka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "head")).WithTranslation(NewTranslation(Russian, "глава").WithPhonetic(NewPhonetic(English, "glava")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "hill")).WithTranslation(NewTranslation(Russian, "холм").WithPhonetic(NewPhonetic(English, "kholm")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "home")).WithTranslation(NewTranslation(Russian, "Главная").WithPhonetic(NewPhonetic(English, "Glavnaya")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "horse")).WithTranslation(NewTranslation(Russian, "лошадь").WithPhonetic(NewPhonetic(English, "loshad'")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "house")).WithTranslation(NewTranslation(Russian, "дом").WithPhonetic(NewPhonetic(English, "dom")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "kitty")).WithTranslation(NewTranslation(Russian, "Китти").WithPhonetic(NewPhonetic(English, "Kitti")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "leg")).WithTranslation(NewTranslation(Russian, "ножка").WithPhonetic(NewPhonetic(English, "nozhka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "letter")).WithTranslation(NewTranslation(Russian, "письмо").WithPhonetic(NewPhonetic(English, "pis'mo")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "man")).WithTranslation(NewTranslation(Russian, "человек").WithPhonetic(NewPhonetic(English, "chelovek")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "men")).WithTranslation(NewTranslation(Russian, "люди").WithPhonetic(NewPhonetic(English, "lyudi")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "milk")).WithTranslation(NewTranslation(Russian, "молоко").WithPhonetic(NewPhonetic(English, "moloko")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "money")).WithTranslation(NewTranslation(Russian, "Деньги").WithPhonetic(NewPhonetic(English, "Den'gi")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "morning")).WithTranslation(NewTranslation(Russian, "утро").WithPhonetic(NewPhonetic(English, "utro")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "mother")).WithTranslation(NewTranslation(Russian, "мама").WithPhonetic(NewPhonetic(English, "mama")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "name")).WithTranslation(NewTranslation(Russian, "имя").WithPhonetic(NewPhonetic(English, "imya")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "nest")).WithTranslation(NewTranslation(Russian, "гнездо").WithPhonetic(NewPhonetic(English, "gnezdo")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "night")).WithTranslation(NewTranslation(Russian, "ночь").WithPhonetic(NewPhonetic(English, "noch'")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "paper")).WithTranslation(NewTranslation(Russian, "бумага").WithPhonetic(NewPhonetic(English, "bumaga")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "party")).WithTranslation(NewTranslation(Russian, "вечеринка").WithPhonetic(NewPhonetic(English, "vecherinka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "picture")).WithTranslation(NewTranslation(Russian, "картина").WithPhonetic(NewPhonetic(English, "kartina")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "pig")).WithTranslation(NewTranslation(Russian, "свинья").WithPhonetic(NewPhonetic(English, "svin'ya")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "rabbit")).WithTranslation(NewTranslation(Russian, "кролик").WithPhonetic(NewPhonetic(English, "krolik")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "rain")).WithTranslation(NewTranslation(Russian, "дождь").WithPhonetic(NewPhonetic(English, "dozhd'")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "ring")).WithTranslation(NewTranslation(Russian, "кольцо").WithPhonetic(NewPhonetic(English, "kol'tso")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "robin")).WithTranslation(NewTranslation(Russian, "Робин").WithPhonetic(NewPhonetic(English, "Robin")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "Santa Claus")).WithTranslation(NewTranslation(Russian, "Дед Мороз").WithPhonetic(NewPhonetic(English, "Ded Moroz")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "school")).WithTranslation(NewTranslation(Russian, "школа").WithPhonetic(NewPhonetic(English, "shkola")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "seed")).WithTranslation(NewTranslation(Russian, "семя").WithPhonetic(NewPhonetic(English, "semya")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "sheep")).WithTranslation(NewTranslation(Russian, "овца").WithPhonetic(NewPhonetic(English, "ovtsa")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "shoe")).WithTranslation(NewTranslation(Russian, "башмак").WithPhonetic(NewPhonetic(English, "bashmak")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "sister")).WithTranslation(NewTranslation(Russian, "сестра").WithPhonetic(NewPhonetic(English, "sestra")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "snow")).WithTranslation(NewTranslation(Russian, "снег").WithPhonetic(NewPhonetic(English, "sneg")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "song")).WithTranslation(NewTranslation(Russian, "песня").WithPhonetic(NewPhonetic(English, "pesnya")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "squirrel")).WithTranslation(NewTranslation(Russian, "белка").WithPhonetic(NewPhonetic(English, "belka")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "stick")).WithTranslation(NewTranslation(Russian, "придерживаться").WithPhonetic(NewPhonetic(English, "priderzhivat'sya")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "street")).WithTranslation(NewTranslation(Russian, "улица").WithPhonetic(NewPhonetic(English, "ulitsa")))))
	deck.WithCard(NewCard(NewPhrase().WithTranslation(NewTranslation(English, "sun")).WithTranslation(NewTranslation(Russian, "солнце").WithPhonetic(NewPhonetic(English, "solntse")))))
	deck.SaveState()
}
