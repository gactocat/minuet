package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"
	"github.com/gactocat/minuet/pkg"
)

var (
	listStyle = lip.NewStyle().
			BorderStyle(lip.NormalBorder())

	cardListStyle = lip.NewStyle().
			BorderStyle(lip.NormalBorder())
	cardStyle = lip.NewStyle().
			Width(16).
			Height(11).
			BorderStyle(lip.NormalBorder())
)

type combo struct {
	hands []pkg.Hand
	field pkg.Field
}

func (c combo) Title() string {
	b := strings.Builder{}
	for i, hand := range c.hands {
		if i != 0 {
			b.WriteString(" + ")
		}
		b.WriteString("No")
		b.WriteString(strconv.Itoa(hand.CardNo))
	}
	return b.String()
}

func (c combo) Description() string {
	b := strings.Builder{}
	for i, hand := range c.hands {
		if i != 0 {
			b.WriteRune('\n')
		}
		b.WriteString(pkg.NoCardMap[hand.CardNo].Name)
	}
	return b.String()
}

func (c combo) FilterValue() string {
	b := strings.Builder{}
	for i, hand := range c.hands {
		if i != 0 {
			b.WriteRune('+')
		}
		b.WriteString(strconv.Itoa(hand.CardNo))
	}
	return b.String()
}

type model struct {
	width  int
	height int

	cards []pkg.Card
	list  list.Model
}

type tickMsg time.Time

func main() {
	result := map[int]map[int]pkg.Field{}
	combos := []list.Item{}
	for i, c1 := range pkg.Cards {
		f := pkg.CreateField(30, 30)
		f, _ = f.Put(pkg.Pos{X: 12, Y: 12}, c1)
		for j := i + 1; j < len(pkg.Cards); j++ {
			c2 := pkg.Cards[j]
			for _, fp := range f.GetSidePoints() {
				for _, d := range pkg.AllDir {
					tc := c2.Turn(d)
					for _, cp := range tc.GetFilledPoint() {
						p := pkg.Pos{X: fp.X - cp.X, Y: fp.Y - cp.Y}
						f, err := f.Put(p, tc)
						if err != nil {
							continue
						}
						if f.GetSpecialPoint() < 2 {
							continue
						}
						if _, ok := result[c1.No]; !ok {
							result[c1.No] = map[int]pkg.Field{}
						}
						if _, ok := result[c1.No][c2.No]; !ok {
							result[c1.No][c2.No] = f
							combos = append(combos, combo{
								hands: []pkg.Hand{
									{
										Pos:    pkg.Pos{X: 12, Y: 12},
										CardNo: c1.No,
										Dir:    pkg.UP,
									},
									{
										Pos:    p,
										CardNo: c2.No,
										Dir:    d,
									},
								},
								field: f,
							})
						}
					}
				}
			}
		}
	}

	m := model{
		cards: pkg.Cards[0:4],
		list:  list.New(combos, list.NewDefaultDelegate(), 0, 0),
	}
	m.list.Title = "Combos"
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		h, v := listStyle.GetFrameSize()
		m.list.SetSize((msg.Width)/3-h, msg.Height-v)
		cardListStyle.Width(msg.Width / 3 * 2)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	doc := strings.Builder{}

	renderedCards := make([]string, 0, len(m.cards))
	combo := m.list.SelectedItem().(combo)
	for _, hand := range combo.hands {
		card := pkg.NoCardMap[hand.CardNo]
		renderedCards = append(renderedCards, cardStyle.Render(card.String()))
	}

	doc.WriteString(
		lip.JoinHorizontal(lip.Top,
			listStyle.Render(m.list.View()),
			lip.JoinVertical(lip.Left,
				lip.JoinHorizontal(lip.Top,
					renderedCards...,
				),
				lip.NewStyle().
					BorderStyle(lip.NormalBorder()).
					Render(combo.field.String()),
			),
		))
	return doc.String()
}
