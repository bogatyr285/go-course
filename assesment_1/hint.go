/*
Реализуйте запись сообщений от разных пользователей в файлы следующим образом:

Выделите канал для записи сообщений каждому файлу.
При получении сообщения с корректным токеном, запишите сообщение в выделенный
канал файла. Чтение из каждого канала реализуйте в отдельной горутине.
Запишите данные из канала в кеш файла.
Отдельный воркер WriteFiles должен каждую секунду проходится по кешу файлов и записывать
пришедшие изменения в файл.

Несколько пользователей могут писать в один и тот же файл.

АПИ вашего приложения - функции AddUser и SendMsg, которые будут использоваться в
разных вариациях при тестировании вашего приложения.
*/
package main

import "context"

type User struct {
	token string
	file  string
}

var ctx context.Context

func main() {
	ctx = context.TODO()

	go WriteFiles(ctx)

	// АПИ приложения
	AddUser(User{"123456789", "file1.txt"})
	SendMsg("123456789", "test message")
}

// AddUser добавляет писателя в кеш, выделяет канал для записи в файл
func AddUser(wr User) {

	// save User to cache

	go WriteMsgs2Cache(ctx, getFileCh(wr.file))
}

// SendMsg проверяет токен пользователя и записывает пришедшее сообщение
// в канал соответствующего файла
func SendMsg(token string, msg string) {

}

// WriteMsgs2Cache читает канал файла и записывает сообщение в кеш
// соответствующего файла
func WriteMsgs2Cache(ctx context.Context, ch <-chan string) {

}

// WriteFiles проверяет кеши всех файлов и записывает пришедшие изменения в
// указанный файл писателя
func WriteFiles(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
		}
	}
}

func getFileCh(filename string) <-chan string {
	return nil
}
