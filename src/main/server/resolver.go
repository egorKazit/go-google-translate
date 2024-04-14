package server

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"strings"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func resolveTranslation(originWord string, response []byte) (*Translation, error) {
	// covert string and split by newline
	bodyString := string(response)
	allFindings := strings.Split(bodyString, "\n")
	// 3rd line should contain json with data -> return error if any
	if allFindings == nil || len(allFindings) < 4 {
		log.Printf("Google returns less rows then expected")
		return nil, errors.New("provided google result is not correct")
	}
	// unmarshal 3rd line -> return error if any
	var responseArray []any
	err := json.Unmarshal([]byte(allFindings[3]), &responseArray)
	if err != nil {
		log.Printf("Can not unmarshal result: %s", err.Error())
		return nil, err
	}
	// at least one line has to be in data
	if responseArray == nil || len(responseArray) < 1 {
		log.Printf("Google returns less rows then expected")
		return nil, errors.New("provided google result is not correct")
	}
	// for now 2 interfaces expected
	if _, ok := responseArray[0].([]interface{}); ok && len(responseArray[0].([]interface{})) > 2 {
		err = json.Unmarshal([]byte(responseArray[0].([]interface{})[2].(string)), &responseArray)
		if err != nil {
			log.Printf("Can not unmarshal result: %s", err.Error())
			return nil, err
		}

		result := &Translation{}
		result.Origin = getOriginalWord(responseArray)
		result.Translation = getTranslation(responseArray[1])
		result.OtherTranslations = getOtherTranslations(responseArray)

		if strings.ToLower(result.Translation.Word) == strings.ToLower(originWord) &&
			!digitCheck.MatchString(originWord) && len(result.OtherTranslations) > 0 {
			result.Translation, result.OtherTranslations = result.OtherTranslations[0], result.OtherTranslations[1:]
		}

		if result.Origin.Word == "" {
			result.Origin.Word = originWord
		}

		return result, nil
	} else {
		return nil, errors.New("can not parse response")
	}

}

func getOriginalWord(translationObject any) WordDefinition {

	result := WordDefinition{}
	// if not a list -> return empty Translation
	if _, ok := translationObject.([]interface{}); !ok {
		log.Printf("can not covert type %v. Please create an incident if critical", translationObject)
		return result
	}
	// get nested. Translation should be at 0,0
	translation := getNested(translationObject.([]interface{}), []int{0, 0})

	// there should be a string at the end of the path
	if translationValue, ok := translation.(string); ok {
		// fill Translation
		result.Word = translationValue
	}

	// fill Meaning as well
	result.WordMeanings = getMeanings(getNested(translationObject, []int{3, 1, 0}))
	result.Examples = getExamples(getNested(translationObject, []int{3, 2, 0}))
	return result

}

func getTranslation(translationObject any) WordWithPathOfSpeech {

	result := WordWithPathOfSpeech{}
	wordCandidate := getNested(translationObject, []int{0, 0, 5, 0, 0})
	// there should be a string at the end of the path
	if wordCandidateValue, ok := wordCandidate.(string); ok {
		result.Word = wordCandidateValue
	}
	// fill other information
	result.PartOfSpeeches = getPartOfSpeech(getNested(translationObject, []int{0, 0, 9, 0}))
	return result

}

func getOtherTranslations(translationObject any) []WordWithPathOfSpeech {

	result := make([]WordWithPathOfSpeech, 0)
	// array expected
	if _, ok := translationObject.([]interface{}); !ok || len(translationObject.([]interface{})) < 6 {
		if !ok {
			log.Printf("can not covert type %v. Please create an incident if critical", translationObject)
		}
		return result
	}

	// go through array and fill Translation if possible
	if translationObjectValues, ok := translationObject.([]interface{})[6].([]interface{}); ok {
		for _, translation := range translationObjectValues {
			result = append(result, getTranslation(translation))
		}
	}
	return result

}

