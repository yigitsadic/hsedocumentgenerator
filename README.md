# HSE Group için PDF Belgesi Üreten CLI

Google Sheet üzerinden verileri okuyup çıktı olarak PDF belgelerini ve eklenecek
kayıtları CSV olarak oluşturan komut satırı uygulaması.

Örnek kullanım:
```
./hsedocumentgenerator

🚀 Google Sheets üzerinden okuma başlatıldı.
📗 Google Sheets üzerinden 30 kayıt okundu.

🤔 Eklenenler için çıktı dosyası ne olacak?: olusanlar.csv
🤔 Oluşturulan PDFleri içeren ZIP dosyası nereye depolanacak?: belgeler.zip

⏳ PDF belge üretme işlemi başlandı...
👍 [1/30] PDF belgesi üretildi.
👍 [2/30] PDF belgesi üretildi.
...
👍 [30/30] PDF belgesi üretildi.

✅ PDF belgeleri "belgeler.zip" olarak sıkıştırıldı ve eklendi.
✅ Eklenen kayıtlar "olusanlar.csv" olarak eklendi.

💫 İşlem tamamlandı. İyi günler!
```
