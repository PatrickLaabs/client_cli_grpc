/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/PatrickLaabs/grpc-qs/helloworld"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// eitcoCmd represents the eitco command
var eitcoCmd = &cobra.Command{
	Use:   "eitco",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		// Set up a connection to the server.
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewEitcoClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		//r, err := c.EitcoHello(ctx, &pb.EitcoRequest{Name: *name})
		r, err := c.SayEitco(ctx, &pb.HelloEitcoRequest{Name: *name})
		if err != nil {
			log.Fatalf("could not send eitco: %v", err)
		}
		log.Printf("Hallo: %s", r.GetMessage())
		fmt.Println("eitco called")
	},
}

func init() {
	rootCmd.AddCommand(eitcoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eitcoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eitcoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
