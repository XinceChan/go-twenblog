package config

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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
	cfg.Port = viper.GetString("port")                                    // port
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

func Initial() {
	if _, err := exec.LookPath("git"); err != nil {
		fmt.Println("请先安装git")
		panic(err)
	}
	if !utils.IsDir(Cfg.DocumentDir) {
		fmt.Println("正在克隆文档仓库，请稍等...")
		out, err := utils.RunCmdByDir(Cfg.CurrentDir, "git", "clone", Cfg.DocumentGitUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println(out)
	} else {
		out, err := utils.RunCmdByDir(Cfg.DocumentDir, "git", "pull")
		fmt.Println(out)
		if err != nil {
			panic(err)
		}
	}
	if err := checkDocDirAndBindConfig(&Cfg); err != nil {
		fmt.Println("文档缺少必要的目录")
		panic(err)
	}

	imgDir := Cfg.CurrentDir + "/images"
	if !utils.IsDir(imgDir) {
		if os.Mkdir(imgDir, os.ModePerm) != nil {
			panic("生成images目录失败!")
		}
	}
}

func checkDocDirAndBindConfig(cfg *Config) error {
	dirs := []string{"assets", "content", "extra_nav"}
	for _, dir := range dirs {
		absoluteDir := Cfg.DocumentDir + "/" + dir
		if !utils.IsDir(absoluteDir) {
			return errors.New("documents cannot lack " + absoluteDir + " dir")
		}
	}
	cfg.DocumentAssetsDir = cfg.DocumentDir + "/assets"
	cfg.DocumentContentDir = cfg.DocumentDir + "/content"
	cfg.DocumentExtraNavDir = cfg.DocumentDir + "/extra_nav"
	return nil
}
