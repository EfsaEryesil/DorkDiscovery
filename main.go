package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type dork_bilgisi struct {
	Id          string
	Sorgu       string
	Aciklama_tr string
	Aciklama_en string
}

type dork_grubu struct {
	KategoriAdi_tr string
	KategoriAdi_en string
	Dorklar        []dork_bilgisi
}

type sayfa_yapisi struct {
	HedefSite    string
	DorkGruplari []dork_grubu
	ToplamSayi   int
	Lang         string
}

var veriler = []dork_grubu{
	{
		KategoriAdi_tr: "1. Giriş ve Login Panelleri",
		KategoriAdi_en: "1. Login & Admin Panels",
		Dorklar: []dork_bilgisi{
			{"GHDB-01", "inurl:adminlogin.php", "Admin giriş sayfaları tespiti", "Detect admin login pages"},
			{"GHDB-02", "intitle:\"Login\" \"admin\"", "Başlığında admin ve login geçen sayfalar", "Pages with 'admin' and 'login' in title"},
			{"GHDB-03", "inurl:wp-admin", "Wordpress yönetim panelleri", "WordPress admin panels"},
			{"GHDB-04", "inurl:cpanel/main.php", "cPanel yönetim ekranları", "cPanel management interfaces"},
			{"GHDB-05", "intitle:\"Control Panel\" \"login\"", "Genel kontrol paneli girişleri", "Generic control panel logins"},
			{"GHDB-06", "inurl:/auth/login", "Oturum açma servisleri", "Authentication services"},
			{"GHDB-07", "inurl:signin", "Kullanıcı giriş sayfaları", "User sign-in pages"},
			{"GHDB-08", "intitle:\"Dashboard\" inurl:login", "Dashboard giriş ekranları", "Dashboard login screens"},
			{"GHDB-09", "inurl:/phpmyadmin/index.php", "Veritabanı yönetim paneli (phpmyadmin)", "Database management panel (phpMyAdmin)"},
		},
	},
	{
		KategoriAdi_tr: "2. Hassas Dosyalar ve Sızıntılar",
		KategoriAdi_en: "2. Sensitive Files & Leaks",
		Dorklar: []dork_bilgisi{
			{"GHDB-10", "filetype:env \"DB_PASSWORD\"", "Veritabanı şifreleri içeren .env dosyaları", ".env files containing database passwords"},
			{"GHDB-11", "filetype:sql \"INSERT INTO\"", "Veritabanı yedekleri ve SQL dump'lar", "Database backups and SQL dumps"},
			{"GHDB-12", "filetype:bak \"config\"", "Yapılandırma dosyası yedekleri", "Configuration file backups"},
			{"GHDB-13", "filetype:log \"password\"", "Log dosyalarındaki şifre kalıntıları", "Password remnants in log files"},
			{"GHDB-14", "filetype:ini \"password\"", "Sistem ayarlarındaki parolalar", "Passwords in system INI files"},
			{"GHDB-15", "filetype:pdf \"gizli\"", "Gizli içerikli PDF belgeleri", "PDF documents marked as confidential"},
			{"GHDB-16", "filetype:xlsx \"maas\"", "Maaş listeleri ve Excel tabloları", "Salary lists and Excel sheets"},
			{"GHDB-17", "filetype:txt \"private key\"", "Özel anahtar (Private Key) dosyaları", "Private key text files"},
			{"GHDB-18", "extension:pfx", "Sertifika ve anahtar paketleri", "Certificate and key packages"},
			{"GHDB-19", "filename:connections.xml", "Veritabanı bağlantı detayları dosyası", "Database connection details XML"},
			{"GHDB-20", "filetype:json \"api_key\"", "API anahtarı içeren JSON dosyaları", "JSON files with API keys"},
		},
	},
	{
		KategoriAdi_tr: "3. DevOps ve Bulut Sızıntıları",
		KategoriAdi_en: "3. DevOps & Cloud Leaks",
		Dorklar: []dork_bilgisi{
			{"GHDB-21", "intitle:\"Dashboard [Jenkins]\"", "Açık Jenkins otomasyon sunucuları", "Exposed Jenkins automation servers"},
			{"GHDB-22", "inurl:docker-compose.yml", "Docker yapılandırma dosyaları", "Docker compose configuration files"},
			{"GHDB-23", "inurl:/.git/config", "Açık unutulmuş .git dizini", "Exposed .git directory"},
			{"GHDB-24", "inurl:swagger-ui.html", "API endpoint dökümanları", "API endpoint documentation"},
			{"GHDB-25", "intitle:\"Sonarqube\"", "Kod analizi ekranları", "Code quality dashboards (SonarQube)"},
			{"GHDB-26", "inurl:/.vscode/settings.json", "VS Code ayar dosyaları", "VS Code settings files"},
			{"GHDB-27", "inurl:/.gitlab-ci.yml", "Gitlab CI/CD boru hattı dosyaları", "GitLab CI/CD pipeline files"},
			{"GHDB-28", "intitle:\"index of\" \"/node_modules/\"", "NodeJS kütüphane dizinleri", "NodeJS library directories"},
			{"GHDB-29", "site:s3.amazonaws.com", "AWS S3 bucket sızıntıları", "AWS S3 bucket leaks"},
			{"GHDB-30", "site:firebaseio.com", "Açık Firebase veritabanı kayıtları", "Exposed Firebase database records"},
		},
	},
	{
		KategoriAdi_tr: "4. Dizin Listeleme (Open Directories)",
		KategoriAdi_en: "4. Directory Listing (Open Directories)",
		Dorklar: []dork_bilgisi{
			{"GHDB-31", "intitle:\"index of\" \"/etc/\"", "Linux sistem dosyaları dizini", "Linux system configuration directory"},
			{"GHDB-32", "intitle:\"index of\" \"/backup/\"", "Yedek dosyalarının tutulduğu klasör", "Backup folders"},
			{"GHDB-33", "intitle:\"index of\" \"/mail/\"", "E-posta arşivi klasörleri", "Email archive directories"},
			{"GHDB-34", "intitle:\"index of\" \"/db/\"", "Veritabanı dosyaları klasörü", "Database directories"},
			{"GHDB-35", "intitle:\"index of\" \"/private/\"", "Özel dosya dizinleri", "Private file directories"},
			{"GHDB-36", "inurl:ftp:// \"index of\"", "Açık FTP dizinleri", "Open FTP directories"},
			{"GHDB-37", "intitle:\"index of\" \"/logs/\"", "Sistem günlükleri dizini", "System logs directory"},
			{"GHDB-38", "intitle:\"index of\" \"/uploads/\"", "Yüklenen dosyaların dizini", "Uploads directory"},
			{"GHDB-39", "intitle:\"index of\" \"/conf/\"", "Yapılandırma dosyaları klasörü", "Configuration files directory"},
		},
	},
	{
		KategoriAdi_tr: "5. Hata Mesajları ve Loglar",
		KategoriAdi_en: "5. Error Messages & Logs",
		Dorklar: []dork_bilgisi{
			{"GHDB-40", "\"SQL syntax error\"", "SQL injection zafiyet izleri tespiti", "SQL injection vulnerability traces"},
			{"GHDB-41", "\"PHP Parse error\"", "PHP kodlama hataları", "PHP parsing errors"},
			{"GHDB-42", "\"Traceback (most recent call last)\"", "Python hata yığını çıktıları", "Python traceback outputs"},
			{"GHDB-43", "\"ORA-00933\"", "Oracle DB syntax hataları", "Oracle DB syntax errors"},
			{"GHDB-44", "\"MySQL Error\"", "MySQL veritabanı hata mesajları", "MySQL database errors"},
			{"GHDB-45", "\"warning: mysql_connect()\"", "Veritabanı bağlantı uyarıları", "Database connection warnings"},
			{"GHDB-46", "\"Exception details:\"", "Sistem istisna ayrıntıları", "System exception details"},
			{"GHDB-47", "\"access denied for user\"", "Erişim engellenen kullanıcı kayıtları", "Access denied user records"},
			{"GHDB-48", "\"Fatal error:\"", "Kritik çalışma zamanı hataları", "Fatal runtime errors"},
		},
	},
	{
		KategoriAdi_tr: "6. Donanım ve IoT Cihazları",
		KategoriAdi_en: "6. Hardware & IoT Devices",
		Dorklar: []dork_bilgisi{
			{"GHDB-49", "intitle:\"Network Printer Status\"", "Yazıcı yönetim panelleri", "Printer management panels"},
			{"GHDB-50", "inurl:axis-cgi/jpg", "Kamera canlı görüntüleri", "Live camera feeds"},
			{"GHDB-51", "intitle:\"webcamXP 5\"", "Webcam yönetim ekranları", "Webcam management interfaces"},
			{"GHDB-52", "intitle:\"D-Link\" inurl:config", "D-Link cihaz yapılandırması", "D-Link device configuration"},
			{"GHDB-53", "intitle:\"SonicWALL\" \"login\"", "Sonicwall giriş ekranları", "SonicWall login screens"},
			{"GHDB-54", "inurl:\"/view/index.shtml\"", "Ağ kamerası yayınları", "Network camera streams"},
			{"GHDB-55", "intitle:\"Toshiba Network Camera\"", "Toshiba kamera arayüzleri", "Toshiba camera interfaces"},
		},
	},
}

