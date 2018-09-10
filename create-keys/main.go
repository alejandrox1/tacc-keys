package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Configurations stores the credentials necessary to interact with the Agave API.
type Configurations struct {
	TenantId     string `mapstructure:"tenantid" json:"tenantid"`
	BaseUrl      string `mapstructure:"baseurl" json:"baseurl"`
	ApiSecret    string `mapstructure:"apisecret" json:"apisecret"`
	ApiKey       string `mapstructure:"apikey" json:"apikey"`
	Username     string `mapstructure:"username" json:"username"`
	AccessToken  string `mapstructure:"access_token" json:"access_token"`
	RefreshToken string `mapstructure:"refresh_token" json:"refresh_token"`
	CreatedAt    string `mapstructure:"created_at" json:"created_at"`
	ExpiresIn    string `mapstructure:"expires_in" json:"expires_in"`
	ExpiresAt    string `mapstructure:"expires_at" json:"expires_at"`

	ConfigFile string
}

// SaveConfig updates the value of the configuration file based on the
// contents fo the Configurations struct.
func (c *Configurations) SaveConfig() error {
	// Open config file.
	configFile, err := os.OpenFile(c.ConfigFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("SaveConfig error opening file '%s' for read and write: %s  %s\n", c.ConfigFile, err)
		return err
	}
	defer configFile.Close()

	// Write values to file.
	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(c); err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return err
	}

	return err
}

var (
	username     string
	keysFilename string
)

func init() {
	pflag.StringVar(&username, "user", "", "Owner of public keys")
	pflag.StringVar(&keysFilename, "keys", "id_rsa", "Name of file to save rsa keys to")
}

func main() {
	// Parse command line arguments.
	pflag.Parse()

	// Read config file.
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	var conf Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading configuration file: %s\n", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Printf("Error decoding into struct: %s\n", err)
		os.Exit(1)
	}
	conf.ConfigFile = viper.ConfigFileUsed()

	// Refresh Token.
	createdAt, err := strconv.ParseInt(conf.CreatedAt, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ExpiresIn, err := strconv.ParseInt(conf.ExpiresIn, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	now := time.Now().Unix() - 100
	// Check if token needs to be refreshed.
	if (createdAt + ExpiresIn) < now {
		fmt.Fprintln(os.Stderr, "Refreshing token...")
		if err := conf.RefreshAPIToken(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Create rsa key pair.
	if err := SaveRSAKeysToFile(keysFilename); err != nil {
		fmt.Printf("Error creating RSA keys: %s\n", err)
		os.Exit(1)
	}

	// Post public keys to key server.
	publicKey, err := ioutil.ReadFile(keysFilename + ".pub")
	if err != nil {
		fmt.Printf("Error opening pubkey file: %s\n", err)
		os.Exit(1)
	}

	if err := conf.PostUserPubKey(username, string(publicKey)); err != nil {
		fmt.Printf("Error posting public key to server: %s\n", err)
		os.Exit(1)
	}
}
