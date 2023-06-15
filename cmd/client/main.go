package main

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/client/api"
	"github.com/vasiliyantufev/gophkeeper/internal/client/config"
	"github.com/vasiliyantufev/gophkeeper/internal/client/gui"
	"github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//---------------------------------------------------------------------- client application init
	log := logrus.New()
	ctx := context.Background()
	config := config.NewConfig(log)
	log.SetLevel(config.DebugLevel)
	conn, err := grpc.Dial(config.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	grpc := gophkeeper.NewGophkeeperClient(conn)
	client := api.NewClient(ctx, log, grpc)
	ping, err := client.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(ping)
	//---------------------------------------------------------------------- fyne.app init
	application := app.New()
	application.Settings().SetTheme(theme.LightTheme())
	window := application.NewWindow("GophKeeper")
	window.Resize(fyne.NewSize(250, 80))
	gui.InitGUI(log, window, client)
}
