# Todo Yönetim API'si

Bu proje, kullanıcıların yapılacaklar listelerini (Todo listesi) yönetmelerine olanak tanıyan bir RESTful API'dir. Proje, Go dilinde yazılmıştır ve yüksek performanslı bir web çerçevesi olan Fiber kullanılarak geliştirilmiştir. Veritabanı olarak MongoDB kullanılmıştır. Proje, temiz kod yazımını ve test edilebilirliği sağlamak amacıyla bağımlılık enjeksiyonu prensiplerini de içermektedir.

## Proje Amacı

Bu projenin temel amacı, yapılacaklar listelerini yönetmek için bir API sağlamaktır. Kullanıcılar, bu API aracılığıyla yeni yapılacaklar (Todo) ekleyebilir, mevcut yapılacakları listeleyebilir ve tamamlananları veya istemedikleri maddeleri silebilirler. Bu API:

- **Veri Yönetimi**: Todo öğelerinin MongoDB'de depolanmasını sağlar.
- **RESTful Hizmetler**: HTTP istekleri aracılığıyla CRUD (Create, Read, Update, Delete) işlemlerini gerçekleştirir.
- **Test Edilebilirlik**: Mock'lar ve birim testlerle, sistemin kararlılığı ve güvenilirliği sağlanır.
- **Kolay Yönetim**: Bağımlılık enjeksiyonu sayesinde servis ve repository katmanlarının kolayca yönetilmesi ve test edilmesi sağlanır.

## Projenin Sağladığı İşlevler

- **Todo Oluşturma**: Kullanıcılar yeni yapılacaklar ekleyebilir.
- **Tüm Todo'ları Listeleme**: Veritabanında kayıtlı olan tüm yapılacaklar listelenir.
- **Todo Silme**: Kullanıcılar belirli bir Todo öğesini ID'si üzerinden silebilir.
- **Veritabanı Entegrasyonu**: MongoDB kullanılarak Todo öğeleri güvenli bir şekilde depolanır ve yönetilir.

## Proje Nasıl Çalışır?

Proje, Go programlama dilinde geliştirilmiş ve Fiber web çerçevesi kullanılarak yapılandırılmıştır. Aşağıda proje çalışmasının detayları bulunmaktadır:

1. **Bağlantı ve Yapılandırma**: Uygulama başlatıldığında, `configs` paketindeki `ConnectDB` fonksiyonu çağrılarak MongoDB veritabanına bağlanılır. Bağlantı URI'si, `.env` dosyasında belirtilen `MONGOURI` ortam değişkeni üzerinden alınır.
2. **Veritabanı Bağlantısı**: `ConnectDB` fonksiyonu, veritabanı bağlantısını gerçekleştirir ve bağlantıyı `GetCollection` fonksiyonu kullanılarak belirli bir koleksiyona (örneğin, `todos`) yönlendirir.
3. **API Rotaları**: Uygulama, Fiber kullanarak HTTP rotalarını yapılandırır:
   - **POST /api/todo**: Yeni bir Todo oluşturur ve veritabanına ekler.
   - **GET /api/todos**: Tüm Todo öğelerini getirir.
   - **DELETE /api/todo/:id**: Belirtilen ID'ye sahip Todo öğesini siler.
4. **Servis ve Repository Katmanı**: `services` ve `repository` katmanları, uygulamanın iş mantığını ve veri erişim işlevlerini içerir. Repository katmanı, MongoDB'ye erişim sağlar ve Todo öğelerini yönetir. Servis katmanı, iş kurallarını içerir ve verilerin doğruluğunu kontrol eder.
5. **Mock ve Birim Testleri**: `gomock` kullanılarak servis ve repository katmanlarının mock'ları oluşturulmuştur. Bu sayede birim testlerde gerçek veritabanı erişimi olmadan uygulamanın işlevselliği test edilebilir.
6. **Uygulamanın Başlatılması**: `main.go` dosyasındaki `main` fonksiyonu, Fiber uygulamasını başlatır ve API'nin belirtilen port üzerinden dinlemeye başlamasını sağlar (`http://localhost:8080`).

### API Endpointleri

*Yöntem*       *Yol*              *Açıklama*
- POST	    `/api/todo`	        Yeni bir Todo oluşturur.
- GET	    `/api/todos`	    Tüm Todo öğelerini getirir.
- DELETE	`/api/todo/:id` 	Belirtilen ID'ye sahip Todo'yu siler.


## Kullanılan Teknolojiler

- **Go**: API'yi geliştirmek için kullanılan programlama dili.
- **Fiber**: Go için hızlı ve esnek bir web çerçevesi.
- **MongoDB**: NoSQL tabanlı veritabanı, Todo öğelerinin depolanmasında kullanılır.
- **GoMock**: Testler için mock nesneler oluşturmak ve iş mantığını test etmek için kullanılır.
- **Testify**: Testlerde doğrulama yapmak için kullanılır.
- **Godotenv**: `.env` dosyasından ortam değişkenlerini yüklemek için kullanılır.

## Kurulum ve Çalıştırma

Projeyi kurmak ve çalıştırmak için aşağıdaki adımları takip edebilirsiniz:

### Gereksinimler

- Go 1.20 veya daha yeni bir sürüm
- MongoDB kurulu ve erişilebilir durumda olmalıdır.
- Git

### Adım 1: Depoyu Klonlayın

```bash
git clone https://github.com/erenerdogmus/todo-management-api.git
cd todo-management-api


├── todo-management-api/
                │   ├── app/
                │   │   ├── handlers.go          # API handler'ları ve işlevleri
                │   │   ├── handlers_test.go     # Handler testleri
                │   ├── configs/
                │   │   ├── db.go                # MongoDB bağlantı ayarları
                │   │   ├── env.go               # Ortam değişkenlerini okuma
                │   ├── models/
                │   │   ├── todo.go              # Todo model tanımı
                │   ├── repository/
                │   │   ├── todo_repository.go   # Todo veritabanı işlemleri
                │   │   ├── todo_repository_test.go # Veritabanı mock testleri
                │   ├── services/
                │   │   ├── todo_service.go      # Todo iş kuralları ve servisler
                │   │   ├── todo_service_test.go # Servis katmanı testleri
                │   ├── dto/
                │   │   ├── todo_dto.go          # Data transfer objeleri
                │   ├── main.go                  # Giriş noktası
                │   ├── .env                     # Ortam değişkenleri
                │   ├── go.mod                   # Go modülü tanımı ve bağımlılıklar
```

### Ortam Değişkenlerini Ayarlayın
 - Projenin kök dizininde bir .env dosyası oluşturun ve MongoDB bağlantı URI'sini belirtin:

```bash
MONGOURI=MONGOURI=mongodb+srv://username:password@Clusters.mongodb.net/
```

### Testler
 - Birim testleri çalıştırmak için şu komutu kullanabilirsiniz:

```bash
go test ./...
```

### Run

```bash
go run main.go
```
