package server

var wordTypes = [...]string{
	"Noun",
	"Verb",
	"Adjective",
	"Adverb",
	"Preposition",
	"Abbreviation",
	"Conjunction",
	"Pronoun",
	"Interjection",
	"Phrase",
	"Prefix",
	"Suffix",
	"Article",
	"Combining form",
	"Numeral",
	"Auxiliary verb",
	"Exclamation",
	"Plural",
	"Particle"}

type Translation struct {
	Origin            WordDefinition
	Translation       WordWithPathOfSpeech
	OtherTranslations []WordWithPathOfSpeech
}

type WordWithPathOfSpeech struct {
	Word           string
	PartOfSpeeches map[float64]string
}

type WordDefinition struct {
	Word         string
	WordMeanings []WordMeaning
	Examples     []string
}

type WordMeaning struct {
	Meaning          string
	Usage            string
	Synonyms         []string
	PartOfSpeech     float64
	PartOfSpeechName string
}
