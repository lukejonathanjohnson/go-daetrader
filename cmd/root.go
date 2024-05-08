/*
Copyright © 2024 luckydaem0n <@luckydaem0n>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

const logo = `
██████╗░░█████╗░███████╗████████╗██████╗░░█████╗░██████╗░███████╗██████╗░
██╔══██╗██╔══██╗██╔════╝╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗
██║░░██║███████║█████╗░░░░░██║░░░██████╔╝███████║██║░░██║█████╗░░██████╔╝
██║░░██║██╔══██║██╔══╝░░░░░██║░░░██╔══██╗██╔══██║██║░░██║██╔══╝░░██╔══██╗
██████╔╝██║░░██║███████╗░░░██║░░░██║░░██║██║░░██║██████╔╝███████╗██║░░██║
╚═════╝░╚═╝░░╚═╝╚══════╝░░░╚═╝░░░╚═╝░░╚═╝╚═╝░░╚═╝╚═════╝░╚══════╝╚═╝░░╚═╝
`

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF41")).Bold(true)
	//tipMsgStyle    = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Italic(true)
	//endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

var titleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFDF5")).
	Background(lipgloss.Color("#25A065")).
	Padding(0, 1).
	Bold(true)

type item struct {
	title, description string
}

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func triggerStart() tea.Msg {
	return startMsg{}
}

type startMsg struct{}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			selectedItem := m.list.SelectedItem().(item) // Ensure this type assertion is safe
			if selectedItem.title == "Start" {
				// If 'Start' is selected, trigger the start command
				return m, triggerStart
			}
		}
		if msg.String() == "q" || msg.String() == "esc" {
			return m, tea.Quit
		}

	case startMsg:
		// Handle start command execution
		go func() {
			fmt.Println("executing start command...")
			if err := startCmd.Execute(); err != nil {
				fmt.Printf("Error executing start command: %v\n", err)
			}
		}()
		return m, tea.Quit // Exit the Bubble Tea program after starting the command

	case tea.WindowSizeMsg:
		height := msg.Height - 45
		m.list.SetSize(msg.Width, height)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "daetrader",
	Short: "EVM blockchain CLI interface",
	Long:  `Daetrader is a Go CLI for executing EVM transactions and custom trading stategies`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("%s\n", logoStyle.Render(logo))
		//fmt.Printf("%s\n", tipMsgStyle.Render("Access daetrader CLI: go-daetrader start"))
		//fmt.Printf("%s\n", endingMsgStyle.Render("luckydaem0n © 2024"))
		items := []list.Item{
			item{title: "Start", description: "Start the application"},
			item{title: "Version", description: "Display version information"},
			item{title: "Help", description: "Show help options"},
		}

		listModel := list.New(items, list.NewDefaultDelegate(), 0, 0)
		listModel.Title = "daetrader/"
		listModel.Styles.Title = titleStyle

		p := tea.NewProgram(model{list: listModel})

		fmt.Printf("%s\n", logoStyle.Render(logo))
		fmt.Println("Welcome to the Daetrader CLI!")
		fmt.Println("Please use the arrow keys to navigate the menu, and press enter to select an option.")
		fmt.Println("")
		if err, exitMsg := p.Run(); err != nil {
			//fmt.Printf("Error running program: %s\n", err)
			//os.Exit(1)
		} else {
			fmt.Printf("Program exited: %s\n", exitMsg)
		}
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-daetrader.yaml)")
	//fmt.Printf("%s\n", logoStyle.Render(logo))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
