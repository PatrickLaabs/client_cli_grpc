/*
Package cmd

Copyright © 2022 Patrick Laabs <patrick.laabs@me.com>

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

const (
	defaultName = "world"
)

// byeCmd represents the bye command
var byeCmd = &cobra.Command{
	Use:   "bye",
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
		c := pb.NewByeClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayBye(ctx, &pb.ByeRequest{Name: *name})
		if err != nil {
			log.Fatalf("could not send bye: %v", err)
		}
		log.Printf("Bye: %s", r.GetMessage())
		fmt.Println("bye called")
	},
}

var (
	name = flag.String("name", defaultName, "Name to greet")
)

func init() {
	rootCmd.AddCommand(byeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
