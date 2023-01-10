package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/XinceChan/go-blog-backend/utils"
	"github.com/spf13/viper"
)

type Config struct {
	systemConfig
	userConfig
}

var Cfg Config

func init() {
	var err error

	Cfg.CurrentDir, err = os.Getwd()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.SetConfigFile(Cfg.CurrentDir + "/config.json")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	initConfig(&Cfg)

	if Cfg.DashboardEntrance == "" || !strings.HasPrefix(Cfg.DashboardEntrance, "/") {
		Cfg.DashboardEntrance = "/admin"
	}

	repoName, err := utils.GetRepoName(Cfg.DocumentGitUrl)
	if err != nil {
		panic(err)
	}

	Cfg.AppName = "TwenBlog"
	Cfg.Version = 3.0
	Cfg.DocumentDir = Cfg.CurrentDir + "/" + repoName
	Cfg.GitHookUrl = "/api/git_push_hook"
	Cfg.AppRepository = "https://github.com/XinceChan/TwenBlog"
}

func initConfig(cfg *Config) {
	cfg.SiteName = viper.GetString("siteName")                            // sitename
	cfg.Author = viper.GetString("author")                                // author
	cfg.Icp = viper.GetString("icp")                                      // icp
	cfg.TimeLayout = viper.GetString("timeLayout")                        // timeLayout
	cfg.Port = viper.GetInt("port")                                       // port
	cfg.WebHookSecret = viper.GetString("webHookSecret")                  // webHookSecret
	cfg.CategoryDisplayQuantity = viper.GetInt("categoryDisplayQuantity") // categoryDisplayQuantity
	cfg.TagDisplayQuantity = viper.GetInt("tagDisplayQuantity")           // tagDisplayQuantity
	cfg.UtterancesRepo = viper.GetString("utterancesRepo")                // utterancesRepo
	cfg.PageSize = viper.GetInt("pageSize")                               // pageSize
	cfg.DescriptionLen = viper.GetInt("descriptionLen")                   // descriptionLen
	cfg.DocumentGitUrl = viper.GetString("documentGitUrl")                // documentGitUrl
	cfg.HtmlKeywords = viper.GetString("htmlKeywords")                    // htmlKeywords
	cfg.ThemeColor = viper.GetString("themeColor")                        // themeColor
	cfg.ThemeOption = viper.GetStringSlice("themeOption")                 // themeOption
	cfg.DashboardEntrance = viper.GetString("dashboardEntrance")          // dashboardEntrance
}
