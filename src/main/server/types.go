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
	origin            WordDefinition
	translation       WordWithPathOfSpeech
	otherTranslations []WordWithPathOfSpeech
}

type WordWithPathOfSpeech struct {
	word           string
	partOfSpeeches map[float64]string
}

type WordDefinition struct {
	word         string
	wordMeanings []WordMeaning
	examples     []string
}

type WordMeaning struct {
	meaning          string
	usage            string
	synonyms         []string
	partOfSpeech     float64
	partOfSpeechName string
}
