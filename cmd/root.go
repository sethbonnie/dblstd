package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sethbonnie/dblstd/shape"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "dblstd <repo_path>",
	Version: "0.1.0",
	Short:   "checks if a repo conforms to a given standard",
	Long: `dblstd - short for DoubleStandards - checks if a project repo
conforms to a given standard (in the form of a "shape" file).

Prints any of the files in the shape that are missing from the repo.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		shapeFilename, err := cmd.Flags().GetString("shape-file")
		if err != nil {
			return err
		}
		if shapeFilename == "" {
			return fmt.Errorf("shape file must be specified")
		}

		var shapeSpec []byte

		// Check if the filename is actually a URL

		_, err = url.Parse(shapeFilename)
		if strings.HasPrefix(shapeFilename, "http") && err == nil {
			res, err := http.Get(shapeFilename)
			if err != nil {
				return err
			}

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}

			shapeSpec = data

		} else {
			data, err := ioutil.ReadFile(shapeFilename)
			if err != nil {
				return err
			}

			shapeSpec = data
		}

		s, err := shape.NewShape(shapeSpec)
		if err != nil {
			return err
		}
		rootDir := args[0]
		missing, err := s.Missing(rootDir)
		if err != nil {
			return err
		}
		for pathName, isDir := range missing {
			if isDir {
				fmt.Fprintf(os.Stderr, "⚠️ Missing required directory: %s\n", pathName)
			} else {
				fmt.Fprintf(os.Stderr, "⚠️ Missing required file: %s\n", pathName)
			}
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dblstd.yaml)")

	rootCmd.PersistentFlags().StringP("shape-file", "s", "", "(required) path to file containing expected shape of repo")
	rootCmd.MarkPersistentFlagRequired("shape-file")

	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".dblstd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dblstd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
