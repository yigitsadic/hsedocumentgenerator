package translations

import "fmt"

const (
	TurkishStateName        = "Özbekistan Cumhuriyeti"
	TurkishCertificateTitle = "Katılım Belgesi"
	TurkishText             = "%s ve %s tarihleri arasındaki %s saatlik eğitime katılarak ve sonrasında \"Ölçme ve Değerlendirme Testi\"nde başarılı olarak, bu belgeyi almaya hak kazanmıştır."

	EnglishStateName        = "Republic of Uzbekistan"
	EnglishCertificateTitle = "Certificate of Participation"
	EnglishText             = "By participating in the %s-hour training dated between %s-%s and succeeding in the Assessment and Evaluation Test, he/she has been entitled to receive this certificate."

	RussianStateName        = "Республика Узбекистан"
	RussianCertificateTitle = "Cертификат"
	RussianText             = "Проведена стажировка в количестве %s часавого учение с %s до %s сдал экзамен <<Тест по измерениям и оценки >> Допущен к самостоятельной работе и выдан Сертификат"
)

type TranslatedContent struct {
	StateName string
	Title     string
	Content   string
}

func TranslateTo(educationHours, educationDateStart, educationDateEnd, lang string) TranslatedContent {
	switch lang {
	case "TR":
		return TranslatedContent{
			StateName: TurkishStateName,
			Title:     TurkishCertificateTitle,
			Content:   fmt.Sprintf(TurkishText, educationDateStart, educationDateEnd, educationHours),
		}
	case "EN":
		return TranslatedContent{
			StateName: EnglishStateName,
			Title:     EnglishCertificateTitle,
			Content:   fmt.Sprintf(EnglishText, educationHours, educationDateStart, educationDateEnd),
		}
	case "RU":
		return TranslatedContent{
			StateName: RussianStateName,
			Title:     RussianCertificateTitle,
			Content:   fmt.Sprintf(RussianText, educationHours, educationDateStart, educationDateEnd),
		}
	default:
		return TranslatedContent{
			StateName: RussianStateName,
			Title:     RussianCertificateTitle,
			Content:   fmt.Sprintf(RussianText, educationHours, educationDateStart, educationDateEnd),
		}
	}
}
