package grifts

import (
	"github.com/icrowley/fake"
	"github.com/myWebsite/golang/models"
)

var languages = []*models.Language{
	&models.Language{
		Name:          "Golang",
		Color:         "",
		Description:   fake.CharactersN(250),
		Documentation: "http://golang.org",
	},

	&models.Language{
		Name:          "Javascript",
		Color:         "",
		Description:   fake.CharactersN(250),
		Documentation: "https://developer.mozilla.org/ro/docs/Web/JavaScript",
	},

	&models.Language{
		Name:          "PHP",
		Color:         "",
		Description:   fake.CharactersN(250),
		Documentation: "https://php.net/",
	},

	&models.Language{
		Name:          "Ruby",
		Color:         "",
		Description:   fake.CharactersN(250),
		Documentation: "https://ruby-doc.org/",
	},

	&models.Language{
		Name:          "C++",
		Color:         "",
		Description:   fake.CharactersN(250),
		Documentation: "https://devdocs.io/cpp/",
	},

	&models.Language{
		Name:          "C",
		Color:         "",
		Description:   fake.CharactersN(250),
		Documentation: "https://devdocs.io/c/",
	},

	&models.Language{
		Name:          "Python",
		Color:         "",
		Description:   fake.CharactersN(250),
		Documentation: "https://www.python.org/",
	},

	&models.Language{
		Name:          "C#",
		Color:         "",
		Description:   fake.CharactersN(250),
		Documentation: "https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/",
	},
}

var licenses = []*models.License{
	&models.License{
		Name:        fake.Words(),
		Description: fake.CharactersN(250),
		Body:        fake.Paragraphs(),
		Nickname:    fake.CharactersN(4),
		URL:         fake.CharactersN(40),
		Key:         fake.CharactersN(5),
	},

	&models.License{
		Name:        fake.Words(),
		Description: fake.CharactersN(250),
		Body:        fake.Paragraphs(),
		Nickname:    fake.CharactersN(4),
		URL:         fake.CharactersN(40),
		Key:         fake.CharactersN(5),
	},

	&models.License{
		Name:        fake.Words(),
		Description: fake.CharactersN(250),
		Body:        fake.Paragraphs(),
		Nickname:    fake.CharactersN(4),
		URL:         fake.CharactersN(40),
		Key:         fake.CharactersN(5),
	},

	&models.License{
		Name:        fake.Words(),
		Description: fake.CharactersN(250),
		Body:        fake.Paragraphs(),
		Nickname:    fake.CharactersN(4),
		URL:         fake.CharactersN(40),
		Key:         fake.CharactersN(5),
	},
}


