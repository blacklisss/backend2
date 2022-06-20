package main

import (
	"backend2/internal/infrastructure/db/pgstore"
	"backend2/internal/infrastructure/server"
	"backend2/internal/usecases/app/repos/envrepo"
	"backend2/internal/usecases/app/repos/userrepo"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	db, err := pgstore.NewDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	ust := pgstore.NewUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	un := userrepo.NewUsers(ust)

	est := pgstore.NewEnvs(db)
	if err != nil {
		log.Fatal(err)
	}
	en := envrepo.NewEnvs(est)

	enrst := pgstore.NewEnvRels(db)
	if err != nil {
		log.Fatal(err)
	}
	enr := envrepo.NewEnvsRel(enrst)

	//hs := handlers.NewHandlers(ls)

	//h := routeropenapi.Handler(hs)
	srv := server.NewServer(":"+os.Getenv("PORT"), nil)

	srv.Start(un, en, enr)
	log.Print("Start")

	<-ctx.Done()

	srv.Stop()
	cancel()
	db.Close()

	log.Print("Exit")
}
