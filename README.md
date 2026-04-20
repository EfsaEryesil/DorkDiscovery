# DorkDiscovery
Siber Tehdit İstihbaratı ve Pasif Keşif için Çok Dilli Google Dorking Aracı - Multilingual Google Dorking Tool for Cyber ​​Threat Intelligence and Passive Discovery

Bu proje, siber güvenlikte pasif keşif aşamasında kullanılan Google Dorking işlemlerini otomatize eden web tabanlı bir araçtır. Kullanıcıdan alınan hedef domain için 55 farklı dork sorgusu üretir ve tek tıkla Google'da aratır.

## Özellikler

- 55 adet Google dork sorgusu (GHDB tabanlı)
- Türkçe / İngilizce dil desteği
- Otomatik düzeltmeyi engelleyen verbatim mod (`&tbs=li:1`)
- Modern mor tema, mobil uyumlu arayüz

## Teknolojiler

- Go (Golang) 1.20+
- Standart kütüphaneler (`net/http`, `html/template`)

## Kullanım Adımları

1. **Derleme:**  `go run main.go` ile çalıştırılır.
2. **Erişim:** Tarayıcıdan `http://localhost:8080` adresine gidilir.
3. **Analiz:** Arama çubuğuna hedef domain (örn: `google.com`) yazılır ve butona tıklanır.
4. **Sonuçlar:** Kategorize edilmiş sekmelerden ilgili sorguların yanındaki "Google'da Ara" butonuna basılarak canlı sonuçlar incelenir.




---
---


# DorkDiscovery (English)
This project is a web-based tool that automates Google Dorking operations, one of the most critical tools in the passive reconnaissance phase of cybersecurity. It generates 55 different dork queries for a target domain provided by the user and allows one-click searching on Google.

## Features

- 55 Google dork queries (GHDB based)
- Turkish / English language support
- Verbatim mode (`&tbs=li:1`) to prevent automatic spell correction
- Modern purple theme, mobile-friendly interface

## Technologies

- Go (Golang) 1.20+
- Standard libraries (`net/http`, `html/template`)

## Usage Steps

1. **Build:**  Run directly with `go run main.go`.
2. **Access:** Go to `http://localhost:8080` in your browser.
3. **Analyze:** Enter the target domain (e.g., `google.com`) in the search box and click the button.
4. **Results:** Click the "Search on Google" button next to the relevant queries in the categorized sections to view live results.