func sayfa_isleyici(w http.ResponseWriter, r *http.Request) {

	lang := r.URL.Query().Get("lang")
	if lang != "en" && lang != "tr" {
		lang = "tr"
	}
	site_adi := strings.TrimSpace(r.URL.Query().Get("hedef"))

	sablon, _ := template.New("index").Parse(html_kodlari)

	paket := sayfa_yapisi{
		HedefSite:    site_adi,
		DorkGruplari: veriler,
		ToplamSayi:   55,
		Lang:         lang,
	}

	sablon.Execute(w, paket)
}

func main() {
	http.HandleFunc("/", sayfa_isleyici)
	fmt.Println("DorkDiscovery Çok Dilli Sürüm: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

const html_kodlari = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>DorkDiscovery - {{if eq .Lang "en"}}Multi-Language{{else}}Çok Dilli{{end}}</title>
    <style>
        body { background-color: #0b0614; color: #e2e8f0; font-family: 'Segoe UI', sans-serif; padding: 40px; }
        .ana-bolum { max-width: 1100px; margin: auto; }
        h1 { text-align: center; color: #a855f7; font-size: 35px; letter-spacing: -1px; }
        h1 span { color: #fdfdfd; }
        .ust-kisim { background: #1a102e; padding: 25px; border-radius: 15px; margin-bottom: 30px; border: 1px solid #3b2063; box-shadow: 0 4px 15px rgba(0,0,0,0.4); }
        .dil-select { float: right; margin-bottom: 15px; }
        .dil-select a { color: #a855f7; margin-left: 10px; text-decoration: none; font-weight: bold; }
        .dil-select a.active { color: white; text-decoration: underline; }
        input { width: 70%; padding: 14px; border-radius: 8px; border: 1px solid #3b2063; background: #0b0614; color: white; font-size: 15px; outline: none; }
        input:focus { border-color: #a855f7; }
        button { padding: 14px 30px; border-radius: 8px; border: none; background: #a855f7; color: #fff; font-weight: bold; cursor: pointer; transition: 0.3s; }
        button:hover { background: #9333ea; transform: translateY(-2px); }
        .liste { display: flex; flex-wrap: wrap; gap: 20px; justify-content: center; }
        .kategori-ismi { width: 100%; border-left: 5px solid #a855f7; color: #a855f7; padding: 5px 15px; margin-top: 30px; font-weight: bold; text-transform: uppercase; font-size: 14px; }
        .kart { background: #1a102e; border: 1px solid #2e1a4a; border-radius: 12px; padding: 20px; width: 320px; transition: 0.3s; }
        .kart:hover { border-color: #a855f7; background: #24163f; }
        .id-etiket { background: #7e22ce; color: white; font-size: 10px; padding: 4px 8px; border-radius: 5px; font-weight: bold; }
        .aciklama { font-size: 13px; color: #94a3b8; margin: 15px 0; min-height: 35px; }
        code { display: block; background: #000; padding: 12px; border-radius: 8px; color: #d8b4fe; font-size: 11px; word-break: break-all; border: 1px solid #2e1a4a; }
        .link { display: block; margin-top: 15px; text-align: center; background: #3b2063; color: white; text-decoration: none; padding: 10px; border-radius: 8px; font-size: 12px; font-weight: bold; }
        .link:hover { background: #a855f7; }
        .bilgi { text-align: center; color: #4b5563; font-size: 12px; margin-top: 50px; border-top: 1px solid #1a102e; padding-top: 20px; }
    </style>
</head>
<body>
    <div class="ana-bolum">
        <div class="dil-select">
            <a href="?lang=tr{{if .HedefSite}}&hedef={{.HedefSite}}{{end}}" {{if eq .Lang "tr"}}class="active"{{end}}>Türkçe</a> |
            <a href="?lang=en{{if .HedefSite}}&hedef={{.HedefSite}}{{end}}" {{if eq .Lang "en"}}class="active"{{end}}>English</a>
        </div>
        <h1>Dork<span>Discovery</span></h1>
        
        <div class="ust-kisim">
            <form method="GET">
                <input type="hidden" name="lang" value="{{.Lang}}">
                <input type="text" name="hedef" placeholder="{{if eq .Lang "en"}}Target domain.{{else}}Hedef domaini yazın.{{end}}" value="{{.HedefSite}}">
                <button type="submit">{{if eq .Lang "en"}}START ANALYSIS{{else}}ANALİZİ BAŞLAT{{end}}</button>
            </form>
        </div>

        {{if .HedefSite}}
        <p style="text-align: center; color: #a855f7;">{{if eq .Lang "en"}}Target:{{else}}Hedef:{{end}} <b style="color:white">{{.HedefSite}}</b> | {{if eq .Lang "en"}}Total dorks loaded:{{else}}Toplam dork yüklendi:{{end}} <b>{{.ToplamSayi}}</b></p>
        
        <div class="liste">
            {{range .DorkGruplari}}
                <div class="kategori-ismi">{{if eq $.Lang "en"}}{{.KategoriAdi_en}}{{else}}{{.KategoriAdi_tr}}{{end}}</div>
                {{range .Dorklar}}
                <div class="kart">
                    <span class="id-etiket">{{.Id}}</span>
                    <p class="aciklama">{{if eq $.Lang "en"}}{{.Aciklama_en}}{{else}}{{.Aciklama_tr}}{{end}}</p>
                    <code>site:{{$.HedefSite}} {{.Sorgu}}</code>
                    <a href="https://www.google.com/search?q=site:{{$.HedefSite}}+{{.Sorgu}}&tbs=li:1" target="_blank" class="link">{{if eq $.Lang "en"}}SEARCH ON GOOGLE{{else}}GOOGLE'DA GÖSTER{{end}}</a>
                </div>
                {{end}}
            {{end}}
        </div>
        {{else}}
        <p style="text-align: center; color: #6b7280;">{{if eq .Lang "en"}}System ready. {{.ToplamSayi}} GHDB dorks loaded. Enter a target above.{{else}}Sistem hazır. {{.ToplamSayi}} adet GHDB dorku yüklenmiştir. Yukarıya bir hedef girin.{{end}}</p>
        {{end}}

        <div class="bilgi">
            {{if eq .Lang "en"}}YILDIZ CTI PROJECT - 2026<br>Educational purpose only.{{else}}YILDIZ CTI PROJESİ - 2026<br>Siber Vatan Programı Eğitim Materyalidir.{{end}}
        </div>
    </div>
</body>
</html>
`