func getPartOfSpeech(translationObject any) map[float64]string {
	result := make(map[float64]string)
	// return if impossible to get part of speech
	if _, ok := translationObject.([]interface{}); !ok {
		log.Printf("can not covert type %v. Please create an incident if critical", translationObject)
		return result
	}
	// go through and fill
	for _, partOfSpeech := range translationObject.([]interface{}) {
		// if type is float, then fill in part of speech
		if _, ok := partOfSpeech.(float64); ok {
			result[partOfSpeech.(float64)] = wordTypes[int(partOfSpeech.(float64))-1]
		}
	}
	return result
}

func getSynonyms(translationObject any) []string {

	result := make([]string, 0)
	// get Synonyms at path
	if _, ok := translationObject.([]interface{}); !ok {
		log.Printf("can not covert type %v. Please create an incident if critical", translationObject)
		return result
	}
	// fill in if any
	for _, synonymObject := range translationObject.([]interface{}) {
		if synonymsObjectArrayRecord, ok := synonymObject.([]interface{}); !ok || len(synonymsObjectArrayRecord) == 0 {
			if !ok {
				log.Printf("can not covert type %v. Please create an incident if critical", synonymObject)
			}
			continue
		}
		if synonymValue, ok := synonymObject.([]interface{})[0].(string); ok {
			result = append(result, synonymValue)
		}
	}

	return result

}

func getMeanings(meaningsObject any) []WordMeaning {
	result := make([]WordMeaning, 0)
	if _, ok := meaningsObject.([]interface{}); !ok {
		log.Printf("can not covert type %v. Please create an incident if critical", meaningsObject)
		return result
	}
	for _, meaningObject := range meaningsObject.([]interface{}) {
		if _, ok := meaningObject.([]interface{}); !ok || len(meaningObject.([]interface{})) < 2 {
			if !ok {
				log.Printf("can not covert type %v. Please create an incident if critical", meaningObject)
			}
			continue
		}
		meaningsPerPartOfSpeech := meaningObject.([]interface{})[1]
		if _, ok := meaningsPerPartOfSpeech.([]interface{}); !ok {
			log.Printf("can not covert type %v. Please create an incident if critical", meaningsPerPartOfSpeech)
			continue
		}
		for _, meaningPerPartOfSpeech := range meaningsPerPartOfSpeech.([]interface{}) {
			if _, ok := meaningPerPartOfSpeech.([]interface{}); !ok {
				log.Printf("can not covert type %v. Please create an incident if critical", meaningPerPartOfSpeech)
				continue
			}
			wordMeaning := WordMeaning{}
			// if Meaning is presented, then fill in
			if meaning, ok := meaningPerPartOfSpeech.([]interface{})[0].(string); ok {
				wordMeaning.Meaning = meaning
			}
			if len(meaningPerPartOfSpeech.([]interface{})) > 1 {
				// the same for example
				if meaningUsage, ok := meaningPerPartOfSpeech.([]interface{})[1].(string); ok {
					wordMeaning.Usage = meaningUsage
				}
			}
			if len(meaningPerPartOfSpeech.([]interface{})) > 3 {
				// and part of speech
				if partOfSpeech, ok := meaningObject.([]interface{})[3].(float64); ok {
					wordMeaning.PartOfSpeech = partOfSpeech
					wordMeaning.PartOfSpeechName = wordTypes[int(wordMeaning.PartOfSpeech)-1]
				}
			}
			if len(meaningPerPartOfSpeech.([]interface{})) > 5 {
				wordMeaning.Synonyms = getSynonyms(getNested(meaningPerPartOfSpeech.([]interface{})[5], []int{0, 0}))
			}
			result = append(result, wordMeaning)
		}
	}
	return result
}

func getExamples(examplesObject any) []string {
	result := make([]string, 0)
	if _, ok := examplesObject.([]interface{}); !ok {
		log.Printf("can not covert type %v. Please create an incident if critical", examplesObject)
		return result
	}
	for _, exampleObject := range examplesObject.([]interface{}) {
		if _, ok := exampleObject.([]interface{}); !ok || len(exampleObject.([]interface{})) < 2 {
			if !ok {
				log.Printf("can not covert type %v. Please create an incident if critical", examplesObject)
			}
			continue
		}
		if example, ok := exampleObject.([]interface{})[1].(string); ok {
			result = append(result, example)
		}
	}
	return result
}
