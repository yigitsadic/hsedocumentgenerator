package translations

import "fmt"

const (
	TurkishStateName        = "Özbekistan Cumhuriyeti"
	TurkishCertificateTitle = "Katılım Belgesi"
	TurkishText             = "%s tarihli %s saatlik eğitime katılarak ve sonrasında \"Ölçme ve Değerlendirme Testi\"nde başarılı olarak, bu belgeyi almaya hak kazanmıştır."

	EnglishStateName        = "Republic of Uzbekistan"
	EnglishCertificateTitle = "Certificate of Participation"
	EnglishText             = "By participating in the %s-hour training dated %s and succeeding in the Assessment and Evaluation Test, he/she has been entitled to receive this certificate."

	RussianStateName        = "Республика Узбекистан"
	RussianCertificateTitle = "Cертификат"
	RussianText             = "Участник в %s часовых тренинге, проведенном в %s, и по окончании тренинга он успешно сдал «Тест по измерениям и оценке» и имел право на получение этого сертификата."
)

type TranslatedContent struct {
	StateName string
	Title     string
	Content   string
}

func TranslateTo(educationHours, educationDate, lang string) TranslatedContent {
	switch lang {
	case "TR":
		return TranslatedContent{
			StateName: TurkishStateName,
			Title:     TurkishCertificateTitle,
			Content:   fmt.Sprintf(TurkishText, educationDate, educationHours),
		}
	case "EN":
		return TranslatedContent{
			StateName: EnglishStateName,
			Title:     EnglishCertificateTitle,
			Content:   fmt.Sprintf(EnglishText, educationHours, educationDate),
		}
	case "RU":
		return TranslatedContent{
			StateName: RussianStateName,
			Title:     RussianCertificateTitle,
			Content:   fmt.Sprintf(RussianText, educationHours, educationDate),
		}
	default:
		return TranslatedContent{
			StateName: RussianStateName,
			Title:     RussianCertificateTitle,
			Content:   fmt.Sprintf(RussianText, educationHours, educationDate),
		}
	}
}
