package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

type Limiter struct {
	cond  sync.Cond
	cache sync.Map
	limit int
}

func New(limit int) *Limiter {

	return &Limiter{
		cond:  sync.Cond{L: new(sync.Mutex)},
		limit: limit,
		cache: sync.Map{},
	}
}
func (lim *Limiter) LimitLock(key string) {
	lim.cond.L.Lock()
	for !lim.CheckAndStore(key) {
		lim.cond.Wait()
	}
	lim.cond.L.Unlock()
}
func (lim *Limiter) LimitUnLock(key string) {
	lim.cond.L.Lock()
	lim.CheckAndDelete(key)
	lim.cond.Signal()
	lim.cond.L.Unlock()
}
func (lim *Limiter) CheckAndDelete(key string) {
	if value, load := lim.cache.Load(key); load {
		if value.(int) <= 1 {
			lim.cache.Delete(key)
			return
		}
		lim.cache.Store(key, value.(int)-1)
	}
}
func (lim *Limiter) CheckAndStore(key string) bool {
	actual := 0
	if value, load := lim.cache.Load(key); load {
		if value.(int) >= lim.limit {
			return false
		}
		actual = value.(int)
	}
	lim.cache.Store(key, actual+1)
	return true
}

// Middleware is the middleware for gin.
type Middleware struct {
	Limiter *Limiter
}

// NewBeeMiddleware return a new instance of a beego middleware.
func NewBeeMiddleware(limiter *Limiter) *Middleware {

	middleware := &Middleware{
		Limiter: limiter,
	}
	return middleware
}
func (middleware *Middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ind := strings.LastIndex(r.RemoteAddr, ":")
		key := r.RemoteAddr[:ind]
		middleware.Limiter.LimitLock(key)
		next.ServeHTTP(w, r)
		middleware.Limiter.LimitUnLock(key)

	})
}

const (
	stdLim  = 5
	stdAddr = "localhost:7777"
	ginLim  = 4
	ginAddr = "localhost:8888"
	beeLim  = 3
	beeAddr = "localhost:9999"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	//standart http server
	go func() {
		stdMiddleware := NewBeeMiddleware(New(stdLim))
		http.Handle("/", stdMiddleware.Handler(http.HandlerFunc(HandleTask)))
		fmt.Printf("Standart server is running on %s...", stdAddr)
		log.Fatal(http.ListenAndServe(stdAddr, nil))
	}()
	<-ctx.Done()
}

func HandleTask(writer http.ResponseWriter, request *http.Request) {
	// Simulate a resource-intensive task
	// time.Sleep(2 * time.Second)
	_, _ = fmt.Fprintf(writer, "Task from %s Completed", request.RemoteAddr)
}
