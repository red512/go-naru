package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/red512/go-naru/internal"
	"github.com/red512/go-naru/pkg"
	"github.com/red512/go-naru/utils"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
)

var (
	Client = pkg.Client()
	V      = utils.Config
)

// func init() {
// 	v := viper.New()
// 	v.SetConfigName("config")
// 	v.SetConfigType("yaml")
// 	v.AddConfigPath("./")
// 	err := v.ReadInConfig()
// 	fmt.Println(v.GetString("app.env"))
// 	if err != nil {
// 		fmt.Println("Error when loading config file \n", err)
// 	}
// }

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "naru",
	Short: "naru - K8S data fetcher",
	Long:  `naru - a CLI to fetch K8S custom data and send to DB`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println(args[0])
	// },
}

var runServerCmd = &cobra.Command{
	Use:   "server",
	Short: "server - the command will run the server",
	Long:  `The command will run the server with echo`,
	Run: func(cmd *cobra.Command, args []string) {
		s := gocron.NewScheduler(time.UTC)
		s.Every(5).Seconds().Do(func() {
			ns, _ := pkg.GetNamespaces()
			data, _ := pkg.GetNamespacesData(ns)
			pkg.SendK8sDataToMongoDB(data, pkg.Client())
		})
		s.StartAsync()

		port := os.Getenv("MY_APP_PORT")
		if port == "" {
			port = "8080"
		}
		e := echo.New()
		e.GET("/namespaces", internal.GetDataFromMongoDBHandler)
		e.Logger.Print("Listening on port %s", port)
		e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
		utils.CmdExecutor("kubectl", "get", "po", "kubectl get po -ojson")
	},
}

var sendDataToMongo = &cobra.Command{
	Use:   "sendDataToMongo",
	Short: "sendDataToMongo - will send the namespace data to mongoDB",
	Long:  `The command will run will send the data to mongo`,
	Run: func(cmd *cobra.Command, args []string) {
		ns, _ := pkg.GetNamespaces()
		data, _ := pkg.GetNamespacesData(ns)
		pkg.SendK8sDataToMongoDB(data, pkg.Client())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(runServerCmd)
	rootCmd.AddCommand(sendDataToMongo)
}
