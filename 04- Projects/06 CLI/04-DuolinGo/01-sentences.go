package main

type Sentences struct {
	Sentences map[string]string 
} 


func (sentences Sentences) ReturnSentences() map[string]string {
	sentences.Sentences = map[string]string{
		"Common German Slang":"Allgemeiner deutscher Slang",
		"Moin, moin":"Morning - Hi - Hello - Good day - How are you?",
		"Geil":"Awesome - Cool - Sexy",
		"Dit jefällt ma":"I like it",
		"Na?":"Hey, what's up? - How are you?",
		"Quatsch":"Nonsense - That's ridiculous",
		"Ich habe die Nase voll":"Meaning: I'm fed up - I'm sick of it\nLiteral: I have a full nose",
		"Das ist nicht mein Bier":"Meaning: Not my problem\nLiteral: That's not my beer",
		"das ist dein Bier":"this is your beer - It's all yours",
		"Abwarten und Tee trinken":"Just wait and see",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
		"":"",
	}
}


