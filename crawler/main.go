package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	// максимально допустимое число ошибок при парсинге
	errorsLimit = 100000

	// число результатов, которые хотим получить
	resultsLimit = 10000
)

var (
	// адрес в интернете (например, https://en.wikipedia.org/wiki/Lionel_Messi)
	url string

	// насколько глубоко нам надо смотреть (например, 10)
	depthLimit int

	// общий таймаут
	defaultTimeout time.Duration = 1 * time.Second
)

// Как вы помните, функция инициализации стартует первой
func init() {
	// задаём и парсим флаги
	flag.StringVar(&url, "url", "", "url address")
	flag.IntVar(&depthLimit, "depth", 3, "max depth for run")
	flag.Parse()

	// Проверяем обязательное условие
	if url == "" {
		log.Print("no url set by flag")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	fmt.Printf("defaultTimeout: %v", defaultTimeout)
	started := time.Now()

	crawler := newCrawler(depthLimit)

	ctx, cancel := context.WithCancel(context.Background())
	go watchSignals(cancel, crawler)
	defer cancel()

	// создаём канал для результатов
	results := make(chan crawlResult)

	// запускаем горутину для чтения из каналов
	done := watchCrawler(ctx, results, errorsLimit, resultsLimit)

	// запуск основной логики
	// внутри есть рекурсивные запуски анализа в других горутинах
	crawler.run(ctx, url, results, 0)

	// ждём завершения работы чтения в своей горутине
	<-done

	log.Println(time.Since(started))
}

// ловим сигналы выключения
func watchSignals(cancel context.CancelFunc, c *crawler) {
	osSignalChan := make(chan os.Signal)

	signal.Notify(osSignalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGUSR1)

	for {

		sig := <-osSignalChan
		log.Printf("got signal %q", sig.String())

		if sig.(syscall.Signal) == syscall.SIGUSR1 {
			c.SetMaxDepth(10)
			log.Println("MaxDepth = 10")
			continue
		}
		// если сигнал получен, отменяем контекст работы
		cancel()
		break
	}
}

func watchCrawler(ctx context.Context, results <-chan crawlResult, maxErrors, maxResults int) chan struct{} {
	readersDone := make(chan struct{})

	go func() {
		defer close(readersDone)
		for {
			select {
			case <-ctx.Done():
				return

			case result := <-results:
				if result.err != nil {
					maxErrors--
					if maxErrors <= 0 {
						log.Println("max errors exceeded")
						return
					}
					continue
				}

				log.Printf("crawling result: %v", result.msg)
				maxResults--
				if maxResults <= 0 {
					log.Println("got max results")
					return
				}
			}
		}
	}()

	return readersDone
}
