# HSE Group için PDF Belgesi Üreten CLI

Google Sheet üzerinden verileri okuyup çıktı olarak PDF belgelerini ve eklenecek
kayıtları CSV olarak oluşturan komut satırı uygulaması.

Örnek kullanım:
```
./hsedocumentgenerator

🚀      Google Sheets üzerinden okuma başlatıldı.
📗      Google Sheets üzerinden 2 kayıt okundu.
🤔      Oluşturulan PDFleri içeren ZIP dosyası nereye depolanacak?:     belgeler.zip
⏳       PDF belge üretme işlemi başlandı...
👍      [abE1Ec1-A.pdf] Yiğit   Sadıç   için PDF belgesi üretildi.
👍      [FE234-qZ.pdf]  Aycan   Çotoy   için PDF belgesi üretildi.
✅       PDF belgeleri "belgeler.zip" olarak sıkıştırıldı ve okunan kayıtlar Google Sheets içine eklendi.
💫      İşlem tamamlandı. İyi günler!
```


### Gerekenler

ENV variables:

- GOTENBERG_URL (default: http://localhost:3000)
